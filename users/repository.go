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

func (ur *UserRepository) GetAll(dto *UsersSearchQueryDto) ([]*ent.User, error) {
	return ur.db.Client.User.Query().
		Where(
			user.FirstNameContains(dto.FirstName),
			user.LastNameContains(dto.LastName),
			user.EmailContains(dto.Email),
		).
		Offset(int(dto.From)).
		Limit(int(dto.Size)).
		All(ur.db.Context)
}

func (ur *UserRepository) GetOne(id uint32) (*ent.User, error) {
	return ur.db.Client.User.Query().
		Where(user.ID(id)).
		Select(user.FieldEmail, user.FieldFirstName, user.FieldLastName, user.FieldID).
		Only(ur.db.Context)

}

func (ur *UserRepository) HasWithThisEmail(email string) (bool, error) {
	return ur.db.Client.User.Query().Where(user.Email(email)).Exist(ur.db.Context)
}

func (ur *UserRepository) Create(dto *CreateUserDto) (*ent.User, error) {
	return ur.db.Client.User.Create().
		SetEmail(dto.Email).
		SetFirstName(dto.FirstName).
		SetLastName(dto.LastName).
		SetPassword(dto.Password).
		Save(ur.db.Context)

}

func (ur *UserRepository) Update(id uint32, dto *UpdateUserDto) (*ent.User, error) {
	return ur.db.Client.User.
		UpdateOneID(id).
		SetEmail(dto.Email).
		SetFirstName(dto.FirstName).
		SetLastName(dto.LastName).
		SetPassword(dto.Password).
		Save(ur.db.Context)
}

func (ur *UserRepository) Remove(id uint32) error {
	return ur.db.Client.User.DeleteOneID(id).Exec(ur.db.Context)
}
