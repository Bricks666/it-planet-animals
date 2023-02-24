package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AnimalTag holds the schema definition for the AnimalTag entity.
type AnimalTag struct {
	ent.Schema
}

// Fields of the AnimalType.
func (AnimalTag) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Positive(),
		field.String("type").MinLen(1).Unique(),
	}
}

// Edges of the AnimalType.
func (AnimalTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("animals", Animal.Type).
			Ref("types").
			Through("animal_types", AnimalType.Type),
	}
}
