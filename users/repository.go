package users

import (
	"animals/ent"
	"animals/ent/user"
	"animals/shared"
)

type UserRepository struct {
	db *shared.DB
}

var Repository UserRepository

func NewUserRepository(db *shared.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func init() {
	Repository = *NewUserRepository(&shared.Database)
}

func (this *UserRepository) GetAll(dto *UsersSearchQueryDto) ([]*ent.User, error) {
	return this.db.Client.User.
		Query().
		Where(
			user.FirstNameContainsFold(dto.FirstName),
			user.LastNameContainsFold(dto.LastName),
			user.EmailContainsFold(dto.Email),
		).
		Order(ent.Asc(user.FieldID)).
		Offset(int(dto.From)).
		Limit(int(dto.Size)).
		All(this.db.Context)
}

func (this *UserRepository) GetOne(id uint32) (*ent.User, error) {
	return this.db.Client.User.
		Query().
		Where(user.ID(id)).
		Select(user.FieldEmail, user.FieldFirstName, user.FieldLastName, user.FieldID, user.FieldRole).
		Only(this.db.Context)

}

func (this *UserRepository) GetByEmail(email string) (*ent.User, error) {
	return this.db.Client.User.
		Query().
		Where(user.Email(email)).
		Only(this.db.Context)
}

func (this *UserRepository) Create(dto *CreateUserDto) (*ent.User, error) {
	var query = this.db.Client.User.Create()

	if dto.Role != nil {
		query.SetRole(user.Role(*dto.Role))
	} else {
		query.SetRole(user.Role(USER_ROLE))
	}

	return query.
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
		SetRole(user.Role(dto.Role)).
		Save(this.db.Context)
}

func (this *UserRepository) Remove(id uint32) error {
	return this.db.Client.User.
		DeleteOneID(id).
		Exec(this.db.Context)
}
