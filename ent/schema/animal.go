package schema

import "entgo.io/ent"

// Animal holds the schema definition for the Animal entity.
type Animal struct {
	ent.Schema
}

// Fields of the Animal.
func (Animal) Fields() []ent.Field {
	return nil
}

// Edges of the Animal.
func (Animal) Edges() []ent.Edge {
	return nil
}
