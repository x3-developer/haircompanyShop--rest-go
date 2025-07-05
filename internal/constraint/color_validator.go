package constraint

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func validateHexColor(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	if len(value) != 7 || value[0] != '#' {
		return false
	}

	hexPattern := regexp.MustCompile(`^#[0-9A-Fa-f]{6}$`)
	return hexPattern.MatchString(value)
}
