// Code generated by ent, DO NOT EDIT.

package ent

import (
	"animals/ent/animal"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Animal is the model entity for the Animal schema.
type Animal struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Animal) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case animal.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Animal", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Animal fields.
func (a *Animal) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case animal.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Animal.
// Note that you need to call Animal.Unwrap() before calling this method if this Animal
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Animal) Update() *AnimalUpdateOne {
	return NewAnimalClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Animal entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Animal) Unwrap() *Animal {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Animal is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Animal) String() string {
	var builder strings.Builder
	builder.WriteString("Animal(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Animals is a parsable slice of Animal.
type Animals []*Animal
