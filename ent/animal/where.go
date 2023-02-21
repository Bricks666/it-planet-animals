// Code generated by ent, DO NOT EDIT.

package animal

import (
	"animals/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Animal {
	return predicate.Animal(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Animal {
	return predicate.Animal(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Animal {
	return predicate.Animal(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Animal {
	return predicate.Animal(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Animal {
	return predicate.Animal(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Animal {
	return predicate.Animal(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Animal {
	return predicate.Animal(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Animal {
	return predicate.Animal(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Animal {
	return predicate.Animal(sql.FieldLTE(FieldID, id))
}

// Weight applies equality check predicate on the "weight" field. It's identical to WeightEQ.
func Weight(v float32) predicate.Animal {
	return predicate.Animal(sql.FieldEQ(FieldWeight, v))
}

// Length applies equality check predicate on the "length" field. It's identical to LengthEQ.
func Length(v float32) predicate.Animal {
	return predicate.Animal(sql.FieldEQ(FieldLength, v))
}

// Height applies equality check predicate on the "height" field. It's identical to HeightEQ.
func Height(v float32) predicate.Animal {
	return predicate.Animal(sql.FieldEQ(FieldHeight, v))
}

// ChippingDateTime applies equality check predicate on the "chippingDateTime" field. It's identical to ChippingDateTimeEQ.
func ChippingDateTime(v time.Time) predicate.Animal {
	return predicate.Animal(sql.FieldEQ(FieldChippingDateTime, v))
}

// ChipperId applies equality check predicate on the "chipperId" field. It's identical to ChipperIdEQ.
func ChipperId(v uint32) predicate.Animal {
	return predicate.Animal(sql.FieldEQ(FieldChipperId, v))
}

// ChippingLocationId applies equality check predicate on the "chippingLocationId" field. It's identical to ChippingLocationIdEQ.
func ChippingLocationId(v uint64) predicate.Animal {
	return predicate.Animal(sql.FieldEQ(FieldChippingLocationId, v))
}

// DeathDateTime applies equality check predicate on the "deathDateTime" field. It's identical to DeathDateTimeEQ.
func DeathDateTime(v time.Time) predicate.Animal {
	return predicate.Animal(sql.FieldEQ(FieldDeathDateTime, v))
}

// WeightEQ applies the EQ predicate on the "weight" field.
func WeightEQ(v float32) predicate.Animal {
	return predicate.Animal(sql.FieldEQ(FieldWeight, v))
}

// WeightNEQ applies the NEQ predicate on the "weight" field.
func WeightNEQ(v float32) predicate.Animal {
	return predicate.Animal(sql.FieldNEQ(FieldWeight, v))
}

// WeightIn applies the In predicate on the "weight" field.
func WeightIn(vs ...float32) predicate.Animal {
	return predicate.Animal(sql.FieldIn(FieldWeight, vs...))
}

// WeightNotIn applies the NotIn predicate on the "weight" field.
func WeightNotIn(vs ...float32) predicate.Animal {
	return predicate.Animal(sql.FieldNotIn(FieldWeight, vs...))
}

// WeightGT applies the GT predicate on the "weight" field.
func WeightGT(v float32) predicate.Animal {
	return predicate.Animal(sql.FieldGT(FieldWeight, v))
}

// WeightGTE applies the GTE predicate on the "weight" field.
func WeightGTE(v float32) predicate.Animal {
	return predicate.Animal(sql.FieldGTE(FieldWeight, v))
}

// WeightLT applies the LT predicate on the "weight" field.
func WeightLT(v float32) predicate.Animal {
	return predicate.Animal(sql.FieldLT(FieldWeight, v))
}

// WeightLTE applies the LTE predicate on the "weight" field.
func WeightLTE(v float32) predicate.Animal {
	return predicate.Animal(sql.FieldLTE(FieldWeight, v))
}

// LengthEQ applies the EQ predicate on the "length" field.
func LengthEQ(v float32) predicate.Animal {
	return predicate.Animal(sql.FieldEQ(FieldLength, v))
}

// LengthNEQ applies the NEQ predicate on the "length" field.
func LengthNEQ(v float32) predicate.Animal {
	return predicate.Animal(sql.FieldNEQ(FieldLength, v))
}

// LengthIn applies the In predicate on the "length" field.
func LengthIn(vs ...float32) predicate.Animal {
	return predicate.Animal(sql.FieldIn(FieldLength, vs...))
}

// LengthNotIn applies the NotIn predicate on the "length" field.
func LengthNotIn(vs ...float32) predicate.Animal {
	return predicate.Animal(sql.FieldNotIn(FieldLength, vs...))
}

// LengthGT applies the GT predicate on the "length" field.
func LengthGT(v float32) predicate.Animal {
	return predicate.Animal(sql.FieldGT(FieldLength, v))
}

// LengthGTE applies the GTE predicate on the "length" field.
func LengthGTE(v float32) predicate.Animal {
	return predicate.Animal(sql.FieldGTE(FieldLength, v))
}

// LengthLT applies the LT predicate on the "length" field.
func LengthLT(v float32) predicate.Animal {
	return predicate.Animal(sql.FieldLT(FieldLength, v))
}

// LengthLTE applies the LTE predicate on the "length" field.
func LengthLTE(v float32) predicate.Animal {
	return predicate.Animal(sql.FieldLTE(FieldLength, v))
}

// HeightEQ applies the EQ predicate on the "height" field.
func HeightEQ(v float32) predicate.Animal {
	return predicate.Animal(sql.FieldEQ(FieldHeight, v))
}

// HeightNEQ applies the NEQ predicate on the "height" field.
func HeightNEQ(v float32) predicate.Animal {
	return predicate.Animal(sql.FieldNEQ(FieldHeight, v))
}

// HeightIn applies the In predicate on the "height" field.
func HeightIn(vs ...float32) predicate.Animal {
	return predicate.Animal(sql.FieldIn(FieldHeight, vs...))
}

// HeightNotIn applies the NotIn predicate on the "height" field.
func HeightNotIn(vs ...float32) predicate.Animal {
	return predicate.Animal(sql.FieldNotIn(FieldHeight, vs...))
}

// HeightGT applies the GT predicate on the "height" field.
func HeightGT(v float32) predicate.Animal {
	return predicate.Animal(sql.FieldGT(FieldHeight, v))
}

// HeightGTE applies the GTE predicate on the "height" field.
func HeightGTE(v float32) predicate.Animal {
	return predicate.Animal(sql.FieldGTE(FieldHeight, v))
}

// HeightLT applies the LT predicate on the "height" field.
func HeightLT(v float32) predicate.Animal {
	return predicate.Animal(sql.FieldLT(FieldHeight, v))
}

// HeightLTE applies the LTE predicate on the "height" field.
func HeightLTE(v float32) predicate.Animal {
	return predicate.Animal(sql.FieldLTE(FieldHeight, v))
}

// GenderEQ applies the EQ predicate on the "gender" field.
func GenderEQ(v Gender) predicate.Animal {
	return predicate.Animal(sql.FieldEQ(FieldGender, v))
}

// GenderNEQ applies the NEQ predicate on the "gender" field.
func GenderNEQ(v Gender) predicate.Animal {
	return predicate.Animal(sql.FieldNEQ(FieldGender, v))
}

// GenderIn applies the In predicate on the "gender" field.
func GenderIn(vs ...Gender) predicate.Animal {
	return predicate.Animal(sql.FieldIn(FieldGender, vs...))
}

// GenderNotIn applies the NotIn predicate on the "gender" field.
func GenderNotIn(vs ...Gender) predicate.Animal {
	return predicate.Animal(sql.FieldNotIn(FieldGender, vs...))
}

// LifestatusEQ applies the EQ predicate on the "lifestatus" field.
func LifestatusEQ(v Lifestatus) predicate.Animal {
	return predicate.Animal(sql.FieldEQ(FieldLifestatus, v))
}

// LifestatusNEQ applies the NEQ predicate on the "lifestatus" field.
func LifestatusNEQ(v Lifestatus) predicate.Animal {
	return predicate.Animal(sql.FieldNEQ(FieldLifestatus, v))
}

// LifestatusIn applies the In predicate on the "lifestatus" field.
func LifestatusIn(vs ...Lifestatus) predicate.Animal {
	return predicate.Animal(sql.FieldIn(FieldLifestatus, vs...))
}

// LifestatusNotIn applies the NotIn predicate on the "lifestatus" field.
func LifestatusNotIn(vs ...Lifestatus) predicate.Animal {
	return predicate.Animal(sql.FieldNotIn(FieldLifestatus, vs...))
}

// ChippingDateTimeEQ applies the EQ predicate on the "chippingDateTime" field.
func ChippingDateTimeEQ(v time.Time) predicate.Animal {
	return predicate.Animal(sql.FieldEQ(FieldChippingDateTime, v))
}

// ChippingDateTimeNEQ applies the NEQ predicate on the "chippingDateTime" field.
func ChippingDateTimeNEQ(v time.Time) predicate.Animal {
	return predicate.Animal(sql.FieldNEQ(FieldChippingDateTime, v))
}

// ChippingDateTimeIn applies the In predicate on the "chippingDateTime" field.
func ChippingDateTimeIn(vs ...time.Time) predicate.Animal {
	return predicate.Animal(sql.FieldIn(FieldChippingDateTime, vs...))
}

// ChippingDateTimeNotIn applies the NotIn predicate on the "chippingDateTime" field.
func ChippingDateTimeNotIn(vs ...time.Time) predicate.Animal {
	return predicate.Animal(sql.FieldNotIn(FieldChippingDateTime, vs...))
}

// ChippingDateTimeGT applies the GT predicate on the "chippingDateTime" field.
func ChippingDateTimeGT(v time.Time) predicate.Animal {
	return predicate.Animal(sql.FieldGT(FieldChippingDateTime, v))
}

// ChippingDateTimeGTE applies the GTE predicate on the "chippingDateTime" field.
func ChippingDateTimeGTE(v time.Time) predicate.Animal {
	return predicate.Animal(sql.FieldGTE(FieldChippingDateTime, v))
}

// ChippingDateTimeLT applies the LT predicate on the "chippingDateTime" field.
func ChippingDateTimeLT(v time.Time) predicate.Animal {
	return predicate.Animal(sql.FieldLT(FieldChippingDateTime, v))
}

// ChippingDateTimeLTE applies the LTE predicate on the "chippingDateTime" field.
func ChippingDateTimeLTE(v time.Time) predicate.Animal {
	return predicate.Animal(sql.FieldLTE(FieldChippingDateTime, v))
}

// ChipperIdEQ applies the EQ predicate on the "chipperId" field.
func ChipperIdEQ(v uint32) predicate.Animal {
	return predicate.Animal(sql.FieldEQ(FieldChipperId, v))
}

// ChipperIdNEQ applies the NEQ predicate on the "chipperId" field.
func ChipperIdNEQ(v uint32) predicate.Animal {
	return predicate.Animal(sql.FieldNEQ(FieldChipperId, v))
}

// ChipperIdIn applies the In predicate on the "chipperId" field.
func ChipperIdIn(vs ...uint32) predicate.Animal {
	return predicate.Animal(sql.FieldIn(FieldChipperId, vs...))
}

// ChipperIdNotIn applies the NotIn predicate on the "chipperId" field.
func ChipperIdNotIn(vs ...uint32) predicate.Animal {
	return predicate.Animal(sql.FieldNotIn(FieldChipperId, vs...))
}

// ChipperIdIsNil applies the IsNil predicate on the "chipperId" field.
func ChipperIdIsNil() predicate.Animal {
	return predicate.Animal(sql.FieldIsNull(FieldChipperId))
}

// ChipperIdNotNil applies the NotNil predicate on the "chipperId" field.
func ChipperIdNotNil() predicate.Animal {
	return predicate.Animal(sql.FieldNotNull(FieldChipperId))
}

// ChippingLocationIdEQ applies the EQ predicate on the "chippingLocationId" field.
func ChippingLocationIdEQ(v uint64) predicate.Animal {
	return predicate.Animal(sql.FieldEQ(FieldChippingLocationId, v))
}

// ChippingLocationIdNEQ applies the NEQ predicate on the "chippingLocationId" field.
func ChippingLocationIdNEQ(v uint64) predicate.Animal {
	return predicate.Animal(sql.FieldNEQ(FieldChippingLocationId, v))
}

// ChippingLocationIdIn applies the In predicate on the "chippingLocationId" field.
func ChippingLocationIdIn(vs ...uint64) predicate.Animal {
	return predicate.Animal(sql.FieldIn(FieldChippingLocationId, vs...))
}

// ChippingLocationIdNotIn applies the NotIn predicate on the "chippingLocationId" field.
func ChippingLocationIdNotIn(vs ...uint64) predicate.Animal {
	return predicate.Animal(sql.FieldNotIn(FieldChippingLocationId, vs...))
}

// ChippingLocationIdIsNil applies the IsNil predicate on the "chippingLocationId" field.
func ChippingLocationIdIsNil() predicate.Animal {
	return predicate.Animal(sql.FieldIsNull(FieldChippingLocationId))
}

// ChippingLocationIdNotNil applies the NotNil predicate on the "chippingLocationId" field.
func ChippingLocationIdNotNil() predicate.Animal {
	return predicate.Animal(sql.FieldNotNull(FieldChippingLocationId))
}

// DeathDateTimeEQ applies the EQ predicate on the "deathDateTime" field.
func DeathDateTimeEQ(v time.Time) predicate.Animal {
	return predicate.Animal(sql.FieldEQ(FieldDeathDateTime, v))
}

// DeathDateTimeNEQ applies the NEQ predicate on the "deathDateTime" field.
func DeathDateTimeNEQ(v time.Time) predicate.Animal {
	return predicate.Animal(sql.FieldNEQ(FieldDeathDateTime, v))
}

// DeathDateTimeIn applies the In predicate on the "deathDateTime" field.
func DeathDateTimeIn(vs ...time.Time) predicate.Animal {
	return predicate.Animal(sql.FieldIn(FieldDeathDateTime, vs...))
}

// DeathDateTimeNotIn applies the NotIn predicate on the "deathDateTime" field.
func DeathDateTimeNotIn(vs ...time.Time) predicate.Animal {
	return predicate.Animal(sql.FieldNotIn(FieldDeathDateTime, vs...))
}

// DeathDateTimeGT applies the GT predicate on the "deathDateTime" field.
func DeathDateTimeGT(v time.Time) predicate.Animal {
	return predicate.Animal(sql.FieldGT(FieldDeathDateTime, v))
}

// DeathDateTimeGTE applies the GTE predicate on the "deathDateTime" field.
func DeathDateTimeGTE(v time.Time) predicate.Animal {
	return predicate.Animal(sql.FieldGTE(FieldDeathDateTime, v))
}

// DeathDateTimeLT applies the LT predicate on the "deathDateTime" field.
func DeathDateTimeLT(v time.Time) predicate.Animal {
	return predicate.Animal(sql.FieldLT(FieldDeathDateTime, v))
}

// DeathDateTimeLTE applies the LTE predicate on the "deathDateTime" field.
func DeathDateTimeLTE(v time.Time) predicate.Animal {
	return predicate.Animal(sql.FieldLTE(FieldDeathDateTime, v))
}

// DeathDateTimeIsNil applies the IsNil predicate on the "deathDateTime" field.
func DeathDateTimeIsNil() predicate.Animal {
	return predicate.Animal(sql.FieldIsNull(FieldDeathDateTime))
}

// DeathDateTimeNotNil applies the NotNil predicate on the "deathDateTime" field.
func DeathDateTimeNotNil() predicate.Animal {
	return predicate.Animal(sql.FieldNotNull(FieldDeathDateTime))
}

// HasChipperAnimal applies the HasEdge predicate on the "chipper_animal" edge.
func HasChipperAnimal() predicate.Animal {
	return predicate.Animal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ChipperAnimalTable, ChipperAnimalColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChipperAnimalWith applies the HasEdge predicate on the "chipper_animal" edge with a given conditions (other predicates).
func HasChipperAnimalWith(preds ...predicate.User) predicate.Animal {
	return predicate.Animal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChipperAnimalInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ChipperAnimalTable, ChipperAnimalColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAnimalTypeAnimal applies the HasEdge predicate on the "animal_type_animal" edge.
func HasAnimalTypeAnimal() predicate.Animal {
	return predicate.Animal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, AnimalTypeAnimalTable, AnimalTypeAnimalPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAnimalTypeAnimalWith applies the HasEdge predicate on the "animal_type_animal" edge with a given conditions (other predicates).
func HasAnimalTypeAnimalWith(preds ...predicate.AnimalType) predicate.Animal {
	return predicate.Animal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AnimalTypeAnimalInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, AnimalTypeAnimalTable, AnimalTypeAnimalPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChippingLocation applies the HasEdge predicate on the "chipping_location" edge.
func HasChippingLocation() predicate.Animal {
	return predicate.Animal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ChippingLocationTable, ChippingLocationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChippingLocationWith applies the HasEdge predicate on the "chipping_location" edge with a given conditions (other predicates).
func HasChippingLocationWith(preds ...predicate.Location) predicate.Animal {
	return predicate.Animal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChippingLocationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ChippingLocationTable, ChippingLocationColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVisitedLocations applies the HasEdge predicate on the "visited_locations" edge.
func HasVisitedLocations() predicate.Animal {
	return predicate.Animal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, VisitedLocationsTable, VisitedLocationsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVisitedLocationsWith applies the HasEdge predicate on the "visited_locations" edge with a given conditions (other predicates).
func HasVisitedLocationsWith(preds ...predicate.Location) predicate.Animal {
	return predicate.Animal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VisitedLocationsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, VisitedLocationsTable, VisitedLocationsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAnimals applies the HasEdge predicate on the "animals" edge.
func HasAnimals() predicate.Animal {
	return predicate.Animal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, AnimalsTable, AnimalsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAnimalsWith applies the HasEdge predicate on the "animals" edge with a given conditions (other predicates).
func HasAnimalsWith(preds ...predicate.AnimalsLocations) predicate.Animal {
	return predicate.Animal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AnimalsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, AnimalsTable, AnimalsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Animal) predicate.Animal {
	return predicate.Animal(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Animal) predicate.Animal {
	return predicate.Animal(func(s *sql.Selector) {
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
func Not(p predicate.Animal) predicate.Animal {
	return predicate.Animal(func(s *sql.Selector) {
		p(s.Not())
	})
}
