// Code generated by ent, DO NOT EDIT.

package ent

import (
	"animals/ent/animal"
	"animals/ent/animalslocations"
	"animals/ent/animaltype"
	"animals/ent/location"
	"animals/ent/predicate"
	"animals/ent/user"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AnimalQuery is the builder for querying Animal entities.
type AnimalQuery struct {
	config
	ctx                  *QueryContext
	order                []OrderFunc
	inters               []Interceptor
	predicates           []predicate.Animal
	withChipperAnimal    *UserQuery
	withAnimalTypeAnimal *AnimalTypeQuery
	withChippingLocation *LocationQuery
	withVisitedLocations *LocationQuery
	withAnimals          *AnimalsLocationsQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AnimalQuery builder.
func (aq *AnimalQuery) Where(ps ...predicate.Animal) *AnimalQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *AnimalQuery) Limit(limit int) *AnimalQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *AnimalQuery) Offset(offset int) *AnimalQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AnimalQuery) Unique(unique bool) *AnimalQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *AnimalQuery) Order(o ...OrderFunc) *AnimalQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryChipperAnimal chains the current query on the "chipper_animal" edge.
func (aq *AnimalQuery) QueryChipperAnimal() *UserQuery {
	query := (&UserClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(animal.Table, animal.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, animal.ChipperAnimalTable, animal.ChipperAnimalColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAnimalTypeAnimal chains the current query on the "animal_type_animal" edge.
func (aq *AnimalQuery) QueryAnimalTypeAnimal() *AnimalTypeQuery {
	query := (&AnimalTypeClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(animal.Table, animal.FieldID, selector),
			sqlgraph.To(animaltype.Table, animaltype.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, animal.AnimalTypeAnimalTable, animal.AnimalTypeAnimalPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChippingLocation chains the current query on the "chipping_location" edge.
func (aq *AnimalQuery) QueryChippingLocation() *LocationQuery {
	query := (&LocationClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(animal.Table, animal.FieldID, selector),
			sqlgraph.To(location.Table, location.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, animal.ChippingLocationTable, animal.ChippingLocationColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVisitedLocations chains the current query on the "visited_locations" edge.
func (aq *AnimalQuery) QueryVisitedLocations() *LocationQuery {
	query := (&LocationClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(animal.Table, animal.FieldID, selector),
			sqlgraph.To(location.Table, location.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, animal.VisitedLocationsTable, animal.VisitedLocationsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAnimals chains the current query on the "animals" edge.
func (aq *AnimalQuery) QueryAnimals() *AnimalsLocationsQuery {
	query := (&AnimalsLocationsClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(animal.Table, animal.FieldID, selector),
			sqlgraph.To(animalslocations.Table, animalslocations.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, animal.AnimalsTable, animal.AnimalsColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Animal entity from the query.
// Returns a *NotFoundError when no Animal was found.
func (aq *AnimalQuery) First(ctx context.Context) (*Animal, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{animal.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AnimalQuery) FirstX(ctx context.Context) *Animal {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Animal ID from the query.
// Returns a *NotFoundError when no Animal ID was found.
func (aq *AnimalQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{animal.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AnimalQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Animal entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Animal entity is found.
// Returns a *NotFoundError when no Animal entities are found.
func (aq *AnimalQuery) Only(ctx context.Context) (*Animal, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{animal.Label}
	default:
		return nil, &NotSingularError{animal.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AnimalQuery) OnlyX(ctx context.Context) *Animal {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Animal ID in the query.
// Returns a *NotSingularError when more than one Animal ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AnimalQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{animal.Label}
	default:
		err = &NotSingularError{animal.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AnimalQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Animals.
func (aq *AnimalQuery) All(ctx context.Context) ([]*Animal, error) {
	ctx = setContextOp(ctx, aq.ctx, "All")
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Animal, *AnimalQuery]()
	return withInterceptors[[]*Animal](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *AnimalQuery) AllX(ctx context.Context) []*Animal {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Animal IDs.
func (aq *AnimalQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, "IDs")
	if err = aq.Select(animal.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AnimalQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AnimalQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, "Count")
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*AnimalQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AnimalQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AnimalQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, "Exist")
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AnimalQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AnimalQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AnimalQuery) Clone() *AnimalQuery {
	if aq == nil {
		return nil
	}
	return &AnimalQuery{
		config:               aq.config,
		ctx:                  aq.ctx.Clone(),
		order:                append([]OrderFunc{}, aq.order...),
		inters:               append([]Interceptor{}, aq.inters...),
		predicates:           append([]predicate.Animal{}, aq.predicates...),
		withChipperAnimal:    aq.withChipperAnimal.Clone(),
		withAnimalTypeAnimal: aq.withAnimalTypeAnimal.Clone(),
		withChippingLocation: aq.withChippingLocation.Clone(),
		withVisitedLocations: aq.withVisitedLocations.Clone(),
		withAnimals:          aq.withAnimals.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithChipperAnimal tells the query-builder to eager-load the nodes that are connected to
// the "chipper_animal" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AnimalQuery) WithChipperAnimal(opts ...func(*UserQuery)) *AnimalQuery {
	query := (&UserClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withChipperAnimal = query
	return aq
}

// WithAnimalTypeAnimal tells the query-builder to eager-load the nodes that are connected to
// the "animal_type_animal" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AnimalQuery) WithAnimalTypeAnimal(opts ...func(*AnimalTypeQuery)) *AnimalQuery {
	query := (&AnimalTypeClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withAnimalTypeAnimal = query
	return aq
}

// WithChippingLocation tells the query-builder to eager-load the nodes that are connected to
// the "chipping_location" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AnimalQuery) WithChippingLocation(opts ...func(*LocationQuery)) *AnimalQuery {
	query := (&LocationClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withChippingLocation = query
	return aq
}

// WithVisitedLocations tells the query-builder to eager-load the nodes that are connected to
// the "visited_locations" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AnimalQuery) WithVisitedLocations(opts ...func(*LocationQuery)) *AnimalQuery {
	query := (&LocationClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withVisitedLocations = query
	return aq
}

// WithAnimals tells the query-builder to eager-load the nodes that are connected to
// the "animals" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AnimalQuery) WithAnimals(opts ...func(*AnimalsLocationsQuery)) *AnimalQuery {
	query := (&AnimalsLocationsClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withAnimals = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Weight float32 `json:"weight,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Animal.Query().
//		GroupBy(animal.FieldWeight).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AnimalQuery) GroupBy(field string, fields ...string) *AnimalGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AnimalGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = animal.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Weight float32 `json:"weight,omitempty"`
//	}
//
//	client.Animal.Query().
//		Select(animal.FieldWeight).
//		Scan(ctx, &v)
func (aq *AnimalQuery) Select(fields ...string) *AnimalSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &AnimalSelect{AnimalQuery: aq}
	sbuild.label = animal.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AnimalSelect configured with the given aggregations.
func (aq *AnimalQuery) Aggregate(fns ...AggregateFunc) *AnimalSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AnimalQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !animal.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AnimalQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Animal, error) {
	var (
		nodes       = []*Animal{}
		_spec       = aq.querySpec()
		loadedTypes = [5]bool{
			aq.withChipperAnimal != nil,
			aq.withAnimalTypeAnimal != nil,
			aq.withChippingLocation != nil,
			aq.withVisitedLocations != nil,
			aq.withAnimals != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Animal).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Animal{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withChipperAnimal; query != nil {
		if err := aq.loadChipperAnimal(ctx, query, nodes, nil,
			func(n *Animal, e *User) { n.Edges.ChipperAnimal = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withAnimalTypeAnimal; query != nil {
		if err := aq.loadAnimalTypeAnimal(ctx, query, nodes,
			func(n *Animal) { n.Edges.AnimalTypeAnimal = []*AnimalType{} },
			func(n *Animal, e *AnimalType) { n.Edges.AnimalTypeAnimal = append(n.Edges.AnimalTypeAnimal, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withChippingLocation; query != nil {
		if err := aq.loadChippingLocation(ctx, query, nodes, nil,
			func(n *Animal, e *Location) { n.Edges.ChippingLocation = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withVisitedLocations; query != nil {
		if err := aq.loadVisitedLocations(ctx, query, nodes,
			func(n *Animal) { n.Edges.VisitedLocations = []*Location{} },
			func(n *Animal, e *Location) { n.Edges.VisitedLocations = append(n.Edges.VisitedLocations, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withAnimals; query != nil {
		if err := aq.loadAnimals(ctx, query, nodes,
			func(n *Animal) { n.Edges.Animals = []*AnimalsLocations{} },
			func(n *Animal, e *AnimalsLocations) { n.Edges.Animals = append(n.Edges.Animals, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AnimalQuery) loadChipperAnimal(ctx context.Context, query *UserQuery, nodes []*Animal, init func(*Animal), assign func(*Animal, *User)) error {
	ids := make([]uint32, 0, len(nodes))
	nodeids := make(map[uint32][]*Animal)
	for i := range nodes {
		fk := nodes[i].ChipperId
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "chipperId" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AnimalQuery) loadAnimalTypeAnimal(ctx context.Context, query *AnimalTypeQuery, nodes []*Animal, init func(*Animal), assign func(*Animal, *AnimalType)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uint64]*Animal)
	nids := make(map[uint64]map[*Animal]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(animal.AnimalTypeAnimalTable)
		s.Join(joinT).On(s.C(animaltype.FieldID), joinT.C(animal.AnimalTypeAnimalPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(animal.AnimalTypeAnimalPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(animal.AnimalTypeAnimalPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := uint64(values[0].(*sql.NullInt64).Int64)
				inValue := uint64(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Animal]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*AnimalType](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "animal_type_animal" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (aq *AnimalQuery) loadChippingLocation(ctx context.Context, query *LocationQuery, nodes []*Animal, init func(*Animal), assign func(*Animal, *Location)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*Animal)
	for i := range nodes {
		fk := nodes[i].ChippingLocationId
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(location.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "chippingLocationId" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AnimalQuery) loadVisitedLocations(ctx context.Context, query *LocationQuery, nodes []*Animal, init func(*Animal), assign func(*Animal, *Location)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uint64]*Animal)
	nids := make(map[uint64]map[*Animal]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(animal.VisitedLocationsTable)
		s.Join(joinT).On(s.C(location.FieldID), joinT.C(animal.VisitedLocationsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(animal.VisitedLocationsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(animal.VisitedLocationsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := uint64(values[0].(*sql.NullInt64).Int64)
				inValue := uint64(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Animal]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Location](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "visited_locations" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (aq *AnimalQuery) loadAnimals(ctx context.Context, query *AnimalsLocationsQuery, nodes []*Animal, init func(*Animal), assign func(*Animal, *AnimalsLocations)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*Animal)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.AnimalsLocations(func(s *sql.Selector) {
		s.Where(sql.InValues(animal.AnimalsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.AnimalId
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "animalId" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (aq *AnimalQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AnimalQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(animal.Table, animal.Columns, sqlgraph.NewFieldSpec(animal.FieldID, field.TypeUint64))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, animal.FieldID)
		for i := range fields {
			if fields[i] != animal.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AnimalQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(animal.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = animal.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AnimalGroupBy is the group-by builder for Animal entities.
type AnimalGroupBy struct {
	selector
	build *AnimalQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AnimalGroupBy) Aggregate(fns ...AggregateFunc) *AnimalGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *AnimalGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, "GroupBy")
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AnimalQuery, *AnimalGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *AnimalGroupBy) sqlScan(ctx context.Context, root *AnimalQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AnimalSelect is the builder for selecting fields of Animal entities.
type AnimalSelect struct {
	*AnimalQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AnimalSelect) Aggregate(fns ...AggregateFunc) *AnimalSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AnimalSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, "Select")
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AnimalQuery, *AnimalSelect](ctx, as.AnimalQuery, as, as.inters, v)
}

func (as *AnimalSelect) sqlScan(ctx context.Context, root *AnimalQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
