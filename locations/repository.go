package locations

import (
	"animals/ent"
	"animals/ent/location"
	"animals/shared"
)

var Repository LocationRepository

type LocationRepository struct {
	db *shared.DB
}

func init() {
	Repository = LocationRepository{
		db: &shared.Database,
	}
}

func (lr *LocationRepository) GetOne(id uint64) (*ent.Location, error) {
	return lr.db.Client.Location.
		Query().
		Where(location.ID(id)).
		Only(lr.db.Context)
}

func (lr *LocationRepository) Create(dto *CreateLocationDto) (*ent.Location, error) {
	return lr.db.Client.Location.
		Create().
		SetLatitude(dto.Latitude).
		SetLongitude(dto.Longitude).
		Save(lr.db.Context)
}

func (lr *LocationRepository) Update(id uint64, dto *UpdateLocationDto) (*ent.Location, error) {
	return lr.db.Client.Location.
		UpdateOneID(id).
		SetLatitude(dto.Latitude).
		SetLongitude(dto.Longitude).
		Save(lr.db.Context)
}

func (lr *LocationRepository) Remove(id uint64) error {
	return lr.db.Client.Location.
		DeleteOneID(id).
		Exec(lr.db.Context)
}
