package areas

import "animals/shared"

type AreaPointDto struct {
	Latitude  float64 `json:"latitude" validate:"any-number,latitude"`
	Longitude float64 `json:"longitude" validate:"any-number,longitude"`
}

type AreaDto struct {
	Id         uint64          `json:"id"`
	Name       string          `json:"name"`
	AreaPoints []*AreaPointDto `json:"areaPoints"`
}

type AreaBodyDto struct {
	Name       string          `json:"name" validate:"required,notblank,min=1"`
	AreaPoints []*AreaPointDto `json:"areaPoints" validate:"required,min=3"`
}

type InnerAreaDto struct {
	Name         string
	AreaPointIds []uint64
}

type AreaParamsDto struct {
	Id uint64 `params:"id" validate:"any-number,min=1"`
}

var NewAreaPointDto = shared.BaseNew[AreaPointDto]()
var NewAreaDto = shared.BaseNew[AreaDto]()
var NewAreaBodyDto = shared.BaseNew[AreaBodyDto]()
var NewInnerAreaDto = shared.BaseNew[InnerAreaDto]()
var NewAreaParamsDto = shared.BaseNew[AreaParamsDto]()
