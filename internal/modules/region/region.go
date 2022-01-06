package region

type Region struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null;size:50"`
}
