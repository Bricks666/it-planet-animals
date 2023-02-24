package users

import (
	"animals/ent"
	"animals/shared"
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

func (this *UsersService) GetAll(dto *UsersSearchQueryDto) ([]*SecurityUserDto, error) {
	users, err := this.usersRepository.GetAll(dto)

	if err != nil {
		return nil, err
	}

	var securityUsers = []*SecurityUserDto{}

	for _, user := range users {
		securityUsers = append(securityUsers, prepareSecurityUser(user))
	}

	return securityUsers, err
}

func (this *UsersService) GetOne(id uint32) (*SecurityUserDto, error) {
	user, err := this.usersRepository.GetOne(id)

	if err != nil {
		return nil, err
	}

	return prepareSecurityUser(user), nil
}

func (this *UsersService) Create(dto *CreateUserDto) (*SecurityUserDto, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), shared.ROUND_COUNT)

	// May be should replace this mutation
	dto.Password = string(hashed)

	user, err := this.usersRepository.Create(dto)

	if err != nil {
		return nil, err
	}

	return prepareSecurityUser(user), nil
}

func (this *UsersService) Update(id uint32, dto *UpdateUserDto) (*SecurityUserDto, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), shared.ROUND_COUNT)

	dto.Password = string(hashed)

	user, err := this.usersRepository.Update(id, dto)

	if err != nil {
		return nil, err
	}

	return prepareSecurityUser(user), nil
}

func (this *UsersService) Remove(id uint32) error {
	user, _ := this.usersRepository.GetOne(id)

	if user == nil {
		return errors.New("User doesn't exist")
	}

	return this.usersRepository.Remove(id)
}

func (this *UsersService) VerifyUser(email string, password string) (*SecurityUserDto, error) {
	user, err := this.usersRepository.GetByEmail(email)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return prepareSecurityUser(user), nil
}

func prepareSecurityUser(user *ent.User) *SecurityUserDto {
	return &SecurityUserDto{
		Id:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
