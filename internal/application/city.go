package application

import (
	"github.com/vfilipovsky/geo-service/internal/domain/city"
	"gorm.io/gorm"
)

type CityService struct {
	r city.Repository
}

func NewCityService(db *gorm.DB) *CityService {
	return &CityService{
		r: nil,
	}
}
