package region

import (
	"github.com/vfilipovsky/geo-service/internal/country"
	"gorm.io/gorm"
)

type Region struct {
	gorm.Model
	Countries []country.Country
	Name      string
}
