package users

import "log"

var Service UsersService

func init() {
	Service = UsersService{
		usersRepository: &Repository,
	}
}

type UsersService struct {
	usersRepository *UserRepository
}

func (ac *UsersService) GetAll() {}

func (ac *UsersService) GetOne() {}

func (ac *UsersService) Create() {
	log.Println(ac)
	ac.usersRepository.Create()
}

func (ac *UsersService) Update() {}

func (ac *UsersService) Remove() {}
