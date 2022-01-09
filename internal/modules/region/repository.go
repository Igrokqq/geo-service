package region

import "gorm.io/gorm"

type Regions []Region

type Repository interface {
	Find(id uint) (*Region, error)
	FindByName(name string) (*Region, error)
	FindAll() (*Regions, error)
	Save(region *Region) error
	SaveMany(regions *Regions) error
	Delete(id uint) error
}

type RegionRepository struct {
	db *gorm.DB
}

func NewRegionRepository(db *gorm.DB) Repository {
	return &RegionRepository{db: db}
}

func (r *RegionRepository) Find(id uint) (*Region, error) {
	var rn Region

	result := r.db.First(&rn, id)

	return &rn, result.Error
}

func (r *RegionRepository) FindByName(name string) (*Region, error) {
	var rn Region

	result := r.db.Where("name = ?", name).First(&rn)

	return &rn, result.Error
}

func (r *RegionRepository) Save(rn *Region) error {
	return r.db.Create(rn).Error
}

func (r *RegionRepository) SaveMany(regions *Regions) error {
	return r.db.Create(regions).Error
}

func (r *RegionRepository) FindAll() (*Regions, error) {
	var rns Regions

	result := r.db.Find(&rns)

	return &rns, result.Error
}

func (r *RegionRepository) Delete(id uint) error {
	return r.db.Delete(&Region{}, id).Error
}
