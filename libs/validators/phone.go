package validators

import "github.com/go-playground/validator/v10"

func Phone() validator.Func {
	return func(fl validator.FieldLevel) bool {
		if value, ok := fl.Field().Interface().(string); ok {
			return ValidateRegex(PhoneRegex, value)
		}
		return true
	}
}
