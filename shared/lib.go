package shared

type mappable interface {
	uint | uint16 | uint32 | uint64 | uint8 |
		int | int16 | int32 | int64 | int8 |
		float32 | float64 | string | bool
}

func HasOnlyUnique[T mappable](array []T) bool {
	var occurred = map[T]bool{}

	for _, e := range array {
		if occurred[e] {
			return false
		}

		occurred[e] = true
	}

	return true
}
