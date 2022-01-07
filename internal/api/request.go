package api

import (
	"errors"
	"strconv"

	"github.com/go-playground/validator"
)

var (
	ErrIdMustBeAPositiveNum = errors.New("id must be a positive number")
	Validate                = validator.New()
)

func ParseIdToUint(id string) (uint, error) {
	i, err := strconv.Atoi(id)

	if err != nil {
		return 0, ErrIdMustBeAPositiveNum
	}

	return uint(i), nil
}
