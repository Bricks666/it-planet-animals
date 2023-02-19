package users

import (
	"animals/ent"
	"animals/ent/user"
	"animals/shared"
)

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

func (ur *UserRepository) GetOne(id uint32) (*ent.User, error) {
	return ur.db.Client.User.Query().
		Where(user.ID(id)).
		Select(user.FieldEmail, user.FieldFirstName, user.FieldLastName, user.FieldID).
		Only(ur.db.Context)

}

func (ur *UserRepository) Create(dto *CreateUserDto) (*ent.User, error) {
	return ur.db.Client.User.Create().
		SetEmail(dto.Email).
		SetFirstName(dto.FirstName).
		SetLastName(dto.LastName).
		SetPassword(dto.Password).
		Save(ur.db.Context)

}

func (ur *UserRepository) Update() {}

func (ur *UserRepository) Remove() {}
