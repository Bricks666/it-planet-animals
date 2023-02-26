package animals

import (
	"animals/shared"
)

type AnimalParamsDto struct {
	Id uint64 `param:"id" validate:"required,gt=0"`
}

type AnimalDto struct {
	ID                 uint64   `json:"id"`
	AnimalTypes        []uint64 `json:"animalTypes"`
	Weight             float32  `json:"weight"`
	Length             float32  `json:"length"`
	Height             float32  `json:"height"`
	Gender             string   `json:"gender"`
	LifeStatus         string   `json:"lifeStatus"`
	ChippingDateTime   string   `json:"chippingDateTime"`
	ChipperId          uint32   `json:"chipperId"`
	ChippingLocationId uint64   `json:"chippingLocationId"`
	VisitedLocations   []uint64 `json:"visitedLocations"`
	DeathDateTime      *string  `json:"deathDateTime"`
}

type AnimalsSearchQueryDto struct {
	shared.PaginationDto
	shared.TimeSearchQueryDto
	ChipperId          uint32 `query:"chipperid" validate:"omitempty,number,gt=0"`
	ChippingLocationId uint64 `query:"chippinglocationid" validate:"omitempty,number,gt=0"`
	LifeStatus         string `query:"lifestatus" validate:"omitempty,oneof=ALIVE DEAD"`
	Gender             string `query:"gender" validate:"omitempty,oneof=MALE FEMALE OTHER"`
}

type CreateAnimalDto struct {
	AnimalTypes        []uint64 `json:"animaltypes" validate:"required,gt=0,dive,number,gt=0"`
	Weight             float32  `json:"weight" validate:"required,gt=0"`
	Length             float32  `json:"length" validate:"required,gt=0"`
	Height             float32  `json:"height" validate:"required,gt=0"`
	Gender             string   `json:"gender" validate:"required,oneof=MALE FEMALE OTHER"`
	ChipperId          uint32   `json:"chipperid" validate:"required,number,gt=0"`
	ChippingLocationId uint64   `json:"chippinglocationid" validate:"required,number,gt=0"`
}

type UpdateAnimalDto struct {
	Weight             float32 `json:"weight" validate:"required,gt=0"`
	Length             float32 `json:"length" validate:"required,gt=0"`
	Height             float32 `json:"height" validate:"required,gt=0"`
	Gender             string  `json:"gender" validate:"required,oneof=MALE FEMALE OTHER"`
	ChipperId          uint32  `json:"chipperid" validate:"required,number,gt=0"`
	LifeStatus         string  `json:"lifestatus" validate:"omitempty,oneof=ALIVE DEAD"`
	ChippingLocationId uint64  `json:"chippinglocationid" validate:"required,number,gt=0"`
}

type AnimalTypeParamsDto struct {
	Id     uint64 `param:"id" validate:"required,gt=0"`
	TypeId uint64 `param:"typeid" validate:"required,gt=0"`
}

type ReplaceAnimalTypeDto struct {
	OldTypeId uint64 `json:"oldtypeid" validate:"required,gt=0"`
	NewTypeId uint64 `json:"newtypeid" validate:"required,gt=0"`
}

var NewAnimalParamsDto = shared.BaseNew[AnimalParamsDto]()
var NewAnimalDto = shared.BaseNew[AnimalDto]()
var NewAnimalsSearchQueryDto = func() *AnimalsSearchQueryDto {
	return &AnimalsSearchQueryDto{
		PaginationDto:      *shared.NewPaginationDto(),
		TimeSearchQueryDto: *shared.NewTimeSearchQueryDto(),
	}
}
var NewCreateAnimalDto = shared.BaseNew[CreateAnimalDto]()
var NewUpdateAnimalDto = shared.BaseNew[UpdateAnimalDto]()
var NewAnimalTypeParamsDto = shared.BaseNew[AnimalTypeParamsDto]()
var NewReplaceAnimalTypeDto = shared.BaseNew[ReplaceAnimalTypeDto]()
