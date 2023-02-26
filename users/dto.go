package users

import (
	"animals/shared"
)

type CreateUserDto struct {
	FirstName string `json:"firstName" validate:"required,notblank,min=1"`
	LastName  string `json:"lastName" validate:"required,notblank,min=1"`
	Email     string `json:"email" validate:"required,email,notblank,min=1"`
	Password  string `json:"password" validate:"required,notblank,min=1"`
}

type UpdateUserDto struct {
	FirstName string `json:"firstName" validate:"required,notblank,min=1"`
	LastName  string `json:"lastName" validate:"required,notblank,min=1"`
	Email     string `json:"email" validate:"required,email,notblank,min=1"`
	Password  string `json:"password" validate:"required,notblank,min=1"`
}

type SecurityUserDto struct {
	Id        uint32 `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type UserParamsDto struct {
	Id uint32 `params:"id" validate:"required,number,gt=0"`
}

type UsersSearchQueryDto struct {
	shared.PaginationDto
	shared.TimeSearchQueryDto
	FirstName string `query:"firstName" validate:"omitempty,notblank"`
	LastName  string `query:"lastName" validate:"omitempty,notblank"`
	Email     string `query:"email" validate:"omitempty,notblank"`
}

var NewCreateUserDto = shared.BaseNew[CreateUserDto]()

var NewUpdateUserDto = shared.BaseNew[UpdateUserDto]()

var NewSecurityUserDto = shared.BaseNew[SecurityUserDto]()

var NewUserParamsDto = shared.BaseNew[UserParamsDto]()

func NewUsersSearchQueryDto() *UsersSearchQueryDto {
	return &UsersSearchQueryDto{
		PaginationDto:      *shared.NewPaginationDto(),
		TimeSearchQueryDto: *shared.NewTimeSearchQueryDto(),
	}
}
