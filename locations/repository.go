package locations

import (
	"animals/ent"
	"animals/ent/location"
	"animals/shared"
)

type LocationRepository struct {
	db *shared.DB
}

var Repository LocationRepository

func NewLocationRepository(db *shared.DB) *LocationRepository {
	return &LocationRepository{
		db: db,
	}
}

func init() {
	Repository = *NewLocationRepository(
		&shared.Database,
	)
}

func (this *LocationRepository) GetOne(id uint64) (*ent.Location, error) {
	return this.db.Client.Location.
		Query().
		Where(location.ID(id)).
		Only(this.db.Context)
}

func (this *LocationRepository) Create(dto *LocationBodyDto) (*ent.Location, error) {
	return this.db.Client.Location.
		Create().
		SetLatitude(dto.Latitude).
		SetLongitude(dto.Longitude).
		Save(this.db.Context)
}

func (this *LocationRepository) Update(id uint64, dto *LocationBodyDto) (*ent.Location, error) {
	return this.db.Client.Location.
		UpdateOneID(id).
		SetLatitude(dto.Latitude).
		SetLongitude(dto.Longitude).
		Save(this.db.Context)
}

func (this *LocationRepository) Remove(id uint64) error {
	return this.db.Client.Location.
		DeleteOneID(id).
		Exec(this.db.Context)
}

func (this *LocationRepository) GetOrCreate(params *LocationBodyDto) (*ent.Location, error) {
	var response *ent.Location
	var err error

	response, err = this.db.Client.Location.
		Query().
		Where(location.Latitude(params.Latitude),
			location.Longitude(params.Longitude)).
		Only(this.db.Context)

	if err != nil {
		return this.Create(params)
	}

	return response, nil
}
