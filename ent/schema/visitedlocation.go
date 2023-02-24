package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// VisitedLocation holds the schema definition for the VisitedLocation entity.
type VisitedLocation struct {
	ent.Schema
}

func (VisitedLocation) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}

// Fields of the AnimalsLocations.
func (VisitedLocation) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Positive(),
		field.Uint64("animal_id"),
		field.Uint64("location_id"),
		field.Time("date_time_of_visit_location_point").
			Optional(),
	}
}

// Edges of the AnimalsLocations.
func (VisitedLocation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("animals", Animal.Type).
			Unique().
			Required().
			Field("animal_id"),
		edge.To("locations", Location.Type).
			Unique().
			Required().
			Field("location_id"),
	}
}
