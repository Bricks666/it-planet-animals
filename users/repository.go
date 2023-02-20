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

func (this *UserRepository) GetAll(dto *UsersSearchQueryDto) ([]*ent.User, error) {
	return this.db.Client.User.
		Query().
		Where(
			user.FirstNameContains(dto.FirstName),
			user.LastNameContains(dto.LastName),
			user.EmailContains(dto.Email),
		).
		Offset(int(dto.From)).
		Limit(int(dto.Size)).
		All(this.db.Context)
}

func (this *UserRepository) GetOne(id uint32) (*ent.User, error) {
	return this.db.Client.User.
		Query().
		Where(user.ID(id)).
		Select(user.FieldEmail, user.FieldFirstName, user.FieldLastName, user.FieldID).
		Only(this.db.Context)

}

func (this *UserRepository) GetByEmail(email string) (*ent.User, error) {
	return this.db.Client.User.
		Query().
		Where(user.Email(email)).
		Only(this.db.Context)
}

func (this *UserRepository) Create(dto *CreateUserDto) (*ent.User, error) {
	return this.db.Client.User.
		Create().
		SetEmail(dto.Email).
		SetFirstName(dto.FirstName).
		SetLastName(dto.LastName).
		SetPassword(dto.Password).
		Save(this.db.Context)

}

func (this *UserRepository) Update(id uint32, dto *UpdateUserDto) (*ent.User, error) {
	return this.db.Client.User.
		UpdateOneID(id).
		SetEmail(dto.Email).
		SetFirstName(dto.FirstName).
		SetLastName(dto.LastName).
		SetPassword(dto.Password).
		Save(this.db.Context)
}

func (this *UserRepository) Remove(id uint32) error {
	return this.db.Client.User.
		DeleteOneID(id).
		Exec(this.db.Context)
}
