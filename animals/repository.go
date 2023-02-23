package animals

import (
	"animals/ent"
	"animals/ent/animal"
	"animals/ent/predicate"
	"animals/shared"
)

var Repository AnimalsRepository

type AnimalsRepository struct {
	db *shared.DB
}

func init() {
	Repository = AnimalsRepository{
		db: &shared.Database,
	}
}

func (this *AnimalsRepository) GetAll(dto *AnimalsSearchQueryDto) ([]*ent.Animal, error) {
	query := this.db.Client.Animal.Query()

	if dto.ChipperId != 0 {
		query = query.Where(animal.ChipperID(dto.ChipperId))
	}
	if dto.Gender != "" {
		query = query.Where(animal.GenderEQ(animal.Gender(dto.Gender)))
	}
	if dto.LifeStatus != "" {
		query = query.Where(animal.LifeStatusEQ(animal.LifeStatus(dto.LifeStatus)))
	}
	if dto.ChippingLocationId != 0 {
		query = query.Where(animal.ChippingLocationID(dto.ChippingLocationId))
	}
	var conditions []predicate.Animal

	if !dto.AfterDateTime.IsZero() {
		conditions = append(conditions, animal.ChippingDateTimeGTE(dto.AfterDateTime))
	}

	if !dto.BeforeDateTime.IsZero() {
		conditions = append(conditions, animal.ChippingDateTimeLTE(dto.BeforeDateTime))
	}

	if len(conditions) != 0 {
		query = query.Where(animal.Or(conditions...))
	}

	return query.
		WithAnimalTypeAnimal().
		WithVisitedLocations().
		Offset(int(dto.From)).
		Limit(int(dto.Size)).
		All(this.db.Context)
}

func (this *AnimalsRepository) GetOne(id uint64) (*ent.Animal, error) {
	return this.db.Client.Animal.
		Query().
		Where(animal.ID(id)).
		WithAnimalTypeAnimal().
		WithVisitedLocations().
		WithAnimals().
		Only(this.db.Context)
}

func (this *AnimalsRepository) Create(dto *CreateAnimalDto) (*ent.Animal, error) {
	animal, err := this.db.Client.Animal.
		Create().
		SetChipperAnimalID(dto.ChipperId).
		SetChippingLocationID(dto.ChippingLocationId).
		SetWeight(dto.Weight).
		SetLength(dto.Length).
		SetHeight(dto.Height).
		SetGender(animal.Gender(dto.Gender)).
		AddAnimalTypeAnimalIDs(dto.AnimalTypes...).
		AddVisitedLocationIDs([]uint64{}...).
		Save(this.db.Context)

	if err != nil {
		return nil, err
	}

	return this.GetOne(animal.ID)
}

func (this *AnimalsRepository) Update(id uint64, dto *UpdateAnimalDto) (*ent.Animal, error) {
	return this.db.Client.Animal.
		UpdateOneID(id).
		SetChipperAnimalID(dto.ChipperId).
		SetChippingLocationID(dto.ChippingLocationId).
		SetWeight(dto.Weight).
		SetLength(dto.Length).
		SetHeight(dto.Height).
		SetGender(animal.Gender(dto.Gender)).
		SetLifeStatus(animal.LifeStatus(dto.LifeStatus)).
		Save(this.db.Context)
}

func (this *AnimalsRepository) Remove() error {
	return nil
}

func (this *AnimalsRepository) AddType() (*ent.Animal, error) {
	return nil, nil
}

func (this *AnimalsRepository) RemoveType() (*ent.Animal, error) {
	return nil, nil
}
