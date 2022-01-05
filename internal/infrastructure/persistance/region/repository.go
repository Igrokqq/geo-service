package timezone

import (
	"github.com/vfilipovsky/geo-service/internal/domain/region"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Find(id uint) (*region.Region, error) {
	var rn region.Region
	result := r.db.First(&rn, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &rn, nil
}

func (r *Repository) FindByName(name string) (*region.Region, error) {
	var rn region.Region
	result := r.db.Where("name = ?", name).First(&rn)

	if result.Error != nil {
		return nil, result.Error
	}

	return &rn, nil
}

func (r *Repository) Save(rn *region.Region) error {
	result := r.db.Create(rn)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) SaveMany(regions *region.Regions) error {
	result := r.db.Create(regions)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) FindAll() (*region.Regions, error) {
	var rns *region.Regions

	result := r.db.Find(rns)

	if result.Error != nil {
		return nil, result.Error
	}

	return rns, nil
}
