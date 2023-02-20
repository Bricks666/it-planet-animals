// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"animals/ent/migrate"

	"animals/ent/animal"
	"animals/ent/animaltype"
	"animals/ent/location"
	"animals/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Animal is the client for interacting with the Animal builders.
	Animal *AnimalClient
	// AnimalType is the client for interacting with the AnimalType builders.
	AnimalType *AnimalTypeClient
	// Location is the client for interacting with the Location builders.
	Location *LocationClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Animal = NewAnimalClient(c.config)
	c.AnimalType = NewAnimalTypeClient(c.config)
	c.Location = NewLocationClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Animal:     NewAnimalClient(cfg),
		AnimalType: NewAnimalTypeClient(cfg),
		Location:   NewLocationClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Animal:     NewAnimalClient(cfg),
		AnimalType: NewAnimalTypeClient(cfg),
		Location:   NewLocationClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Animal.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Animal.Use(hooks...)
	c.AnimalType.Use(hooks...)
	c.Location.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Animal.Intercept(interceptors...)
	c.AnimalType.Intercept(interceptors...)
	c.Location.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AnimalMutation:
		return c.Animal.mutate(ctx, m)
	case *AnimalTypeMutation:
		return c.AnimalType.mutate(ctx, m)
	case *LocationMutation:
		return c.Location.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// AnimalClient is a client for the Animal schema.
type AnimalClient struct {
	config
}

