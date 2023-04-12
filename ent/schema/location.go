package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Location holds the schema definition for the Location entity.
type Location struct {
	ent.Schema
}

// Fields of the Location.
func (Location) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Positive(),
		field.Float("latitude").Range(-90, 90),
		field.Float("longitude").Range(-180, 180),
	}
}

// Edges of the Location.
func (Location) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chipped_animals", Animal.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.NoAction,
			}),
		edge.From("animals", Animal.Type).
			Ref("locations").
			Annotations(entsql.Annotation{
				OnDelete: entsql.NoAction,
			}).
			Through("having_animals", VisitedLocation.Type),
		edge.From("areas", Area.Type).
			Ref("points"),
	}
}

func (Location) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("latitude", "longitude").
			Unique(),
	}
}
