package city

import (
	"github.com/jinzhu/gorm"
	"github.com/vfilipovsky/geo-service/internal/country"
)

type City struct {
	gorm.Model
	Country country.Country
	Code    string
	Name    string
}
