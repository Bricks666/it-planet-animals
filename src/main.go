package main

import (
	"fmt"

	"animals.com/src/shared"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

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
