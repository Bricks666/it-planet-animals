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

func (this *LocationsService) GetOne(id uint64) (*ent.Location, error) {
	return this.locationRepository.GetOne(id)
}

func (this *LocationsService) Create(dto *CreateLocationDto) (*ent.Location, error) {
	return this.locationRepository.Create(dto)
}

func (this *LocationsService) Update(id uint64, dto *UpdateLocationDto) (*ent.Location, error) {
	return this.locationRepository.Update(id, dto)
}

func (this *LocationsService) Remove(id uint64) error {
	return this.locationRepository.Remove(id)
}
