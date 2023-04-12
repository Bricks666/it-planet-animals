package locations

import (
	"animals/ent"
	"animals/shared"

	"github.com/gofiber/fiber/v2"
)

type LocationsController struct {
	locationsService *LocationsService
}

var Controller LocationsController

func NewLocationsController(locationsService *LocationsService) *LocationsController {
	return &LocationsController{
		locationsService: locationsService,
	}
}

func init() {
	Controller = *NewLocationsController(
		&Service,
	)
}

func (lc *LocationsController) GetOne(ct *fiber.Ctx) error {
	var params = NewLocationParams()
	var err error

	err = shared.GetParams(ct, params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var location = NewLocationDto()
	location, err = lc.locationsService.GetOne(params.Id)
	if err != nil {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(location)
}
func (lc *LocationsController) Create(ct *fiber.Ctx) error {
	var body = NewLocationBodyDto()
	var err error

	if bodyContainsNull(ct.Body()) {
		return ct.Status(fiber.StatusBadRequest).JSON("")
	}

	err = shared.GetBody(ct, body)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var location = NewLocationDto()
	location, err = lc.locationsService.Create(body)
	if ent.IsConstraintError(err) {
		return ct.Status(fiber.StatusConflict).JSON(err.Error())
	}
	if ent.IsNotFound(err) {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return ct.Status(fiber.StatusCreated).JSON(location)
}

func (lc *LocationsController) Update(ct *fiber.Ctx) error {
	var params = NewLocationParams()
	var body = NewLocationBodyDto()
	var err error

	err = shared.GetParams(ct, params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if bodyContainsNull(ct.Body()) {
		return ct.Status(fiber.StatusBadRequest).JSON("")
	}

	err = shared.GetBody(ct, body)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var location = NewLocationDto()
	location, err = lc.locationsService.Update(params.Id, body)
	if ent.IsNotFound(err) {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}
	if ent.IsConstraintError(err) {
		return ct.Status(fiber.StatusConflict).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(location)
}

func (lc *LocationsController) Remove(ct *fiber.Ctx) error {
	var params = NewLocationParams()
	var err error

	err = shared.GetParams(ct, params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = lc.locationsService.Remove(params.Id)
	if ent.IsNotFound(err) {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}
	if ent.IsConstraintError(err) {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON("Ok")
}
