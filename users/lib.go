package users

import "github.com/gofiber/fiber/v2"

func AuthHasSameId(ct *fiber.Ctx, id uint32) error {
	var user = ct.Locals(USER_LOCALS).(*SecurityUserDto)

	if user.Id != id {
		return fiber.ErrForbidden
	}

	return nil
}
