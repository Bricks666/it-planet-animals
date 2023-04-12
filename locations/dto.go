package locations

import "animals/shared"

type LocationDto struct {
	Id        uint64  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type LocationBodyDto struct {
	Latitude  float64 `json:"latitude" validate:"any-number,latitude"`
	Longitude float64 `json:"longitude" validate:"any-number,longitude"`
}

type LocationParams struct {
	Id uint64 `params:"id" validate:"any-number,min=1"`
}

var NewLocationDto = shared.BaseNew[LocationDto]()
var NewLocationBodyDto = shared.BaseNew[LocationBodyDto]()
var NewLocationParams = shared.BaseNew[LocationParams]()
