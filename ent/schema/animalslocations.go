package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AnimalsLocations holds the schema definition for the AnimalsLocations entity.
type AnimalsLocations struct {
	ent.Schema
}

func (AnimalsLocations) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}

// Fields of the AnimalsLocations.
func (AnimalsLocations) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Positive(),
		field.Uint64("animalId"),
		field.Uint64("locationId"),
		field.Time("dateTimeOfVisitLocationPoint").Default(time.Now),
	}
}

// Edges of the AnimalsLocations.
func (AnimalsLocations) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("animals_locations_animal", Animal.Type).
			Unique().
			Required().
			Field("animalId"),
		edge.To("animals_locations_location", Location.Type).
			Unique().
			Required().
			Field("locationId"),
	}
}
