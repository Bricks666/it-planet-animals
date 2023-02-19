package users

import (
	"animals/shared"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Controller UsersController

func init() {
	Controller = UsersController{
		usersService: &Service,
	}
}

type UsersController struct {
	usersService *UsersService
}

func (uc *UsersController) GetAll(ct *gin.Context) {}

func (uc *UsersController) GetOne(ct *gin.Context) {
	var stringId = ct.Params.ByName("id")
	id, err := strconv.ParseUint(stringId, 10, 0)

	if err != nil || id == 0 {
		ct.JSON(http.StatusBadRequest, "")
		return
	}

	user, err := uc.usersService.GetOne(uint32(id))

	if err != nil {
		ct.JSON(http.StatusNotFound, err)
	}

	ct.JSON(http.StatusOK, user)

}

func (uc *UsersController) Create(ct *gin.Context) {
	var dto CreateUserDto

	err := ct.ShouldBindJSON(&dto)

	if err != nil {
		ct.JSON(http.StatusBadRequest, err)
		return
	}

	if shared.IsEmpty(dto.Email) ||
		shared.IsEmpty(dto.FirstName) ||
		shared.IsEmpty(dto.LastName) ||
		shared.IsEmpty(dto.Password) {
		ct.JSON(http.StatusBadRequest, "Some field is empty")
		return
	}

	var user *SecurityUserDto

	user, err = uc.usersService.Create(&dto)

	if err != nil {
		ct.JSON(http.StatusConflict, "")
		return
	}

	ct.JSON(http.StatusCreated, user)
}

func (uc *UsersController) Update(ct *gin.Context) {}

func (uc *UsersController) Remove(ct *gin.Context) {}
