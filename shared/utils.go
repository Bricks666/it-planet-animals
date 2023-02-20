package shared

import "reflect"

func IsInstanceOf[T interface{}](instance *T, typePtr interface{}) bool {
	return reflect.TypeOf(instance).Name() == reflect.TypeOf(typePtr).Name()
}
