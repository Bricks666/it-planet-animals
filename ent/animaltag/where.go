// Code generated by ent, DO NOT EDIT.

package animaltag

import (
	"animals/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldLTE(FieldID, id))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldEQ(FieldType, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.AnimalTag {
	return predicate.AnimalTag(sql.FieldContainsFold(FieldType, v))
}

// HasAnimals applies the HasEdge predicate on the "animals" edge.
func HasAnimals() predicate.AnimalTag {
	return predicate.AnimalTag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, AnimalsTable, AnimalsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAnimalsWith applies the HasEdge predicate on the "animals" edge with a given conditions (other predicates).
func HasAnimalsWith(preds ...predicate.Animal) predicate.AnimalTag {
	return predicate.AnimalTag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AnimalsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, AnimalsTable, AnimalsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAnimalTypes applies the HasEdge predicate on the "animal_types" edge.
func HasAnimalTypes() predicate.AnimalTag {
	return predicate.AnimalTag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, AnimalTypesTable, AnimalTypesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAnimalTypesWith applies the HasEdge predicate on the "animal_types" edge with a given conditions (other predicates).
func HasAnimalTypesWith(preds ...predicate.AnimalType) predicate.AnimalTag {
	return predicate.AnimalTag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AnimalTypesInverseTable, AnimalTypesColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, AnimalTypesTable, AnimalTypesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AnimalTag) predicate.AnimalTag {
	return predicate.AnimalTag(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AnimalTag) predicate.AnimalTag {
	return predicate.AnimalTag(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AnimalTag) predicate.AnimalTag {
	return predicate.AnimalTag(func(s *sql.Selector) {
		p(s.Not())
	})
}
