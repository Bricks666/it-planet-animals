package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").Positive(),
		field.String("email").Unique().MinLen(1),
		field.String("password").MinLen(1),
		field.String("firstName").MinLen(1),
		field.String("lastName").MinLen(1),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
