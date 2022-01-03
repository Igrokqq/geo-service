package region

import "github.com/vfilipovsky/geo-service/internal/country"

type Region struct {
	Countries []country.Country
	Name      string
}
