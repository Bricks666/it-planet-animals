// Code generated by ent, DO NOT EDIT.

package ent

import (
	"animals/ent/animal"
	"animals/ent/animaltype"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AnimalTypeCreate is the builder for creating a AnimalType entity.
type AnimalTypeCreate struct {
	config
	mutation *AnimalTypeMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (atc *AnimalTypeCreate) SetType(s string) *AnimalTypeCreate {
	atc.mutation.SetType(s)
	return atc
}

// SetID sets the "id" field.
func (atc *AnimalTypeCreate) SetID(u uint64) *AnimalTypeCreate {
	atc.mutation.SetID(u)
	return atc
}

// AddAnimalTypeTypeIDs adds the "animal_type_type" edge to the Animal entity by IDs.
func (atc *AnimalTypeCreate) AddAnimalTypeTypeIDs(ids ...uint64) *AnimalTypeCreate {
	atc.mutation.AddAnimalTypeTypeIDs(ids...)
	return atc
}

// AddAnimalTypeType adds the "animal_type_type" edges to the Animal entity.
func (atc *AnimalTypeCreate) AddAnimalTypeType(a ...*Animal) *AnimalTypeCreate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atc.AddAnimalTypeTypeIDs(ids...)
}

// Mutation returns the AnimalTypeMutation object of the builder.
func (atc *AnimalTypeCreate) Mutation() *AnimalTypeMutation {
	return atc.mutation
}

// Save creates the AnimalType in the database.
func (atc *AnimalTypeCreate) Save(ctx context.Context) (*AnimalType, error) {
	return withHooks[*AnimalType, AnimalTypeMutation](ctx, atc.sqlSave, atc.mutation, atc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (atc *AnimalTypeCreate) SaveX(ctx context.Context) *AnimalType {
	v, err := atc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atc *AnimalTypeCreate) Exec(ctx context.Context) error {
	_, err := atc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atc *AnimalTypeCreate) ExecX(ctx context.Context) {
	if err := atc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atc *AnimalTypeCreate) check() error {
	if _, ok := atc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "AnimalType.type"`)}
	}
	if v, ok := atc.mutation.GetType(); ok {
		if err := animaltype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "AnimalType.type": %w`, err)}
		}
	}
	if v, ok := atc.mutation.ID(); ok {
		if err := animaltype.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "AnimalType.id": %w`, err)}
		}
	}
	return nil
}

func (atc *AnimalTypeCreate) sqlSave(ctx context.Context) (*AnimalType, error) {
	if err := atc.check(); err != nil {
		return nil, err
	}
	_node, _spec := atc.createSpec()
	if err := sqlgraph.CreateNode(ctx, atc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	atc.mutation.id = &_node.ID
	atc.mutation.done = true
	return _node, nil
}

func (atc *AnimalTypeCreate) createSpec() (*AnimalType, *sqlgraph.CreateSpec) {
	var (
		_node = &AnimalType{config: atc.config}
		_spec = sqlgraph.NewCreateSpec(animaltype.Table, sqlgraph.NewFieldSpec(animaltype.FieldID, field.TypeUint64))
	)
	if id, ok := atc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := atc.mutation.GetType(); ok {
		_spec.SetField(animaltype.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if nodes := atc.mutation.AnimalTypeTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   animaltype.AnimalTypeTypeTable,
			Columns: animaltype.AnimalTypeTypePrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AnimalTypeCreateBulk is the builder for creating many AnimalType entities in bulk.
type AnimalTypeCreateBulk struct {
	config
	builders []*AnimalTypeCreate
}

// Save creates the AnimalType entities in the database.
func (atcb *AnimalTypeCreateBulk) Save(ctx context.Context) ([]*AnimalType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(atcb.builders))
	nodes := make([]*AnimalType, len(atcb.builders))
	mutators := make([]Mutator, len(atcb.builders))
	for i := range atcb.builders {
		func(i int, root context.Context) {
			builder := atcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AnimalTypeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, atcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, atcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, atcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (atcb *AnimalTypeCreateBulk) SaveX(ctx context.Context) []*AnimalType {
	v, err := atcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atcb *AnimalTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := atcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atcb *AnimalTypeCreateBulk) ExecX(ctx context.Context) {
	if err := atcb.Exec(ctx); err != nil {
		panic(err)
	}
}
