package startup

import (
	"accommodation_booking/accommodation_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var accommodations = []*domain.Accommodation{
	{
		Id: getObjectId("bbbbbbbb9900000099000000"),
		Host: domain.Host{
			HostId:        getObjectId("aaaaaaaa0123456789000000"),
			Username:      "host1",
			PhoneNumber:   "060000000",
			IsOutstanding: false,
		},
		Name:               "Villa Woox",
		Location:           domain.Location{Country: "Bosna i Hercegovina", City: "Doboj", Street: "Marka Kraljevica 22"},
		HasWifi:            true,
		HasAirConditioning: true,
		HasFreeParking:     true,
		HasKitchen:         false,
		HasWashingMachine:  false,
		HasBathtub:         false,
		HasBalcony:         true,
		MinNumberOfGuests:  2,
		MaxNumberOfGuests:  7,
		Availability: []domain.AvailableDate{
			{
				Beginning:       time.Date(2023, time.June, 1, 0, 0, 0, 0, time.UTC),
				Ending:          time.Date(2023, time.June, 31, 23, 59, 59, 999, time.UTC),
				Price:           50,
				IsPricePerGuest: false,
			},
			{
				Beginning:       time.Date(2023, time.July, 1, 0, 0, 0, 0, time.UTC),
				Ending:          time.Date(2023, time.July, 31, 23, 59, 59, 999, time.UTC),
				Price:           60,
				IsPricePerGuest: false,
			},
		},
		IsReservationAcceptenceManual: false,
	},
	{
		Id: getObjectId("bbbbbbbb9900000099000001"),
		Host: domain.Host{
			HostId:        getObjectId("aaaaaaaa0123456789000000"),
			Username:      "host1",
			PhoneNumber:   "060000000",
			IsOutstanding: false,
		},
		Name:               "Villa Trnjanac",
		Location:           domain.Location{Country: "Bosna i Hercegovina", City: "Teslic", Street: "Glavna 31"},
		HasWifi:            true,
		HasAirConditioning: true,
		HasFreeParking:     false,
		HasKitchen:         true,
		HasWashingMachine:  false,
		HasBathtub:         false,
		HasBalcony:         false,
		MinNumberOfGuests:  3,
		MaxNumberOfGuests:  5,
		Availability: []domain.AvailableDate{
			{
				Beginning:       time.Date(2023, time.June, 1, 0, 0, 0, 0, time.UTC),
				Ending:          time.Date(2023, time.June, 30, 23, 59, 59, 999, time.UTC),
				Price:           40,
				IsPricePerGuest: false,
			},
			{
				Beginning:       time.Date(2023, time.July, 1, 0, 0, 0, 0, time.UTC),
				Ending:          time.Date(2023, time.July, 10, 23, 59, 59, 999, time.UTC),
				Price:           50,
				IsPricePerGuest: false,
			},
			{
				Beginning:       time.Date(2023, time.July, 21, 0, 0, 0, 0, time.UTC),
				Ending:          time.Date(2023, time.July, 31, 23, 59, 59, 999, time.UTC),
				Price:           50,
				IsPricePerGuest: false,
			},
		},
		IsReservationAcceptenceManual: false,
	},
	{
		Id: getObjectId("bbbbbbbb9900000199000002"),
		Host: domain.Host{
			HostId:        getObjectId("aaaaaaaa0123456789000001"),
			Username:      "host2",
			PhoneNumber:   "060111111",
			IsOutstanding: false,
		},
		Name:               "Villa Arta",
		Location:           domain.Location{Country: "Srbija", City: "Novi Sad", Street: "Maksima Gorkog 78"},
		HasWifi:            true,
		HasAirConditioning: true,
		HasFreeParking:     false,
		HasKitchen:         true,
		HasWashingMachine:  true,
		HasBathtub:         true,
		HasBalcony:         true,
		MinNumberOfGuests:  2,
		MaxNumberOfGuests:  3,
		Availability: []domain.AvailableDate{
			{
				Beginning:       time.Date(2023, time.June, 1, 0, 0, 0, 0, time.UTC),
				Ending:          time.Date(2023, time.June, 30, 23, 59, 59, 999, time.UTC),
				Price:           50,
				IsPricePerGuest: true,
			},
			{
				Beginning:       time.Date(2023, time.July, 1, 0, 0, 0, 0, time.UTC),
				Ending:          time.Date(2023, time.July, 10, 23, 59, 59, 999, time.UTC),
				Price:           60,
				IsPricePerGuest: true,
			},
		},
		IsReservationAcceptenceManual: true,
	},
	{
		Id: getObjectId("bbbbbbbb9900000299000003"),
		Host: domain.Host{
			HostId:        getObjectId("aaaaaaaa0123456789000002"),
			Username:      "host3",
			PhoneNumber:   "060222222",
			IsOutstanding: false,
		},
		Name:               "Villa Moskva",
		Location:           domain.Location{Country: "Srbija", City: "Beograd", Street: "Partizanska 138"},
		HasWifi:            true,
		HasAirConditioning: true,
		HasFreeParking:     false,
		HasKitchen:         true,
		HasWashingMachine:  true,
		HasBathtub:         true,
		HasBalcony:         true,
		MinNumberOfGuests:  1,
		MaxNumberOfGuests:  4,
		Availability: []domain.AvailableDate{
			{
				Beginning:       time.Date(2023, time.June, 1, 0, 0, 0, 0, time.UTC),
				Ending:          time.Date(2023, time.June, 30, 23, 59, 59, 999, time.UTC),
				Price:           100,
				IsPricePerGuest: false,
			},
			{
				Beginning:       time.Date(2023, time.July, 1, 0, 0, 0, 0, time.UTC),
				Ending:          time.Date(2023, time.July, 10, 23, 59, 59, 999, time.UTC),
				Price:           120,
				IsPricePerGuest: false,
			},
		},
		IsReservationAcceptenceManual: true,
	},
	{
		Id: getObjectId("bbbbbbbb9900000399000004"),
		Host: domain.Host{
			HostId:        getObjectId("aaaaaaaa0123456789000003"),
			Username:      "host4",
			PhoneNumber:   "060333333",
			IsOutstanding: false,
		},
		Name:               "Villa Maslacak",
		Location:           domain.Location{Country: "Srbija", City: "Valjevo", Street: "Ive Andrica 77"},
		HasWifi:            true,
		HasAirConditioning: true,
		HasFreeParking:     false,
		HasKitchen:         false,
		HasWashingMachine:  false,
		HasBathtub:         false,
		HasBalcony:         false,
		MinNumberOfGuests:  2,
		MaxNumberOfGuests:  5,
		Availability: []domain.AvailableDate{
			{
				Beginning:       time.Date(2023, time.June, 1, 0, 0, 0, 0, time.UTC),
				Ending:          time.Date(2023, time.June, 30, 23, 59, 59, 999, time.UTC),
				Price:           20,
				IsPricePerGuest: false,
			},
			{
				Beginning:       time.Date(2023, time.July, 1, 0, 0, 0, 0, time.UTC),
				Ending:          time.Date(2023, time.July, 10, 23, 59, 59, 999, time.UTC),
				Price:           25,
				IsPricePerGuest: false,
			},
		},
		IsReservationAcceptenceManual: true,
	},
	{
		Id: getObjectId("bbbbbbbb9900000499000005"),
		Host: domain.Host{
			HostId:        getObjectId("aaaaaaaa0123456789000004"),
			Username:      "host5",
			PhoneNumber:   "060444444",
			IsOutstanding: false,
		},
		Name:               "Villa Rodman",
		Location:           domain.Location{Country: "Srbija", City: "Pozarevac", Street: "Ulicnih sviraca 51"},
		HasWifi:            true,
		HasAirConditioning: true,
		HasFreeParking:     true,
		HasKitchen:         true,
		HasWashingMachine:  true,
		HasBathtub:         true,
		HasBalcony:         true,
		MinNumberOfGuests:  2,
		MaxNumberOfGuests:  10,
		Availability: []domain.AvailableDate{
			{
				Beginning:       time.Date(2023, time.June, 1, 0, 0, 0, 0, time.UTC),
				Ending:          time.Date(2023, time.June, 30, 23, 59, 59, 999, time.UTC),
				Price:           10,
				IsPricePerGuest: true,
			},
			{
				Beginning:       time.Date(2023, time.July, 1, 0, 0, 0, 0, time.UTC),
				Ending:          time.Date(2023, time.July, 10, 23, 59, 59, 999, time.UTC),
				Price:           12,
				IsPricePerGuest: true,
			},
		},
		IsReservationAcceptenceManual: true,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
