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

func (uc *UsersController) GetAll(ct *fiber.Ctx) error {
	var query UsersSearchQueryDto
	var validationErrors []*shared.ErrorResponse
	var err = ct.QueryParser(&query)
	var size = ct.QueryInt("size", 10)

	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if query.Size == 0 {
		query.Size = uint32(size)
	}

	validationErrors = shared.ValidateStruct(&query)

	if validationErrors != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(validationErrors)

	}

	var users []*SecurityUserDto

	users, err = uc.usersService.GetAll(&query)

	if err != nil {
		return ct.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(users)
}

func (uc *UsersController) GetOne(ct *fiber.Ctx) error {
	var params UserParamsDto
	var validationErrors []*shared.ErrorResponse
	var err = ct.ParamsParser(&params)

	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	validationErrors = shared.ValidateStruct(&params)

	if validationErrors != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(validationErrors)

	}

	user, err := uc.usersService.GetOne(params.Id)

	if err != nil {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(user)

}

func (uc *UsersController) Create(ct *fiber.Ctx) error {
	var dto CreateUserDto

	err := ct.BodyParser(&dto)

	if err != nil {
		return fiber.ErrBadRequest
	}

	validationErrors := shared.ValidateStruct(&dto)

	if validationErrors != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(validationErrors)

	}

	var user *SecurityUserDto

	user, err = uc.usersService.Create(&dto)

	if err != nil {
		return fiber.ErrConflict
	}

	return ct.Status(http.StatusCreated).JSON(user)
}

func (uc *UsersController) Update(ct *fiber.Ctx) error {
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

	user, err = uc.usersService.Update(params.Id, &dto)

	if err != nil {
		if shared.IsInstanceOf(&err, new(*ent.ConstraintError)) {
			return ct.Status(fiber.StatusConflict).JSON("Email already exists")
		}

		return ct.Status(fiber.StatusBadRequest).JSON("")
	}

	return ct.Status(fiber.StatusOK).JSON(user)
}

func (uc *UsersController) Remove(ct *fiber.Ctx) error {
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

	err = uc.usersService.Remove(params.Id)

	if err != nil {
		return ct.Status(fiber.StatusForbidden).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON("")
}
