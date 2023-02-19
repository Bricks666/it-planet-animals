package locations

import (
	"animals/ent"
	"errors"
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

func (ls *LocationsService) GetOne(id uint64) (*ent.Location, error) {
	return ls.locationRepository.GetOne(id)
}

func (ls *LocationsService) Create(dto *CreateLocationDto) (*ent.Location, error) {
	return ls.locationRepository.Create(dto)
}

func (ls *LocationsService) Update(id uint64, dto *UpdateLocationDto) (*ent.Location, error) {
	location, _ := ls.locationRepository.GetOne(id)

	if location != nil {
		return nil, errors.New("404")
	}

	return ls.locationRepository.Update(id, dto)
}

func (ls *LocationsService) Remove(id uint64) error {
	return ls.locationRepository.Remove(id)
}
