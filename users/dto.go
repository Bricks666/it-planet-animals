package users

import "animals/shared"

type CreateUserDto struct {
	FirstName string `json:"firstname" validate:"required,notblank,min=1"`
	LastName  string `json:"lastname" validate:"required,notblank,min=1"`
	Email     string `json:"email" validate:"required,email,notblank,min=1"`
	Password  string `json:"password" validate:"required,notblank,min=1"`
}

type UpdateUserDto struct {
	FirstName string `json:"firstname" validate:"required,notblank,min=1"`
	LastName  string `json:"lastname" validate:"required,notblank,min=1"`
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
	FirstName string `query:"firstname" validate:"omitempty,notblank"`
	LastName  string `query:"lastname" validate:"omitempty,notblank"`
	Email     string `query:"email" validate:"omitempty,notblank"`
}
