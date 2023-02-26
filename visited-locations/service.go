package visitedlocation

import (
	"animals/animals"
	"animals/ent"
	Animal "animals/ent/animal"
	"animals/shared"

	"golang.org/x/exp/slices"
)

type VisitedLocationsService struct {
	visitedLocationsRepository *VisitedLocationsRepository
	animalsRepository          *animals.AnimalsRepository
}

var Service VisitedLocationsService

func NewVisitedLocationsService(
	visitedLocationsRepository *VisitedLocationsRepository,
	animalsRepository *animals.AnimalsRepository) *VisitedLocationsService {
	return &VisitedLocationsService{
		visitedLocationsRepository: visitedLocationsRepository,
		animalsRepository:          animalsRepository,
	}
}

func init() {
	Service = *NewVisitedLocationsService(
		&Repository,
		&animals.Repository,
	)
}

func (this *VisitedLocationsService) GetAll(animalId uint64, dto *VisitedLocationSearchQueryDto) ([]*VisitedLocationDto, error) {

	locations, err := this.visitedLocationsRepository.GetAll(animalId, dto)

	if err != nil {
		return nil, err
	}

	preparedLocations := []*VisitedLocationDto{}

	for _, location := range locations {
		preparedLocations = append(preparedLocations, prepareAnimalsLocation(location))
	}

	return preparedLocations, nil
}

func (this *VisitedLocationsService) Create(animalId uint64, locationId uint64) (*VisitedLocationDto, error) {
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
		visitedLocations[visitedLocationsCount-1].LocationID == locationId {
		return nil, &ent.ConstraintError{}
	}

	animalLocation, err := this.visitedLocationsRepository.Create(animalId, locationId)

	if err != nil {
		return nil, &ent.NotFoundError{}
	}

	return prepareAnimalsLocation(animalLocation), nil
}

func (this *VisitedLocationsService) Update(animalId uint64, dto *UpdateVisitedLocationDto) (*VisitedLocationDto, error) {
	animal, err := this.animalsRepository.GetOne(animalId)

	if err != nil {
		return nil, &ent.NotFoundError{}
	}

	visitedLocations := animal.Edges.VisitedLocations
	visitedLocationsCount := len(visitedLocations)

	if visitedLocationsCount == 0 {
		return nil, &ent.NotFoundError{}
	}

	if visitedLocations[0].ID == dto.VisitedAnimalLocationId &&
		dto.NewLocationId == animal.ChippingLocationID {
		return nil, &ent.ConstraintError{}
	}

	var pointIndex = slices.IndexFunc(visitedLocations, func(element *ent.VisitedLocation) bool {
		return element.ID == dto.VisitedAnimalLocationId
	})

	if pointIndex == -1 {
		return nil, &ent.NotFoundError{}
	}

	var element = visitedLocations[pointIndex]

	if element.LocationID == dto.NewLocationId {
		return nil, &ent.ConstraintError{}
	}

	var prev, next *ent.VisitedLocation

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

	var point *ent.VisitedLocation
	point, err = this.visitedLocationsRepository.Update(animalId, dto)

	if err != nil {
		return nil, &ent.NotFoundError{}
	}

	return prepareAnimalsLocation(point), nil
}

func (this *VisitedLocationsService) Remove(animalId uint64, visitedId uint64) error {
	var err = this.visitedLocationsRepository.Remove(animalId, visitedId)

	if err != nil {
		return err
	}

	var animal *ent.Animal
	animal, err = this.animalsRepository.GetOne(animalId)

	if err != nil {
		return err
	}

	visitedLocations := animal.Edges.VisitedLocations

	if len(visitedLocations) == 0 {
		return nil
	}

	if visitedLocations[0].LocationID == animal.ChippingLocationID {
		return this.Remove(animalId, visitedLocations[0].ID)
	}

	return nil
}

func prepareAnimalsLocation(animalLocation *ent.VisitedLocation) *VisitedLocationDto {
	return &VisitedLocationDto{
		Id:                           animalLocation.ID,
		LocationId:                   animalLocation.LocationID,
		DateTimeOfVisitLocationPoint: animalLocation.DateTimeOfVisitLocationPoint.Format(shared.ISO8601_PATTER),
	}
}
