package animalslocations

import (
	"animals/shared"
)

type AnimalsLocationsDto struct {
	Id                           uint64 `json:"id"`
	DateTimeOfVisitLocationPoint string `json:"dateTimeOfVisitLocationPoint"`
	LocationId                   uint64 `json:"locationPointId"`
}

type AnimalsLocationParamsDto struct {
	AnimalId uint64 `params:"animalId" validate:"required,number,gt=0"`
}

type AnimalsLocationSearchQueryDto struct {
	shared.PaginationDto
	shared.TimeSearchQueryDto
}

type AnimalLocationParamsDto struct {
	AnimalsLocationParamsDto
	LocationId uint64 `param:"locationId" validate:"required,number,gt=0"`
}

type UpdateAnimalsLocationDto struct {
	VisitedAnimalLocationId uint64 `json:"visitedLocationPointId" validate:"required,number,gt=0"`
	NewLocationId           uint64 `json:"locationPointId" validate:"required,number,gt=0"`
}
