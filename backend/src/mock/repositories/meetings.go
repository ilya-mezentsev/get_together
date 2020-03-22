package repositories

import (
	"models"
	"time"
)

var (
	FirstUserStatuses = []string{
		"invited", "not-invited", "not-invited",
	}
	TimeCheckData = models.UserTimeCheckData{
		UserId:    3,
		MeetingId: 2,
	}
	ExpectedDate = time.Date(2020, 3, 2, 20, 0, 0, 0, time.UTC)
)

func GetFirstLabeledPlace() *models.LabeledPlace {
	return &models.LabeledPlace{
		Label: MeetingsPlaces[0]["label"].(string),
		PublicPlace: models.PublicPlace{
			Latitude:  models.Latitude(MeetingsPlaces[0]["latitude"].(float64)),
			Longitude: models.Longitude(MeetingsPlaces[0]["longitude"].(float64)),
		},
	}
}

func GetNotExistsMeetingId() uint {
	return uint(len(Meetings) + 1)
}

func GetPlaceLongitudeById(idx int) models.Longitude {
	return models.Longitude(MeetingsPlaces[idx]["longitude"].(float64))
}

func GetPlaceLatitudeById(idx int) models.Latitude {
	return models.Latitude(MeetingsPlaces[idx]["latitude"].(float64))
}
