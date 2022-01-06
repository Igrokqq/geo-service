package country

type Countries []Country

type Repository interface {
	Find(id uint) (*Country, error)
	FindByName(name string) (*Country, error)
	FindAll() (*Countries, error)
	Save(country *Country) error
	SaveMany(countries *Countries) error
}
