package timezone

type Timezone struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique;not null;size:50" json:"name"`
}
