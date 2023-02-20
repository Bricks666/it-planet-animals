package main

import (
	"animals/animals"
	"animals/animaltypes"
	"animals/locations"
	"animals/users"

	"github.com/gofiber/fiber/v2"
)

func setupRouter() *fiber.App {
	router := fiber.New()

	router.Post("/registration", users.Controller.Create)

	usersRouter := router.Group("accounts")

	usersRouter.Get("/search", users.DisableAuthCheck, users.AuthMiddleware, users.Controller.GetAll)
	usersRouter.Get("/:id", users.DisableAuthCheck, users.AuthMiddleware, users.Controller.GetOne)
	usersRouter.Put("/:id", users.AuthMiddleware, users.Controller.Update)
	usersRouter.Delete("/:id", users.AuthMiddleware, users.Controller.Remove)

	locationsRouter := router.Group("locations")

	locationsRouter.Get("/:id", users.DisableAuthCheck, users.AuthMiddleware, locations.Controller.GetOne)
	locationsRouter.Post("/", users.AuthMiddleware, locations.Controller.Create)
	locationsRouter.Put("/:id", users.AuthMiddleware, locations.Controller.Update)
	locationsRouter.Delete("/:id", users.AuthMiddleware, locations.Controller.Remove)

	animalsRouter := router.Group("animals")

	animalsRouter.Get("/search", users.DisableAuthCheck, users.AuthMiddleware, animals.Controller.GetAll)
	animalsRouter.Get("/:id", users.DisableAuthCheck, users.AuthMiddleware, animals.Controller.GetOne)
	animalsRouter.Post("/", users.AuthMiddleware, animals.Controller.Create)
	animalsRouter.Put("/:id", users.AuthMiddleware, animals.Controller.Update)
	animalsRouter.Delete("/:id", users.AuthMiddleware, animals.Controller.Remove)

	animalsTypesRouter := animalsRouter.Group(":id/types")

	animalsTypesRouter.Post("/:typeId", users.AuthMiddleware, animals.Controller.AddType)
	animalsTypesRouter.Put("/", users.AuthMiddleware, animals.Controller.ReplaceType)
	animalsTypesRouter.Delete("/:typeId", users.AuthMiddleware, animals.Controller.RemoveType)

	animalTypesRouter := animalsRouter.Group("types")

	animalTypesRouter.Get("/:id", users.DisableAuthCheck, users.AuthMiddleware, animaltypes.Controller.GetOne)
	animalTypesRouter.Post("/", users.AuthMiddleware, animaltypes.Controller.Create)
	animalTypesRouter.Put("/:id", users.AuthMiddleware, animaltypes.Controller.Update)
	animalTypesRouter.Delete("/:id", users.AuthMiddleware, animaltypes.Controller.Remove)

	return router
}

func main() {
	router := setupRouter()

	router.Get("/ping", func(ct *fiber.Ctx) error {
		ct.Status(200).SendString("pong!")

		return nil
	})

	router.Listen(":5000")
}
