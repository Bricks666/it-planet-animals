package animalslocations

import (
	"animals/ent"
	"animals/ent/animalslocations"
	"animals/ent/predicate"
	"animals/shared"
)

type AnimalsLocationsRepository struct {
	db *shared.DB
}

var Repository AnimalsLocationsRepository

func init() {
	Repository = AnimalsLocationsRepository{
		db: &shared.Database,
	}
}

func (this *AnimalsLocationsRepository) GetAll(animalId uint64, dto *AnimalsLocationSearchQueryDto) ([]*ent.AnimalsLocations, error) {
	query := this.db.Client.AnimalsLocations.Query().
		Where(animalslocations.AnimalId(animalId))

	var conditions []predicate.AnimalsLocations

	if !dto.AfterDateTime.IsZero() {
		conditions = append(conditions, animalslocations.DateTimeOfVisitLocationPointGTE(dto.AfterDateTime))
	}

	if !dto.BeforeDateTime.IsZero() {
		conditions = append(conditions, animalslocations.DateTimeOfVisitLocationPointLTE(dto.BeforeDateTime))
	}

	if len(conditions) != 0 {
		query = query.Where(animalslocations.Or(conditions...))
	}

	return query.
		Offset(int(dto.From)).
		Limit(int(dto.Size)).
		Order(ent.Asc(animalslocations.FieldDateTimeOfVisitLocationPoint)).
		All(this.db.Context)
}

func (this *AnimalsLocationsRepository) Create() (*ent.AnimalsLocations, error) {
	return nil, nil
}

func (this *AnimalsLocationsRepository) Update() (*ent.AnimalsLocations, error) {
	return nil, nil
}

func (this *AnimalsLocationsRepository) Remove() error {
	return nil
}
