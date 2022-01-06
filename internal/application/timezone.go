package application

import (
	"errors"

	"github.com/vfilipovsky/geo-service/internal/domain/timezone"
	"github.com/vfilipovsky/geo-service/internal/infrastructure/persistance"
	"gorm.io/gorm"
)

type TimezoneService struct {
	r timezone.Repository
}

func (s *TimezoneService) Delete(timezone *timezone.Timezone) error {
	return errors.New("not implemented yet")
}

func (s *TimezoneService) Update() (*timezone.Timezone, error) {
	return nil, errors.New("not implemented yet")
}

func (s *TimezoneService) List() (*timezone.Timezones, error) {
	return nil, errors.New("not implemented yet")
}

func (s *TimezoneService) Create(timezone *timezone.Timezone) (*timezone.Timezone, error) {
	if err := s.r.Save(timezone); err != nil {
		return nil, err
	}

	return timezone, nil
}

func (s *TimezoneService) GetById(id uint) (*timezone.Timezone, error) {
	tz, err := s.r.Find(id)

	if err != nil {
		return nil, err
	}

	return tz, nil
}

func (s *TimezoneService) GetByName(name string) error {
	return errors.New("not implemented yet")
}

func NewTimezoneService(db *gorm.DB) *TimezoneService {
	return &TimezoneService{
		r: persistance.NewTimezoneRepository(db),
	}
}
