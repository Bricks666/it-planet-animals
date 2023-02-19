package main

import (
	"fmt"

	"animals/shared"
	"animals/users"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/registration", users.Controller.Create)

	usersRouter := router.Group("accounts")

	usersRouter.GET("/:id", users.Controller.GetOne)
	usersRouter.GET("search", users.Controller.GetAll)

	return router
}

func main() {
	router := setupRouter()

	router.GET("/ping", func(ct *gin.Context) {
		ct.JSON(200, "pong!")
	})

	address := fmt.Sprintf(":%s", shared.ENV["PORT"])

	router.Run(address)

}
