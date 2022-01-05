package timezone

import (
	"github.com/vfilipovsky/geo-service/internal/domain/timezone"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Find(id uint) (*timezone.Timezone, error) {
	var tz timezone.Timezone
	result := r.db.First(&tz, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &tz, nil
}

func (r *Repository) FindByName(name string) (*timezone.Timezone, error) {
	var tz timezone.Timezone
	result := r.db.Where("name = ?", name).First(&tz)

	if result.Error != nil {
		return nil, result.Error
	}

	return &tz, nil
}

func (r *Repository) Save(tz *timezone.Timezone) error {
	result := r.db.Create(tz)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) SaveMany(timezones []*timezone.Timezone) error {
	result := r.db.Create(timezones)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) FindAll() ([]*timezone.Timezone, error) {
	var tzs []*timezone.Timezone

	result := r.db.Find(tzs)

	if result.Error != nil {
		return nil, result.Error
	}

	return tzs, nil
}
