package animaltypes

import (
	"animals/ent"
	"animals/shared"

	"github.com/gofiber/fiber/v2"
)

type AnimalTypesController struct {
	animalTypesService *AnimalTypesService
}

var Controller AnimalTypesController

func NewAnimalTypesController(animalTypesService *AnimalTypesService) *AnimalTypesController {
	return &AnimalTypesController{
		animalTypesService: animalTypesService,
	}
}

func init() {
	Controller = *NewAnimalTypesController(
		&Service,
	)
}

func (this *AnimalTypesController) GetOne(ct *fiber.Ctx) error {
	var params = NewAnimalTypeParamsDto()
	var err error

	err = shared.GetParams(ct, params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var animalType = NewAnimalTypeDto()
	animalType, err = this.animalTypesService.GetOne(params.Id)
	if ent.IsNotFound(err) {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(animalType)

}

func (this *AnimalTypesController) Create(ct *fiber.Ctx) error {
	var body = NewCreateAnimaTypeDto()
	var err error

	err = shared.GetBody(ct, body)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var animalType = NewAnimalTypeDto()
	animalType, err = this.animalTypesService.Create(body)
	if ent.IsConstraintError(err) {
		return ct.Status(fiber.StatusConflict).JSON(err.Error())
	}

	return ct.Status(fiber.StatusCreated).JSON(animalType)
}

func (this *AnimalTypesController) Update(ct *fiber.Ctx) error {
	var params = NewAnimalTypeParamsDto()
	var body = NewUpdateAnimaTypeDto()
	var err error

	err = shared.GetParams(ct, params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = shared.GetBody(ct, body)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var animalType = NewAnimalTypeDto()
	animalType, err = this.animalTypesService.Update(params.Id, body)
	if ent.IsConstraintError(err) {
		return ct.Status(fiber.StatusConflict).JSON(err.Error())
	}

	if ent.IsNotFound(err) {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(animalType)
}

func (this *AnimalTypesController) Remove(ct *fiber.Ctx) error {
	var params = NewAnimalTypeParamsDto()
	var err error

	err = shared.GetParams(ct, params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = this.animalTypesService.Remove(params.Id)
	if ent.IsNotFound(err) {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}
	if ent.IsConstraintError(err) {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON("")
}
