package region

type Regions []Region

type Repository interface {
	Find(id uint) (*Region, error)
	FindByName(name string) (*Region, error)
	FindAll() (*Regions, error)
	Save(region *Region) error
	SaveMany(regions *Regions) error
}
