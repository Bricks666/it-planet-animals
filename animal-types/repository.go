package animaltypes

import (
	"animals/ent"
	"animals/ent/animaltype"
	"animals/shared"
)

var Repository AnimalTypeRepository

type AnimalTypeRepository struct {
	db *shared.DB
}

func init() {
	Repository = AnimalTypeRepository{
		db: &shared.Database,
	}
}

func (this *AnimalTypeRepository) GetOne(id uint64) (*ent.AnimalType, error) {
	return this.db.Client.AnimalType.
		Query().
		Where(animaltype.ID(id)).
		Only(this.db.Context)
}

func (this *AnimalTypeRepository) Create(dto *CreateAnimaTypeDto) (*ent.AnimalType, error) {
	return this.db.Client.AnimalType.
		Create().
		SetType(dto.AnimalType).
		Save(this.db.Context)
}

func (this *AnimalTypeRepository) Update(id uint64, dto *UpdateAnimaTypeDto) (*ent.AnimalType, error) {
	return this.db.Client.AnimalType.
		UpdateOneID(id).
		SetType(dto.AnimalType).
		Save(this.db.Context)
}

func (this *AnimalTypeRepository) Remove(id uint64) error {
	return this.db.Client.AnimalType.
		DeleteOneID(id).
		Exec(this.db.Context)
}
