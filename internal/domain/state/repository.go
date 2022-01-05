package state

type States []State

type Repository interface {
	Find(id uint) (*State, error)
	FindByName(name string) (*State, error)
	FindAll() (*States, error)
	Save(state *State) error
	SaveMany(states *States) error
}
