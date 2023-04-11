package users

import (
	"animals/ent"
	"animals/shared"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UsersController struct {
	usersService *UsersService
}

var Controller UsersController

func NewUsersController(usersService *UsersService) *UsersController {
	return &UsersController{
		usersService: &Service,
	}
}

func init() {
	Controller = *NewUsersController(&Service)
}

func (this *UsersController) GetAll(ct *fiber.Ctx) error {
	var query = NewUsersSearchQueryDto()
	var err error

	err = shared.GetQuery(ct, query)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var users = []*SecurityUserDto{}
	users, err = this.usersService.GetAll(query)

	if err != nil {
		return ct.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(users)
}

func (this *UsersController) GetOne(ct *fiber.Ctx) error {
	var params = NewUserParamsDto()
	var err error

	err = shared.GetParams(ct, params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var authUser = ct.Locals(USER_LOCALS).(*SecurityUserDto)
	if authUser.Id != params.Id && authUser.Role != ADMIN_ROLE {
		return ct.Status(fiber.StatusForbidden).JSON("invalid id")
	}

	user, err := this.usersService.GetOne(params.Id)
	if err != nil {
		return ct.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON(user)

}

func (this *UsersController) Create(ct *fiber.Ctx) error {
	var body = NewCreateUserDto()
	var err error

	err = shared.GetBody(ct, body)
	if err != nil {
		return fiber.ErrBadRequest
	}

	var user *SecurityUserDto
	user, err = this.usersService.Create(body)
	if err != nil {
		return fiber.ErrConflict
	}

	return ct.Status(http.StatusCreated).JSON(user)
}

func (this *UsersController) Update(ct *fiber.Ctx) error {
	var params = NewUserParamsDto()
	var body = NewUpdateUserDto()
	var err error

	err = shared.GetParams(ct, params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var authUser = ct.Locals(USER_LOCALS).(*SecurityUserDto)

	if authUser.Id != params.Id && authUser.Role != ADMIN_ROLE {
		return ct.Status(fiber.StatusForbidden).JSON("invalid id")
	}

	err = shared.GetBody(ct, body)
	if err != nil {
		return fiber.ErrBadRequest
	}

	var user = NewSecurityUserDto()
	user, err = this.usersService.Update(params.Id, body)
	if ent.IsConstraintError(err) {
		return ct.Status(fiber.StatusConflict).JSON("Email already exists")
	}

	if ent.IsNotFound(err) {
		return ct.Status(fiber.StatusBadRequest).JSON("")
	}

	return ct.Status(fiber.StatusOK).JSON(user)
}

func (this *UsersController) Remove(ct *fiber.Ctx) error {
	var params = NewUserParamsDto()
	var err error

	err = shared.GetParams(ct, params)
	if err != nil {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var authUser = ct.Locals(USER_LOCALS).(*SecurityUserDto)

	if authUser.Id != params.Id && authUser.Role != ADMIN_ROLE {
		return ct.Status(fiber.StatusForbidden).JSON("invalid id")
	}

	err = this.usersService.Remove(params.Id)
	if ent.IsNotFound(err) {
		return ct.Status(fiber.StatusForbidden).JSON(err.Error())
	}

	if ent.IsConstraintError(err) {
		return ct.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return ct.Status(fiber.StatusOK).JSON("")
}
