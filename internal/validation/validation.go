package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var UserIDPattern *regexp.Regexp = regexp.MustCompile("[^0-9a-z._]+")

func UserIDValidator(fl validator.FieldLevel) bool {
	if UserIDPattern.Copy().MatchString(fl.Field().String()) {
		return true
	} else {
		return false
	}
}
