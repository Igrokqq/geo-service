package state

import "github.com/vfilipovsky/geo-service/internal/country"

type State struct {
	Country country.Country
	Name    string
	Code    string
}
