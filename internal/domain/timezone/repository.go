package timezone

type Repository interface {
	Find(id uint) (*Timezone, error)
	FindByName(name string) (*Timezone, error)
	FindAll() ([]*Timezone, error)
	Save(timezone *Timezone) error
	SaveMany(timezones []*Timezone) error
}
