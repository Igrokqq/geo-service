package city

type Cities []City

type Repository interface {
	Find(id uint) (*City, error)
	FindByName(name string) (*City, error)
	FindAll() (*Cities, error)
	Save(city *City) error
	SaveMany(cities *Cities) error
}
