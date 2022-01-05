package application

import (
	"github.com/vfilipovsky/geo-service/internal/domain/country"
	"gorm.io/gorm"
)

type CountryService struct {
	r country.Repository
}

func NewCountryService(db *gorm.DB) *CountryService {
	return &CountryService{
		r: nil,
	}
}
