package country

import (
	"github.com/vfilipovsky/geo-service/internal/city"
	"github.com/vfilipovsky/geo-service/internal/region"
	"github.com/vfilipovsky/geo-service/internal/state"
)

type Country struct {
	Region region.Region
	Cities []city.City
	States []state.State
	Code   string
	Name   string
}
