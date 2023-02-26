package animals

import (
	"animals/ent"
	"animals/shared"

	"github.com/gofiber/fiber/v2"
)

type AnimalsController struct {
	animalsService *AnimalsService
}

var Controller AnimalsController

func NewAnimalsController(animalsService *AnimalsService) *AnimalsController {
	return &AnimalsController{
		animalsService: animalsService,
	}
}

func init() {
	Controller = *NewAnimalsController(
		&Service,
	)
}

func (this *AnimalsController) GetAll(ct *fiber.Ctx) error {
	var query = NewAnimalsSearchQueryDto()
	var err error

	err = shared.GetQuery(ct, query)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var animals []*AnimalDto
	animals, err = this.animalsService.GetAll(query)
	if err != nil {
		return ct.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(animals)
}

func (this *AnimalsController) GetOne(ct *fiber.Ctx) error {
	var params = NewAnimalParamsDto()
	var err error

	err = shared.GetParams(ct, params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var animal = NewAnimalDto()
	animal, err = this.animalsService.GetOne(params.Id)

	if err != nil {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(animal)
}

func (this *AnimalsController) Create(ct *fiber.Ctx) error {
	var body = NewCreateAnimalDto()
	var err error

	err = shared.GetBody(ct, body)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	onlyUnique := shared.HasOnlyUnique(body.AnimalTypes)
	if !onlyUnique {
		return ct.Status(fiber.StatusConflict).JSON(onlyUnique)
	}

	var animal = NewAnimalDto()
	animal, err = this.animalsService.Create(body)

	if err != nil {

		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return ct.Status(fiber.StatusCreated).JSON(animal)
}

func (this *AnimalsController) Update(ct *fiber.Ctx) error {
	var params = NewAnimalParamsDto()
	var body = NewUpdateAnimalDto()
	var err error

	err = shared.GetParams(ct, params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = shared.GetBody(ct, body)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var animal = NewAnimalDto()
	animal, err = this.animalsService.Update(params.Id, body)

	if ent.IsConstraintError(err) {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if ent.IsNotFound(err) {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(animal)
}

func (this *AnimalsController) Remove(ct *fiber.Ctx) error {
	var params = NewAnimalParamsDto()
	var err error

	err = shared.GetParams(ct, params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = this.animalsService.Remove(params.Id)
	if ent.IsNotFound(err) {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}
	if ent.IsConstraintError(err) {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON("deleted")
}

func (this *AnimalsController) AddType(ct *fiber.Ctx) error {
	var params = NewAnimalTypeParamsDto()
	var err error

	err = shared.GetParams(ct, params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var animal = NewAnimalDto()
	animal, err = this.animalsService.AddType(params.Id, params.TypeId)

	if ent.IsNotFound(err) {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	if ent.IsConstraintError(err) {
		return ct.Status(fiber.StatusConflict).JSON(err.Error())
	}

	return ct.Status(fiber.StatusCreated).JSON(animal)
}

func (this *AnimalsController) ReplaceType(ct *fiber.Ctx) error {
	var params = NewAnimalParamsDto()
	var body = NewReplaceAnimalTypeDto()
	var err error

	err = shared.GetParams(ct, params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = shared.GetBody(ct, body)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var animal = NewAnimalDto()
	animal, err = this.animalsService.ReplaceType(params.Id, body)

	if ent.IsNotFound(err) {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	if ent.IsConstraintError(err) {
		return ct.Status(fiber.StatusConflict).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(animal)
}

func (this *AnimalsController) RemoveType(ct *fiber.Ctx) error {
	var params = NewAnimalTypeParamsDto()
	var err error

	err = shared.GetParams(ct, params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var animal = NewAnimalDto()
	animal, err = this.animalsService.RemoveType(params.Id, params.TypeId)

	if ent.IsNotFound(err) {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	if ent.IsConstraintError(err) {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(animal)
}
