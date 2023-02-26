package visitedlocation

import (
	"animals/ent"
	"animals/ent/predicate"
	"animals/ent/visitedlocation"
	"animals/shared"
	"time"
)

type VisitedLocationsRepository struct {
	db *shared.DB
}

var Repository VisitedLocationsRepository

func NewVisitedLocationsRepository(db *shared.DB) *VisitedLocationsRepository {
	return &VisitedLocationsRepository{
		db: db,
	}
}

func init() {
	Repository = *NewVisitedLocationsRepository(&shared.Database)
}

func (this *VisitedLocationsRepository) GetAll(animalId uint64, dto *VisitedLocationSearchQueryDto) ([]*ent.VisitedLocation, error) {
	query := this.db.Client.VisitedLocation.Query().
		Where(visitedlocation.AnimalID(animalId))

	var conditions []predicate.VisitedLocation

	if !dto.AfterDateTime.IsZero() {
		conditions = append(conditions, visitedlocation.DateTimeOfVisitLocationPointGTE(dto.AfterDateTime))
	}

	if !dto.BeforeDateTime.IsZero() {
		conditions = append(conditions, visitedlocation.DateTimeOfVisitLocationPointLTE(dto.BeforeDateTime))
	}

	if len(conditions) != 0 {
		query = query.Where(visitedlocation.Or(conditions...))
	}

	return query.
		Offset(int(dto.From)).
		Limit(int(dto.Size)).
		Order(ent.Asc(visitedlocation.FieldDateTimeOfVisitLocationPoint)).
		All(this.db.Context)
}

func (this *VisitedLocationsRepository) Create(animalId uint64, locationId uint64) (*ent.VisitedLocation, error) {
	return this.db.Client.VisitedLocation.
		Create().
		SetAnimalID(animalId).
		SetLocationID(locationId).
		SetDateTimeOfVisitLocationPoint(time.Now()).
		Save(this.db.Context)
}

func (this *VisitedLocationsRepository) Update(animalId uint64, dto *UpdateVisitedLocationDto) (*ent.VisitedLocation, error) {
	return this.db.Client.VisitedLocation.
		UpdateOneID(dto.VisitedAnimalLocationId).
		Where(visitedlocation.AnimalID(animalId)).
		SetLocationID(dto.NewLocationId).
		SetDateTimeOfVisitLocationPoint(time.Now()).
		Save(this.db.Context)
}

func (this *VisitedLocationsRepository) Remove(animalId uint64, visitedPointId uint64) error {
	return this.db.Client.VisitedLocation.
		DeleteOneID(visitedPointId).
		Where(visitedlocation.AnimalID(animalId)).
		Exec(this.db.Context)
}
