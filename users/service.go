package users

import (
	"animals/ent"
	"errors"

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

func (us *UsersService) GetAll(dto *UsersSearchQueryDto) ([]*SecurityUserDto, error) {
	users, err := us.usersRepository.GetAll(dto)

	if err != nil {
		return nil, err
	}

	var securityUsers = []*SecurityUserDto{}

	for _, user := range users {
		securityUsers = append(securityUsers, prepareSecurityUser(user))
	}

	return securityUsers, err
}

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

func (us *UsersService) Update(id uint32, dto *UpdateUserDto) (*SecurityUserDto, error) {
	isExists, _ := us.usersRepository.HasWithThisEmail(dto.Email)

	if isExists {
		return nil, errors.New("email")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), 8)

	dto.Password = string(hashed)

	user, err := us.usersRepository.Update(id, dto)

	if err != nil {
		return nil, err
	}

	return prepareSecurityUser(user), nil
}

func (us *UsersService) Remove(id uint32) error {
	user, _ := us.usersRepository.GetOne(id)

	if user == nil {
		return errors.New("User doesn't exist")
	}

	return us.usersRepository.Remove(id)
}

func prepareSecurityUser(user *ent.User) *SecurityUserDto {
	return &SecurityUserDto{
		Id:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
