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
		return ct.Status(fiber.StatusBadRequest).JSON(err)
	}

	validationErrors = shared.ValidateStruct(&params)
	if validationErrors != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err)
	}

	var animalType *ent.AnimalType
	animalType, err = this.animalTypesService.GetOne(params.Id)
	if err != nil {
		if shared.IsInstanceOf(&err, new(*ent.NotFoundError)) {
			return ct.Status(fiber.StatusNotFound).JSON(err)
		}

		return ct.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ct.Status(fiber.StatusOK).JSON(animalType)

}

func (this *AnimalTypesController) Create(ct *fiber.Ctx) error {
	var dto CreateAnimaTypeDto
	var validationErrors []*shared.ErrorResponse
	var err error

	err = ct.BodyParser(&dto)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err)
	}

	validationErrors = shared.ValidateStruct(&dto)
	if validationErrors != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err)
	}

	var animalType *ent.AnimalType
	animalType, err = this.animalTypesService.Create(&dto)
	if err != nil {
		if shared.IsInstanceOf(&err, new(*ent.ConstraintError)) {
			return ct.Status(fiber.StatusConflict).JSON(err)
		}

		return ct.Status(fiber.StatusInternalServerError).JSON(err)
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
		return ct.Status(fiber.StatusBadRequest).JSON(err)
	}

	validationErrors = shared.ValidateStruct(&params)
	if validationErrors != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err)
	}

	err = ct.BodyParser(&dto)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err)
	}

	validationErrors = shared.ValidateStruct(&dto)
	if validationErrors != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err)
	}

	var animalType *ent.AnimalType
	animalType, err = this.animalTypesService.Update(params.Id, &dto)
	if err != nil {
		if shared.IsInstanceOf(&err, new(*ent.ConstraintError)) {
			return ct.Status(fiber.StatusConflict).JSON(err)
		}

		if shared.IsInstanceOf(&err, new(*ent.NotFoundError)) {
			return ct.Status(fiber.StatusNotFound).JSON(err)
		}

		return ct.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ct.Status(fiber.StatusOK).JSON(animalType)
}

func (this *AnimalTypesController) Remove(ct *fiber.Ctx) error {
	var params AnimalTypeParamsDto
	var validationErrors []*shared.ErrorResponse
	var err error

	err = ct.ParamsParser(&params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err)
	}

	validationErrors = shared.ValidateStruct(&params)
	if validationErrors != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err)
	}

	err = this.animalTypesService.Remove(params.Id)

	if err != nil {
		if shared.IsInstanceOf(&err, new(*ent.NotFoundError)) {
			return ct.Status(fiber.StatusNotFound).JSON(err)
		}

		return ct.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ct.Status(fiber.StatusOK).JSON("")
}
