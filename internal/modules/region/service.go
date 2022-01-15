package region

import (
	"errors"
	"net/http"

	"gorm.io/gorm"

	"github.com/vfilipovsky/geo-service/internal/service"
)

var ErrRegionNotFound = errors.New("region not found")

type Service struct {
	r Repository
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		r: NewRegionRepository(db),
	}
}

func (s *Service) GetById(id uint) *service.Response {
	region, err := s.r.Find(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &service.Response{Code: http.StatusNotFound, Error: ErrRegionNotFound}
	} else if err != nil {
		return &service.Response{Code: http.StatusInternalServerError, Error: err}
	}

	return &service.Response{Code: http.StatusOK, Value: &region}
}

func (s *Service) GetByName(name string) *service.Response {
	region, err := s.r.FindByName(name)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &service.Response{Code: http.StatusNotFound, Error: ErrRegionNotFound}
	} else if err != nil {
		return &service.Response{Code: http.StatusInternalServerError, Error: err}
	}

	return &service.Response{Code: http.StatusOK, Value: &region}
}

func (s *Service) List() *service.Response {
	list, err := s.r.FindAll()

	if err != nil {
		return &service.Response{Code: http.StatusInternalServerError, Error: err}
	}

	return &service.Response{Code: http.StatusOK, Value: &list}
}

func (s *Service) Create(req *createRequest) *service.Response {
	region := Region{Name: req.Name}
	exist, err := s.r.FindByName(req.Name)

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return &service.Response{Code: http.StatusOK, Value: exist}
	} else if err != nil {
		return &service.Response{Code: http.StatusInternalServerError, Error: err}
	}

	err = s.r.Save(&region)

	if err != nil {
		return &service.Response{Code: http.StatusInternalServerError, Error: err}
	}

	return &service.Response{Code: http.StatusOK, Value: &region}
}

func (s *Service) Update(id uint, req *updateRequest) *service.Response {
	exist, err := s.r.Find(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &service.Response{Code: http.StatusNotFound, Error: ErrRegionNotFound}
	} else if err != nil {
		return &service.Response{Code: http.StatusInternalServerError, Error: err}
	}

	exist.Name = req.Name
	err = s.r.Save(exist)

	if err != nil {
		return &service.Response{Code: http.StatusInternalServerError, Error: err}
	}

	return &service.Response{Code: http.StatusOK, Value: &exist}
}

func (s *Service) Delete(id uint) *service.Response {
	err := s.r.Delete(id)

	if err != nil {
		return &service.Response{Error: err, Code: http.StatusInternalServerError}
	}

	return &service.Response{Code: http.StatusOK}
}
