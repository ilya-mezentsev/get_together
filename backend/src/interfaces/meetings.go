package interfaces

import "models"

type (
  Place interface {
    GetLatitude() models.Latitude
    SetLatitude(l models.Latitude)
    GetLongitude() models.Longitude
    SetLongitude(l models.Longitude)
  }
)
