package timezone

import (
	"errors"

	"gorm.io/gorm"
)

type Service struct {
	r Repository
}

func (s *Service) Delete(id uint) error {
	return s.r.Delete(id)
}

func (s *Service) Update() (*Timezone, error) {
	return nil, errors.New("not implemented yet")
}

func (s *Service) List() (*Timezones, error) {
	return s.r.FindAll()
}

func (s *Service) Create(timezone *Timezone) (*Timezone, error) {
	err := s.r.Save(timezone)

	return timezone, err
}

func (s *Service) GetById(id uint) (*Timezone, error) {
	return s.r.Find(id)
}

func (s *Service) GetByName(name string) error {
	return errors.New("not implemented yet")
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		r: NewRepository(db),
	}
}
