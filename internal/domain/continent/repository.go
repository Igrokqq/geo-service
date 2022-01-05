package continent

type Continents []Continent

type Repository interface {
	Find(id uint) (*Continent, error)
	FindByName(name string) (*Continent, error)
	FindAll() (*Continents, error)
	Save(continent *Continent) error
	SaveMany(continents *Continents) error
}
