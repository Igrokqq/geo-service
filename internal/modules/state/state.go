package state

import "github.com/vfilipovsky/geo-service/internal/modules/country"

type State struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"unique;not null;size:3"`

	CountryId uint `gorm:"not null"`
	Country   country.Country
}
