package animaltypes

type CreateAnimaTypeDto struct {
	AnimalType string `json:"type" validate:"required,notblank,min=1"`
}

type UpdateAnimaTypeDto struct {
	AnimalType string `json:"type" validate:"required,notblank,min=1"`
}

type AnimalTypeParamsDto struct {
	Id uint64 `param:"id" validate:"number,gt=0"`
}
