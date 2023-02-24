package animaltypes

type AnimalTypeDto struct {
	Id   uint64 `json:"id"`
	Type string `json:"type"`
}

type CreateAnimaTypeDto struct {
	AnimalType string `json:"type" validate:"required,notblank,min=1"`
}

type UpdateAnimaTypeDto struct {
	Type string `json:"type" validate:"required,notblank,min=1"`
}

type AnimalTypeParamsDto struct {
	Id uint64 `param:"id" validate:"number,gt=0"`
}
