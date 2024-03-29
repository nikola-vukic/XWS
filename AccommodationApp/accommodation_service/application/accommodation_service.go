package application

import (
	"accommodation_booking/accommodation_service/domain"
	"context"
	"sort"
	"time"
)

type AccommodationService struct {
	store domain.AccommodationStore
}

func NewAccommodationService(store domain.AccommodationStore) *AccommodationService {
	return &AccommodationService{
		store: store,
	}
}

func (service *AccommodationService) Get(ctx context.Context, accommodationId string) (*domain.Accommodation, error) {
	return service.store.Get(ctx, accommodationId)
}

func (service *AccommodationService) GetByHost(ctx context.Context, hostId string) ([]*domain.Accommodation, error) {
	return service.store.GetByHost(ctx, hostId)
}

func (service *AccommodationService) GetAll(ctx context.Context) ([]*domain.Accommodation, error) {
	return service.store.GetAll(ctx)
}

func (service *AccommodationService) GetAllFiltered(ctx context.Context, benefits domain.Benefits) ([]*domain.Accommodation, error) {
	return service.store.GetAllFiltered(ctx, benefits)
}

func (service *AccommodationService) GetAllSearched(ctx context.Context, location domain.Location, numberOfGuests int) ([]*domain.Accommodation, error) {
	return service.store.GetAllSearched(ctx, location, numberOfGuests)
}

func (service *AccommodationService) Create(ctx context.Context, accommodation *domain.Accommodation) (*domain.Accommodation, error) {
	return service.store.Create(ctx, accommodation)
}

func (service *AccommodationService) Update(ctx context.Context, acommodationId string, accommodation *domain.Accommodation) (*domain.Accommodation, error) {
	return service.store.Update(ctx, acommodationId, accommodation)
}

func (service *AccommodationService) UpdateAvailability(ctx context.Context, accommodationId string, newAvailableDate domain.AvailableDate) (*domain.Accommodation, error) {
	accommodation, err := service.Get(ctx, accommodationId)
	if err != nil {
		return nil, err
	}

	currentlyAvailableDates := accommodation.Availability
	if len(currentlyAvailableDates) == 0 {
		accommodation.Availability = append(accommodation.Availability, newAvailableDate)
		accommodation, err := service.Update(context.TODO(), accommodationId, accommodation)
		if err != nil {
			return nil, err
		}
		return accommodation, err
	}

	sort.Slice(currentlyAvailableDates, func(i, j int) bool {
		return currentlyAvailableDates[i].Beginning.Before(currentlyAvailableDates[j].Beginning)
	})

	// nova je podskup postojece
	var underlappingCompletely []int
	// nova je nadskup postojece
	var overlappingCompletely []int
	// overlapping beggining -> postojeći početak upada u novi
	var overlappingBeginning []int
	// overlapping ending -> postojeći kraj upada u novi
	var overlappingEnding []int
	var equals []int

	for i, currentlyAvailableDate := range currentlyAvailableDates {
		if newAvailableDate.Beginning.After(currentlyAvailableDate.Beginning) && newAvailableDate.Ending.Before(currentlyAvailableDate.Ending) {
			underlappingCompletely = append(underlappingCompletely, i)
		}
		if newAvailableDate.Beginning.Before(currentlyAvailableDate.Beginning) && newAvailableDate.Ending.After(currentlyAvailableDate.Ending) {
			overlappingCompletely = append(overlappingCompletely, i)
		}
		if newAvailableDate.Beginning.Before(currentlyAvailableDate.Beginning) && newAvailableDate.Ending.After(currentlyAvailableDate.Beginning) {
			overlappingBeginning = append(overlappingBeginning, i)
		}
		if newAvailableDate.Beginning.Before(currentlyAvailableDate.Ending) && newAvailableDate.Ending.After(currentlyAvailableDate.Ending) {
			overlappingEnding = append(overlappingEnding, i)
		}
		if newAvailableDate.Beginning.Equal(currentlyAvailableDate.Beginning) && newAvailableDate.Ending.Equal(currentlyAvailableDate.Ending) {
			equals = append(equals, i)
		}
	}

	if len(overlappingCompletely) == 0 && len(overlappingBeginning) == 0 && len(overlappingEnding) == 0 && len(equals) == 0 && len(underlappingCompletely) == 0 {
		accommodation.Availability = append(accommodation.Availability, newAvailableDate)
		accommodation, err := service.Update(context.TODO(), accommodationId, accommodation)
		if err != nil {
			return nil, err
		}
		return accommodation, err
	}

	if len(equals) != 0 {
		accommodation.Availability[equals[0]].Price = newAvailableDate.Price
		accommodation.Availability[equals[0]].IsPricePerGuest = newAvailableDate.IsPricePerGuest
		accommodation, err := service.Update(context.TODO(), accommodationId, accommodation)
		if err != nil {
			return nil, err
		}
		return accommodation, err
	}
	var additional domain.AvailableDate
	usedAdditional := false
	for _, currIdx := range underlappingCompletely {
		temp := domain.AvailableDate{
			Beginning:       newAvailableDate.Ending,
			Ending:          accommodation.Availability[currIdx].Ending,
			Price:           accommodation.Availability[currIdx].Price,
			IsPricePerGuest: accommodation.Availability[currIdx].IsPricePerGuest,
		}
		accommodation.Availability[currIdx].Ending = newAvailableDate.Beginning
		additional = temp
		usedAdditional = true
	}
	for _, currIdx := range overlappingBeginning {
		accommodation.Availability[currIdx].Beginning = newAvailableDate.Ending
	}

	for _, currIdx := range overlappingEnding {
		accommodation.Availability[currIdx].Ending = newAvailableDate.Beginning
	}

	for _, currIdx := range overlappingCompletely {
		accommodation.Availability[currIdx].Price = -1
	}

	var indicesToRemove []int
	for i, currDate := range accommodation.Availability {
		if currDate.Price == -1 {
			indicesToRemove = append(indicesToRemove, i)
		}
	}

	for i := len(indicesToRemove) - 1; i >= 0; i-- {
		index := indicesToRemove[i]
		accommodation.Availability = append(accommodation.Availability[:index], accommodation.Availability[index+1:]...)
	}

	accommodation.Availability = append(accommodation.Availability, newAvailableDate)

	if usedAdditional == true {
		accommodation.Availability = append(accommodation.Availability, additional)
	}

	sort.Slice(accommodation.Availability, func(i, j int) bool {
		return accommodation.Availability[i].Beginning.Before(accommodation.Availability[j].Beginning)
	})
	accommodation, err = service.Update(context.TODO(), accommodationId, accommodation)
	if err != nil {
		return nil, err
	}
	return accommodation, err
}

func (service *AccommodationService) GetAccommodationAvailableDatesForTimePeriod(ctx context.Context, accommodationId string, beginning time.Time, ending time.Time) ([]domain.AvailableDate, error) {
	accommodation, err := service.store.Get(ctx, accommodationId)
	if err != nil {
		return nil, err
	}
	var result []domain.AvailableDate
	for _, currDate := range accommodation.Availability {
		beginningChecksOut := currDate.Beginning.Before(beginning) || currDate.Beginning.Equal(beginning)
		endingChecksOut := currDate.Ending.After(ending) || currDate.Ending.Equal(ending)
		if beginningChecksOut && endingChecksOut {
			result = append(result, currDate)
		}
	}
	return result, nil
}
func (service *AccommodationService) Delete(ctx context.Context, id string) error {
	return service.store.Delete(ctx, id)
}
