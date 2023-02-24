// Code generated by ent, DO NOT EDIT.

package ent

import (
	"animals/ent/animal"
	"animals/ent/animaltag"
	"animals/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AnimalTagUpdate is the builder for updating AnimalTag entities.
type AnimalTagUpdate struct {
	config
	hooks    []Hook
	mutation *AnimalTagMutation
}

// Where appends a list predicates to the AnimalTagUpdate builder.
func (atu *AnimalTagUpdate) Where(ps ...predicate.AnimalTag) *AnimalTagUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// SetType sets the "type" field.
func (atu *AnimalTagUpdate) SetType(s string) *AnimalTagUpdate {
	atu.mutation.SetType(s)
	return atu
}

// AddAnimalIDs adds the "animals" edge to the Animal entity by IDs.
func (atu *AnimalTagUpdate) AddAnimalIDs(ids ...uint64) *AnimalTagUpdate {
	atu.mutation.AddAnimalIDs(ids...)
	return atu
}

// AddAnimals adds the "animals" edges to the Animal entity.
func (atu *AnimalTagUpdate) AddAnimals(a ...*Animal) *AnimalTagUpdate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atu.AddAnimalIDs(ids...)
}

// Mutation returns the AnimalTagMutation object of the builder.
func (atu *AnimalTagUpdate) Mutation() *AnimalTagMutation {
	return atu.mutation
}

// ClearAnimals clears all "animals" edges to the Animal entity.
func (atu *AnimalTagUpdate) ClearAnimals() *AnimalTagUpdate {
	atu.mutation.ClearAnimals()
	return atu
}

// RemoveAnimalIDs removes the "animals" edge to Animal entities by IDs.
func (atu *AnimalTagUpdate) RemoveAnimalIDs(ids ...uint64) *AnimalTagUpdate {
	atu.mutation.RemoveAnimalIDs(ids...)
	return atu
}

// RemoveAnimals removes "animals" edges to Animal entities.
func (atu *AnimalTagUpdate) RemoveAnimals(a ...*Animal) *AnimalTagUpdate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atu.RemoveAnimalIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *AnimalTagUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, AnimalTagMutation](ctx, atu.sqlSave, atu.mutation, atu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atu *AnimalTagUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *AnimalTagUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *AnimalTagUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atu *AnimalTagUpdate) check() error {
	if v, ok := atu.mutation.GetType(); ok {
		if err := animaltag.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "AnimalTag.type": %w`, err)}
		}
	}
	return nil
}

func (atu *AnimalTagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := atu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(animaltag.Table, animaltag.Columns, sqlgraph.NewFieldSpec(animaltag.FieldID, field.TypeUint64))
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atu.mutation.GetType(); ok {
		_spec.SetField(animaltag.FieldType, field.TypeString, value)
	}
	if atu.mutation.AnimalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   animaltag.AnimalsTable,
			Columns: animaltag.AnimalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: animal.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.RemovedAnimalsIDs(); len(nodes) > 0 && !atu.mutation.AnimalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   animaltag.AnimalsTable,
			Columns: animaltag.AnimalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: animal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.AnimalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   animaltag.AnimalsTable,
			Columns: animaltag.AnimalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: animal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{animaltag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	atu.mutation.done = true
	return n, nil
}

// AnimalTagUpdateOne is the builder for updating a single AnimalTag entity.
type AnimalTagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AnimalTagMutation
}

// SetType sets the "type" field.
func (atuo *AnimalTagUpdateOne) SetType(s string) *AnimalTagUpdateOne {
	atuo.mutation.SetType(s)
	return atuo
}

// AddAnimalIDs adds the "animals" edge to the Animal entity by IDs.
func (atuo *AnimalTagUpdateOne) AddAnimalIDs(ids ...uint64) *AnimalTagUpdateOne {
	atuo.mutation.AddAnimalIDs(ids...)
	return atuo
}

// AddAnimals adds the "animals" edges to the Animal entity.
func (atuo *AnimalTagUpdateOne) AddAnimals(a ...*Animal) *AnimalTagUpdateOne {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atuo.AddAnimalIDs(ids...)
}

// Mutation returns the AnimalTagMutation object of the builder.
func (atuo *AnimalTagUpdateOne) Mutation() *AnimalTagMutation {
	return atuo.mutation
}

// ClearAnimals clears all "animals" edges to the Animal entity.
func (atuo *AnimalTagUpdateOne) ClearAnimals() *AnimalTagUpdateOne {
	atuo.mutation.ClearAnimals()
	return atuo
}

// RemoveAnimalIDs removes the "animals" edge to Animal entities by IDs.
func (atuo *AnimalTagUpdateOne) RemoveAnimalIDs(ids ...uint64) *AnimalTagUpdateOne {
	atuo.mutation.RemoveAnimalIDs(ids...)
	return atuo
}

// RemoveAnimals removes "animals" edges to Animal entities.
func (atuo *AnimalTagUpdateOne) RemoveAnimals(a ...*Animal) *AnimalTagUpdateOne {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atuo.RemoveAnimalIDs(ids...)
}

// Where appends a list predicates to the AnimalTagUpdate builder.
func (atuo *AnimalTagUpdateOne) Where(ps ...predicate.AnimalTag) *AnimalTagUpdateOne {
	atuo.mutation.Where(ps...)
	return atuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *AnimalTagUpdateOne) Select(field string, fields ...string) *AnimalTagUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated AnimalTag entity.
func (atuo *AnimalTagUpdateOne) Save(ctx context.Context) (*AnimalTag, error) {
	return withHooks[*AnimalTag, AnimalTagMutation](ctx, atuo.sqlSave, atuo.mutation, atuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *AnimalTagUpdateOne) SaveX(ctx context.Context) *AnimalTag {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *AnimalTagUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *AnimalTagUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atuo *AnimalTagUpdateOne) check() error {
	if v, ok := atuo.mutation.GetType(); ok {
		if err := animaltag.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "AnimalTag.type": %w`, err)}
		}
	}
	return nil
}

func (atuo *AnimalTagUpdateOne) sqlSave(ctx context.Context) (_node *AnimalTag, err error) {
	if err := atuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(animaltag.Table, animaltag.Columns, sqlgraph.NewFieldSpec(animaltag.FieldID, field.TypeUint64))
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AnimalTag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, animaltag.FieldID)
		for _, f := range fields {
			if !animaltag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != animaltag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atuo.mutation.GetType(); ok {
		_spec.SetField(animaltag.FieldType, field.TypeString, value)
	}
	if atuo.mutation.AnimalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   animaltag.AnimalsTable,
			Columns: animaltag.AnimalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: animal.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.RemovedAnimalsIDs(); len(nodes) > 0 && !atuo.mutation.AnimalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   animaltag.AnimalsTable,
			Columns: animaltag.AnimalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: animal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.AnimalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   animaltag.AnimalsTable,
			Columns: animaltag.AnimalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: animal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AnimalTag{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{animaltag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	atuo.mutation.done = true
	return _node, nil
}
