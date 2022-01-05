package persistance

import (
	"github.com/vfilipovsky/geo-service/internal/domain/timezone"
	"gorm.io/gorm"
)

type TimezoneRepository struct {
	db *gorm.DB
}

func NewTimezoneRepository(db *gorm.DB) *TimezoneRepository {
	return &TimezoneRepository{db: db}
}

func (r *TimezoneRepository) Find(id uint) (*timezone.Timezone, error) {
	var tz timezone.Timezone
	result := r.db.First(&tz, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &tz, nil
}

func (r *TimezoneRepository) FindByName(name string) (*timezone.Timezone, error) {
	var tz timezone.Timezone
	result := r.db.Where("name = ?", name).First(&tz)

	if result.Error != nil {
		return nil, result.Error
	}

	return &tz, nil
}

func (r *TimezoneRepository) Save(tz *timezone.Timezone) error {
	result := r.db.Create(tz)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *TimezoneRepository) SaveMany(timezones *timezone.Timezones) error {
	result := r.db.Create(timezones)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *TimezoneRepository) FindAll() (*timezone.Timezones, error) {
	var tzs *timezone.Timezones

	result := r.db.Find(tzs)

	if result.Error != nil {
		return nil, result.Error
	}

	return tzs, nil
}
