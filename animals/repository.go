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
		query = query.Where(animal.ChipperId(dto.ChipperId))
	}
	if dto.Gender != "" {
		query = query.Where(animal.GenderEQ(animal.Gender(dto.Gender)))
	}
	if dto.LifeStatus != "" {
		query = query.Where(animal.LifestatusEQ(animal.Lifestatus(dto.LifeStatus)))
	}
	if dto.ChippingLocationId != 0 {
		query = query.Where(animal.ChippingLocationId(dto.ChippingLocationId))
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
		WithAnimalTypeAnimal(func(query *ent.AnimalTypeQuery) {
			query.OnlyID(this.db.Context)
		}).
		Offset(int(dto.From)).
		Limit(int(dto.Size)).
		All(this.db.Context)
}

func (this *AnimalsRepository) GetOne(id uint64) (*ent.Animal, error) {
	return this.db.Client.Animal.
		Query().Where(animal.ID(id)).Only(this.db.Context)
}

func (this *AnimalsRepository) Create() (*ent.Animal, error) {
	return nil, nil
}

func (this *AnimalsRepository) Update() (*ent.Animal, error) {
	return nil, nil
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
