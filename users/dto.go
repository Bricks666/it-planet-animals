package users

type CreateUserDto struct {
	FirstName string `json:"firstname" validate:"required,min=1"`
	LastName  string `json:"lastname" validate:"required,min=1"`
	Email     string `json:"email" validate:"required,email,min=1"`
	Password  string `json:"password" validate:"required,min=1"`
}

type UpdateUserDto struct {
	FirstName string `json:"firstname" validate:"required,min=1"`
	LastName  string `json:"lastname" validate:"required,min=1"`
	Email     string `json:"email" validate:"required,email,min=1"`
	Password  string `json:"password" validate:"required,min=1"`
}

type SecurityUserDto struct {
	Id        uint32 `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

type UserParamsDto struct {
	Id uint32 `params:"id" validate:"required,number,gt=0"`
}

type UsersSearchQueryDto struct {
	FirstName string `query:"firstname"`
	LastName  string `query:"lastname"`
	Email     string `query:"email"`
	From      uint32 `query:"from" validate:"number,min=0"`
	Size      uint32 `query:"size" validate:"number,min=1"`
}
