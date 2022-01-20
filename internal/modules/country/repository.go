package country

import "gorm.io/gorm"

type Countries []Country

type Repository interface {
	Find(id uint) (*Country, error)
	FindByName(name string) (*Country, error)
	FindAll() (*Countries, error)
	Save(country *Country) error
	SaveMany(countries *Countries) error
	Delete(id uint) error
}

type CountryRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &CountryRepository{
		db: db,
	}
}

func (r *CountryRepository) Find(id uint) (*Country, error) {
	var c Country

	result := r.db.Find(&c, id)

	return &c, result.Error
}

func (r *CountryRepository) FindByName(name string) (*Country, error) {
	var c Country

	result := r.db.Where("name = ?", name).First(&c)

	return &c, result.Error
}

func (r *CountryRepository) FindAll() (*Countries, error) {
	var cs Countries

	result := r.db.Find(&cs)

	return &cs, result.Error
}

func (r *CountryRepository) Save(country *Country) error {
	return r.db.Save(country).Error
}

func (r *CountryRepository) SaveMany(countries *Countries) error {
	return r.db.Save(countries).Error
}

func (r *CountryRepository) Delete(id uint) error {
	return r.db.Delete(&Country{}, id).Error
}
