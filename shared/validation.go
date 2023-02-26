package shared

import (
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       interface{}
}

var Validator = validator.New()

func init() {
	Validator.RegisterValidation("notblank", NotBlank)
	Validator.RegisterValidation("iso-8601", ISO8601)
	Validator.RegisterValidation("any-number", AnyNumber)
}

func ValidateStruct(obj interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := Validator.Struct(obj)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Value()
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

func ISO8601(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		_, err := time.Parse(ISO8601_PATTER, field.String())
		return err == nil
	case reflect.Struct:
		if field.Type() == reflect.ValueOf(time.Time{}).Type() {
			return true
		}

		return false
	default:
		return false
	}

}

func AnyNumber(fl validator.FieldLevel) bool {
	field := fl.Field()
	switch field.Kind() {
	case reflect.Float32,
		reflect.Float64,
		reflect.Uint,
		reflect.Uint32,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint64,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		return true
	default:
		return false
	}
}
