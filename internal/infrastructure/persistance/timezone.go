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

	return &tz, result.Error
}

func (r *TimezoneRepository) FindByName(name string) (*timezone.Timezone, error) {
	var tz timezone.Timezone
	result := r.db.Where("name = ?", name).First(&tz)

	return &tz, result.Error
}

func (r *TimezoneRepository) Save(tz *timezone.Timezone) error {
	return r.db.Create(tz).Error
}

func (r *TimezoneRepository) SaveMany(timezones *timezone.Timezones) error {
	return r.db.Create(timezones).Error
}

func (r *TimezoneRepository) FindAll() (*timezone.Timezones, error) {
	var tzs timezone.Timezones

	result := r.db.Find(&tzs)

	return &tzs, result.Error
}

func (r *TimezoneRepository) Delete(id uint) error {
	return r.db.Delete(&timezone.Timezone{}, id).Error
}
