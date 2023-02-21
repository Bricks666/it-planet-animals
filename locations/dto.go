package locations

type LocationDto struct {
	Id        uint64  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

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
