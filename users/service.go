package users

import (
	"animals/ent"

	"golang.org/x/crypto/bcrypt"
)

var Service UsersService

func init() {
	Service = UsersService{
		usersRepository: &Repository,
	}
}

type UsersService struct {
	usersRepository *UserRepository
}

func (us *UsersService) GetAll() {}

func (us *UsersService) GetOne(id uint32) (*SecurityUserDto, error) {
	user, err := us.usersRepository.GetOne(id)
	if err != nil {
		return nil, err
	}

	return prepareSecurityUser(user), nil
}

func (us *UsersService) Create(dto *CreateUserDto) (*SecurityUserDto, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), 8)

	// May be should replace this mutation
	dto.Password = string(hashed)

	user, err := us.usersRepository.Create(dto)

	if err != nil {
		return nil, err
	}

	return prepareSecurityUser(user), nil
}

func (us *UsersService) Update() {}

func (us *UsersService) Remove() {}

func prepareSecurityUser(user *ent.User) *SecurityUserDto {
	return &SecurityUserDto{
		Id:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
