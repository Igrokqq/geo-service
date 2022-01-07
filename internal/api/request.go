package api

import (
	"strconv"

	"github.com/go-playground/validator"
)

var Validate = validator.New()

func ParseIdToUint(id string) (uint, error) {
	i, err := strconv.Atoi(id)

	if err != nil {
		return 0, ErrIdMustBeAPositiveNum
	}

	return uint(i), nil
}
