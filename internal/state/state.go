package state

import (
	"github.com/vfilipovsky/geo-service/internal/country"
	"gorm.io/gorm"
)

type State struct {
	gorm.Model
	Country country.Country
	Name    string
	Code    string
}
