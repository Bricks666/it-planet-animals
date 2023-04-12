package areas

import (
	"animals/shared"

	"github.com/gofiber/fiber/v2"
)

type AreasController struct {
	areasService *AreasService
}

var Controller AreasController

func NewAreasController(usersService *AreasService) *AreasController {
	return &AreasController{
		areasService: &Service,
	}
}

func init() {
	Controller = *NewAreasController(&Service)
}

func (this *AreasController) GetOne(ct *fiber.Ctx) error {
	var params = NewAreaParamsDto()
	var err error

	err = shared.GetParams(ct, params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var area *AreaDto

	area, err = this.areasService.GetOne(params)

	if err != nil {
		return ct.Status(fiber.StatusNotFound).JSON("")
	}

	return ct.Status(fiber.StatusOK).JSON(area)
}

func (this *AreasController) Create(ct *fiber.Ctx) error {
	var body = NewAreaBodyDto()
	var err error

	err = shared.GetBody(ct, body)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var area *AreaDto
	area, err = this.areasService.Create(body)

	if err != nil {
		return ct.Status(fiber.StatusConflict).JSON("")
	}

	return ct.Status(fiber.StatusCreated).JSON(area)
}

func (this *AreasController) Update(ct *fiber.Ctx) error {
	var params = NewAreaParamsDto()
	var body = NewAreaBodyDto()
	var err error

	err = shared.GetParams(ct, params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = shared.GetBody(ct, body)
	if err != nil {
		return fiber.ErrBadRequest
	}

	var area *AreaDto
	area, err = this.areasService.Update(params, body)
	if err != nil {
		return ct.Status(fiber.StatusConflict).JSON("")
	}

	return ct.Status(fiber.StatusOK).JSON(area)
}

func (this *AreasController) Remove(ct *fiber.Ctx) error {
	var params = NewAreaParamsDto()
	var err error

	err = shared.GetParams(ct, params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = this.areasService.Remove(params)

	if err != nil {
		return ct.Status(fiber.StatusNotFound).JSON("")
	}

	return ct.Status(fiber.StatusOK).JSON("OK")
}
