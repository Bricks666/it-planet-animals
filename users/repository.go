package users

import "animals/shared"

var Repository UserRepository

type UserRepository struct {
	db *shared.DB
}

func init() {
	Repository = UserRepository{
		db: &shared.Database,
	}
}

func (ur *UserRepository) GetAll() {}

func (ur *UserRepository) GetOne() {}

func (ur *UserRepository) Create() {
	ur.db.Client.User.Create().SetEmail("email@gmail.com").SetFirstName("First Name").SetLastName("LastName").SetPassword("Password").Save(ur.db.Context)
}

func (ur *UserRepository) Update() {}

func (ur *UserRepository) Remove() {}
