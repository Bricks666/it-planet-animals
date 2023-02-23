package animals

import (
	"animals/ent"
	Animal "animals/ent/animal"
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

	return prepareAnimal(animal), nil
}

func (this *AnimalsService) Update(id uint64, dto *UpdateAnimalDto) (*AnimalDto, error) {
	animal, err := this.animalsRepository.GetOne(id)

	if err != nil {
		return nil, &ent.NotFoundError{}
	}

	if animal.LifeStatus == Animal.LifeStatusDEAD &&
		dto.LifeStatus == Animal.LifeStatusALIVE.String() {
		return nil, ent.ConstraintError{}
	}

	count := len(animal.Edges.VisitedLocations)
	if count > 0 && animal.Edges.VisitedLocations[count-1].ID == dto.ChippingLocationId {
		return nil, ent.ConstraintError{}
	}

	animal, err = this.animalsRepository.Update(id, dto)

	if err != nil {
		return nil, &ent.NotFoundError{}
	}

	return prepareAnimal(animal), nil
}

func (this *AnimalsService) Remove() {}

func (this *AnimalsService) AddType() (*ent.Animal, error) {
	return nil, nil
}

func (this *AnimalsService) ReplaceType() (*ent.Animal, error) {
	return nil, nil
}

func (this *AnimalsService) RemoveType() (*ent.Animal, error) {
	return nil, nil
}

func prepareAnimal(animal *ent.Animal) *AnimalDto {
	var typeIds = []uint64{}
	var locationIds = []uint64{}

	for _, t := range animal.Edges.AnimalTypeAnimal {
		typeIds = append(typeIds, t.ID)
	}

	for _, l := range animal.Edges.VisitedLocations {
		locationIds = append(locationIds, l.ID)
	}

	return &AnimalDto{
		ID:                 animal.ID,
		AnimalTypes:        typeIds,
		Weight:             animal.Weight,
		Length:             animal.Length,
		Height:             animal.Height,
		Gender:             string(animal.Gender),
		Lifestatus:         string(animal.LifeStatus),
		ChippingDateTime:   animal.ChippingDateTime,
		ChipperId:          animal.ChipperID,
		ChippingLocationId: animal.ChippingLocationID,
		DeathDateTime:      animal.DeathDateTime,
		VisitedLocations:   locationIds,
	}
}
