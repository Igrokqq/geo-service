package application

import (
	"github.com/vfilipovsky/geo-service/internal/domain/airport"
	"gorm.io/gorm"
)

type AirportService struct {
	r airport.Repository
}

func NewAirportService(db *gorm.DB) *AirportService {
	return &AirportService{
		r: nil,
	}
}
