package shared

type Vector struct {
	X float64
	Y float64
}

var NewVector = func(x float64, y float64) *Vector {
	return &Vector{
		X: x,
		Y: y,
	}
}
