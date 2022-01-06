package timezone

import "gorm.io/gorm"

type Timezones []Timezone

type Repository interface {
	Find(id uint) (*Timezone, error)
	FindByName(name string) (*Timezone, error)
	FindAll() (*Timezones, error)
	Save(timezone *Timezone) error
	SaveMany(timezones *Timezones) error
	Delete(id uint) error
}

type TimezoneRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &TimezoneRepository{db: db}
}

func (r *TimezoneRepository) Find(id uint) (*Timezone, error) {
	var tz Timezone

	result := r.db.First(&tz, id)

	return &tz, result.Error
}

func (r *TimezoneRepository) FindByName(name string) (*Timezone, error) {
	var tz Timezone

	result := r.db.Where("name = ?", name).First(&tz)

	return &tz, result.Error
}

func (r *TimezoneRepository) Save(tz *Timezone) error {
	return r.db.Create(tz).Error
}

func (r *TimezoneRepository) SaveMany(timezones *Timezones) error {
	return r.db.Create(timezones).Error
}

func (r *TimezoneRepository) FindAll() (*Timezones, error) {
	var tzs Timezones

	result := r.db.Find(&tzs)

	return &tzs, result.Error
}

func (r *TimezoneRepository) Delete(id uint) error {
	return r.db.Delete(&Timezone{}, id).Error
}
