package region

import "gorm.io/gorm"

type Service struct {
	r Repository
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		r: NewRegionRepository(db),
	}
}
