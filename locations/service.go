package locations

import (
	"animals/ent"
)

var Service LocationsService

type LocationsService struct {
	locationRepository *LocationRepository
}

func init() {
	Service = LocationsService{
		locationRepository: &Repository,
	}
}

func (this *LocationsService) GetOne(id uint64) (*LocationDto, error) {
	location, err := this.locationRepository.GetOne(id)

	if err != nil {
		return nil, err
	}

	return prepareLocation(location), nil
}

func (this *LocationsService) Create(dto *CreateLocationDto) (*LocationDto, error) {
	location, err := this.locationRepository.Create(dto)

	if err != nil {
		return nil, err
	}

	return prepareLocation(location), nil
}

func (this *LocationsService) Update(id uint64, dto *UpdateLocationDto) (*LocationDto, error) {
	location, err := this.locationRepository.Update(id, dto)

	if err != nil {
		return nil, err
	}

	return prepareLocation(location), nil
}

func (this *LocationsService) Remove(id uint64) error {
	return this.locationRepository.Remove(id)
}

func prepareLocation(location *ent.Location) *LocationDto {
	return &LocationDto{
		Id:        location.ID,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
	}
}
