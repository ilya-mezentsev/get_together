package repositories

import "models"

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
  }
  FirstUserStatuses = []string{
    "invited", "not-invited",
  }
)
