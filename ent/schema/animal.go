package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Animal holds the schema definition for the Animal entity.
type Animal struct {
	ent.Schema
}

// Fields of the Animal.
func (Animal) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Positive(),
		field.Float32("weight").Positive(),
		field.Float32("length").Positive(),
		field.Float32("height").Positive(),
		field.Enum("gender").Values("MALE", "FEMALE", "OTHER"),
		field.Enum("life_status").Values("ALIVE", "DEAD").Default("ALIVE"),
		field.Time("chipping_date_time").Default(time.Now),
		field.Uint32("chipper_id").Optional(),
		field.Uint64("chipping_location_id").Optional(),
		field.Time("death_date_time").Optional().Nillable(),
	}
}

// Edges of the Animal.
func (Animal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chipper", User.Type).
			Ref("animals").
			Field("chipper_id").
			Unique(),
		edge.To("types", AnimalTag.Type).
			Through("animal_types", AnimalType.Type),
		edge.From("chipping_location", Location.Type).
			Ref("chipped_animals").
			Field("chipping_location_id").
			Unique(),
		edge.To("locations", Location.Type).
			Through("visited_locations", VisitedLocation.Type),
	}
}
