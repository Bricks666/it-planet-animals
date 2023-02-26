package visitedlocation

import (
	"animals/shared"
)

type VisitedLocationDto struct {
	Id                           uint64 `json:"id"`
	DateTimeOfVisitLocationPoint string `json:"dateTimeOfVisitLocationPoint"`
	LocationId                   uint64 `json:"locationPointId"`
}

type VisitedLocationParamsDto struct {
	AnimalId uint64 `params:"animalId" validate:"required,number,gt=0"`
}

type VisitedLocationSearchQueryDto struct {
	shared.TimeSearchQueryDto
	shared.PaginationDto
}

type VisitedLocationMutationParamsDto struct {
	VisitedLocationParamsDto
	LocationId uint64 `param:"locationId" validate:"required,number,gt=0"`
}

type UpdateVisitedLocationDto struct {
	VisitedAnimalLocationId uint64 `json:"visitedLocationPointId" validate:"required,number,gt=0"`
	NewLocationId           uint64 `json:"locationPointId" validate:"required,number,gt=0"`
}

var NewVisitedLocationDto = shared.BaseNew[VisitedLocationDto]()
var NewVisitedLocationParamsDto = shared.BaseNew[VisitedLocationParamsDto]()
var NewVisitedLocationSearchQueryDto = func() *VisitedLocationSearchQueryDto {
	return &VisitedLocationSearchQueryDto{
		PaginationDto:      *shared.NewPaginationDto(),
		TimeSearchQueryDto: *shared.NewTimeSearchQueryDto(),
	}

}
var NewVisitedLocationMutationParamsDto = shared.BaseNew[VisitedLocationMutationParamsDto]()
var NewUpdateVisitedLocationDto = shared.BaseNew[UpdateVisitedLocationDto]()