// NewAnimalClient returns a client for the Animal from the given config.
func NewAnimalClient(c config) *AnimalClient {
	return &AnimalClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `animal.Hooks(f(g(h())))`.
func (c *AnimalClient) Use(hooks ...Hook) {
	c.hooks.Animal = append(c.hooks.Animal, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `animal.Intercept(f(g(h())))`.
func (c *AnimalClient) Intercept(interceptors ...Interceptor) {
	c.inters.Animal = append(c.inters.Animal, interceptors...)
}

// Create returns a builder for creating a Animal entity.
func (c *AnimalClient) Create() *AnimalCreate {
	mutation := newAnimalMutation(c.config, OpCreate)
	return &AnimalCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Animal entities.
func (c *AnimalClient) CreateBulk(builders ...*AnimalCreate) *AnimalCreateBulk {
	return &AnimalCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Animal.
func (c *AnimalClient) Update() *AnimalUpdate {
	mutation := newAnimalMutation(c.config, OpUpdate)
	return &AnimalUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AnimalClient) UpdateOne(a *Animal) *AnimalUpdateOne {
	mutation := newAnimalMutation(c.config, OpUpdateOne, withAnimal(a))
	return &AnimalUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AnimalClient) UpdateOneID(id int) *AnimalUpdateOne {
	mutation := newAnimalMutation(c.config, OpUpdateOne, withAnimalID(id))
	return &AnimalUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Animal.
func (c *AnimalClient) Delete() *AnimalDelete {
	mutation := newAnimalMutation(c.config, OpDelete)
	return &AnimalDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AnimalClient) DeleteOne(a *Animal) *AnimalDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AnimalClient) DeleteOneID(id int) *AnimalDeleteOne {
	builder := c.Delete().Where(animal.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AnimalDeleteOne{builder}
}

// Query returns a query builder for Animal.
func (c *AnimalClient) Query() *AnimalQuery {
	return &AnimalQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAnimal},
		inters: c.Interceptors(),
	}
}

// Get returns a Animal entity by its id.
func (c *AnimalClient) Get(ctx context.Context, id int) (*Animal, error) {
	return c.Query().Where(animal.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AnimalClient) GetX(ctx context.Context, id int) *Animal {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AnimalClient) Hooks() []Hook {
	return c.hooks.Animal
}

// Interceptors returns the client interceptors.
func (c *AnimalClient) Interceptors() []Interceptor {
	return c.inters.Animal
}

func (c *AnimalClient) mutate(ctx context.Context, m *AnimalMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AnimalCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AnimalUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AnimalUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AnimalDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Animal mutation op: %q", m.Op())
	}
}

// AnimalTypeClient is a client for the AnimalType schema.
type AnimalTypeClient struct {
	config
}

// NewAnimalTypeClient returns a client for the AnimalType from the given config.
func NewAnimalTypeClient(c config) *AnimalTypeClient {
	return &AnimalTypeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `animaltype.Hooks(f(g(h())))`.
func (c *AnimalTypeClient) Use(hooks ...Hook) {
	c.hooks.AnimalType = append(c.hooks.AnimalType, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `animaltype.Intercept(f(g(h())))`.
func (c *AnimalTypeClient) Intercept(interceptors ...Interceptor) {
	c.inters.AnimalType = append(c.inters.AnimalType, interceptors...)
}

// Create returns a builder for creating a AnimalType entity.
func (c *AnimalTypeClient) Create() *AnimalTypeCreate {
	mutation := newAnimalTypeMutation(c.config, OpCreate)
	return &AnimalTypeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AnimalType entities.
func (c *AnimalTypeClient) CreateBulk(builders ...*AnimalTypeCreate) *AnimalTypeCreateBulk {
	return &AnimalTypeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AnimalType.
func (c *AnimalTypeClient) Update() *AnimalTypeUpdate {
	mutation := newAnimalTypeMutation(c.config, OpUpdate)
	return &AnimalTypeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AnimalTypeClient) UpdateOne(at *AnimalType) *AnimalTypeUpdateOne {
	mutation := newAnimalTypeMutation(c.config, OpUpdateOne, withAnimalType(at))
	return &AnimalTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AnimalTypeClient) UpdateOneID(id uint64) *AnimalTypeUpdateOne {
	mutation := newAnimalTypeMutation(c.config, OpUpdateOne, withAnimalTypeID(id))
	return &AnimalTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AnimalType.
func (c *AnimalTypeClient) Delete() *AnimalTypeDelete {
	mutation := newAnimalTypeMutation(c.config, OpDelete)
	return &AnimalTypeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AnimalTypeClient) DeleteOne(at *AnimalType) *AnimalTypeDeleteOne {
	return c.DeleteOneID(at.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AnimalTypeClient) DeleteOneID(id uint64) *AnimalTypeDeleteOne {
	builder := c.Delete().Where(animaltype.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AnimalTypeDeleteOne{builder}
}

// Query returns a query builder for AnimalType.
func (c *AnimalTypeClient) Query() *AnimalTypeQuery {
	return &AnimalTypeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAnimalType},
		inters: c.Interceptors(),
	}
}

// Get returns a AnimalType entity by its id.
func (c *AnimalTypeClient) Get(ctx context.Context, id uint64) (*AnimalType, error) {
	return c.Query().Where(animaltype.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AnimalTypeClient) GetX(ctx context.Context, id uint64) *AnimalType {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AnimalTypeClient) Hooks() []Hook {
	return c.hooks.AnimalType
}

// Interceptors returns the client interceptors.
func (c *AnimalTypeClient) Interceptors() []Interceptor {
	return c.inters.AnimalType
}

func (c *AnimalTypeClient) mutate(ctx context.Context, m *AnimalTypeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AnimalTypeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AnimalTypeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AnimalTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AnimalTypeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown AnimalType mutation op: %q", m.Op())
	}
}

// LocationClient is a client for the Location schema.
type LocationClient struct {
	config
}

// NewLocationClient returns a client for the Location from the given config.
func NewLocationClient(c config) *LocationClient {
	return &LocationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `location.Hooks(f(g(h())))`.
func (c *LocationClient) Use(hooks ...Hook) {
	c.hooks.Location = append(c.hooks.Location, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `location.Intercept(f(g(h())))`.
func (c *LocationClient) Intercept(interceptors ...Interceptor) {
	c.inters.Location = append(c.inters.Location, interceptors...)
}

// Create returns a builder for creating a Location entity.
func (c *LocationClient) Create() *LocationCreate {
	mutation := newLocationMutation(c.config, OpCreate)
	return &LocationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Location entities.
func (c *LocationClient) CreateBulk(builders ...*LocationCreate) *LocationCreateBulk {
	return &LocationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Location.
func (c *LocationClient) Update() *LocationUpdate {
	mutation := newLocationMutation(c.config, OpUpdate)
	return &LocationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LocationClient) UpdateOne(l *Location) *LocationUpdateOne {
	mutation := newLocationMutation(c.config, OpUpdateOne, withLocation(l))
	return &LocationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LocationClient) UpdateOneID(id uint64) *LocationUpdateOne {
	mutation := newLocationMutation(c.config, OpUpdateOne, withLocationID(id))
	return &LocationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Location.
func (c *LocationClient) Delete() *LocationDelete {
	mutation := newLocationMutation(c.config, OpDelete)
	return &LocationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LocationClient) DeleteOne(l *Location) *LocationDeleteOne {
	return c.DeleteOneID(l.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *LocationClient) DeleteOneID(id uint64) *LocationDeleteOne {
	builder := c.Delete().Where(location.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LocationDeleteOne{builder}
}

// Query returns a query builder for Location.
func (c *LocationClient) Query() *LocationQuery {
	return &LocationQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeLocation},
		inters: c.Interceptors(),
	}
}

// Get returns a Location entity by its id.
func (c *LocationClient) Get(ctx context.Context, id uint64) (*Location, error) {
	return c.Query().Where(location.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LocationClient) GetX(ctx context.Context, id uint64) *Location {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *LocationClient) Hooks() []Hook {
	return c.hooks.Location
}

// Interceptors returns the client interceptors.
func (c *LocationClient) Interceptors() []Interceptor {
	return c.inters.Location
}

func (c *LocationClient) mutate(ctx context.Context, m *LocationMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&LocationCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&LocationUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&LocationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&LocationDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Location mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uint32) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uint32) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uint32) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uint32) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}
