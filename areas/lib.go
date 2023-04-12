package areas

import (
	"animals/ent"
	"animals/shared"
	"math"
)

func prepareAreaDto(area *ent.Area) *AreaDto {
	var dto = NewAreaDto()

	dto.Id = area.ID
	dto.Name = area.Name

	dto.AreaPoints = shared.Map(area.Edges.Points, func(element *ent.Location) *AreaPointDto {
		var areaPointDto = NewAreaPointDto()

		areaPointDto.Latitude = element.Latitude
		areaPointDto.Longitude = element.Longitude

		return areaPointDto
	})

	return dto
}

type EquationFunction func(x float64) float64

func stubLineEquation(x float64) float64 {
	return math.Inf(1)
}

func getLineEquation(point1 *shared.Vector, point2 *shared.Vector) EquationFunction {
	var xDifferent = point2.X - point1.X
	var yDifferent = point2.Y - point1.Y

	if xDifferent == 0 || yDifferent == 0 {
		return stubLineEquation
	}

	return func(x float64) float64 {
		return (x-point1.X)/xDifferent*yDifferent + point1.Y
	}
}
