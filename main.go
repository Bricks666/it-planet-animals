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

	usersRouter.Get("/search", users.Controller.GetAll)
	usersRouter.Get("/:id", users.Controller.GetOne)
	usersRouter.Put("/:id", users.Controller.Update)
	usersRouter.Delete("/:id", users.Controller.Remove)

	locationsRouter := router.Group("locations")

	locationsRouter.Get("/:id", locations.Controller.GetOne)
	locationsRouter.Post("/", locations.Controller.Create)
	locationsRouter.Put("/:id", locations.Controller.Update)
	locationsRouter.Delete("/:id", locations.Controller.Remove)

	animalsRouter := router.Group("animals")

	animalsRouter.Get("/search", animals.Controller.GetAll)
	animalsRouter.Get("/:id", animals.Controller.GetOne)
	animalsRouter.Post("/", animals.Controller.Create)
	animalsRouter.Put("/:id", animals.Controller.Update)
	animalsRouter.Delete("/:id", animals.Controller.Remove)

	animalsTypesRouter := animalsRouter.Group(":id/types")

	animalsTypesRouter.Post("/:typeId", animals.Controller.AddType)
	animalsTypesRouter.Put("/", animals.Controller.ReplaceType)
	animalsTypesRouter.Delete("/:typeId", animals.Controller.RemoveType)

	animalTypesRouter := animalsRouter.Group("types")

	animalTypesRouter.Get("/:id", animaltypes.Controller.GetOne)
	animalTypesRouter.Post("/", animaltypes.Controller.Create)
	animalTypesRouter.Put("/:id", animaltypes.Controller.Update)
	animalTypesRouter.Delete("/:id", animaltypes.Controller.Remove)

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
