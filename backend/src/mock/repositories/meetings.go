package repositories

import (
  "models"
  "time"
)

var (
  FirstLabeledPlace = models.LabeledPlace{
    Label: "221b baker street",
    PublicPlace: models.PublicPlace{
      Latitude: 51.5207,
      Longitude: -0.1550,
    },
  }
  PublicPlaces = []models.PublicPlace{
    {
      Latitude: 51.5207,
      Longitude: -0.155,
    },
    {
      Latitude: 0,
      Longitude: 0,
    },
    {
      Latitude: 51.5207,
      Longitude: -0.155,
    },
  }
  FirstUserStatuses = []string{
    "invited", "not-invited", "not-invited",
  }
  TimeCheckData = models.UserTimeCheckData{
    UserId: 3,
    MeetingId: 2,
  }
  ExpectedDate = time.Date(2020, 3, 2, 20, 0, 0, 0, time.UTC)
)
