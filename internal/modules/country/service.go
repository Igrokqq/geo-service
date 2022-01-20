package country

import (
	"gorm.io/gorm"

	"github.com/vfilipovsky/geo-service/internal/service"
)

type Service struct {
	r Repository
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		r: NewRepository(db),
	}
}

func (s *Service) GetById(id uint) *service.Response {
	return &service.Response{}
}
