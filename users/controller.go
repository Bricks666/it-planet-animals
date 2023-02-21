package users

import (
	"animals/ent"
	"animals/shared"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var Controller UsersController

func init() {
	Controller = UsersController{
		usersService: &Service,
	}
}

type UsersController struct {
	usersService *UsersService
}

func (this *UsersController) GetAll(ct *fiber.Ctx) error {
	var query UsersSearchQueryDto
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

	var users []*SecurityUserDto
	users, err = this.usersService.GetAll(&query)

	if err != nil {
		return ct.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(users)
}

func (this *UsersController) GetOne(ct *fiber.Ctx) error {
	var params UserParamsDto
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

	user, err := this.usersService.GetOne(params.Id)
	if err != nil {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(user)

}

func (this *UsersController) Create(ct *fiber.Ctx) error {
	var dto CreateUserDto
	var err error

	err = ct.BodyParser(&dto)
	if err != nil {
		return fiber.ErrBadRequest
	}

	validationErrors := shared.ValidateStruct(&dto)
	if validationErrors != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(validationErrors)

	}

	var user *SecurityUserDto
	user, err = this.usersService.Create(&dto)
	if err != nil {
		return fiber.ErrConflict
	}

	return ct.Status(http.StatusCreated).JSON(user)
}

func (this *UsersController) Update(ct *fiber.Ctx) error {
	var params UserParamsDto
	var dto UpdateUserDto
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
		return fiber.ErrBadRequest
	}

	validationErrors = shared.ValidateStruct(&dto)
	if validationErrors != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(validationErrors)

	}

	var user *SecurityUserDto
	user, err = this.usersService.Update(params.Id, &dto)
	if err != nil {
		if ent.IsConstraintError(err) {
			return ct.Status(fiber.StatusConflict).JSON("Email already exists")
		}

		return ct.Status(fiber.StatusBadRequest).JSON("")
	}

	return ct.Status(fiber.StatusOK).JSON(user)
}

func (this *UsersController) Remove(ct *fiber.Ctx) error {
	var params UserParamsDto
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

	err = this.usersService.Remove(params.Id)
	if err != nil {
		return ct.Status(fiber.StatusForbidden).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON("")
}
