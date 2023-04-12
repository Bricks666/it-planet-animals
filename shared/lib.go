package shared

import "github.com/gofiber/fiber/v2"

type primitives interface {
	uint | uint16 | uint32 | uint64 | uint8 |
		int | int16 | int32 | int64 | int8 |
		float32 | float64 | string | bool
}

func HasOnlyUnique[T primitives](array []T) bool {
	var occurred = map[T]bool{}

	for _, e := range array {
		if occurred[e] {
			return false
		}

		occurred[e] = true
	}

	return true
}
func GetBody(ct *fiber.Ctx, dto interface{}) error {
	var err = ct.BodyParser(dto)
	if err != nil {
		return err
	}

	var validationError = ValidateStruct(dto)
	if validationError != nil {
		return fiber.ErrBadRequest
	}

	return nil
}

func GetParams(ct *fiber.Ctx, params interface{}) error {

	var err = ct.ParamsParser(params)
	if err != nil {
		return err
	}

	var validationError = ValidateStruct(params)
	if validationError != nil {
		return fiber.ErrBadRequest
	}

	return nil
}

func GetQuery(ct *fiber.Ctx, query interface{}) error {
	var err = ct.QueryParser(query)
	if err != nil {
		return err
	}
	var validationError = ValidateStruct(query)
	if validationError != nil {
		return fiber.ErrBadRequest
	}
	return nil
}

func Contains[T primitives](array []T, element T) bool {
	for _, e := range array {
		if e == element {
			return true
		}
	}

	return false
}

type ElementHandler[T interface{}, O interface{}] func(element T) O

func Map[T interface{}, O interface{}](array []T, handler ElementHandler[T, O]) []O {
	var response = []O{}
	for _, e := range array {
		response = append(response, handler(e))
	}

	return response
}
