package areas

import (
	"animals/ent"
	"animals/ent/area"
	"animals/shared"
)

type AreasRepository struct {
	db *shared.DB
}

var Repository AreasRepository

func NewAreasRepository(db *shared.DB) *AreasRepository {
	return &AreasRepository{
		db: db,
	}
}

func init() {
	Repository = *NewAreasRepository(&shared.Database)
}

func (this *AreasRepository) GetAll() ([]*ent.Area, error) {
	return this.db.Client.Area.Query().
		WithPoints().
		All(this.db.Context)
}

func (this *AreasRepository) GetOne(params *AreaParamsDto) (*ent.Area, error) {
	return this.db.Client.Area.Query().
		Where(area.ID(params.Id)).
		WithPoints().
		Only(this.db.Context)
}

func (this *AreasRepository) Create(dto *InnerAreaDto) (*ent.Area, error) {
	return this.db.Client.Area.Create().
		SetName(dto.Name).
		AddPointIDs(dto.AreaPointIds...).
		Save(this.db.Context)
}

func (this *AreasRepository) Update(params *AreaParamsDto, dto *InnerAreaDto) (*ent.Area, error) {
	return this.db.Client.Area.UpdateOneID(params.Id).
		SetName(dto.Name).
		AddPointIDs(dto.AreaPointIds...).
		Save(this.db.Context)
}

func (this *AreasRepository) Remove(params *AreaParamsDto) error {
	return this.db.Client.Area.DeleteOneID(params.Id).
		Exec(this.db.Context)
}
