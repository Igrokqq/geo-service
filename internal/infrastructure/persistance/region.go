package persistance

import (
	"github.com/vfilipovsky/geo-service/internal/domain/region"
	"gorm.io/gorm"
)

type RegionRepository struct {
	db *gorm.DB
}

func NewRegionRepository(db *gorm.DB) *RegionRepository {
	return &RegionRepository{db: db}
}

func (r *RegionRepository) Find(id uint) (*region.Region, error) {
	var rn region.Region
	result := r.db.First(&rn, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &rn, nil
}

func (r *RegionRepository) FindByName(name string) (*region.Region, error) {
	var rn region.Region
	result := r.db.Where("name = ?", name).First(&rn)

	if result.Error != nil {
		return nil, result.Error
	}

	return &rn, nil
}

func (r *RegionRepository) Save(rn *region.Region) error {
	result := r.db.Create(rn)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *RegionRepository) SaveMany(regions *region.Regions) error {
	result := r.db.Create(regions)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *RegionRepository) FindAll() (*region.Regions, error) {
	var rns *region.Regions

	result := r.db.Find(rns)

	if result.Error != nil {
		return nil, result.Error
	}

	return rns, nil
}
