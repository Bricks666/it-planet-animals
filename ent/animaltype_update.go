// Code generated by ent, DO NOT EDIT.

package ent

import (
	"animals/ent/animal"
	"animals/ent/animaltype"
	"animals/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AnimalTypeUpdate is the builder for updating AnimalType entities.
type AnimalTypeUpdate struct {
	config
	hooks    []Hook
	mutation *AnimalTypeMutation
}

// Where appends a list predicates to the AnimalTypeUpdate builder.
func (atu *AnimalTypeUpdate) Where(ps ...predicate.AnimalType) *AnimalTypeUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// SetType sets the "type" field.
func (atu *AnimalTypeUpdate) SetType(s string) *AnimalTypeUpdate {
	atu.mutation.SetType(s)
	return atu
}

// AddAnimalTagsTypeIDs adds the "animal_tags_types" edge to the Animal entity by IDs.
func (atu *AnimalTypeUpdate) AddAnimalTagsTypeIDs(ids ...uint64) *AnimalTypeUpdate {
	atu.mutation.AddAnimalTagsTypeIDs(ids...)
	return atu
}

// AddAnimalTagsTypes adds the "animal_tags_types" edges to the Animal entity.
func (atu *AnimalTypeUpdate) AddAnimalTagsTypes(a ...*Animal) *AnimalTypeUpdate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atu.AddAnimalTagsTypeIDs(ids...)
}

// Mutation returns the AnimalTypeMutation object of the builder.
func (atu *AnimalTypeUpdate) Mutation() *AnimalTypeMutation {
	return atu.mutation
}

// ClearAnimalTagsTypes clears all "animal_tags_types" edges to the Animal entity.
func (atu *AnimalTypeUpdate) ClearAnimalTagsTypes() *AnimalTypeUpdate {
	atu.mutation.ClearAnimalTagsTypes()
	return atu
}

// RemoveAnimalTagsTypeIDs removes the "animal_tags_types" edge to Animal entities by IDs.
func (atu *AnimalTypeUpdate) RemoveAnimalTagsTypeIDs(ids ...uint64) *AnimalTypeUpdate {
	atu.mutation.RemoveAnimalTagsTypeIDs(ids...)
	return atu
}

