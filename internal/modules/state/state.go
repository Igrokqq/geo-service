package state

type State struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"unique;not null"`
}
