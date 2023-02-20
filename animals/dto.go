package animals

import (
	"time"
)

type AnimalParamsDto struct {
	Id uint64 `param:"id" validate:"required,gt=0"`
}

type AnimalResponseDto struct {
	ID                 uint64     `json:"id"`
	AnimalTypes        []uint64   `json:"animalTypes"`
	Weight             float32    `json:"weight"`
	Length             float32    `json:"length"`
	Height             float32    `json:"height"`
	Gender             string     `json:"gender"`
	Lifestatus         string     `json:"lifeStatus"`
	ChippingDateTime   time.Time  `json:"chippingDateTime"`
	ChipperId          uint32     `json:"chipperId"`
	ChippingLocationId uint64     `json:"chippingLocationId"`
	VisitedLocations   []uint64   `json:"visitedLocations"`
	DeathDateTime      *time.Time `json:"deathDateTime"`
}

type AnimalsSearchQueryDto struct {
	AfterChippingDate  time.Time `query:"startdatetime" validate:"omitempty,iso-8601"`
	BeforeChippingDate time.Time `query:"enddatetime" validate:"omitempty,iso-8601"`
	ChipperId          uint32    `query:"chipperid" validate:"omitempty,number,gt=0"`
	ChippingLocationId uint64    `query:"chippinglocationid" validate:"omitempty,number,gt=0"`
	LifeStatus         string    `query:"lifestatus" validate:"omitempty,oneof=alive dead"`
	Gender             string    `query:"gender" validate:"omitempty,oneof=male female other"`
	From               uint32    `query:"from" validate:"number,gte=0"`
	Size               uint32    `query:"size" validate:"number,gt=0"`
}

type CreateAnimalDto struct {
	AnimalTypes        uint64  `json:"animaltypes" validate:"required,gt=0,dive,number,gt=0"`
	Weight             float32 `json:"weight" validate:"required,gt=0"`
	Length             float32 `json:"length" validate:"required,gt=0"`
	Height             float32 `json:"height" validate:"required,gt=0"`
	Gender             string  `json:"gender" validate:"required,oneof=male female other"`
	ChipperId          uint32  `json:"chipperid" validate:"required,number,gt=0"`
	ChippingLocationId uint64  `json:"chippinglocationid" validate:"required,number,gt=0"`
}

type UpdateAnimalDto struct {
	Weight             float32 `json:"weight" validate:"required,gt=0"`
	Length             float32 `json:"length" validate:"required,gt=0"`
	Height             float32 `json:"height" validate:"required,gt=0"`
	Gender             string  `json:"gender" validate:"required,oneof=male female other"`
	ChipperId          uint32  `json:"chipperid" validate:"required,number,gt=0"`
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
