package timezone

type Timezones []Timezone

type Repository interface {
	Find(id uint) (*Timezone, error)
	FindByName(name string) (*Timezone, error)
	FindAll() (*Timezones, error)
	Save(timezone *Timezone) error
	SaveMany(timezones *Timezones) error
}
