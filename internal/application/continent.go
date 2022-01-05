package application

import (
	"github.com/vfilipovsky/geo-service/internal/domain/continent"
	"github.com/vfilipovsky/geo-service/internal/infrastructure/persistance"
	"gorm.io/gorm"
)

type ContinentService struct {
	r continent.Repository
}

func NewContinentService(db *gorm.DB) *ContinentService {
	return &ContinentService{
		r: persistance.NewContinentRepository(db),
	}
}
