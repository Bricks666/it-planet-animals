// Code generated by ent, DO NOT EDIT.

package ent

import (
	"animals/ent/animal"
	"animals/ent/location"
	"animals/ent/visitedlocation"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LocationCreate is the builder for creating a Location entity.
type LocationCreate struct {
	config
	mutation *LocationMutation
	hooks    []Hook
}

// SetLatitude sets the "latitude" field.
func (lc *LocationCreate) SetLatitude(f float64) *LocationCreate {
	lc.mutation.SetLatitude(f)
	return lc
}

// SetLongitude sets the "longitude" field.
func (lc *LocationCreate) SetLongitude(f float64) *LocationCreate {
	lc.mutation.SetLongitude(f)
	return lc
}

// SetID sets the "id" field.
func (lc *LocationCreate) SetID(u uint64) *LocationCreate {
	lc.mutation.SetID(u)
	return lc
}

// AddChippedAnimalIDs adds the "chipped_animals" edge to the Animal entity by IDs.
func (lc *LocationCreate) AddChippedAnimalIDs(ids ...uint64) *LocationCreate {
	lc.mutation.AddChippedAnimalIDs(ids...)
	return lc
}

// AddChippedAnimals adds the "chipped_animals" edges to the Animal entity.
func (lc *LocationCreate) AddChippedAnimals(a ...*Animal) *LocationCreate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return lc.AddChippedAnimalIDs(ids...)
}

// AddAnimalIDs adds the "animals" edge to the Animal entity by IDs.
func (lc *LocationCreate) AddAnimalIDs(ids ...uint64) *LocationCreate {
	lc.mutation.AddAnimalIDs(ids...)
	return lc
}

// AddAnimals adds the "animals" edges to the Animal entity.
func (lc *LocationCreate) AddAnimals(a ...*Animal) *LocationCreate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return lc.AddAnimalIDs(ids...)
}

// AddHavingAnimalIDs adds the "having_animals" edge to the VisitedLocation entity by IDs.
func (lc *LocationCreate) AddHavingAnimalIDs(ids ...uint64) *LocationCreate {
	lc.mutation.AddHavingAnimalIDs(ids...)
	return lc
}

// AddHavingAnimals adds the "having_animals" edges to the VisitedLocation entity.
func (lc *LocationCreate) AddHavingAnimals(v ...*VisitedLocation) *LocationCreate {
	ids := make([]uint64, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return lc.AddHavingAnimalIDs(ids...)
}

// Mutation returns the LocationMutation object of the builder.
func (lc *LocationCreate) Mutation() *LocationMutation {
	return lc.mutation
}

// Save creates the Location in the database.
func (lc *LocationCreate) Save(ctx context.Context) (*Location, error) {
	return withHooks[*Location, LocationMutation](ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LocationCreate) SaveX(ctx context.Context) *Location {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LocationCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LocationCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LocationCreate) check() error {
	if _, ok := lc.mutation.Latitude(); !ok {
		return &ValidationError{Name: "latitude", err: errors.New(`ent: missing required field "Location.latitude"`)}
	}
	if v, ok := lc.mutation.Latitude(); ok {
		if err := location.LatitudeValidator(v); err != nil {
			return &ValidationError{Name: "latitude", err: fmt.Errorf(`ent: validator failed for field "Location.latitude": %w`, err)}
		}
	}
	if _, ok := lc.mutation.Longitude(); !ok {
		return &ValidationError{Name: "longitude", err: errors.New(`ent: missing required field "Location.longitude"`)}
	}
	if v, ok := lc.mutation.Longitude(); ok {
		if err := location.LongitudeValidator(v); err != nil {
			return &ValidationError{Name: "longitude", err: fmt.Errorf(`ent: validator failed for field "Location.longitude": %w`, err)}
		}
	}
	if v, ok := lc.mutation.ID(); ok {
		if err := location.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Location.id": %w`, err)}
		}
	}
	return nil
}

func (lc *LocationCreate) sqlSave(ctx context.Context) (*Location, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LocationCreate) createSpec() (*Location, *sqlgraph.CreateSpec) {
	var (
		_node = &Location{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(location.Table, sqlgraph.NewFieldSpec(location.FieldID, field.TypeUint64))
	)
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := lc.mutation.Latitude(); ok {
		_spec.SetField(location.FieldLatitude, field.TypeFloat64, value)
		_node.Latitude = value
	}
	if value, ok := lc.mutation.Longitude(); ok {
		_spec.SetField(location.FieldLongitude, field.TypeFloat64, value)
		_node.Longitude = value
	}
	if nodes := lc.mutation.ChippedAnimalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ChippedAnimalsTable,
			Columns: []string{location.ChippedAnimalsColumn},
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
	if nodes := lc.mutation.AnimalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   location.AnimalsTable,
			Columns: location.AnimalsPrimaryKey,
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
	if nodes := lc.mutation.HavingAnimalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   location.HavingAnimalsTable,
			Columns: []string{location.HavingAnimalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: visitedlocation.FieldID,
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

// LocationCreateBulk is the builder for creating many Location entities in bulk.
type LocationCreateBulk struct {
	config
	builders []*LocationCreate
}

// Save creates the Location entities in the database.
func (lcb *LocationCreateBulk) Save(ctx context.Context) ([]*Location, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Location, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LocationMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LocationCreateBulk) SaveX(ctx context.Context) []*Location {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LocationCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LocationCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
