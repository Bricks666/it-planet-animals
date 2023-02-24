// Code generated by ent, DO NOT EDIT.

package ent

import (
	"animals/ent/animal"
	"animals/ent/location"
	"animals/ent/predicate"
	"animals/ent/visitedlocation"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LocationQuery is the builder for querying Location entities.
type LocationQuery struct {
	config
	ctx                *QueryContext
	order              []OrderFunc
	inters             []Interceptor
	predicates         []predicate.Location
	withChippedAnimals *AnimalQuery
	withAnimals        *AnimalQuery
	withHavingAnimals  *VisitedLocationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LocationQuery builder.
func (lq *LocationQuery) Where(ps ...predicate.Location) *LocationQuery {
	lq.predicates = append(lq.predicates, ps...)
	return lq
}

// Limit the number of records to be returned by this query.
func (lq *LocationQuery) Limit(limit int) *LocationQuery {
	lq.ctx.Limit = &limit
	return lq
}

// Offset to start from.
func (lq *LocationQuery) Offset(offset int) *LocationQuery {
	lq.ctx.Offset = &offset
	return lq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lq *LocationQuery) Unique(unique bool) *LocationQuery {
	lq.ctx.Unique = &unique
	return lq
}

// Order specifies how the records should be ordered.
func (lq *LocationQuery) Order(o ...OrderFunc) *LocationQuery {
	lq.order = append(lq.order, o...)
	return lq
}

// QueryChippedAnimals chains the current query on the "chipped_animals" edge.
func (lq *LocationQuery) QueryChippedAnimals() *AnimalQuery {
	query := (&AnimalClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(location.Table, location.FieldID, selector),
			sqlgraph.To(animal.Table, animal.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, location.ChippedAnimalsTable, location.ChippedAnimalsColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAnimals chains the current query on the "animals" edge.
func (lq *LocationQuery) QueryAnimals() *AnimalQuery {
	query := (&AnimalClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(location.Table, location.FieldID, selector),
			sqlgraph.To(animal.Table, animal.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, location.AnimalsTable, location.AnimalsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryHavingAnimals chains the current query on the "having_animals" edge.
func (lq *LocationQuery) QueryHavingAnimals() *VisitedLocationQuery {
	query := (&VisitedLocationClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(location.Table, location.FieldID, selector),
			sqlgraph.To(visitedlocation.Table, visitedlocation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, location.HavingAnimalsTable, location.HavingAnimalsColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Location entity from the query.
// Returns a *NotFoundError when no Location was found.
func (lq *LocationQuery) First(ctx context.Context) (*Location, error) {
	nodes, err := lq.Limit(1).All(setContextOp(ctx, lq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{location.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lq *LocationQuery) FirstX(ctx context.Context) *Location {
	node, err := lq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Location ID from the query.
// Returns a *NotFoundError when no Location ID was found.
func (lq *LocationQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = lq.Limit(1).IDs(setContextOp(ctx, lq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{location.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lq *LocationQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := lq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Location entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Location entity is found.
// Returns a *NotFoundError when no Location entities are found.
func (lq *LocationQuery) Only(ctx context.Context) (*Location, error) {
	nodes, err := lq.Limit(2).All(setContextOp(ctx, lq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{location.Label}
	default:
		return nil, &NotSingularError{location.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lq *LocationQuery) OnlyX(ctx context.Context) *Location {
	node, err := lq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Location ID in the query.
// Returns a *NotSingularError when more than one Location ID is found.
// Returns a *NotFoundError when no entities are found.
func (lq *LocationQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = lq.Limit(2).IDs(setContextOp(ctx, lq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{location.Label}
	default:
		err = &NotSingularError{location.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lq *LocationQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := lq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Locations.
func (lq *LocationQuery) All(ctx context.Context) ([]*Location, error) {
	ctx = setContextOp(ctx, lq.ctx, "All")
	if err := lq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Location, *LocationQuery]()
	return withInterceptors[[]*Location](ctx, lq, qr, lq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lq *LocationQuery) AllX(ctx context.Context) []*Location {
	nodes, err := lq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Location IDs.
func (lq *LocationQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if lq.ctx.Unique == nil && lq.path != nil {
		lq.Unique(true)
	}
	ctx = setContextOp(ctx, lq.ctx, "IDs")
	if err = lq.Select(location.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lq *LocationQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := lq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lq *LocationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lq.ctx, "Count")
	if err := lq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lq, querierCount[*LocationQuery](), lq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lq *LocationQuery) CountX(ctx context.Context) int {
	count, err := lq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lq *LocationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lq.ctx, "Exist")
	switch _, err := lq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lq *LocationQuery) ExistX(ctx context.Context) bool {
	exist, err := lq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LocationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lq *LocationQuery) Clone() *LocationQuery {
	if lq == nil {
		return nil
	}
	return &LocationQuery{
		config:             lq.config,
		ctx:                lq.ctx.Clone(),
		order:              append([]OrderFunc{}, lq.order...),
		inters:             append([]Interceptor{}, lq.inters...),
		predicates:         append([]predicate.Location{}, lq.predicates...),
		withChippedAnimals: lq.withChippedAnimals.Clone(),
		withAnimals:        lq.withAnimals.Clone(),
		withHavingAnimals:  lq.withHavingAnimals.Clone(),
		// clone intermediate query.
		sql:  lq.sql.Clone(),
		path: lq.path,
	}
}

// WithChippedAnimals tells the query-builder to eager-load the nodes that are connected to
// the "chipped_animals" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LocationQuery) WithChippedAnimals(opts ...func(*AnimalQuery)) *LocationQuery {
	query := (&AnimalClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withChippedAnimals = query
	return lq
}

// WithAnimals tells the query-builder to eager-load the nodes that are connected to
// the "animals" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LocationQuery) WithAnimals(opts ...func(*AnimalQuery)) *LocationQuery {
	query := (&AnimalClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withAnimals = query
	return lq
}

// WithHavingAnimals tells the query-builder to eager-load the nodes that are connected to
// the "having_animals" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LocationQuery) WithHavingAnimals(opts ...func(*VisitedLocationQuery)) *LocationQuery {
	query := (&VisitedLocationClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withHavingAnimals = query
	return lq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Latitude float64 `json:"latitude,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Location.Query().
//		GroupBy(location.FieldLatitude).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lq *LocationQuery) GroupBy(field string, fields ...string) *LocationGroupBy {
	lq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LocationGroupBy{build: lq}
	grbuild.flds = &lq.ctx.Fields
	grbuild.label = location.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Latitude float64 `json:"latitude,omitempty"`
//	}
//
//	client.Location.Query().
//		Select(location.FieldLatitude).
//		Scan(ctx, &v)
func (lq *LocationQuery) Select(fields ...string) *LocationSelect {
	lq.ctx.Fields = append(lq.ctx.Fields, fields...)
	sbuild := &LocationSelect{LocationQuery: lq}
	sbuild.label = location.Label
	sbuild.flds, sbuild.scan = &lq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LocationSelect configured with the given aggregations.
func (lq *LocationQuery) Aggregate(fns ...AggregateFunc) *LocationSelect {
	return lq.Select().Aggregate(fns...)
}

func (lq *LocationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lq); err != nil {
				return err
			}
		}
	}
	for _, f := range lq.ctx.Fields {
		if !location.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lq.path != nil {
		prev, err := lq.path(ctx)
		if err != nil {
			return err
		}
		lq.sql = prev
	}
	return nil
}

func (lq *LocationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Location, error) {
	var (
		nodes       = []*Location{}
		_spec       = lq.querySpec()
		loadedTypes = [3]bool{
			lq.withChippedAnimals != nil,
			lq.withAnimals != nil,
			lq.withHavingAnimals != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Location).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Location{config: lq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lq.withChippedAnimals; query != nil {
		if err := lq.loadChippedAnimals(ctx, query, nodes,
			func(n *Location) { n.Edges.ChippedAnimals = []*Animal{} },
			func(n *Location, e *Animal) { n.Edges.ChippedAnimals = append(n.Edges.ChippedAnimals, e) }); err != nil {
			return nil, err
		}
	}
	if query := lq.withAnimals; query != nil {
		if err := lq.loadAnimals(ctx, query, nodes,
			func(n *Location) { n.Edges.Animals = []*Animal{} },
			func(n *Location, e *Animal) { n.Edges.Animals = append(n.Edges.Animals, e) }); err != nil {
			return nil, err
		}
	}
	if query := lq.withHavingAnimals; query != nil {
		if err := lq.loadHavingAnimals(ctx, query, nodes,
			func(n *Location) { n.Edges.HavingAnimals = []*VisitedLocation{} },
			func(n *Location, e *VisitedLocation) { n.Edges.HavingAnimals = append(n.Edges.HavingAnimals, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lq *LocationQuery) loadChippedAnimals(ctx context.Context, query *AnimalQuery, nodes []*Location, init func(*Location), assign func(*Location, *Animal)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*Location)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Animal(func(s *sql.Selector) {
		s.Where(sql.InValues(location.ChippedAnimalsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ChippingLocationID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "chipping_location_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (lq *LocationQuery) loadAnimals(ctx context.Context, query *AnimalQuery, nodes []*Location, init func(*Location), assign func(*Location, *Animal)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uint64]*Location)
	nids := make(map[uint64]map[*Location]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(location.AnimalsTable)
		s.Join(joinT).On(s.C(animal.FieldID), joinT.C(location.AnimalsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(location.AnimalsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(location.AnimalsPrimaryKey[1]))
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
					nids[inValue] = map[*Location]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Animal](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "animals" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (lq *LocationQuery) loadHavingAnimals(ctx context.Context, query *VisitedLocationQuery, nodes []*Location, init func(*Location), assign func(*Location, *VisitedLocation)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*Location)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.VisitedLocation(func(s *sql.Selector) {
		s.Where(sql.InValues(location.HavingAnimalsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.LocationID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "location_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (lq *LocationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lq.querySpec()
	_spec.Node.Columns = lq.ctx.Fields
	if len(lq.ctx.Fields) > 0 {
		_spec.Unique = lq.ctx.Unique != nil && *lq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lq.driver, _spec)
}

func (lq *LocationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(location.Table, location.Columns, sqlgraph.NewFieldSpec(location.FieldID, field.TypeUint64))
	_spec.From = lq.sql
	if unique := lq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lq.path != nil {
		_spec.Unique = true
	}
	if fields := lq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, location.FieldID)
		for i := range fields {
			if fields[i] != location.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lq *LocationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lq.driver.Dialect())
	t1 := builder.Table(location.Table)
	columns := lq.ctx.Fields
	if len(columns) == 0 {
		columns = location.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lq.sql != nil {
		selector = lq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lq.ctx.Unique != nil && *lq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lq.predicates {
		p(selector)
	}
	for _, p := range lq.order {
		p(selector)
	}
	if offset := lq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LocationGroupBy is the group-by builder for Location entities.
type LocationGroupBy struct {
	selector
	build *LocationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgb *LocationGroupBy) Aggregate(fns ...AggregateFunc) *LocationGroupBy {
	lgb.fns = append(lgb.fns, fns...)
	return lgb
}

// Scan applies the selector query and scans the result into the given value.
func (lgb *LocationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lgb.build.ctx, "GroupBy")
	if err := lgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LocationQuery, *LocationGroupBy](ctx, lgb.build, lgb, lgb.build.inters, v)
}

func (lgb *LocationGroupBy) sqlScan(ctx context.Context, root *LocationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lgb.fns))
	for _, fn := range lgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lgb.flds)+len(lgb.fns))
		for _, f := range *lgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LocationSelect is the builder for selecting fields of Location entities.
type LocationSelect struct {
	*LocationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ls *LocationSelect) Aggregate(fns ...AggregateFunc) *LocationSelect {
	ls.fns = append(ls.fns, fns...)
	return ls
}

// Scan applies the selector query and scans the result into the given value.
func (ls *LocationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ls.ctx, "Select")
	if err := ls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LocationQuery, *LocationSelect](ctx, ls.LocationQuery, ls, ls.inters, v)
}

func (ls *LocationSelect) sqlScan(ctx context.Context, root *LocationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ls.fns))
	for _, fn := range ls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
