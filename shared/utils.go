package shared

type IsNotEmptyType interface {
	int | float64 | uint | bool | string
}

func IsEmpty[T IsNotEmptyType](value T) bool {
	var empty T
	return value == empty
}
