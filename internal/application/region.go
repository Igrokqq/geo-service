package application

import (
	"github.com/vfilipovsky/geo-service/internal/domain/region"
	"github.com/vfilipovsky/geo-service/internal/infrastructure/persistance"
	"gorm.io/gorm"
)

type RegionService struct {
	r region.Repository
}

func NewRegionService(db *gorm.DB) *RegionService {
	return &RegionService{
		r: persistance.NewRegionRepository(db),
	}
}
