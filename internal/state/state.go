package state

import (
	"github.com/jinzhu/gorm"
	"github.com/vfilipovsky/geo-service/internal/country"
)

type State struct {
	gorm.Model
	Country country.Country
	Name    string
	Code    string
}
