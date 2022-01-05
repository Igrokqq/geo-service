package application

import (
	"github.com/vfilipovsky/geo-service/internal/domain/timezone"
	"github.com/vfilipovsky/geo-service/internal/infrastructure/persistance"
	"gorm.io/gorm"
)

type TimezoneService struct {
	r timezone.Repository
}

func NewService(db *gorm.DB) *TimezoneService {
	return &TimezoneService{
		r: persistance.NewTimezoneRepository(db),
	}
}
