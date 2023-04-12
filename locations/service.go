package locations

import (
	"animals/ent"
)

type LocationsService struct {
	locationRepository *LocationRepository
}

var Service LocationsService

func NewLocationsService(locationRepository *LocationRepository) *LocationsService {
	return &LocationsService{
		locationRepository: locationRepository,
	}
}

func init() {
	Service = *NewLocationsService(
		&Repository,
	)
}

func (this *LocationsService) GetOne(id uint64) (*LocationDto, error) {
	location, err := this.locationRepository.GetOne(id)

	if err != nil {
		return nil, err
	}

	return prepareLocation(location), nil
}

func (this *LocationsService) Create(dto *LocationBodyDto) (*LocationDto, error) {
	location, err := this.locationRepository.Create(dto)

	if err != nil {
		return nil, err
	}

	return prepareLocation(location), nil
}

func (this *LocationsService) Update(id uint64, dto *LocationBodyDto) (*LocationDto, error) {
	location, err := this.locationRepository.Update(id, dto)

	if err != nil {
		return nil, err
	}

	return prepareLocation(location), nil
}

func (this *LocationsService) Remove(id uint64) error {
	return this.locationRepository.Remove(id)
}

func (this *LocationsService) GetOrCreate(params *LocationBodyDto) *LocationDto {
	var location, err = this.locationRepository.GetOrCreate(params)

	if err != nil {
		return nil
	}

	return prepareLocation(location)
}

func prepareLocation(location *ent.Location) *LocationDto {
	return &LocationDto{
		Id:        location.ID,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
	}
}
