package country

import (
	"errors"
	"gorm.io/gorm"
	"net/http"

	"github.com/vfilipovsky/geo-service/internal/service"
)

var ErrCountryNotFound = errors.New("country not found")

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

func (s *Service) GetByName(name string) *service.Response {
	country, err := s.r.FindByName(name)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &service.Response{Code: http.StatusNotFound, Error: ErrCountryNotFound}
	} else if err != nil {
		return &service.Response{Code: http.StatusInternalServerError, Error: err}
	}

	return &service.Response{Code: http.StatusOK, Value: &country}
}

func (s *Service) List() *service.Response {
	list, err := s.r.FindAll()

	if err != nil {
		return &service.Response{Code: http.StatusInternalServerError, Error: err}
	}

	return &service.Response{Code: http.StatusOK, Value: &list}
}
