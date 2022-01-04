package airport

import (
	"github.com/vfilipovsky/geo-service/internal/modules/city"
	"github.com/vfilipovsky/geo-service/internal/modules/country"
	"github.com/vfilipovsky/geo-service/internal/modules/state"
)

type Airport struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"unique;not null"`
	Name string `gorm:"unique;not null"`

	Country country.Country
	State   state.State
	City    city.City
}
