package schema

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AnimalType holds the schema definition for the AnimalTag entity.
type AnimalType struct {
	ent.Schema
}

func (AnimalType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("animal_id", "type_id"),
	}
}

// Fields of the AnimalType.
func (AnimalType) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("animal_id"),
		field.Uint64("type_id"),
	}
}

// Edges of the AnimalType.
func (AnimalType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("animals", Animal.Type).
			Unique().
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}).
			Required().
			Field("animal_id"),
		edge.To("types", AnimalTag.Type).
			Unique().
			Required().
			Annotations(entsql.Annotation{
				OnDelete: entsql.NoAction,
			}).
			Field("type_id"),
	}
}
