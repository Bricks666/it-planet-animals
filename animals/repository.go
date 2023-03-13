package animals

import (
	"animals/ent"
	"animals/ent/animal"
	"animals/ent/predicate"
	"animals/shared"
	"time"
)

type AnimalsRepository struct {
	db *shared.DB
}

var Repository AnimalsRepository

func NewAnimalsRepository(db *shared.DB) *AnimalsRepository {
	return &AnimalsRepository{
		db: db,
	}
}

func init() {
	Repository = *NewAnimalsRepository(
		&shared.Database,
	)
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
		query = query.Where(animal.And(conditions...))
	}

	return query.
		WithTypes().
		WithVisitedLocations().
		Offset(int(dto.From)).
		Limit(int(dto.Size)).
		All(this.db.Context)
}

func (this *AnimalsRepository) GetOne(id uint64) (*ent.Animal, error) {
	return this.db.Client.Animal.
		Query().
		Where(animal.ID(id)).
		WithTypes().
		WithVisitedLocations().
		Only(this.db.Context)
}

func (this *AnimalsRepository) Create(dto *CreateAnimalDto) (*ent.Animal, error) {
	animal, err := this.db.Client.Animal.
		Create().
		SetChipperID(dto.ChipperId).
		SetChippingLocationID(dto.ChippingLocationId).
		SetWeight(dto.Weight).
		SetLength(dto.Length).
		SetHeight(dto.Height).
		SetGender(animal.Gender(dto.Gender)).
		AddTypeIDs(dto.AnimalTypes...).
		AddVisitedLocationIDs([]uint64{}...).
		Save(this.db.Context)

	if err != nil {
		return nil, err
	}

	return this.GetOne(animal.ID)
}

func (this *AnimalsRepository) Update(id uint64, dto *UpdateAnimalDto) (*ent.Animal, error) {
	var query = this.db.Client.Animal.UpdateOneID(id)

	if dto.LifeStatus == animal.LifeStatusDEAD.String() {
		query = query.SetDeathDateTime(time.Now())
	}

	return query.
		SetChipperID(dto.ChipperId).
		SetChippingLocationID(dto.ChippingLocationId).
		SetWeight(dto.Weight).
		SetLength(dto.Length).
		SetHeight(dto.Height).
		SetGender(animal.Gender(dto.Gender)).
		SetLifeStatus(animal.LifeStatus(dto.LifeStatus)).
		Save(this.db.Context)
}

func (this *AnimalsRepository) Remove(id uint64) error {
	return this.db.Client.Animal.
		DeleteOneID(id).
		Exec(this.db.Context)
}

func (this *AnimalsRepository) AddType(id uint64, typeId uint64) (*ent.Animal, error) {
	return this.db.Client.Animal.
		UpdateOneID(id).
		AddTypeIDs(typeId).
		Save(this.db.Context)
}

func (this *AnimalsRepository) RemoveType(id uint64, typeId uint64) (*ent.Animal, error) {
	return this.db.Client.Animal.
		UpdateOneID(id).
		RemoveTypeIDs(typeId).
		Save(this.db.Context)
}
