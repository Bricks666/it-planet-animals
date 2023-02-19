package main

import (
	"fmt"

	"animals/shared"
	"animals/users"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	usersR := router.Group("users")

	usersR.POST("/", users.Controller.Create)

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
