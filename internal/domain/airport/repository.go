package airport

type Airports []Airport

type Repository interface {
	Find(id uint) (*Airport, error)
	FindByName(name string) (*Airport, error)
	FindAll() (*Airports, error)
	Save(airport *Airport) error
	SaveMany(airports *Airports) error
}
