package city

import (
	"github.com/vfilipovsky/geo-service/internal/country"
	"gorm.io/gorm"
)

type City struct {
	gorm.Model
	Country country.Country
	Code    string
	Name    string
}
