// Code generated by ent, DO NOT EDIT.

package ent

import (
	"animals/ent/animal"
	"animals/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AnimalUpdate is the builder for updating Animal entities.
type AnimalUpdate struct {
	config
	hooks    []Hook
	mutation *AnimalMutation
}

// Where appends a list predicates to the AnimalUpdate builder.
func (au *AnimalUpdate) Where(ps ...predicate.Animal) *AnimalUpdate {
	au.mutation.Where(ps...)
	return au
}

// Mutation returns the AnimalMutation object of the builder.
func (au *AnimalUpdate) Mutation() *AnimalMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AnimalUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, AnimalMutation](ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AnimalUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AnimalUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AnimalUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AnimalUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(animal.Table, animal.Columns, sqlgraph.NewFieldSpec(animal.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{animal.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AnimalUpdateOne is the builder for updating a single Animal entity.
type AnimalUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AnimalMutation
}

// Mutation returns the AnimalMutation object of the builder.
func (auo *AnimalUpdateOne) Mutation() *AnimalMutation {
	return auo.mutation
}

// Where appends a list predicates to the AnimalUpdate builder.
func (auo *AnimalUpdateOne) Where(ps ...predicate.Animal) *AnimalUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AnimalUpdateOne) Select(field string, fields ...string) *AnimalUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Animal entity.
func (auo *AnimalUpdateOne) Save(ctx context.Context) (*Animal, error) {
	return withHooks[*Animal, AnimalMutation](ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AnimalUpdateOne) SaveX(ctx context.Context) *Animal {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AnimalUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AnimalUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AnimalUpdateOne) sqlSave(ctx context.Context) (_node *Animal, err error) {
	_spec := sqlgraph.NewUpdateSpec(animal.Table, animal.Columns, sqlgraph.NewFieldSpec(animal.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Animal.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, animal.FieldID)
		for _, f := range fields {
			if !animal.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != animal.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &Animal{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{animal.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
