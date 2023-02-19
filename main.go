package main

import (
	"fmt"

	"animals/shared"
	"animals/users"

	"github.com/gofiber/fiber/v2"
)

func setupRouter() *fiber.App {
	router := fiber.New()

	router.Post("/registration", users.Controller.Create)

	usersRouter := router.Group("accounts")

	usersRouter.Get("/search", users.Controller.GetAll)
	usersRouter.Get("/:id", users.Controller.GetOne)
	usersRouter.Put("/:id", users.Controller.Update)
	usersRouter.Delete("/:id", users.Controller.Remove)

	return router
}

func main() {
	router := setupRouter()

	router.Get("/ping", func(ct *fiber.Ctx) error {
		ct.Status(200).SendString("pong!")

		return nil
	})

	address := fmt.Sprintf(":%s", shared.ENV["PORT"])

	router.Listen(address)

}
