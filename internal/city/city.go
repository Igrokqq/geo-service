package city

import "github.com/vfilipovsky/geo-service/internal/country"

type City struct {
	Country country.Country
	Code    string
	Name    string
}
