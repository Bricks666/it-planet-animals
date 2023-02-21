package animaltypes

import (
	"animals/ent"
	"animals/shared"

	"github.com/gofiber/fiber/v2"
)

var Controller AnimalTypesController

type AnimalTypesController struct {
	animalTypesService *AnimalTypesService
}

func init() {
	Controller = AnimalTypesController{
		animalTypesService: &Service,
	}
}

func (this *AnimalTypesController) GetOne(ct *fiber.Ctx) error {
	var params AnimalTypeParamsDto
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

	var animalType *AnimalTypeDto
	animalType, err = this.animalTypesService.GetOne(params.Id)
	if err != nil {
		if ent.IsNotFound(err) {
			return ct.Status(fiber.StatusNotFound).JSON(err.Error())
		}

		return ct.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(animalType)

}

func (this *AnimalTypesController) Create(ct *fiber.Ctx) error {
	var dto CreateAnimaTypeDto
	var validationErrors []*shared.ErrorResponse
	var err error

	err = ct.BodyParser(&dto)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	validationErrors = shared.ValidateStruct(&dto)
	if validationErrors != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	var animalType *AnimalTypeDto
	animalType, err = this.animalTypesService.Create(&dto)
	if err != nil {
		if ent.IsConstraintError(err) {
			return ct.Status(fiber.StatusConflict).JSON(err.Error())
		}

		return ct.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return ct.Status(fiber.StatusCreated).JSON(animalType)
}

func (this *AnimalTypesController) Update(ct *fiber.Ctx) error {
	var params AnimalTypeParamsDto
	var dto UpdateAnimaTypeDto
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

	err = ct.BodyParser(&dto)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	validationErrors = shared.ValidateStruct(&dto)
	if validationErrors != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	var animalType *AnimalTypeDto
	animalType, err = this.animalTypesService.Update(params.Id, &dto)
	if err != nil {

		if ent.IsConstraintError(err) {
			return ct.Status(fiber.StatusConflict).JSON(err.Error())
		}

		if ent.IsNotFound(err) {
			return ct.Status(fiber.StatusNotFound).JSON(err.Error())
		}

		return ct.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(animalType)
}

func (this *AnimalTypesController) Remove(ct *fiber.Ctx) error {
	var params AnimalTypeParamsDto
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

	err = this.animalTypesService.Remove(params.Id)

	if err != nil {
		if ent.IsNotFound(err) {
			return ct.Status(fiber.StatusNotFound).JSON(err.Error())
		}

		return ct.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON("")
}
