package airport

import "gorm.io/gorm"

type Service struct {
	r Repository
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		r: nil,
	}
}
