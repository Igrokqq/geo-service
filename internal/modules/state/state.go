package state

import "github.com/vfilipovsky/geo-service/internal/modules/country"

type State struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"unique;not null"`

	CountryId uint
	Country   country.Country
}
