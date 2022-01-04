package timezone

type Timezone struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null;size:50"`
}
