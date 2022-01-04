package airport

import "github.com/vfilipovsky/geo-service/internal/modules/city"

type Airport struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"unique;not null"`
	Name string `gorm:"unique;not null"`

	CityId uint
	City   city.City
}
