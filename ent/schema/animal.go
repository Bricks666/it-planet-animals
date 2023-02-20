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
		field.Enum("lifestatus").Values("ALIVE", "DEAD").Default("ALIVE"),
		field.Time("chippingDateTime").Default(time.Now),
		field.Uint32("chipperId").Optional(),
		field.Uint64("chippingLocationId").Optional(),
		field.Time("deathDateTime").Optional().Nillable(),
	}
}

// Edges of the Animal.
func (Animal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_animals", User.Type).Field("chipperId").Unique(),
		edge.To("animal_tags_animals", AnimalType.Type),
		edge.To("chipping_location", Location.Type).Field("chippingLocationId").Unique(),
		edge.To("visited_locations_animals", Location.Type),
	}
}
