package persistance

import (
	"github.com/vfilipovsky/geo-service/internal/domain/continent"
	"gorm.io/gorm"
)

type ContinentRepository struct {
	db *gorm.DB
}

func NewContinentRepository(db *gorm.DB) *ContinentRepository {
	return &ContinentRepository{db: db}
}

func (r *ContinentRepository) Find(id uint) (*continent.Continent, error) {
	var ct continent.Continent
	result := r.db.First(&ct, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &ct, nil
}

func (r *ContinentRepository) FindByName(name string) (*continent.Continent, error) {
	var ct continent.Continent
	result := r.db.Where("name = ?", name).First(&ct)

	if result.Error != nil {
		return nil, result.Error
	}

	return &ct, nil
}

func (r *ContinentRepository) Save(ct *continent.Continent) error {
	result := r.db.Create(ct)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ContinentRepository) SaveMany(continents *continent.Continents) error {
	result := r.db.Create(continents)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ContinentRepository) FindAll() (*continent.Continents, error) {
	var cts *continent.Continents

	result := r.db.Find(cts)

	if result.Error != nil {
		return nil, result.Error
	}

	return cts, nil
}
