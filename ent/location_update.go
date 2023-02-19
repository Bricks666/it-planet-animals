// Code generated by ent, DO NOT EDIT.

package ent

import (
	"animals/ent/location"
	"animals/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LocationUpdate is the builder for updating Location entities.
type LocationUpdate struct {
	config
	hooks    []Hook
	mutation *LocationMutation
}

// Where appends a list predicates to the LocationUpdate builder.
func (lu *LocationUpdate) Where(ps ...predicate.Location) *LocationUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetLatitude sets the "latitude" field.
func (lu *LocationUpdate) SetLatitude(f float64) *LocationUpdate {
	lu.mutation.ResetLatitude()
	lu.mutation.SetLatitude(f)
	return lu
}

// AddLatitude adds f to the "latitude" field.
func (lu *LocationUpdate) AddLatitude(f float64) *LocationUpdate {
	lu.mutation.AddLatitude(f)
	return lu
}

// SetLongitude sets the "longitude" field.
func (lu *LocationUpdate) SetLongitude(f float64) *LocationUpdate {
	lu.mutation.ResetLongitude()
	lu.mutation.SetLongitude(f)
	return lu
}

// AddLongitude adds f to the "longitude" field.
func (lu *LocationUpdate) AddLongitude(f float64) *LocationUpdate {
	lu.mutation.AddLongitude(f)
	return lu
}

// Mutation returns the LocationMutation object of the builder.
func (lu *LocationUpdate) Mutation() *LocationMutation {
	return lu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LocationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, LocationMutation](ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LocationUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LocationUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LocationUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *LocationUpdate) check() error {
	if v, ok := lu.mutation.Latitude(); ok {
		if err := location.LatitudeValidator(v); err != nil {
			return &ValidationError{Name: "latitude", err: fmt.Errorf(`ent: validator failed for field "Location.latitude": %w`, err)}
		}
	}
	if v, ok := lu.mutation.Longitude(); ok {
		if err := location.LongitudeValidator(v); err != nil {
			return &ValidationError{Name: "longitude", err: fmt.Errorf(`ent: validator failed for field "Location.longitude": %w`, err)}
		}
	}
	return nil
}

func (lu *LocationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(location.Table, location.Columns, sqlgraph.NewFieldSpec(location.FieldID, field.TypeUint64))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Latitude(); ok {
		_spec.SetField(location.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := lu.mutation.AddedLatitude(); ok {
		_spec.AddField(location.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := lu.mutation.Longitude(); ok {
		_spec.SetField(location.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := lu.mutation.AddedLongitude(); ok {
		_spec.AddField(location.FieldLongitude, field.TypeFloat64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{location.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LocationUpdateOne is the builder for updating a single Location entity.
type LocationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LocationMutation
}

// SetLatitude sets the "latitude" field.
func (luo *LocationUpdateOne) SetLatitude(f float64) *LocationUpdateOne {
	luo.mutation.ResetLatitude()
	luo.mutation.SetLatitude(f)
	return luo
}

// AddLatitude adds f to the "latitude" field.
func (luo *LocationUpdateOne) AddLatitude(f float64) *LocationUpdateOne {
	luo.mutation.AddLatitude(f)
	return luo
}

// SetLongitude sets the "longitude" field.
func (luo *LocationUpdateOne) SetLongitude(f float64) *LocationUpdateOne {
	luo.mutation.ResetLongitude()
	luo.mutation.SetLongitude(f)
	return luo
}

// AddLongitude adds f to the "longitude" field.
func (luo *LocationUpdateOne) AddLongitude(f float64) *LocationUpdateOne {
	luo.mutation.AddLongitude(f)
	return luo
}

// Mutation returns the LocationMutation object of the builder.
func (luo *LocationUpdateOne) Mutation() *LocationMutation {
	return luo.mutation
}

// Where appends a list predicates to the LocationUpdate builder.
func (luo *LocationUpdateOne) Where(ps ...predicate.Location) *LocationUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LocationUpdateOne) Select(field string, fields ...string) *LocationUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Location entity.
func (luo *LocationUpdateOne) Save(ctx context.Context) (*Location, error) {
	return withHooks[*Location, LocationMutation](ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LocationUpdateOne) SaveX(ctx context.Context) *Location {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LocationUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LocationUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *LocationUpdateOne) check() error {
	if v, ok := luo.mutation.Latitude(); ok {
		if err := location.LatitudeValidator(v); err != nil {
			return &ValidationError{Name: "latitude", err: fmt.Errorf(`ent: validator failed for field "Location.latitude": %w`, err)}
		}
	}
	if v, ok := luo.mutation.Longitude(); ok {
		if err := location.LongitudeValidator(v); err != nil {
			return &ValidationError{Name: "longitude", err: fmt.Errorf(`ent: validator failed for field "Location.longitude": %w`, err)}
		}
	}
	return nil
}

func (luo *LocationUpdateOne) sqlSave(ctx context.Context) (_node *Location, err error) {
	if err := luo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(location.Table, location.Columns, sqlgraph.NewFieldSpec(location.FieldID, field.TypeUint64))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Location.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, location.FieldID)
		for _, f := range fields {
			if !location.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != location.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Latitude(); ok {
		_spec.SetField(location.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := luo.mutation.AddedLatitude(); ok {
		_spec.AddField(location.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := luo.mutation.Longitude(); ok {
		_spec.SetField(location.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := luo.mutation.AddedLongitude(); ok {
		_spec.AddField(location.FieldLongitude, field.TypeFloat64, value)
	}
	_node = &Location{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{location.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
