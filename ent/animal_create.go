// Code generated by ent, DO NOT EDIT.

package ent

import (
	"animals/ent/animal"
	"animals/ent/animalslocations"
	"animals/ent/animaltype"
	"animals/ent/location"
	"animals/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AnimalCreate is the builder for creating a Animal entity.
type AnimalCreate struct {
	config
	mutation *AnimalMutation
	hooks    []Hook
}

// SetWeight sets the "weight" field.
func (ac *AnimalCreate) SetWeight(f float32) *AnimalCreate {
	ac.mutation.SetWeight(f)
	return ac
}

// SetLength sets the "length" field.
func (ac *AnimalCreate) SetLength(f float32) *AnimalCreate {
	ac.mutation.SetLength(f)
	return ac
}

// SetHeight sets the "height" field.
func (ac *AnimalCreate) SetHeight(f float32) *AnimalCreate {
	ac.mutation.SetHeight(f)
	return ac
}

// SetGender sets the "gender" field.
func (ac *AnimalCreate) SetGender(a animal.Gender) *AnimalCreate {
	ac.mutation.SetGender(a)
	return ac
}

// SetLifeStatus sets the "life_status" field.
func (ac *AnimalCreate) SetLifeStatus(as animal.LifeStatus) *AnimalCreate {
	ac.mutation.SetLifeStatus(as)
	return ac
}

// SetNillableLifeStatus sets the "life_status" field if the given value is not nil.
func (ac *AnimalCreate) SetNillableLifeStatus(as *animal.LifeStatus) *AnimalCreate {
	if as != nil {
		ac.SetLifeStatus(*as)
	}
	return ac
}

// SetChippingDateTime sets the "chipping_date_time" field.
func (ac *AnimalCreate) SetChippingDateTime(t time.Time) *AnimalCreate {
	ac.mutation.SetChippingDateTime(t)
	return ac
}

// SetNillableChippingDateTime sets the "chipping_date_time" field if the given value is not nil.
func (ac *AnimalCreate) SetNillableChippingDateTime(t *time.Time) *AnimalCreate {
	if t != nil {
		ac.SetChippingDateTime(*t)
	}
	return ac
}

// SetChipperID sets the "chipper_id" field.
func (ac *AnimalCreate) SetChipperID(u uint32) *AnimalCreate {
	ac.mutation.SetChipperID(u)
	return ac
}

// SetNillableChipperID sets the "chipper_id" field if the given value is not nil.
func (ac *AnimalCreate) SetNillableChipperID(u *uint32) *AnimalCreate {
	if u != nil {
		ac.SetChipperID(*u)
	}
	return ac
}

// SetChippingLocationID sets the "chipping_location_id" field.
func (ac *AnimalCreate) SetChippingLocationID(u uint64) *AnimalCreate {
	ac.mutation.SetChippingLocationID(u)
	return ac
}

// SetNillableChippingLocationID sets the "chipping_location_id" field if the given value is not nil.
func (ac *AnimalCreate) SetNillableChippingLocationID(u *uint64) *AnimalCreate {
	if u != nil {
		ac.SetChippingLocationID(*u)
	}
	return ac
}

// SetDeathDateTime sets the "death_date_time" field.
func (ac *AnimalCreate) SetDeathDateTime(t time.Time) *AnimalCreate {
	ac.mutation.SetDeathDateTime(t)
	return ac
}

// SetNillableDeathDateTime sets the "death_date_time" field if the given value is not nil.
func (ac *AnimalCreate) SetNillableDeathDateTime(t *time.Time) *AnimalCreate {
	if t != nil {
		ac.SetDeathDateTime(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AnimalCreate) SetID(u uint64) *AnimalCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetChipperAnimalID sets the "chipper_animal" edge to the User entity by ID.
func (ac *AnimalCreate) SetChipperAnimalID(id uint32) *AnimalCreate {
	ac.mutation.SetChipperAnimalID(id)
	return ac
}

// SetNillableChipperAnimalID sets the "chipper_animal" edge to the User entity by ID if the given value is not nil.
func (ac *AnimalCreate) SetNillableChipperAnimalID(id *uint32) *AnimalCreate {
	if id != nil {
		ac = ac.SetChipperAnimalID(*id)
	}
	return ac
}

// SetChipperAnimal sets the "chipper_animal" edge to the User entity.
func (ac *AnimalCreate) SetChipperAnimal(u *User) *AnimalCreate {
	return ac.SetChipperAnimalID(u.ID)
}

// AddAnimalTypeAnimalIDs adds the "animal_type_animal" edge to the AnimalType entity by IDs.
func (ac *AnimalCreate) AddAnimalTypeAnimalIDs(ids ...uint64) *AnimalCreate {
	ac.mutation.AddAnimalTypeAnimalIDs(ids...)
	return ac
}

// AddAnimalTypeAnimal adds the "animal_type_animal" edges to the AnimalType entity.
func (ac *AnimalCreate) AddAnimalTypeAnimal(a ...*AnimalType) *AnimalCreate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddAnimalTypeAnimalIDs(ids...)
}

// SetChippingLocation sets the "chipping_location" edge to the Location entity.
func (ac *AnimalCreate) SetChippingLocation(l *Location) *AnimalCreate {
	return ac.SetChippingLocationID(l.ID)
}

// AddVisitedLocationIDs adds the "visited_locations" edge to the Location entity by IDs.
func (ac *AnimalCreate) AddVisitedLocationIDs(ids ...uint64) *AnimalCreate {
	ac.mutation.AddVisitedLocationIDs(ids...)
	return ac
}

// AddVisitedLocations adds the "visited_locations" edges to the Location entity.
func (ac *AnimalCreate) AddVisitedLocations(l ...*Location) *AnimalCreate {
	ids := make([]uint64, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ac.AddVisitedLocationIDs(ids...)
}

// AddAnimalIDs adds the "animals" edge to the AnimalsLocations entity by IDs.
func (ac *AnimalCreate) AddAnimalIDs(ids ...uint64) *AnimalCreate {
	ac.mutation.AddAnimalIDs(ids...)
	return ac
}

// AddAnimals adds the "animals" edges to the AnimalsLocations entity.
func (ac *AnimalCreate) AddAnimals(a ...*AnimalsLocations) *AnimalCreate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddAnimalIDs(ids...)
}

// Mutation returns the AnimalMutation object of the builder.
func (ac *AnimalCreate) Mutation() *AnimalMutation {
	return ac.mutation
}

// Save creates the Animal in the database.
func (ac *AnimalCreate) Save(ctx context.Context) (*Animal, error) {
	ac.defaults()
	return withHooks[*Animal, AnimalMutation](ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AnimalCreate) SaveX(ctx context.Context) *Animal {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AnimalCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AnimalCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AnimalCreate) defaults() {
	if _, ok := ac.mutation.LifeStatus(); !ok {
		v := animal.DefaultLifeStatus
		ac.mutation.SetLifeStatus(v)
	}
	if _, ok := ac.mutation.ChippingDateTime(); !ok {
		v := animal.DefaultChippingDateTime()
		ac.mutation.SetChippingDateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AnimalCreate) check() error {
	if _, ok := ac.mutation.Weight(); !ok {
		return &ValidationError{Name: "weight", err: errors.New(`ent: missing required field "Animal.weight"`)}
	}
	if v, ok := ac.mutation.Weight(); ok {
		if err := animal.WeightValidator(v); err != nil {
			return &ValidationError{Name: "weight", err: fmt.Errorf(`ent: validator failed for field "Animal.weight": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Length(); !ok {
		return &ValidationError{Name: "length", err: errors.New(`ent: missing required field "Animal.length"`)}
	}
	if v, ok := ac.mutation.Length(); ok {
		if err := animal.LengthValidator(v); err != nil {
			return &ValidationError{Name: "length", err: fmt.Errorf(`ent: validator failed for field "Animal.length": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Height(); !ok {
		return &ValidationError{Name: "height", err: errors.New(`ent: missing required field "Animal.height"`)}
	}
	if v, ok := ac.mutation.Height(); ok {
		if err := animal.HeightValidator(v); err != nil {
			return &ValidationError{Name: "height", err: fmt.Errorf(`ent: validator failed for field "Animal.height": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Gender(); !ok {
		return &ValidationError{Name: "gender", err: errors.New(`ent: missing required field "Animal.gender"`)}
	}
	if v, ok := ac.mutation.Gender(); ok {
		if err := animal.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "Animal.gender": %w`, err)}
		}
	}
	if _, ok := ac.mutation.LifeStatus(); !ok {
		return &ValidationError{Name: "life_status", err: errors.New(`ent: missing required field "Animal.life_status"`)}
	}
	if v, ok := ac.mutation.LifeStatus(); ok {
		if err := animal.LifeStatusValidator(v); err != nil {
			return &ValidationError{Name: "life_status", err: fmt.Errorf(`ent: validator failed for field "Animal.life_status": %w`, err)}
		}
	}
	if _, ok := ac.mutation.ChippingDateTime(); !ok {
		return &ValidationError{Name: "chipping_date_time", err: errors.New(`ent: missing required field "Animal.chipping_date_time"`)}
	}
	if v, ok := ac.mutation.ID(); ok {
		if err := animal.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Animal.id": %w`, err)}
		}
	}
	return nil
}

func (ac *AnimalCreate) sqlSave(ctx context.Context) (*Animal, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AnimalCreate) createSpec() (*Animal, *sqlgraph.CreateSpec) {
	var (
		_node = &Animal{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(animal.Table, sqlgraph.NewFieldSpec(animal.FieldID, field.TypeUint64))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Weight(); ok {
		_spec.SetField(animal.FieldWeight, field.TypeFloat32, value)
		_node.Weight = value
	}
	if value, ok := ac.mutation.Length(); ok {
		_spec.SetField(animal.FieldLength, field.TypeFloat32, value)
		_node.Length = value
	}
	if value, ok := ac.mutation.Height(); ok {
		_spec.SetField(animal.FieldHeight, field.TypeFloat32, value)
		_node.Height = value
	}
	if value, ok := ac.mutation.Gender(); ok {
		_spec.SetField(animal.FieldGender, field.TypeEnum, value)
		_node.Gender = value
	}
	if value, ok := ac.mutation.LifeStatus(); ok {
		_spec.SetField(animal.FieldLifeStatus, field.TypeEnum, value)
		_node.LifeStatus = value
	}
	if value, ok := ac.mutation.ChippingDateTime(); ok {
		_spec.SetField(animal.FieldChippingDateTime, field.TypeTime, value)
		_node.ChippingDateTime = value
	}
	if value, ok := ac.mutation.DeathDateTime(); ok {
		_spec.SetField(animal.FieldDeathDateTime, field.TypeTime, value)
		_node.DeathDateTime = &value
	}
	if nodes := ac.mutation.ChipperAnimalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   animal.ChipperAnimalTable,
			Columns: []string{animal.ChipperAnimalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint32,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ChipperID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.AnimalTypeAnimalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   animal.AnimalTypeAnimalTable,
			Columns: animal.AnimalTypeAnimalPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: animaltype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.ChippingLocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   animal.ChippingLocationTable,
			Columns: []string{animal.ChippingLocationColumn},
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
		_node.ChippingLocationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.VisitedLocationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   animal.VisitedLocationsTable,
			Columns: animal.VisitedLocationsPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.AnimalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   animal.AnimalsTable,
			Columns: []string{animal.AnimalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: animalslocations.FieldID,
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

// AnimalCreateBulk is the builder for creating many Animal entities in bulk.
type AnimalCreateBulk struct {
	config
	builders []*AnimalCreate
}

// Save creates the Animal entities in the database.
func (acb *AnimalCreateBulk) Save(ctx context.Context) ([]*Animal, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Animal, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AnimalMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AnimalCreateBulk) SaveX(ctx context.Context) []*Animal {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AnimalCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AnimalCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
