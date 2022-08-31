package validators

import "github.com/go-playground/validator/v10"

func Url() validator.Func {
	return func(fl validator.FieldLevel) bool {
		if value, ok := fl.Field().Interface().(string); ok {
			return ValidateRegex(UrlRegex, value)
		}
		return true
	}
}
