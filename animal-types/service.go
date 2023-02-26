package animaltypes

import "animals/ent"

type AnimalTypesService struct {
	animalTypeRepository *AnimalTypeRepository
}

var Service AnimalTypesService

func NewAnimalTypesService(animalTypeRepository *AnimalTypeRepository) *AnimalTypesService {
	return &AnimalTypesService{
		animalTypeRepository: animalTypeRepository,
	}
}

func init() {
	Service = *NewAnimalTypesService(&Repository)
}

func (this *AnimalTypesService) GetOne(id uint64) (*AnimalTypeDto, error) {
	animalType, err := this.animalTypeRepository.GetOne(id)

	if err != nil {
		return nil, err
	}

	return prepareAnimalType(animalType), nil
}

func (this *AnimalTypesService) Create(dto *CreateAnimaTypeDto) (*AnimalTypeDto, error) {
	animalType, err := this.animalTypeRepository.Create(dto)

	if err != nil {
		return nil, err
	}

	return prepareAnimalType(animalType), nil
}

func (this *AnimalTypesService) Update(id uint64, dto *UpdateAnimaTypeDto) (*AnimalTypeDto, error) {
	animalType, err := this.animalTypeRepository.Update(id, dto)

	if err != nil {
		return nil, err
	}

	return prepareAnimalType(animalType), nil
}

func (this *AnimalTypesService) Remove(id uint64) error {
	return this.animalTypeRepository.Remove(id)
}

func prepareAnimalType(animalType *ent.AnimalTag) *AnimalTypeDto {
	return &AnimalTypeDto{
		Id:   animalType.ID,
		Type: animalType.Type,
	}
}
