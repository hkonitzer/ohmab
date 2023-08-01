// OHMAB
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hkonitzer/ohmab/ent/address"
	"github.com/hkonitzer/ohmab/ent/business"
	"github.com/hkonitzer/ohmab/ent/predicate"
	"github.com/hkonitzer/ohmab/ent/publicuser"
	"github.com/hkonitzer/ohmab/ent/tag"
	"github.com/hkonitzer/ohmab/ent/user"
)

// BusinessQuery is the builder for querying Business entities.
type BusinessQuery struct {
	config
	ctx                  *QueryContext
	order                []business.OrderOption
	inters               []Interceptor
	predicates           []predicate.Business
	withAddresses        *AddressQuery
	withTags             *TagQuery
	withUsers            *UserQuery
	withPublicUsers      *PublicUserQuery
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*Business) error
	withNamedAddresses   map[string]*AddressQuery
	withNamedTags        map[string]*TagQuery
	withNamedUsers       map[string]*UserQuery
	withNamedPublicUsers map[string]*PublicUserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BusinessQuery builder.
func (bq *BusinessQuery) Where(ps ...predicate.Business) *BusinessQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit the number of records to be returned by this query.
func (bq *BusinessQuery) Limit(limit int) *BusinessQuery {
	bq.ctx.Limit = &limit
	return bq
}

// Offset to start from.
func (bq *BusinessQuery) Offset(offset int) *BusinessQuery {
	bq.ctx.Offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BusinessQuery) Unique(unique bool) *BusinessQuery {
	bq.ctx.Unique = &unique
	return bq
}