// RemoveAnimalTagsTypes removes "animal_tags_types" edges to Animal entities.
func (atu *AnimalTypeUpdate) RemoveAnimalTagsTypes(a ...*Animal) *AnimalTypeUpdate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atu.RemoveAnimalTagsTypeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *AnimalTypeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, AnimalTypeMutation](ctx, atu.sqlSave, atu.mutation, atu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atu *AnimalTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *AnimalTypeUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *AnimalTypeUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atu *AnimalTypeUpdate) check() error {
	if v, ok := atu.mutation.GetType(); ok {
		if err := animaltype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "AnimalType.type": %w`, err)}
		}
	}
	return nil
}

func (atu *AnimalTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := atu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(animaltype.Table, animaltype.Columns, sqlgraph.NewFieldSpec(animaltype.FieldID, field.TypeUint64))
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atu.mutation.GetType(); ok {
		_spec.SetField(animaltype.FieldType, field.TypeString, value)
	}
	if atu.mutation.AnimalTagsTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   animaltype.AnimalTagsTypesTable,
			Columns: animaltype.AnimalTagsTypesPrimaryKey,
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
	if nodes := atu.mutation.RemovedAnimalTagsTypesIDs(); len(nodes) > 0 && !atu.mutation.AnimalTagsTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   animaltype.AnimalTagsTypesTable,
			Columns: animaltype.AnimalTagsTypesPrimaryKey,
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
	if nodes := atu.mutation.AnimalTagsTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   animaltype.AnimalTagsTypesTable,
			Columns: animaltype.AnimalTagsTypesPrimaryKey,
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
			err = &NotFoundError{animaltype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	atu.mutation.done = true
	return n, nil
}

// AnimalTypeUpdateOne is the builder for updating a single AnimalType entity.
type AnimalTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AnimalTypeMutation
}

// SetType sets the "type" field.
func (atuo *AnimalTypeUpdateOne) SetType(s string) *AnimalTypeUpdateOne {
	atuo.mutation.SetType(s)
	return atuo
}

// AddAnimalTagsTypeIDs adds the "animal_tags_types" edge to the Animal entity by IDs.
func (atuo *AnimalTypeUpdateOne) AddAnimalTagsTypeIDs(ids ...uint64) *AnimalTypeUpdateOne {
	atuo.mutation.AddAnimalTagsTypeIDs(ids...)
	return atuo
}

// AddAnimalTagsTypes adds the "animal_tags_types" edges to the Animal entity.
func (atuo *AnimalTypeUpdateOne) AddAnimalTagsTypes(a ...*Animal) *AnimalTypeUpdateOne {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atuo.AddAnimalTagsTypeIDs(ids...)
}

// Mutation returns the AnimalTypeMutation object of the builder.
func (atuo *AnimalTypeUpdateOne) Mutation() *AnimalTypeMutation {
	return atuo.mutation
}

// ClearAnimalTagsTypes clears all "animal_tags_types" edges to the Animal entity.
func (atuo *AnimalTypeUpdateOne) ClearAnimalTagsTypes() *AnimalTypeUpdateOne {
	atuo.mutation.ClearAnimalTagsTypes()
	return atuo
}

// RemoveAnimalTagsTypeIDs removes the "animal_tags_types" edge to Animal entities by IDs.
func (atuo *AnimalTypeUpdateOne) RemoveAnimalTagsTypeIDs(ids ...uint64) *AnimalTypeUpdateOne {
	atuo.mutation.RemoveAnimalTagsTypeIDs(ids...)
	return atuo
}

// RemoveAnimalTagsTypes removes "animal_tags_types" edges to Animal entities.
func (atuo *AnimalTypeUpdateOne) RemoveAnimalTagsTypes(a ...*Animal) *AnimalTypeUpdateOne {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atuo.RemoveAnimalTagsTypeIDs(ids...)
}

// Where appends a list predicates to the AnimalTypeUpdate builder.
func (atuo *AnimalTypeUpdateOne) Where(ps ...predicate.AnimalType) *AnimalTypeUpdateOne {
	atuo.mutation.Where(ps...)
	return atuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *AnimalTypeUpdateOne) Select(field string, fields ...string) *AnimalTypeUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated AnimalType entity.
func (atuo *AnimalTypeUpdateOne) Save(ctx context.Context) (*AnimalType, error) {
	return withHooks[*AnimalType, AnimalTypeMutation](ctx, atuo.sqlSave, atuo.mutation, atuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *AnimalTypeUpdateOne) SaveX(ctx context.Context) *AnimalType {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *AnimalTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *AnimalTypeUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atuo *AnimalTypeUpdateOne) check() error {
	if v, ok := atuo.mutation.GetType(); ok {
		if err := animaltype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "AnimalType.type": %w`, err)}
		}
	}
	return nil
}

func (atuo *AnimalTypeUpdateOne) sqlSave(ctx context.Context) (_node *AnimalType, err error) {
	if err := atuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(animaltype.Table, animaltype.Columns, sqlgraph.NewFieldSpec(animaltype.FieldID, field.TypeUint64))
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AnimalType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, animaltype.FieldID)
		for _, f := range fields {
			if !animaltype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != animaltype.FieldID {
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
		_spec.SetField(animaltype.FieldType, field.TypeString, value)
	}
	if atuo.mutation.AnimalTagsTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   animaltype.AnimalTagsTypesTable,
			Columns: animaltype.AnimalTagsTypesPrimaryKey,
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
	if nodes := atuo.mutation.RemovedAnimalTagsTypesIDs(); len(nodes) > 0 && !atuo.mutation.AnimalTagsTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   animaltype.AnimalTagsTypesTable,
			Columns: animaltype.AnimalTagsTypesPrimaryKey,
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
	if nodes := atuo.mutation.AnimalTagsTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   animaltype.AnimalTagsTypesTable,
			Columns: animaltype.AnimalTagsTypesPrimaryKey,
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
	_node = &AnimalType{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{animaltype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	atuo.mutation.done = true
	return _node, nil
}
