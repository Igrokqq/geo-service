package continent

import "gorm.io/gorm"

type Continents []Continent

type Repository interface {
	Find(id uint) (*Continent, error)
	FindByName(name string) (*Continent, error)
	FindAll() (*Continents, error)
	Save(continent *Continent) error
	SaveMany(continents *Continents) error
}

type ContinentRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &ContinentRepository{db: db}
}

func (r *ContinentRepository) Find(id uint) (*Continent, error) {
	var ct Continent

	result := r.db.First(&ct, id)

	return &ct, result.Error
}

func (r *ContinentRepository) FindByName(name string) (*Continent, error) {
	var ct Continent

	result := r.db.Where("name = ?", name).First(&ct)

	return &ct, result.Error
}

func (r *ContinentRepository) Save(ct *Continent) error {
	return r.db.Create(ct).Error
}

func (r *ContinentRepository) SaveMany(continents *Continents) error {
	return r.db.Create(continents).Error
}

func (r *ContinentRepository) FindAll() (*Continents, error) {
	var cts Continents

	result := r.db.Find(&cts)

	return &cts, result.Error
}
