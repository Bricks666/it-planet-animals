package areas

import (
	"animals/ent"
	"animals/locations"
	"animals/shared"
)

type AreasService struct {
	areasRepository  *AreasRepository
	locationsService *locations.LocationsService
}

var Service AreasService

func NewAreasService(areasRepository *AreasRepository,
	locationsService *locations.LocationsService) *AreasService {
	return &AreasService{
		areasRepository:  areasRepository,
		locationsService: locationsService,
	}
}

func init() {
	Service = *NewAreasService(&Repository, &locations.Service)
}

func (this *AreasService) GetOne(params *AreaParamsDto) (*AreaDto, error) {
	var area, err = this.areasRepository.GetOne(params)
	if err != nil {
		return nil, err
	}
	return prepareAreaDto(area), nil

}

func (this *AreasService) Create(dto *AreaBodyDto) (*AreaDto, error) {
	var innerDto = NewInnerAreaDto()
	innerDto.Name = dto.Name
	innerDto.AreaPointIds = this.getLocationIds(dto.AreaPoints)

	var area, err = this.areasRepository.Create(innerDto)

	if err != nil {
		return nil, err
	}

	var params = NewAreaParamsDto()
	params.Id = area.ID

	return this.GetOne(params)

}

func (this *AreasService) Update(params *AreaParamsDto, dto *AreaBodyDto) (*AreaDto, error) {
	var innerDto = NewInnerAreaDto()
	innerDto.Name = dto.Name
	innerDto.AreaPointIds = this.getLocationIds(dto.AreaPoints)

	var _, err = this.areasRepository.Update(params, innerDto)

	if err != nil {
		return nil, err
	}

	return this.GetOne(params)
}

func (this *AreasService) Remove(params *AreaParamsDto) error {
	return this.areasRepository.Remove(params)
}

func (this *AreasService) getLocationIds(areaPoints []*AreaPointDto) []uint64 {
	var locs = shared.Map(areaPoints, func(areaPoint *AreaPointDto) *locations.LocationDto {
		var params = locations.NewLocationBodyDto()
		params.Latitude = areaPoint.Latitude
		params.Longitude = areaPoint.Longitude
		return this.locationsService.GetOrCreate(params)
	})
	return shared.Map(locs, func(loc *locations.LocationDto) uint64 {
		return loc.Id
	})
}

func prepareAreaDto(area *ent.Area) *AreaDto {
	var dto = NewAreaDto()

	dto.Id = area.ID
	dto.Name = area.Name

	dto.AreaPoints = shared.Map(area.Edges.Points, func(element *ent.Location) *AreaPointDto {
		var areaPointDto = NewAreaPointDto()

		areaPointDto.Latitude = element.Latitude
		areaPointDto.Longitude = element.Longitude

		return areaPointDto
	})

	return dto
}
