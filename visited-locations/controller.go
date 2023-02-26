package visitedlocation

import (
	"animals/ent"
	"animals/shared"

	"github.com/gofiber/fiber/v2"
)

type VisitedLocationsController struct {
	visitedLocationsService *VisitedLocationsService
}

var Controller VisitedLocationsController

func NewVisitedLocationsController(visitedLocationsService *VisitedLocationsService) *VisitedLocationsController {
	return &VisitedLocationsController{
		visitedLocationsService: visitedLocationsService,
	}
}

func init() {
	Controller = *NewVisitedLocationsController(&Service)
}

func (this *VisitedLocationsController) GetAll(ct *fiber.Ctx) error {
	var query = NewVisitedLocationSearchQueryDto()
	var params = NewVisitedLocationParamsDto()
	var err error

	err = shared.GetQuery(ct, query)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = shared.GetParams(ct, params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var locations = []*VisitedLocationDto{}
	locations, err = this.visitedLocationsService.GetAll(params.AnimalId, query)

	if err != nil {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(locations)
}

func (this *VisitedLocationsController) Create(ct *fiber.Ctx) error {
	var params = NewVisitedLocationMutationParamsDto()
	var err error

	err = shared.GetParams(ct, params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var animalLocation *VisitedLocationDto
	animalLocation, err = this.visitedLocationsService.Create(params.AnimalId, params.LocationId)
	if ent.IsNotFound(err) {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	if ent.IsConstraintError(err) {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return ct.Status(fiber.StatusCreated).JSON(animalLocation)
}

func (this *VisitedLocationsController) Update(ct *fiber.Ctx) error {
	var params = NewVisitedLocationParamsDto()
	var body = NewUpdateVisitedLocationDto()
	var err error

	err = shared.GetParams(ct, params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = shared.GetBody(ct, body)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var animalLocation *VisitedLocationDto
	animalLocation, err = this.visitedLocationsService.Update(params.AnimalId, body)
	if ent.IsConstraintError(err) {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if ent.IsNotFound(err) {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(animalLocation)
}

func (this *VisitedLocationsController) Remove(ct *fiber.Ctx) error {
	var params = NewVisitedLocationMutationParamsDto()
	var err error

	err = shared.GetParams(ct, params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = this.visitedLocationsService.Remove(params.AnimalId, params.LocationId)
	if err != nil {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON("deleted")
}
