package animalslocations

import (
	"animals/shared"

	"github.com/gofiber/fiber/v2"
)

type AnimalsLocationsController struct {
	animalsLocationsService *AnimalsLocationsService
}

var Controller AnimalsLocationsController

func init() {
	Controller = AnimalsLocationsController{
		animalsLocationsService: &Service,
	}
}

func (this *AnimalsLocationsController) GetAll(ct *fiber.Ctx) error {
	var query AnimalsLocationSearchQueryDto
	var params AnimalsLocationParamsDto
	var validationError []*shared.ErrorResponse
	var err error

	err = ct.QueryParser(&query)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var size = ct.QueryInt("size", 10)
	if query.Size == 0 {
		query.Size = uint32(size)
	}

	validationError = shared.ValidateStruct(&query)
	if validationError != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(validationError)
	}

	err = ct.ParamsParser(&params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	validationError = shared.ValidateStruct(&params)
	if validationError != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(validationError)
	}

	var locations = []*AnimalsLocationsDto{}
	locations, err = this.animalsLocationsService.GetAll(params.AnimalId, &query)

	if err != nil {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(locations)
}

func (this *AnimalsLocationsController) Create(ct *fiber.Ctx) error {
	return nil
}

func (this *AnimalsLocationsController) Update(ct *fiber.Ctx) error {
	return nil
}

func (this *AnimalsLocationsController) Remove(ct *fiber.Ctx) error {
	return nil
}
