package locations

import (
	"animals/ent"
	"animals/shared"

	"github.com/gofiber/fiber/v2"
)

var Controller LocationsController

type LocationsController struct {
	locationsService *LocationsService
}

func init() {
	Controller = LocationsController{
		locationsService: &Service,
	}
}

func (lc *LocationsController) GetOne(ct *fiber.Ctx) error {
	var params LocationParams
	var validationErrors []*shared.ErrorResponse
	var err error = ct.ParamsParser(&params)

	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	validationErrors = shared.ValidateStruct(&params)

	if validationErrors != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	var location *ent.Location

	location, err = lc.locationsService.GetOne(params.Id)

	if err != nil {
		return ct.Status(fiber.StatusForbidden).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(location)
}
func (lc *LocationsController) Create(ct *fiber.Ctx) error {
	var dto CreateLocationDto
	var validationErrors []*shared.ErrorResponse
	var err error = ct.BodyParser(&dto)

	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	validationErrors = shared.ValidateStruct(&dto)

	if validationErrors != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	var location *ent.Location

	location, err = lc.locationsService.Create(&dto)

	if err != nil {
		return ct.Status(fiber.StatusForbidden).JSON(err.Error())
	}

	return ct.Status(fiber.StatusCreated).JSON(location)
}

func (lc *LocationsController) Update(ct *fiber.Ctx) error {
	var params LocationParams
	var dto UpdateLocationDto
	var validationErrors []*shared.ErrorResponse

	var err error = ct.ParamsParser(&params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	validationErrors = shared.ValidateStruct(&params)
	if validationErrors != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	err = ct.BodyParser(&dto)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	validationErrors = shared.ValidateStruct(&dto)
	if validationErrors != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	var location *ent.Location

	location, err = lc.locationsService.Update(params.Id, &dto)

	if err != nil {
		if err.Error() == "404" {
			return ct.Status(fiber.StatusNotFound).JSON(err.Error())
		}
		return ct.Status(fiber.StatusConflict).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(location)
}

func (lc *LocationsController) Remove(ct *fiber.Ctx) error {
	var params LocationParams
	var validationErrors []*shared.ErrorResponse
	var err error = ct.ParamsParser(&params)

	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	validationErrors = shared.ValidateStruct(&params)

	if validationErrors != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	err = lc.locationsService.Remove(params.Id)

	if err != nil {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON("Ok")
}
