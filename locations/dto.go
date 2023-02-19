package locations

type CreateLocationDto struct {
	Latitude  float64 `json:"latitude,omitempty" validate:"required,gt=-90,lt=90"`
	Longitude float64 `json:"longitude,omitempty" validate:"required,gt=-180,lt=180"`
}

type UpdateLocationDto struct {
	Latitude  float64 `json:"latitude,omitempty" validate:"required,gt=-90,lt=90"`
	Longitude float64 `json:"longitude,omitempty" validate:"required,gt=-180,lt=180"`
}

type LocationParams struct {
	Id uint64 `params:"id" validate:"required,number,min=1"`
}
