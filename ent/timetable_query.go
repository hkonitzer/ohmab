// OHMAB
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hkonitzer/ohmab/ent/address"
	"github.com/hkonitzer/ohmab/ent/operator"
	"github.com/hkonitzer/ohmab/ent/predicate"
	"github.com/hkonitzer/ohmab/ent/timetable"
)

// TimetableQuery is the builder for querying Timetable entities.
type TimetableQuery struct {
	config
	ctx                      *QueryContext
	order                    []timetable.OrderOption
	inters                   []Interceptor
	predicates               []predicate.Timetable
	withAddress              *AddressQuery
	withOperatorsOnDuty      *OperatorQuery
	withFKs                  bool
	modifiers                []func(*sql.Selector)
	loadTotal                []func(context.Context, []*Timetable) error
	withNamedOperatorsOnDuty map[string]*OperatorQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TimetableQuery builder.
func (tq *TimetableQuery) Where(ps ...predicate.Timetable) *TimetableQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TimetableQuery) Limit(limit int) *TimetableQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TimetableQuery) Offset(offset int) *TimetableQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TimetableQuery) Unique(unique bool) *TimetableQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TimetableQuery) Order(o ...timetable.OrderOption) *TimetableQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryAddress chains the current query on the "address" edge.
func (tq *TimetableQuery) QueryAddress() *AddressQuery {
	query := (&AddressClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(timetable.Table, timetable.FieldID, selector),
			sqlgraph.To(address.Table, address.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, timetable.AddressTable, timetable.AddressColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOperatorsOnDuty chains the current query on the "operators_on_duty" edge.
func (tq *TimetableQuery) QueryOperatorsOnDuty() *OperatorQuery {
	query := (&OperatorClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(timetable.Table, timetable.FieldID, selector),
			sqlgraph.To(operator.Table, operator.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, timetable.OperatorsOnDutyTable, timetable.OperatorsOnDutyPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Timetable entity from the query.
// Returns a *NotFoundError when no Timetable was found.
func (tq *TimetableQuery) First(ctx context.Context) (*Timetable, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{timetable.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TimetableQuery) FirstX(ctx context.Context) *Timetable {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Timetable ID from the query.
// Returns a *NotFoundError when no Timetable ID was found.
func (tq *TimetableQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{timetable.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TimetableQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Timetable entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Timetable entity is found.
// Returns a *NotFoundError when no Timetable entities are found.
func (tq *TimetableQuery) Only(ctx context.Context) (*Timetable, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{timetable.Label}
	default:
		return nil, &NotSingularError{timetable.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TimetableQuery) OnlyX(ctx context.Context) *Timetable {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Timetable ID in the query.
// Returns a *NotSingularError when more than one Timetable ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TimetableQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{timetable.Label}
	default:
		err = &NotSingularError{timetable.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TimetableQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Timetables.
func (tq *TimetableQuery) All(ctx context.Context) ([]*Timetable, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Timetable, *TimetableQuery]()
	return withInterceptors[[]*Timetable](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TimetableQuery) AllX(ctx context.Context) []*Timetable {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Timetable IDs.
func (tq *TimetableQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(timetable.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TimetableQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TimetableQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TimetableQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TimetableQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TimetableQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TimetableQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TimetableQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TimetableQuery) Clone() *TimetableQuery {
	if tq == nil {
		return nil
	}
	return &TimetableQuery{
		config:              tq.config,
		ctx:                 tq.ctx.Clone(),
		order:               append([]timetable.OrderOption{}, tq.order...),
		inters:              append([]Interceptor{}, tq.inters...),
		predicates:          append([]predicate.Timetable{}, tq.predicates...),
		withAddress:         tq.withAddress.Clone(),
		withOperatorsOnDuty: tq.withOperatorsOnDuty.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithAddress tells the query-builder to eager-load the nodes that are connected to
// the "address" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TimetableQuery) WithAddress(opts ...func(*AddressQuery)) *TimetableQuery {
	query := (&AddressClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withAddress = query
	return tq
}

// WithOperatorsOnDuty tells the query-builder to eager-load the nodes that are connected to
// the "operators_on_duty" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TimetableQuery) WithOperatorsOnDuty(opts ...func(*OperatorQuery)) *TimetableQuery {
	query := (&OperatorClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withOperatorsOnDuty = query
	return tq
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
//	client.Timetable.Query().
//		GroupBy(timetable.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TimetableQuery) GroupBy(field string, fields ...string) *TimetableGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TimetableGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = timetable.Label
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
//	client.Timetable.Query().
//		Select(timetable.FieldCreatedAt).
//		Scan(ctx, &v)
func (tq *TimetableQuery) Select(fields ...string) *TimetableSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TimetableSelect{TimetableQuery: tq}
	sbuild.label = timetable.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TimetableSelect configured with the given aggregations.
func (tq *TimetableQuery) Aggregate(fns ...AggregateFunc) *TimetableSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TimetableQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !timetable.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TimetableQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Timetable, error) {
	var (
		nodes       = []*Timetable{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [2]bool{
			tq.withAddress != nil,
			tq.withOperatorsOnDuty != nil,
		}
	)
	if tq.withAddress != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, timetable.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Timetable).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Timetable{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withAddress; query != nil {
		if err := tq.loadAddress(ctx, query, nodes, nil,
			func(n *Timetable, e *Address) { n.Edges.Address = e }); err != nil {
			return nil, err
		}
	}
	if query := tq.withOperatorsOnDuty; query != nil {
		if err := tq.loadOperatorsOnDuty(ctx, query, nodes,
			func(n *Timetable) { n.Edges.OperatorsOnDuty = []*Operator{} },
			func(n *Timetable, e *Operator) { n.Edges.OperatorsOnDuty = append(n.Edges.OperatorsOnDuty, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range tq.withNamedOperatorsOnDuty {
		if err := tq.loadOperatorsOnDuty(ctx, query, nodes,
			func(n *Timetable) { n.appendNamedOperatorsOnDuty(name) },
			func(n *Timetable, e *Operator) { n.appendNamedOperatorsOnDuty(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range tq.loadTotal {
		if err := tq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TimetableQuery) loadAddress(ctx context.Context, query *AddressQuery, nodes []*Timetable, init func(*Timetable), assign func(*Timetable, *Address)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Timetable)
	for i := range nodes {
		if nodes[i].address_timetables == nil {
			continue
		}
		fk := *nodes[i].address_timetables
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(address.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "address_timetables" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tq *TimetableQuery) loadOperatorsOnDuty(ctx context.Context, query *OperatorQuery, nodes []*Timetable, init func(*Timetable), assign func(*Timetable, *Operator)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Timetable)
	nids := make(map[uuid.UUID]map[*Timetable]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(timetable.OperatorsOnDutyTable)
		s.Join(joinT).On(s.C(operator.FieldID), joinT.C(timetable.OperatorsOnDutyPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(timetable.OperatorsOnDutyPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(timetable.OperatorsOnDutyPrimaryKey[0]))
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
					nids[inValue] = map[*Timetable]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Operator](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "operators_on_duty" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (tq *TimetableQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TimetableQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(timetable.Table, timetable.Columns, sqlgraph.NewFieldSpec(timetable.FieldID, field.TypeUUID))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, timetable.FieldID)
		for i := range fields {
			if fields[i] != timetable.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TimetableQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(timetable.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = timetable.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedOperatorsOnDuty tells the query-builder to eager-load the nodes that are connected to the "operators_on_duty"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (tq *TimetableQuery) WithNamedOperatorsOnDuty(name string, opts ...func(*OperatorQuery)) *TimetableQuery {
	query := (&OperatorClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if tq.withNamedOperatorsOnDuty == nil {
		tq.withNamedOperatorsOnDuty = make(map[string]*OperatorQuery)
	}
	tq.withNamedOperatorsOnDuty[name] = query
	return tq
}

// TimetableGroupBy is the group-by builder for Timetable entities.
type TimetableGroupBy struct {
	selector
	build *TimetableQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TimetableGroupBy) Aggregate(fns ...AggregateFunc) *TimetableGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TimetableGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TimetableQuery, *TimetableGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TimetableGroupBy) sqlScan(ctx context.Context, root *TimetableQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TimetableSelect is the builder for selecting fields of Timetable entities.
type TimetableSelect struct {
	*TimetableQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TimetableSelect) Aggregate(fns ...AggregateFunc) *TimetableSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TimetableSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TimetableQuery, *TimetableSelect](ctx, ts.TimetableQuery, ts, ts.inters, v)
}

func (ts *TimetableSelect) sqlScan(ctx context.Context, root *TimetableQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
