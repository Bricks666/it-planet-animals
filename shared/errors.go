package shared

import (
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func init() {
	validate.RegisterValidation("notblank", NotBlank)
	validate.RegisterValidation("iso-8601", ISO8601)
}

func ValidateStruct[T interface{}](obj *T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(obj)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func NotBlank(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		return len(strings.TrimSpace(field.String())) > 0
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Array:
		return field.Len() > 0
	case reflect.Ptr, reflect.Interface, reflect.Func:
		return !field.IsNil()
	default:
		return field.IsValid() && field.Interface() != reflect.Zero(field.Type()).Interface()
	}
}

var ISO8601Layout = "2022-09-27 18:00:00.000"

func ISO8601(fl validator.FieldLevel) bool {
	field := fl.Field()
	if field.Kind() == reflect.String {
		_, err := time.Parse(ISO8601Layout, field.String())

		return err == nil
	}

	return false
}
