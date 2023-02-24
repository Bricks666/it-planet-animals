// Code generated by ent, DO NOT EDIT.

package ent

import (
	"animals/ent/animal"
	"animals/ent/animalslocations"
	"animals/ent/location"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AnimalsLocationsCreate is the builder for creating a AnimalsLocations entity.
type AnimalsLocationsCreate struct {
	config
	mutation *AnimalsLocationsMutation
	hooks    []Hook
}

// SetAnimalID sets the "animal_id" field.
func (alc *AnimalsLocationsCreate) SetAnimalID(u uint64) *AnimalsLocationsCreate {
	alc.mutation.SetAnimalID(u)
	return alc
}

// SetLocationID sets the "location_id" field.
func (alc *AnimalsLocationsCreate) SetLocationID(u uint64) *AnimalsLocationsCreate {
	alc.mutation.SetLocationID(u)
	return alc
}

// SetDateTimeOfVisitLocationPoint sets the "date_time_of_visit_location_point" field.
func (alc *AnimalsLocationsCreate) SetDateTimeOfVisitLocationPoint(t time.Time) *AnimalsLocationsCreate {
	alc.mutation.SetDateTimeOfVisitLocationPoint(t)
	return alc
}

// SetNillableDateTimeOfVisitLocationPoint sets the "date_time_of_visit_location_point" field if the given value is not nil.
func (alc *AnimalsLocationsCreate) SetNillableDateTimeOfVisitLocationPoint(t *time.Time) *AnimalsLocationsCreate {
	if t != nil {
		alc.SetDateTimeOfVisitLocationPoint(*t)
	}
	return alc
}

// SetID sets the "id" field.
func (alc *AnimalsLocationsCreate) SetID(u uint64) *AnimalsLocationsCreate {
	alc.mutation.SetID(u)
	return alc
}

// SetAnimalsLocationsAnimalID sets the "animals_locations_animal" edge to the Animal entity by ID.
func (alc *AnimalsLocationsCreate) SetAnimalsLocationsAnimalID(id uint64) *AnimalsLocationsCreate {
	alc.mutation.SetAnimalsLocationsAnimalID(id)
	return alc
}

// SetAnimalsLocationsAnimal sets the "animals_locations_animal" edge to the Animal entity.
func (alc *AnimalsLocationsCreate) SetAnimalsLocationsAnimal(a *Animal) *AnimalsLocationsCreate {
	return alc.SetAnimalsLocationsAnimalID(a.ID)
}

// SetAnimalsLocationsLocationID sets the "animals_locations_location" edge to the Location entity by ID.
func (alc *AnimalsLocationsCreate) SetAnimalsLocationsLocationID(id uint64) *AnimalsLocationsCreate {
	alc.mutation.SetAnimalsLocationsLocationID(id)
	return alc
}

// SetAnimalsLocationsLocation sets the "animals_locations_location" edge to the Location entity.
func (alc *AnimalsLocationsCreate) SetAnimalsLocationsLocation(l *Location) *AnimalsLocationsCreate {
	return alc.SetAnimalsLocationsLocationID(l.ID)
}

// Mutation returns the AnimalsLocationsMutation object of the builder.
func (alc *AnimalsLocationsCreate) Mutation() *AnimalsLocationsMutation {
	return alc.mutation
}

// Save creates the AnimalsLocations in the database.
func (alc *AnimalsLocationsCreate) Save(ctx context.Context) (*AnimalsLocations, error) {
	return withHooks[*AnimalsLocations, AnimalsLocationsMutation](ctx, alc.sqlSave, alc.mutation, alc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (alc *AnimalsLocationsCreate) SaveX(ctx context.Context) *AnimalsLocations {
	v, err := alc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (alc *AnimalsLocationsCreate) Exec(ctx context.Context) error {
	_, err := alc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alc *AnimalsLocationsCreate) ExecX(ctx context.Context) {
	if err := alc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (alc *AnimalsLocationsCreate) check() error {
	if _, ok := alc.mutation.AnimalID(); !ok {
		return &ValidationError{Name: "animal_id", err: errors.New(`ent: missing required field "AnimalsLocations.animal_id"`)}
	}
	if _, ok := alc.mutation.LocationID(); !ok {
		return &ValidationError{Name: "location_id", err: errors.New(`ent: missing required field "AnimalsLocations.location_id"`)}
	}
	if v, ok := alc.mutation.ID(); ok {
		if err := animalslocations.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "AnimalsLocations.id": %w`, err)}
		}
	}
	if _, ok := alc.mutation.AnimalsLocationsAnimalID(); !ok {
		return &ValidationError{Name: "animals_locations_animal", err: errors.New(`ent: missing required edge "AnimalsLocations.animals_locations_animal"`)}
	}
	if _, ok := alc.mutation.AnimalsLocationsLocationID(); !ok {
		return &ValidationError{Name: "animals_locations_location", err: errors.New(`ent: missing required edge "AnimalsLocations.animals_locations_location"`)}
	}
	return nil
}

func (alc *AnimalsLocationsCreate) sqlSave(ctx context.Context) (*AnimalsLocations, error) {
	if err := alc.check(); err != nil {
		return nil, err
	}
	_node, _spec := alc.createSpec()
	if err := sqlgraph.CreateNode(ctx, alc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	alc.mutation.id = &_node.ID
	alc.mutation.done = true
	return _node, nil
}

func (alc *AnimalsLocationsCreate) createSpec() (*AnimalsLocations, *sqlgraph.CreateSpec) {
	var (
		_node = &AnimalsLocations{config: alc.config}
		_spec = sqlgraph.NewCreateSpec(animalslocations.Table, sqlgraph.NewFieldSpec(animalslocations.FieldID, field.TypeUint64))
	)
	if id, ok := alc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := alc.mutation.DateTimeOfVisitLocationPoint(); ok {
		_spec.SetField(animalslocations.FieldDateTimeOfVisitLocationPoint, field.TypeTime, value)
		_node.DateTimeOfVisitLocationPoint = value
	}
	if nodes := alc.mutation.AnimalsLocationsAnimalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   animalslocations.AnimalsLocationsAnimalTable,
			Columns: []string{animalslocations.AnimalsLocationsAnimalColumn},
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
		_node.AnimalID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := alc.mutation.AnimalsLocationsLocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   animalslocations.AnimalsLocationsLocationTable,
			Columns: []string{animalslocations.AnimalsLocationsLocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: location.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.LocationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AnimalsLocationsCreateBulk is the builder for creating many AnimalsLocations entities in bulk.
type AnimalsLocationsCreateBulk struct {
	config
	builders []*AnimalsLocationsCreate
}

// Save creates the AnimalsLocations entities in the database.
func (alcb *AnimalsLocationsCreateBulk) Save(ctx context.Context) ([]*AnimalsLocations, error) {
	specs := make([]*sqlgraph.CreateSpec, len(alcb.builders))
	nodes := make([]*AnimalsLocations, len(alcb.builders))
	mutators := make([]Mutator, len(alcb.builders))
	for i := range alcb.builders {
		func(i int, root context.Context) {
			builder := alcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AnimalsLocationsMutation)
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
					_, err = mutators[i+1].Mutate(root, alcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, alcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, alcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (alcb *AnimalsLocationsCreateBulk) SaveX(ctx context.Context) []*AnimalsLocations {
	v, err := alcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (alcb *AnimalsLocationsCreateBulk) Exec(ctx context.Context) error {
	_, err := alcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alcb *AnimalsLocationsCreateBulk) ExecX(ctx context.Context) {
	if err := alcb.Exec(ctx); err != nil {
		panic(err)
	}
}
