package animalslocations

import (
	"animals/animals"
	"animals/ent"
	Animal "animals/ent/animal"
	"animals/shared"

	"golang.org/x/exp/slices"
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

func (this *AnimalsLocationsService) Create(animalId uint64, locationId uint64) (*AnimalsLocationsDto, error) {
	animal, err := this.animalsRepository.GetOne(animalId)

	if err != nil {
		return nil, &ent.NotFoundError{}
	}

	if animal.LifeStatus == Animal.LifeStatusDEAD {
		return nil, &ent.ConstraintError{}
	}

	visitedLocations := animal.Edges.VisitedLocations
	visitedLocationsCount := len(visitedLocations)
	if visitedLocationsCount == 0 &&
		animal.ChippingLocationID == locationId {
		return nil, &ent.ConstraintError{}
	}
	if visitedLocationsCount > 0 &&
		visitedLocations[0].ID == locationId {
		return nil, &ent.ConstraintError{}
	}

	animalLocation, err := this.animalsLocationRepository.Create(animalId, locationId)

	if err != nil {
		return nil, &ent.NotFoundError{}
	}

	return prepareAnimalsLocation(animalLocation), nil
}

func (this *AnimalsLocationsService) Update(animalId uint64, dto *UpdateAnimalsLocationDto) (*AnimalsLocationsDto, error) {
	animal, err := this.animalsRepository.GetOne(animalId)

	if err != nil {
		return nil, &ent.NotFoundError{}
	}

	visitedLocations := animal.Edges.Animals
	visitedLocationsCount := len(visitedLocations)

	if visitedLocationsCount == 0 {
		return nil, &ent.NotFoundError{}
	}

	if visitedLocations[0].ID == dto.VisitedAnimalLocationId &&
		dto.NewLocationId == animal.ChippingLocationID {
		return nil, &ent.ConstraintError{}
	}

	var pointIndex = slices.IndexFunc(visitedLocations, func(element *ent.AnimalsLocations) bool {
		return element.ID == dto.VisitedAnimalLocationId
	})

	if pointIndex == -1 {
		return nil, &ent.NotFoundError{}
	}

	var element = visitedLocations[pointIndex]

	if element.LocationID == dto.NewLocationId {
		return nil, &ent.ConstraintError{}
	}

	var prev, next *ent.AnimalsLocations

	if pointIndex > 0 {
		prev = visitedLocations[pointIndex-1]
	}
	if pointIndex < visitedLocationsCount-1 {
		next = visitedLocations[pointIndex+1]
	}

	if (prev != nil && prev.LocationID == dto.NewLocationId) ||
		(next != nil && next.LocationID == dto.NewLocationId) {
		return nil, &ent.ConstraintError{}
	}

	var point *ent.AnimalsLocations
	point, err = this.animalsLocationRepository.Update(animalId, dto)

	if err != nil {
		return nil, &ent.NotFoundError{}
	}

	return prepareAnimalsLocation(point), nil
}

func (this *AnimalsLocationsService) Remove(animalId uint64, visitedId uint64) error {
	var err = this.animalsLocationRepository.Remove(animalId, visitedId)

	if err != nil {
		return err
	}

	var animal *ent.Animal
	animal, err = this.animalsRepository.GetOne(animalId)

	if err != nil {
		return err
	}

	visitedLocations := animal.Edges.Animals

	if len(visitedLocations) == 0 {
		return nil
	}

	if visitedLocations[0].LocationID == animal.ChippingLocationID {
		return this.Remove(animalId, visitedLocations[0].ID)
	}

	return nil
}

func prepareAnimalsLocation(animalLocation *ent.AnimalsLocations) *AnimalsLocationsDto {
	return &AnimalsLocationsDto{
		Id:                           animalLocation.ID,
		LocationId:                   animalLocation.LocationID,
		DateTimeOfVisitLocationPoint: animalLocation.DateTimeOfVisitLocationPoint.Format(shared.ISO8601_PATTER),
	}
}
