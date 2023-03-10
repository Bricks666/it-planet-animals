// Code generated by ent, DO NOT EDIT.

package ent

import (
	"animals/ent/animal"
	"animals/ent/location"
	"animals/ent/visitedlocation"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// VisitedLocation is the model entity for the VisitedLocation schema.
type VisitedLocation struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// AnimalID holds the value of the "animal_id" field.
	AnimalID uint64 `json:"animal_id,omitempty"`
	// LocationID holds the value of the "location_id" field.
	LocationID uint64 `json:"location_id,omitempty"`
	// DateTimeOfVisitLocationPoint holds the value of the "date_time_of_visit_location_point" field.
	DateTimeOfVisitLocationPoint time.Time `json:"date_time_of_visit_location_point,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VisitedLocationQuery when eager-loading is set.
	Edges VisitedLocationEdges `json:"edges"`
}

// VisitedLocationEdges holds the relations/edges for other nodes in the graph.
type VisitedLocationEdges struct {
	// Animals holds the value of the animals edge.
	Animals *Animal `json:"animals,omitempty"`
	// Locations holds the value of the locations edge.
	Locations *Location `json:"locations,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// AnimalsOrErr returns the Animals value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VisitedLocationEdges) AnimalsOrErr() (*Animal, error) {
	if e.loadedTypes[0] {
		if e.Animals == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: animal.Label}
		}
		return e.Animals, nil
	}
	return nil, &NotLoadedError{edge: "animals"}
}

// LocationsOrErr returns the Locations value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VisitedLocationEdges) LocationsOrErr() (*Location, error) {
	if e.loadedTypes[1] {
		if e.Locations == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: location.Label}
		}
		return e.Locations, nil
	}
	return nil, &NotLoadedError{edge: "locations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VisitedLocation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case visitedlocation.FieldID, visitedlocation.FieldAnimalID, visitedlocation.FieldLocationID:
			values[i] = new(sql.NullInt64)
		case visitedlocation.FieldDateTimeOfVisitLocationPoint:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type VisitedLocation", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VisitedLocation fields.
func (vl *VisitedLocation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case visitedlocation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			vl.ID = uint64(value.Int64)
		case visitedlocation.FieldAnimalID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field animal_id", values[i])
			} else if value.Valid {
				vl.AnimalID = uint64(value.Int64)
			}
		case visitedlocation.FieldLocationID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field location_id", values[i])
			} else if value.Valid {
				vl.LocationID = uint64(value.Int64)
			}
		case visitedlocation.FieldDateTimeOfVisitLocationPoint:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_time_of_visit_location_point", values[i])
			} else if value.Valid {
				vl.DateTimeOfVisitLocationPoint = value.Time
			}
		}
	}
	return nil
}

// QueryAnimals queries the "animals" edge of the VisitedLocation entity.
func (vl *VisitedLocation) QueryAnimals() *AnimalQuery {
	return NewVisitedLocationClient(vl.config).QueryAnimals(vl)
}

// QueryLocations queries the "locations" edge of the VisitedLocation entity.
func (vl *VisitedLocation) QueryLocations() *LocationQuery {
	return NewVisitedLocationClient(vl.config).QueryLocations(vl)
}

// Update returns a builder for updating this VisitedLocation.
// Note that you need to call VisitedLocation.Unwrap() before calling this method if this VisitedLocation
// was returned from a transaction, and the transaction was committed or rolled back.
func (vl *VisitedLocation) Update() *VisitedLocationUpdateOne {
	return NewVisitedLocationClient(vl.config).UpdateOne(vl)
}

// Unwrap unwraps the VisitedLocation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vl *VisitedLocation) Unwrap() *VisitedLocation {
	_tx, ok := vl.config.driver.(*txDriver)
	if !ok {
		panic("ent: VisitedLocation is not a transactional entity")
	}
	vl.config.driver = _tx.drv
	return vl
}

// String implements the fmt.Stringer.
func (vl *VisitedLocation) String() string {
	var builder strings.Builder
	builder.WriteString("VisitedLocation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vl.ID))
	builder.WriteString("animal_id=")
	builder.WriteString(fmt.Sprintf("%v", vl.AnimalID))
	builder.WriteString(", ")
	builder.WriteString("location_id=")
	builder.WriteString(fmt.Sprintf("%v", vl.LocationID))
	builder.WriteString(", ")
	builder.WriteString("date_time_of_visit_location_point=")
	builder.WriteString(vl.DateTimeOfVisitLocationPoint.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// VisitedLocations is a parsable slice of VisitedLocation.
type VisitedLocations []*VisitedLocation
