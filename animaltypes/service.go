package animaltypes

import "animals/ent"

var Service AnimalTypesService

type AnimalTypesService struct {
	animalTypeRepository *AnimalTypeRepository
}

func init() {
	Service = AnimalTypesService{
		animalTypeRepository: &Repository,
	}
}

func (this *AnimalTypesService) GetOne(id uint64) (*ent.AnimalType, error) {
	return this.animalTypeRepository.GetOne(id)
}

func (this *AnimalTypesService) Create(dto *CreateAnimaTypeDto) (*ent.AnimalType, error) {
	return this.animalTypeRepository.Create(dto)
}

func (this *AnimalTypesService) Update(id uint64, dto *UpdateAnimaTypeDto) (*ent.AnimalType, error) {
	return this.animalTypeRepository.Update(id, dto)
}

func (this *AnimalTypesService) Remove(id uint64) error {
	return this.animalTypeRepository.Remove(id)
}
