package animaltypes

import (
	"animals/ent"
	"animals/ent/animaltag"
	"animals/shared"
)

type AnimalTypeRepository struct {
	db *shared.DB
}

var Repository AnimalTypeRepository

func NewAnimalTypeRepository(db *shared.DB) *AnimalTypeRepository {
	return &AnimalTypeRepository{
		db: db,
	}
}

func init() {
	Repository = *NewAnimalTypeRepository(&shared.Database)
}

func (this *AnimalTypeRepository) GetOne(id uint64) (*ent.AnimalTag, error) {
	return this.db.Client.AnimalTag.
		Query().
		Where(animaltag.ID(id)).
		Only(this.db.Context)
}

func (this *AnimalTypeRepository) Create(dto *CreateAnimaTypeDto) (*ent.AnimalTag, error) {
	return this.db.Client.AnimalTag.
		Create().
		SetType(dto.AnimalType).
		Save(this.db.Context)
}

func (this *AnimalTypeRepository) Update(id uint64, dto *UpdateAnimaTypeDto) (*ent.AnimalTag, error) {
	return this.db.Client.AnimalTag.
		UpdateOneID(id).
		SetType(dto.Type).
		Save(this.db.Context)
}

func (this *AnimalTypeRepository) Remove(id uint64) error {
	return this.db.Client.AnimalTag.
		DeleteOneID(id).
		Exec(this.db.Context)
}
