package animals

import (
	"animals/shared"

	"github.com/gofiber/fiber/v2"
)

var Controller AnimalsController

type AnimalsController struct {
	animalsService *AnimalsService
}

func init() {
	Controller = AnimalsController{
		animalsService: &Service,
	}
}

func (this *AnimalsController) GetAll(ct *fiber.Ctx) error {
	var query AnimalsSearchQueryDto
	var validationErrors []*shared.ErrorResponse
	var err error

	err = ct.QueryParser(&query)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var size = ct.QueryInt("size", 10)
	if query.Size == 0 {
		query.Size = uint32(size)
	}

	validationErrors = shared.ValidateStruct(&query)
	if validationErrors != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	var animals []*AnimalDto
	animals, err = this.animalsService.GetAll(&query)

	if err != nil {
		return ct.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(animals)
}

func (this *AnimalsController) GetOne(ct *fiber.Ctx) error {
	var params AnimalParamsDto
	var validationErrors []*shared.ErrorResponse
	var err error

	err = ct.ParamsParser(&params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	validationErrors = shared.ValidateStruct(&params)
	if validationErrors != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	var animal *AnimalDto
	animal, err = this.animalsService.GetOne(params.Id)

	if err != nil {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(animal)
}

func (this *AnimalsController) Create(ct *fiber.Ctx) error {
	return nil
}

func (this *AnimalsController) Update(ct *fiber.Ctx) error {
	return nil
}

func (this *AnimalsController) Remove(ct *fiber.Ctx) error {
	return nil
}

func (this *AnimalsController) AddType(ct *fiber.Ctx) error {
	return nil
}

func (this *AnimalsController) ReplaceType(ct *fiber.Ctx) error {
	return nil
}

func (this *AnimalsController) RemoveType(ct *fiber.Ctx) error {
	return nil
}
