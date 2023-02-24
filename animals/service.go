package animals

import (
	"animals/ent"
	Animal "animals/ent/animal"
	"animals/shared"

	"golang.org/x/exp/slices"
)

var Service AnimalsService

type AnimalsService struct {
	animalsRepository *AnimalsRepository
}

func init() {
	Service = AnimalsService{
		animalsRepository: &Repository,
	}
}

func (this *AnimalsService) GetAll(dto *AnimalsSearchQueryDto) ([]*AnimalDto, error) {
	animals, err := this.animalsRepository.GetAll(dto)

	if err != nil {
		return nil, err
	}

	var animalsResponse = []*AnimalDto{}

	for _, animal := range animals {
		animalsResponse = append(animalsResponse, prepareAnimal(animal))
	}

	return animalsResponse, nil

}

func (this *AnimalsService) GetOne(id uint64) (*AnimalDto, error) {
	animal, err := this.animalsRepository.GetOne(id)

	if err != nil {
		return nil, err
	}

	return prepareAnimal(animal), nil
}

func (this *AnimalsService) Create(dto *CreateAnimalDto) (*AnimalDto, error) {
	animal, err := this.animalsRepository.Create(dto)

	if err != nil {
		return nil, err
	}

	return this.GetOne(animal.ID)
}

func (this *AnimalsService) Update(id uint64, dto *UpdateAnimalDto) (*AnimalDto, error) {
	animal, err := this.animalsRepository.GetOne(id)

	if err != nil {
		return nil, &ent.NotFoundError{}
	}

	if animal.LifeStatus == Animal.LifeStatusDEAD &&
		dto.LifeStatus == Animal.LifeStatusALIVE.String() {
		return nil, &ent.ConstraintError{}
	}

	count := len(animal.Edges.VisitedLocations)
	if count > 0 && animal.Edges.VisitedLocations[count-1].ID == dto.ChippingLocationId {
		return nil, &ent.ConstraintError{}
	}

	animal, err = this.animalsRepository.Update(id, dto)

	if err != nil {
		return nil, &ent.NotFoundError{}
	}

	return this.GetOne(animal.ID)
}

func (this *AnimalsService) Remove(id uint64) error {
	var animal, err = this.GetOne(id)

	if err != nil {
		return err
	}

	if len(animal.VisitedLocations) > 0 &&
		animal.ChippingLocationId == animal.VisitedLocations[0] {
		return &ent.ConstraintError{}
	}

	return this.animalsRepository.Remove(id)
}

func (this *AnimalsService) AddType(id uint64, typeId uint64) (*AnimalDto, error) {
	var a, err = this.GetOne(id)
	if err != nil {
		return nil, err
	}

	var isContain = slices.Contains(a.AnimalTypes, typeId)
	if isContain {
		return nil, &ent.ConstraintError{}
	}

	var animal *ent.Animal
	animal, err = this.animalsRepository.AddType(id, typeId)
	if err != nil {
		return nil, &ent.NotFoundError{}
	}

	return this.GetOne(animal.ID)
}

func (this *AnimalsService) ReplaceType(id uint64, dto *ReplaceAnimalTypeDto) (*AnimalDto, error) {
	var a, err = this.GetOne(id)
	if err != nil {
		return nil, err
	}

	var isContainNew = slices.Contains(a.AnimalTypes, dto.NewTypeId)
	if isContainNew {
		return nil, &ent.ConstraintError{}
	}

	var isContainOld = slices.Contains(a.AnimalTypes, dto.OldTypeId)
	if !isContainOld {
		return nil, &ent.NotFoundError{}
	}

	_, err = this.animalsRepository.RemoveType(id, dto.OldTypeId)
	if err != nil {
		return nil, &ent.NotFoundError{}
	}

	var animal *ent.Animal
	animal, err = this.animalsRepository.AddType(id, dto.NewTypeId)
	if err != nil {
		return nil, &ent.NotFoundError{}
	}

	return this.GetOne(animal.ID)
}

func (this *AnimalsService) RemoveType(id uint64, typeId uint64) (*AnimalDto, error) {
	var a, err = this.GetOne(id)
	if err != nil {
		return nil, err
	}

	var isContain = slices.Contains(a.AnimalTypes, typeId)
	if !isContain {
		return nil, &ent.NotFoundError{}
	}
	if isContain && len(a.AnimalTypes) == 1 {
		return nil, &ent.ConstraintError{}
	}

	var animal *ent.Animal
	animal, err = this.animalsRepository.RemoveType(id, typeId)
	if err != nil {
		return nil, &ent.NotFoundError{}
	}

	return this.GetOne(animal.ID)
}

func prepareAnimal(animal *ent.Animal) *AnimalDto {
	var typeIds = []uint64{}
	var locationIds = []uint64{}

	for _, t := range animal.Edges.AnimalTypeAnimal {
		typeIds = append(typeIds, t.ID)
	}

	for _, l := range animal.Edges.Animals {
		locationIds = append(locationIds, l.ID)
	}

	var preparedAnimal = &AnimalDto{
		ID:                 animal.ID,
		AnimalTypes:        typeIds,
		Weight:             animal.Weight,
		Length:             animal.Length,
		Height:             animal.Height,
		Gender:             string(animal.Gender),
		LifeStatus:         string(animal.LifeStatus),
		ChippingDateTime:   animal.ChippingDateTime.Format(shared.ISO8601_PATTER),
		ChipperId:          animal.ChipperID,
		ChippingLocationId: animal.ChippingLocationID,
		DeathDateTime:      nil,
		VisitedLocations:   locationIds,
	}

	if animal.DeathDateTime != nil {
		var formattedDate = animal.DeathDateTime.Format(shared.ISO8601_PATTER)
		preparedAnimal.DeathDateTime = &formattedDate
	}

	return preparedAnimal
}
