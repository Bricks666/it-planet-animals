package animals

import (
	"animals/ent"
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

func (this *AnimalsService) Create() {}

func (this *AnimalsService) Update() {}

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
		Lifestatus:         string(animal.Lifestatus),
		ChippingDateTime:   animal.ChippingDateTime,
		ChipperId:          animal.ChipperId,
		ChippingLocationId: animal.ChippingLocationId,
		DeathDateTime:      animal.DeathDateTime,
		VisitedLocations:   locationIds,
	}
}
