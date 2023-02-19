package users

import "github.com/gin-gonic/gin"

var Controller UsersController

func init() {
	Controller = UsersController{
		usersService: &Service,
	}

}

type UsersController struct {
	usersService *UsersService
}

func (ac *UsersController) GetAll() {}

func (ac *UsersController) GetOne() {}

func (ac *UsersController) Create(ct *gin.Context) {
	ac.usersService.Create()

	ct.JSON(200, "Created")
}

func (ac *UsersController) Update() {}

func (ac *UsersController) Remove() {}
