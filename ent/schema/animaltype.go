package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AnimalType holds the schema definition for the AnimalType entity.
type AnimalType struct {
	ent.Schema
}

// Fields of the AnimalType.
func (AnimalType) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Positive(),
		field.String("type").MinLen(1).Unique(),
	}
}

// Edges of the AnimalType.
func (AnimalType) Edges() []ent.Edge {
	return nil
}
