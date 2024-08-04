package utils

import (
	"github.com/adrg/postcode"
)

func ZipValidation(zipCode string) bool {
	if err := postcode.Validate(zipCode); err != nil {
		return false
	}
	return true
}
