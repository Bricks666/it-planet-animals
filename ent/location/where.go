// Code generated by ent, DO NOT EDIT.

package location

import (
	"animals/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Location {
	return predicate.Location(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Location {
	return predicate.Location(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Location {
	return predicate.Location(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Location {
	return predicate.Location(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Location {
	return predicate.Location(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Location {
	return predicate.Location(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Location {
	return predicate.Location(sql.FieldLTE(FieldID, id))
}

// Latitude applies equality check predicate on the "latitude" field. It's identical to LatitudeEQ.
func Latitude(v float64) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldLatitude, v))
}

// Longitude applies equality check predicate on the "longitude" field. It's identical to LongitudeEQ.
func Longitude(v float64) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldLongitude, v))
}

// LatitudeEQ applies the EQ predicate on the "latitude" field.
func LatitudeEQ(v float64) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldLatitude, v))
}

// LatitudeNEQ applies the NEQ predicate on the "latitude" field.
func LatitudeNEQ(v float64) predicate.Location {
	return predicate.Location(sql.FieldNEQ(FieldLatitude, v))
}

// LatitudeIn applies the In predicate on the "latitude" field.
func LatitudeIn(vs ...float64) predicate.Location {
	return predicate.Location(sql.FieldIn(FieldLatitude, vs...))
}

// LatitudeNotIn applies the NotIn predicate on the "latitude" field.
func LatitudeNotIn(vs ...float64) predicate.Location {
	return predicate.Location(sql.FieldNotIn(FieldLatitude, vs...))
}

// LatitudeGT applies the GT predicate on the "latitude" field.
func LatitudeGT(v float64) predicate.Location {
	return predicate.Location(sql.FieldGT(FieldLatitude, v))
}

// LatitudeGTE applies the GTE predicate on the "latitude" field.
func LatitudeGTE(v float64) predicate.Location {
	return predicate.Location(sql.FieldGTE(FieldLatitude, v))
}

// LatitudeLT applies the LT predicate on the "latitude" field.
func LatitudeLT(v float64) predicate.Location {
	return predicate.Location(sql.FieldLT(FieldLatitude, v))
}

// LatitudeLTE applies the LTE predicate on the "latitude" field.
func LatitudeLTE(v float64) predicate.Location {
	return predicate.Location(sql.FieldLTE(FieldLatitude, v))
}

// LongitudeEQ applies the EQ predicate on the "longitude" field.
func LongitudeEQ(v float64) predicate.Location {
	return predicate.Location(sql.FieldEQ(FieldLongitude, v))
}

// LongitudeNEQ applies the NEQ predicate on the "longitude" field.
func LongitudeNEQ(v float64) predicate.Location {
	return predicate.Location(sql.FieldNEQ(FieldLongitude, v))
}

// LongitudeIn applies the In predicate on the "longitude" field.
func LongitudeIn(vs ...float64) predicate.Location {
	return predicate.Location(sql.FieldIn(FieldLongitude, vs...))
}

// LongitudeNotIn applies the NotIn predicate on the "longitude" field.
func LongitudeNotIn(vs ...float64) predicate.Location {
	return predicate.Location(sql.FieldNotIn(FieldLongitude, vs...))
}

// LongitudeGT applies the GT predicate on the "longitude" field.
func LongitudeGT(v float64) predicate.Location {
	return predicate.Location(sql.FieldGT(FieldLongitude, v))
}

// LongitudeGTE applies the GTE predicate on the "longitude" field.
func LongitudeGTE(v float64) predicate.Location {
	return predicate.Location(sql.FieldGTE(FieldLongitude, v))
}

// LongitudeLT applies the LT predicate on the "longitude" field.
func LongitudeLT(v float64) predicate.Location {
	return predicate.Location(sql.FieldLT(FieldLongitude, v))
}

// LongitudeLTE applies the LTE predicate on the "longitude" field.
func LongitudeLTE(v float64) predicate.Location {
	return predicate.Location(sql.FieldLTE(FieldLongitude, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Location) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Location) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
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
func Not(p predicate.Location) predicate.Location {
	return predicate.Location(func(s *sql.Selector) {
		p(s.Not())
	})
}
