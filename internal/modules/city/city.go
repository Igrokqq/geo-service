package city

import (
	"github.com/vfilipovsky/geo-service/internal/modules/country"
	"github.com/vfilipovsky/geo-service/internal/modules/state"
)

type City struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"unique;not null"`
	Name string `gorm:"unique;not null"`

	CountryId uint
	Country   country.Country
	StateId   uint
	State     state.State
}
