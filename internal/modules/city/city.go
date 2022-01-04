package city

import (
	"github.com/vfilipovsky/geo-service/internal/modules/country"
	"github.com/vfilipovsky/geo-service/internal/modules/state"
)

type City struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"unique;not null"`
	Name string `gorm:"unique;not null"`

	Country country.Country
	State   state.State
}
