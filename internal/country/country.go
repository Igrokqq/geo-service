package country

import (
	"github.com/jinzhu/gorm"
	"github.com/vfilipovsky/geo-service/internal/city"
	"github.com/vfilipovsky/geo-service/internal/region"
	"github.com/vfilipovsky/geo-service/internal/state"
)

type Country struct {
	gorm.Model
	Region region.Region
	Cities []city.City
	States []state.State
	Code   string
	Name   string
}