// Order specifies how the records should be ordered.
func (bq *BusinessQuery) Order(o ...business.OrderOption) *BusinessQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryAddresses chains the current query on the "addresses" edge.
func (bq *BusinessQuery) QueryAddresses() *AddressQuery {
	query := (&AddressClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(business.Table, business.FieldID, selector),
			sqlgraph.To(address.Table, address.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, business.AddressesTable, business.AddressesColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTags chains the current query on the "tags" edge.
func (bq *BusinessQuery) QueryTags() *TagQuery {
	query := (&TagClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(business.Table, business.FieldID, selector),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, business.TagsTable, business.TagsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUsers chains the current query on the "users" edge.
func (bq *BusinessQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(business.Table, business.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, business.UsersTable, business.UsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPublicUsers chains the current query on the "public_users" edge.
func (bq *BusinessQuery) QueryPublicUsers() *PublicUserQuery {
	query := (&PublicUserClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(business.Table, business.FieldID, selector),
			sqlgraph.To(publicuser.Table, publicuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, business.PublicUsersTable, business.PublicUsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Business entity from the query.
// Returns a *NotFoundError when no Business was found.
func (bq *BusinessQuery) First(ctx context.Context) (*Business, error) {
	nodes, err := bq.Limit(1).All(setContextOp(ctx, bq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{business.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BusinessQuery) FirstX(ctx context.Context) *Business {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Business ID from the query.
// Returns a *NotFoundError when no Business ID was found.
func (bq *BusinessQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = bq.Limit(1).IDs(setContextOp(ctx, bq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{business.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BusinessQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Business entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Business entity is found.
// Returns a *NotFoundError when no Business entities are found.
func (bq *BusinessQuery) Only(ctx context.Context) (*Business, error) {
	nodes, err := bq.Limit(2).All(setContextOp(ctx, bq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{business.Label}
	default:
		return nil, &NotSingularError{business.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BusinessQuery) OnlyX(ctx context.Context) *Business {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Business ID in the query.
// Returns a *NotSingularError when more than one Business ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BusinessQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = bq.Limit(2).IDs(setContextOp(ctx, bq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{business.Label}
	default:
		err = &NotSingularError{business.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BusinessQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Businesses.
func (bq *BusinessQuery) All(ctx context.Context) ([]*Business, error) {
	ctx = setContextOp(ctx, bq.ctx, "All")
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Business, *BusinessQuery]()
	return withInterceptors[[]*Business](ctx, bq, qr, bq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bq *BusinessQuery) AllX(ctx context.Context) []*Business {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Business IDs.
func (bq *BusinessQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if bq.ctx.Unique == nil && bq.path != nil {
		bq.Unique(true)
	}
	ctx = setContextOp(ctx, bq.ctx, "IDs")
	if err = bq.Select(business.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BusinessQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BusinessQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bq.ctx, "Count")
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bq, querierCount[*BusinessQuery](), bq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BusinessQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BusinessQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bq.ctx, "Exist")
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BusinessQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BusinessQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BusinessQuery) Clone() *BusinessQuery {
	if bq == nil {
		return nil
	}
	return &BusinessQuery{
		config:          bq.config,
		ctx:             bq.ctx.Clone(),
		order:           append([]business.OrderOption{}, bq.order...),
		inters:          append([]Interceptor{}, bq.inters...),
		predicates:      append([]predicate.Business{}, bq.predicates...),
		withAddresses:   bq.withAddresses.Clone(),
		withTags:        bq.withTags.Clone(),
		withUsers:       bq.withUsers.Clone(),
		withPublicUsers: bq.withPublicUsers.Clone(),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

// WithAddresses tells the query-builder to eager-load the nodes that are connected to
// the "addresses" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BusinessQuery) WithAddresses(opts ...func(*AddressQuery)) *BusinessQuery {
	query := (&AddressClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withAddresses = query
	return bq
}

// WithTags tells the query-builder to eager-load the nodes that are connected to
// the "tags" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BusinessQuery) WithTags(opts ...func(*TagQuery)) *BusinessQuery {
	query := (&TagClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withTags = query
	return bq
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BusinessQuery) WithUsers(opts ...func(*UserQuery)) *BusinessQuery {
	query := (&UserClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withUsers = query
	return bq
}

// WithPublicUsers tells the query-builder to eager-load the nodes that are connected to
// the "public_users" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BusinessQuery) WithPublicUsers(opts ...func(*PublicUserQuery)) *BusinessQuery {
	query := (&PublicUserClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withPublicUsers = query
	return bq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Business.Query().
//		GroupBy(business.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BusinessQuery) GroupBy(field string, fields ...string) *BusinessGroupBy {
	bq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BusinessGroupBy{build: bq}
	grbuild.flds = &bq.ctx.Fields
	grbuild.label = business.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Business.Query().
//		Select(business.FieldCreatedAt).
//		Scan(ctx, &v)
func (bq *BusinessQuery) Select(fields ...string) *BusinessSelect {
	bq.ctx.Fields = append(bq.ctx.Fields, fields...)
	sbuild := &BusinessSelect{BusinessQuery: bq}
	sbuild.label = business.Label
	sbuild.flds, sbuild.scan = &bq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BusinessSelect configured with the given aggregations.
func (bq *BusinessQuery) Aggregate(fns ...AggregateFunc) *BusinessSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BusinessQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bq); err != nil {
				return err
			}
		}
	}
	for _, f := range bq.ctx.Fields {
		if !business.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	if business.Policy == nil {
		return errors.New("ent: uninitialized business.Policy (forgotten import ent/runtime?)")
	}
	if err := business.Policy.EvalQuery(ctx, bq); err != nil {
		return err
	}
	return nil
}

func (bq *BusinessQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Business, error) {
	var (
		nodes       = []*Business{}
		_spec       = bq.querySpec()
		loadedTypes = [4]bool{
			bq.withAddresses != nil,
			bq.withTags != nil,
			bq.withUsers != nil,
			bq.withPublicUsers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Business).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Business{config: bq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(bq.modifiers) > 0 {
		_spec.Modifiers = bq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bq.withAddresses; query != nil {
		if err := bq.loadAddresses(ctx, query, nodes,
			func(n *Business) { n.Edges.Addresses = []*Address{} },
			func(n *Business, e *Address) { n.Edges.Addresses = append(n.Edges.Addresses, e) }); err != nil {
			return nil, err
		}
	}
	if query := bq.withTags; query != nil {
		if err := bq.loadTags(ctx, query, nodes,
			func(n *Business) { n.Edges.Tags = []*Tag{} },
			func(n *Business, e *Tag) { n.Edges.Tags = append(n.Edges.Tags, e) }); err != nil {
			return nil, err
		}
	}
	if query := bq.withUsers; query != nil {
		if err := bq.loadUsers(ctx, query, nodes,
			func(n *Business) { n.Edges.Users = []*User{} },
			func(n *Business, e *User) { n.Edges.Users = append(n.Edges.Users, e) }); err != nil {
			return nil, err
		}
	}
	if query := bq.withPublicUsers; query != nil {
		if err := bq.loadPublicUsers(ctx, query, nodes,
			func(n *Business) { n.Edges.PublicUsers = []*PublicUser{} },
			func(n *Business, e *PublicUser) { n.Edges.PublicUsers = append(n.Edges.PublicUsers, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range bq.withNamedAddresses {
		if err := bq.loadAddresses(ctx, query, nodes,
			func(n *Business) { n.appendNamedAddresses(name) },
			func(n *Business, e *Address) { n.appendNamedAddresses(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range bq.withNamedTags {
		if err := bq.loadTags(ctx, query, nodes,
			func(n *Business) { n.appendNamedTags(name) },
			func(n *Business, e *Tag) { n.appendNamedTags(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range bq.withNamedUsers {
		if err := bq.loadUsers(ctx, query, nodes,
			func(n *Business) { n.appendNamedUsers(name) },
			func(n *Business, e *User) { n.appendNamedUsers(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range bq.withNamedPublicUsers {
		if err := bq.loadPublicUsers(ctx, query, nodes,
			func(n *Business) { n.appendNamedPublicUsers(name) },
			func(n *Business, e *PublicUser) { n.appendNamedPublicUsers(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range bq.loadTotal {
		if err := bq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bq *BusinessQuery) loadAddresses(ctx context.Context, query *AddressQuery, nodes []*Business, init func(*Business), assign func(*Business, *Address)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Business)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Address(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(business.AddressesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.business_addresses
		if fk == nil {
			return fmt.Errorf(`foreign-key "business_addresses" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "business_addresses" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (bq *BusinessQuery) loadTags(ctx context.Context, query *TagQuery, nodes []*Business, init func(*Business), assign func(*Business, *Tag)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Business)
	nids := make(map[uuid.UUID]map[*Business]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(business.TagsTable)
		s.Join(joinT).On(s.C(tag.FieldID), joinT.C(business.TagsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(business.TagsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(business.TagsPrimaryKey[0]))
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
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Business]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Tag](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "tags" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (bq *BusinessQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*Business, init func(*Business), assign func(*Business, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Business)
	nids := make(map[uuid.UUID]map[*Business]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(business.UsersTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(business.UsersPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(business.UsersPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(business.UsersPrimaryKey[0]))
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
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Business]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "users" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (bq *BusinessQuery) loadPublicUsers(ctx context.Context, query *PublicUserQuery, nodes []*Business, init func(*Business), assign func(*Business, *PublicUser)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Business)
	nids := make(map[uuid.UUID]map[*Business]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(business.PublicUsersTable)
		s.Join(joinT).On(s.C(publicuser.FieldID), joinT.C(business.PublicUsersPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(business.PublicUsersPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(business.PublicUsersPrimaryKey[0]))
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
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Business]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*PublicUser](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "public_users" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (bq *BusinessQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	if len(bq.modifiers) > 0 {
		_spec.Modifiers = bq.modifiers
	}
	_spec.Node.Columns = bq.ctx.Fields
	if len(bq.ctx.Fields) > 0 {
		_spec.Unique = bq.ctx.Unique != nil && *bq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BusinessQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(business.Table, business.Columns, sqlgraph.NewFieldSpec(business.FieldID, field.TypeUUID))
	_spec.From = bq.sql
	if unique := bq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bq.path != nil {
		_spec.Unique = true
	}
	if fields := bq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, business.FieldID)
		for i := range fields {
			if fields[i] != business.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BusinessQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(business.Table)
	columns := bq.ctx.Fields
	if len(columns) == 0 {
		columns = business.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.ctx.Unique != nil && *bq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedAddresses tells the query-builder to eager-load the nodes that are connected to the "addresses"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (bq *BusinessQuery) WithNamedAddresses(name string, opts ...func(*AddressQuery)) *BusinessQuery {
	query := (&AddressClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if bq.withNamedAddresses == nil {
		bq.withNamedAddresses = make(map[string]*AddressQuery)
	}
	bq.withNamedAddresses[name] = query
	return bq
}

// WithNamedTags tells the query-builder to eager-load the nodes that are connected to the "tags"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (bq *BusinessQuery) WithNamedTags(name string, opts ...func(*TagQuery)) *BusinessQuery {
	query := (&TagClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if bq.withNamedTags == nil {
		bq.withNamedTags = make(map[string]*TagQuery)
	}
	bq.withNamedTags[name] = query
	return bq
}

// WithNamedUsers tells the query-builder to eager-load the nodes that are connected to the "users"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (bq *BusinessQuery) WithNamedUsers(name string, opts ...func(*UserQuery)) *BusinessQuery {
	query := (&UserClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if bq.withNamedUsers == nil {
		bq.withNamedUsers = make(map[string]*UserQuery)
	}
	bq.withNamedUsers[name] = query
	return bq
}

// WithNamedPublicUsers tells the query-builder to eager-load the nodes that are connected to the "public_users"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (bq *BusinessQuery) WithNamedPublicUsers(name string, opts ...func(*PublicUserQuery)) *BusinessQuery {
	query := (&PublicUserClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if bq.withNamedPublicUsers == nil {
		bq.withNamedPublicUsers = make(map[string]*PublicUserQuery)
	}
	bq.withNamedPublicUsers[name] = query
	return bq
}

// BusinessGroupBy is the group-by builder for Business entities.
type BusinessGroupBy struct {
	selector
	build *BusinessQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BusinessGroupBy) Aggregate(fns ...AggregateFunc) *BusinessGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BusinessGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bgb.build.ctx, "GroupBy")
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BusinessQuery, *BusinessGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BusinessGroupBy) sqlScan(ctx context.Context, root *BusinessQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bgb.flds)+len(bgb.fns))
		for _, f := range *bgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BusinessSelect is the builder for selecting fields of Business entities.
type BusinessSelect struct {
	*BusinessQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BusinessSelect) Aggregate(fns ...AggregateFunc) *BusinessSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BusinessSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bs.ctx, "Select")
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BusinessQuery, *BusinessSelect](ctx, bs.BusinessQuery, bs, bs.inters, v)
}

func (bs *BusinessSelect) sqlScan(ctx context.Context, root *BusinessQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bs.fns))
	for _, fn := range bs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
