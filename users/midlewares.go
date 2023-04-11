package users

import (
	"animals/shared"
	"encoding/base64"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func DisableAuthCheck(ct *fiber.Ctx) error {
	ct.Locals(disableAuthCheck, true)

	return ct.Next()
}

func CheckAuth(ct *fiber.Ctx) error {
	var disableCheck = false
	d := ct.Locals(disableAuthCheck)
	if d != nil {
		disableCheck = d.(bool)
	}
	headers := ct.GetReqHeaders()
	authHeader := headers[fiber.HeaderAuthorization]

	if authHeader == "" {
		if disableCheck == true {
			return ct.Next()
		}

		return ct.Status(fiber.StatusUnauthorized).JSON("Header is empty")
	}

	parts := strings.Split(authHeader, " ")
	base64Str := parts[1]

	credentials, err := base64.StdEncoding.DecodeString(base64Str)

	if err != nil {
		return ct.Status(fiber.StatusUnauthorized).JSON("Invalid credentials")
	}

	userAndPassword := strings.Split(string(credentials), ":")
	if len(userAndPassword) != 2 {
		return ct.Status(fiber.StatusUnauthorized).JSON("Invalid credentials")
	}

	user, err := Service.VerifyUser(userAndPassword[0], userAndPassword[1])

	if err != nil {
		return ct.Status(fiber.StatusUnauthorized).JSON("Invalid credentials")
	}

	ct.Locals(USER_LOCALS, user)

	return ct.Next()
}

func CheckUnauthorized(ct *fiber.Ctx) error {
	headers := ct.GetReqHeaders()
	authHeader := headers[fiber.HeaderAuthorization]

	if authHeader != "" {
		return ct.Status(fiber.StatusForbidden).JSON("Header is empty")
	}
	return ct.Next()
}

func CheckRole(allowedRoles ...string) func(ct *fiber.Ctx) error {
	return func(ct *fiber.Ctx) error {
		var user = ct.Locals(USER_LOCALS).(*SecurityUserDto)

		if !shared.Contains(allowedRoles, user.Role) {
			return ct.Status(fiber.StatusForbidden).JSON("Invalid credentials")
		}

		return ct.Next()
	}
}
