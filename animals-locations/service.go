package animalslocations

import (
	"animals/animals"
	"animals/ent"
)

type AnimalsLocationsService struct {
	animalsLocationRepository *AnimalsLocationsRepository
	animalsRepository         *animals.AnimalsRepository
}

var Service AnimalsLocationsService

func init() {
	Service = AnimalsLocationsService{
		animalsLocationRepository: &Repository,
		animalsRepository:         &animals.Repository,
	}
}

func (this *AnimalsLocationsService) GetAll(animalId uint64, dto *AnimalsLocationSearchQueryDto) ([]*AnimalsLocationsDto, error) {

	locations, err := this.animalsLocationRepository.GetAll(animalId, dto)

	if err != nil {
		return nil, err
	}

	preparedLocations := []*AnimalsLocationsDto{}

	for _, location := range locations {
		preparedLocations = append(preparedLocations, prepareAnimalsLocation(location))
	}

	return preparedLocations, nil
}

func (this *AnimalsLocationsService) Create() (*AnimalsLocationsDto, error) {
	return nil, nil
}

func (this *AnimalsLocationsService) Update() (*AnimalsLocationsDto, error) {
	return nil, nil
}

func (this *AnimalsLocationsService) Remove() error {
	return nil
}

func prepareAnimalsLocation(animalLocation *ent.AnimalsLocations) *AnimalsLocationsDto {
	return &AnimalsLocationsDto{
		Id:                           animalLocation.ID,
		LocationId:                   animalLocation.LocationId,
		DateTimeOfVisitLocationPoint: animalLocation.DateTimeOfVisitLocationPoint,
	}
}
