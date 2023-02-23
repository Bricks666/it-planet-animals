package main

import (
	animaltypes "animals/animal-types"
	"animals/animals"
	animalslocations "animals/animals-locations"
	"animals/locations"
	"animals/users"

	"github.com/gofiber/fiber/v2"
)

func main() {
	router := setupRouter()

	router.Listen(":5000")
}

func setupRouter() *fiber.App {
	router := fiber.New()

	router.Post("/registration", users.CheckUnauthorized, users.Controller.Create)

	usersRouter := router.Group("accounts")

	usersRouter.Get("/search", users.DisableAuthCheck, users.CheckAuth, users.Controller.GetAll)
	usersRouter.Get("/:id", users.DisableAuthCheck, users.CheckAuth, users.Controller.GetOne)
	usersRouter.Put("/:id", users.CheckAuth, users.Controller.Update)
	usersRouter.Delete("/:id", users.CheckAuth, users.Controller.Remove)

	locationsRouter := router.Group("locations")

	locationsRouter.Get("/:id", users.DisableAuthCheck, users.CheckAuth, locations.Controller.GetOne)
	locationsRouter.Post("/", users.CheckAuth, locations.Controller.Create)
	locationsRouter.Put("/:id", users.CheckAuth, locations.Controller.Update)
	locationsRouter.Delete("/:id", users.CheckAuth, locations.Controller.Remove)

	animalsRouter := router.Group("animals")

	animalsRouter.Get("/search", users.DisableAuthCheck, users.CheckAuth, animals.Controller.GetAll)
	animalsRouter.Get("/:id", users.DisableAuthCheck, users.CheckAuth, animals.Controller.GetOne)
	animalsRouter.Post("/", users.CheckAuth, animals.Controller.Create)
	animalsRouter.Put("/:id", users.CheckAuth, animals.Controller.Update)
	animalsRouter.Delete("/:id", users.CheckAuth, animals.Controller.Remove)

	animalsTypesRouter := animalsRouter.Group(":id/types")

	animalsTypesRouter.Post("/:typeId", users.CheckAuth, animals.Controller.AddType)
	animalsTypesRouter.Put("/", users.CheckAuth, animals.Controller.ReplaceType)
	animalsTypesRouter.Delete("/:typeId", users.CheckAuth, animals.Controller.RemoveType)

	animalTypesRouter := animalsRouter.Group("types")

	animalTypesRouter.Get("/:id", users.DisableAuthCheck, users.CheckAuth, animaltypes.Controller.GetOne)
	animalTypesRouter.Post("/", users.CheckAuth, animaltypes.Controller.Create)
	animalTypesRouter.Put("/:id", users.CheckAuth, animaltypes.Controller.Update)
	animalTypesRouter.Delete("/:id", users.CheckAuth, animaltypes.Controller.Remove)

	animalsLocationsRouter := animalsRouter.Group(":animalId/locations")

	animalsLocationsRouter.Get("/", users.DisableAuthCheck, users.CheckAuth, animalslocations.Controller.GetAll)
	animalsLocationsRouter.Post("/:locationId", users.CheckAuth, animalslocations.Controller.Create)
	animalsLocationsRouter.Put("/", users.CheckAuth, animalslocations.Controller.Update)
	animalsLocationsRouter.Delete("/:locationId", users.CheckAuth, animalslocations.Controller.Remove)

	return router
}
