package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/markphelps/optional"
)

func CustomValidationBool(fl validator.FieldLevel) bool {
	var value optional.Bool = fl.Field().Interface().(optional.Bool)
	if !value.Present() {
		return false
	}
	return true
}
func CustomValidationInt(fl validator.FieldLevel) bool {
	var value optional.Int = fl.Field().Interface().(optional.Int)
	if !value.Present() {
		return false
	}
	return true
}
func CustomValidationFloat(fl validator.FieldLevel) bool {
	var value optional.Float64 = fl.Field().Interface().(optional.Float64)
	if !value.Present() {
		return false
	}
	return true
}
