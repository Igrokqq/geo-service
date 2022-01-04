package airport

import "github.com/vfilipovsky/geo-service/internal/city"

type Airport struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"unique;not null;size:3"`
	Name string `gorm:"unique;not null;size:50"`

	CityId uint `gorm:"not null"`
	City   city.City
}
