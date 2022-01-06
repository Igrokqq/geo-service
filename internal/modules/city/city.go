package city

import (
	"github.com/vfilipovsky/geo-service/internal/modules/country"
	"github.com/vfilipovsky/geo-service/internal/modules/state"
)

type City struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"unique;not null;size:3"`
	Name string `gorm:"unique;not null;size:100"`

	CountryId uint `gorm:"not null"`
	Country   country.Country
	StateId   uint `gorm:"default:null"`
	State     state.State
}
