package application

import (
	//auth "accommodation_booking/common/domain"
	"accommodation_booking/reservation_service/domain"
	"context"
	"errors"
	"time"
)

type ReservationService struct {
	store domain.ReservationStore
}

func NewReservationService(store domain.ReservationStore) *ReservationService {
	return &ReservationService{
		store: store,
	}
}

func (service *ReservationService) Get(ctx context.Context, reservationId string) (*domain.Reservation, error) {
	return service.store.Get(ctx, reservationId)
}

func (service *ReservationService) GetAll(ctx context.Context) ([]*domain.Reservation, error) {
	return service.store.GetAll(ctx)
}

func (service *ReservationService) GetBetweenDates(ctx context.Context, beginning time.Time, ending time.Time, accommodationId string) ([]*domain.Reservation, error) {
	return service.store.GetBetweenDates(ctx, beginning, ending, accommodationId)
}

func (service *ReservationService) GetForUser(ctx context.Context, userId string, resType string) ([]*domain.Reservation, error) {
	return service.store.GetForUser(ctx, userId, resType)
}

func (service *ReservationService) GetForAccommodation(ctx context.Context, userId string) ([]*domain.Reservation, error) {
	return service.store.GetForAccommodation(ctx, userId)
}

func (service *ReservationService) GetPending(ctx context.Context) ([]*domain.Reservation, error) {
	return service.store.GetPending(ctx)
}

func (service *ReservationService) GetCanceled(ctx context.Context) ([]*domain.Reservation, error) {
	return service.store.GetCanceled(ctx)
}

func (service *ReservationService) GetApproved(ctx context.Context) ([]*domain.Reservation, error) {
	return service.store.GetApproved(ctx)
}

func (service *ReservationService) GetRejected(ctx context.Context) ([]*domain.Reservation, error) {
	return service.store.GetRejected(ctx)
}

func (service *ReservationService) Create(ctx context.Context, reservation *domain.Reservation) error {
	reservations, err := service.store.GetBetweenDates(ctx, reservation.Beginning, reservation.Ending, reservation.AccommodationId.Hex())
	if err != nil {
		return err
	}
	if len(reservations) != 0 {
		return errors.New("there is reservation overlapping selected time interval")
	}
	return service.store.Create(ctx, reservation)
}

func (service *ReservationService) RollbackUpdate(ctx context.Context, reservation *domain.Reservation) error {
	return service.store.Update(ctx, reservation.Id.Hex(), reservation)
}

func (service *ReservationService) Update(ctx context.Context, reservationId string, reservation *domain.Reservation) error {
	err := service.store.Update(ctx, reservationId, reservation)
	if err != nil {
		return err
	}
	return nil
}

func (service *ReservationService) Approve(ctx context.Context, reservationId string) error {
	reservation, err := service.store.Get(ctx, reservationId)
	if err != nil {
		return err
	}
	reservations, err := service.store.GetBetweenDatesPending(ctx, reservation.Beginning, reservation.Ending, reservation.AccommodationId.Hex())
	err = service.store.Approve(ctx, reservationId)
	if err != nil {
		return err
	}
	for _, res := range reservations {
		if res.Id.Hex() != reservationId {
			err = service.store.Reject(ctx, res.Id.Hex())
		}
		if err != nil {
			return nil
		}
	}
	return nil
}

func (service *ReservationService) Reject(ctx context.Context, reservationId string) error {
	err := service.store.Reject(ctx, reservationId)
	if err != nil {
		return err
	}
	return nil
}

func (service *ReservationService) Cancel(ctx context.Context, reservationId string) error {
	return service.store.Cancel(ctx, reservationId)
}

func (service *ReservationService) Delete(ctx context.Context, id string) error {
	return service.store.Delete(ctx, id)
}
