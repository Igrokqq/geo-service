package region

import (
	"github.com/jinzhu/gorm"
	"github.com/vfilipovsky/geo-service/internal/country"
)

type Region struct {
	gorm.Model
	Countries []country.Country
	Name      string
}
