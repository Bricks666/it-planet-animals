package main

import (
	animaltypes "animals/animal-types"
	"animals/animals"
	"animals/areas"
	"animals/locations"
	"animals/users"
	visitedlocation "animals/visited-locations"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var router = setupRouter()

	writeStartData()

	router.Listen(":5000")
}

func setupRouter() *fiber.App {
	var router = fiber.New()

	router.Post("/registration", users.CheckUnauthorized, users.Controller.Create)

	var authRouter = router.Group("", users.CheckAuth)

	var usersRouter = authRouter.Group("accounts")
	usersRouter.Get("/search", users.CheckRole(users.ADMIN_ROLE), users.Controller.GetAll)
	usersRouter.Get("/:id", users.Controller.GetOne)
	usersRouter.Post("/", users.CheckRole(users.ADMIN_ROLE), users.Controller.Create)
	usersRouter.Put("/:id", users.Controller.Update)
	usersRouter.Delete("/:id", users.Controller.Remove)

	var locationsRouter = authRouter.Group("locations")
	locationsRouter.Get("/:id", locations.Controller.GetOne)
	locationsRouter.Post("/", users.CheckRole(users.ADMIN_ROLE, users.CHIPPER_ROLE), locations.Controller.Create)
	locationsRouter.Put("/:id", users.CheckRole(users.ADMIN_ROLE, users.CHIPPER_ROLE), locations.Controller.Update)
	locationsRouter.Delete("/:id", users.CheckRole(users.ADMIN_ROLE), locations.Controller.Remove)

	var areasRouter = authRouter.Group("areas")
	areasRouter.Get("/:id", areas.Controller.GetOne)
	areasRouter.Post("/", users.CheckRole(users.ADMIN_ROLE), areas.Controller.Create)
	areasRouter.Put("/:id", users.CheckRole(users.ADMIN_ROLE), areas.Controller.Update)
	areasRouter.Delete("/:id", users.CheckRole(users.ADMIN_ROLE), areas.Controller.Remove)

	var animalsRouter = authRouter.Group("animals")
	animalsRouter.Get("/search", animals.Controller.GetAll)
	animalsRouter.Get("/:id", animals.Controller.GetOne)
	animalsRouter.Post("/", users.CheckRole(users.ADMIN_ROLE, users.CHIPPER_ROLE), animals.Controller.Create)
	animalsRouter.Put("/:id", users.CheckRole(users.ADMIN_ROLE, users.CHIPPER_ROLE), animals.Controller.Update)
	animalsRouter.Delete("/:id", users.CheckRole(users.ADMIN_ROLE), animals.Controller.Remove)

	var animalsTypesRouter = animalsRouter.Group(":id/types")
	animalsTypesRouter.Post("/:typeId", users.CheckRole(users.ADMIN_ROLE, users.CHIPPER_ROLE), animals.Controller.AddType)
	animalsTypesRouter.Put("/", users.CheckRole(users.ADMIN_ROLE, users.CHIPPER_ROLE), animals.Controller.ReplaceType)
	animalsTypesRouter.Delete("/:typeId", users.CheckRole(users.ADMIN_ROLE, users.CHIPPER_ROLE), animals.Controller.RemoveType)

	var animalTypesRouter = animalsRouter.Group("types")
	animalTypesRouter.Get("/:id", animaltypes.Controller.GetOne)
	animalTypesRouter.Post("/", users.CheckRole(users.ADMIN_ROLE, users.CHIPPER_ROLE), animaltypes.Controller.Create)
	animalTypesRouter.Put("/:id", users.CheckRole(users.ADMIN_ROLE, users.CHIPPER_ROLE), animaltypes.Controller.Update)
	animalTypesRouter.Delete("/:id", users.CheckRole(users.ADMIN_ROLE), animaltypes.Controller.Remove)

	var animalsLocationsRouter = animalsRouter.Group(":animalId/locations")
	animalsLocationsRouter.Get("/", visitedlocation.Controller.GetAll)
	animalsLocationsRouter.Post("/:locationId", users.CheckRole(users.ADMIN_ROLE, users.CHIPPER_ROLE), visitedlocation.Controller.Create)
	animalsLocationsRouter.Put("/", users.CheckRole(users.ADMIN_ROLE, users.CHIPPER_ROLE), visitedlocation.Controller.Update)
	animalsLocationsRouter.Delete("/:locationId", users.CheckRole(users.ADMIN_ROLE), visitedlocation.Controller.Remove)

	return router
}

func writeStartData() {
	var admin = users.ADMIN_ROLE
	var chipper = users.CHIPPER_ROLE
	var user = users.USER_ROLE

	users.Service.Create(&users.CreateUserDto{
		FirstName: "adminFirstName",
		LastName:  "adminLastName",
		Email:     "admin@simbirsoft.com",
		Password:  "qwerty123",
		Role:      &admin,
	})
	users.Service.Create(&users.CreateUserDto{
		FirstName: "chipperFirstName",
		LastName:  "chipperLastName",
		Email:     "chipper@simbirsoft.com",
		Password:  "qwerty123",
		Role:      &chipper,
	})
	users.Service.Create(&users.CreateUserDto{
		FirstName: "userFirstName",
		LastName:  "userLastName",
		Email:     "user@simbirsoft.com",
		Password:  "qwerty123",
		Role:      &user,
	})
}
