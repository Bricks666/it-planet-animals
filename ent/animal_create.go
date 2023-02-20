// Code generated by ent, DO NOT EDIT.

package ent

import (
	"animals/ent/animal"
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

// SetLifestatus sets the "lifestatus" field.
func (ac *AnimalCreate) SetLifestatus(a animal.Lifestatus) *AnimalCreate {
	ac.mutation.SetLifestatus(a)
	return ac
}

// SetNillableLifestatus sets the "lifestatus" field if the given value is not nil.
func (ac *AnimalCreate) SetNillableLifestatus(a *animal.Lifestatus) *AnimalCreate {
	if a != nil {
		ac.SetLifestatus(*a)
	}
	return ac
}

// SetChippingDateTime sets the "chippingDateTime" field.
func (ac *AnimalCreate) SetChippingDateTime(t time.Time) *AnimalCreate {
	ac.mutation.SetChippingDateTime(t)
	return ac
}

// SetNillableChippingDateTime sets the "chippingDateTime" field if the given value is not nil.
func (ac *AnimalCreate) SetNillableChippingDateTime(t *time.Time) *AnimalCreate {
	if t != nil {
		ac.SetChippingDateTime(*t)
	}
	return ac
}

// SetChipperId sets the "chipperId" field.
func (ac *AnimalCreate) SetChipperId(u uint32) *AnimalCreate {
	ac.mutation.SetChipperId(u)
	return ac
}

// SetNillableChipperId sets the "chipperId" field if the given value is not nil.
func (ac *AnimalCreate) SetNillableChipperId(u *uint32) *AnimalCreate {
	if u != nil {
		ac.SetChipperId(*u)
	}
	return ac
}

// SetChippingLocationId sets the "chippingLocationId" field.
func (ac *AnimalCreate) SetChippingLocationId(u uint64) *AnimalCreate {
	ac.mutation.SetChippingLocationId(u)
	return ac
}

// SetNillableChippingLocationId sets the "chippingLocationId" field if the given value is not nil.
func (ac *AnimalCreate) SetNillableChippingLocationId(u *uint64) *AnimalCreate {
	if u != nil {
		ac.SetChippingLocationId(*u)
	}
	return ac
}

// SetDeathDateTime sets the "deathDateTime" field.
func (ac *AnimalCreate) SetDeathDateTime(t time.Time) *AnimalCreate {
	ac.mutation.SetDeathDateTime(t)
	return ac
}

// SetNillableDeathDateTime sets the "deathDateTime" field if the given value is not nil.
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

// SetUserAnimalsID sets the "user_animals" edge to the User entity by ID.
func (ac *AnimalCreate) SetUserAnimalsID(id uint32) *AnimalCreate {
	ac.mutation.SetUserAnimalsID(id)
	return ac
}

// SetNillableUserAnimalsID sets the "user_animals" edge to the User entity by ID if the given value is not nil.
func (ac *AnimalCreate) SetNillableUserAnimalsID(id *uint32) *AnimalCreate {
	if id != nil {
		ac = ac.SetUserAnimalsID(*id)
	}
	return ac
}

// SetUserAnimals sets the "user_animals" edge to the User entity.
func (ac *AnimalCreate) SetUserAnimals(u *User) *AnimalCreate {
	return ac.SetUserAnimalsID(u.ID)
}

// AddAnimalTagsAnimalIDs adds the "animal_tags_animals" edge to the AnimalType entity by IDs.
func (ac *AnimalCreate) AddAnimalTagsAnimalIDs(ids ...uint64) *AnimalCreate {
	ac.mutation.AddAnimalTagsAnimalIDs(ids...)
	return ac
}

// AddAnimalTagsAnimals adds the "animal_tags_animals" edges to the AnimalType entity.
func (ac *AnimalCreate) AddAnimalTagsAnimals(a ...*AnimalType) *AnimalCreate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddAnimalTagsAnimalIDs(ids...)
}

// SetChippingLocationID sets the "chipping_location" edge to the Location entity by ID.
func (ac *AnimalCreate) SetChippingLocationID(id uint64) *AnimalCreate {
	ac.mutation.SetChippingLocationID(id)
	return ac
}

// SetNillableChippingLocationID sets the "chipping_location" edge to the Location entity by ID if the given value is not nil.
func (ac *AnimalCreate) SetNillableChippingLocationID(id *uint64) *AnimalCreate {
	if id != nil {
		ac = ac.SetChippingLocationID(*id)
	}
	return ac
}

// SetChippingLocation sets the "chipping_location" edge to the Location entity.
func (ac *AnimalCreate) SetChippingLocation(l *Location) *AnimalCreate {
	return ac.SetChippingLocationID(l.ID)
}

// AddVisitedLocationsAnimalIDs adds the "visited_locations_animals" edge to the Location entity by IDs.
func (ac *AnimalCreate) AddVisitedLocationsAnimalIDs(ids ...uint64) *AnimalCreate {
	ac.mutation.AddVisitedLocationsAnimalIDs(ids...)
	return ac
}

// AddVisitedLocationsAnimals adds the "visited_locations_animals" edges to the Location entity.
func (ac *AnimalCreate) AddVisitedLocationsAnimals(l ...*Location) *AnimalCreate {
	ids := make([]uint64, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ac.AddVisitedLocationsAnimalIDs(ids...)
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
	if _, ok := ac.mutation.Lifestatus(); !ok {
		v := animal.DefaultLifestatus
		ac.mutation.SetLifestatus(v)
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
	if _, ok := ac.mutation.Lifestatus(); !ok {
		return &ValidationError{Name: "lifestatus", err: errors.New(`ent: missing required field "Animal.lifestatus"`)}
	}
	if v, ok := ac.mutation.Lifestatus(); ok {
		if err := animal.LifestatusValidator(v); err != nil {
			return &ValidationError{Name: "lifestatus", err: fmt.Errorf(`ent: validator failed for field "Animal.lifestatus": %w`, err)}
		}
	}
	if _, ok := ac.mutation.ChippingDateTime(); !ok {
		return &ValidationError{Name: "chippingDateTime", err: errors.New(`ent: missing required field "Animal.chippingDateTime"`)}
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
	if value, ok := ac.mutation.Lifestatus(); ok {
		_spec.SetField(animal.FieldLifestatus, field.TypeEnum, value)
		_node.Lifestatus = value
	}
	if value, ok := ac.mutation.ChippingDateTime(); ok {
		_spec.SetField(animal.FieldChippingDateTime, field.TypeTime, value)
		_node.ChippingDateTime = value
	}
	if value, ok := ac.mutation.DeathDateTime(); ok {
		_spec.SetField(animal.FieldDeathDateTime, field.TypeTime, value)
		_node.DeathDateTime = &value
	}
	if nodes := ac.mutation.UserAnimalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   animal.UserAnimalsTable,
			Columns: []string{animal.UserAnimalsColumn},
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
		_node.ChipperId = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.AnimalTagsAnimalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   animal.AnimalTagsAnimalsTable,
			Columns: animal.AnimalTagsAnimalsPrimaryKey,
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
		_node.ChippingLocationId = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.VisitedLocationsAnimalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   animal.VisitedLocationsAnimalsTable,
			Columns: animal.VisitedLocationsAnimalsPrimaryKey,
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
