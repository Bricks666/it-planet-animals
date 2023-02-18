package users

var UsersService *usersService

func init() {
	UsersService = new(usersService)
}

type usersService struct{}

func (ac *usersService) GetAll() {}

func (ac *usersService) GetOne() {}

func (ac *usersService) Update() {}

func (ac *usersService) Remove() {}
