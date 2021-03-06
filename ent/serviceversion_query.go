// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"entproject/ent/predicate"
	"entproject/ent/service"
	"entproject/ent/serviceversion"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ServiceVersionQuery is the builder for querying ServiceVersion entities.
type ServiceVersionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ServiceVersion
	// eager-loading edges.
	withService *ServiceQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ServiceVersionQuery builder.
func (svq *ServiceVersionQuery) Where(ps ...predicate.ServiceVersion) *ServiceVersionQuery {
	svq.predicates = append(svq.predicates, ps...)
	return svq
}

// Limit adds a limit step to the query.
func (svq *ServiceVersionQuery) Limit(limit int) *ServiceVersionQuery {
	svq.limit = &limit
	return svq
}

// Offset adds an offset step to the query.
func (svq *ServiceVersionQuery) Offset(offset int) *ServiceVersionQuery {
	svq.offset = &offset
	return svq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (svq *ServiceVersionQuery) Unique(unique bool) *ServiceVersionQuery {
	svq.unique = &unique
	return svq
}

// Order adds an order step to the query.
func (svq *ServiceVersionQuery) Order(o ...OrderFunc) *ServiceVersionQuery {
	svq.order = append(svq.order, o...)
	return svq
}

// QueryService chains the current query on the "service" edge.
func (svq *ServiceVersionQuery) QueryService() *ServiceQuery {
	query := &ServiceQuery{config: svq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := svq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := svq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(serviceversion.Table, serviceversion.FieldID, selector),
			sqlgraph.To(service.Table, service.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, serviceversion.ServiceTable, serviceversion.ServiceColumn),
		)
		fromU = sqlgraph.SetNeighbors(svq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ServiceVersion entity from the query.
// Returns a *NotFoundError when no ServiceVersion was found.
func (svq *ServiceVersionQuery) First(ctx context.Context) (*ServiceVersion, error) {
	nodes, err := svq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{serviceversion.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (svq *ServiceVersionQuery) FirstX(ctx context.Context) *ServiceVersion {
	node, err := svq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ServiceVersion ID from the query.
// Returns a *NotFoundError when no ServiceVersion ID was found.
func (svq *ServiceVersionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = svq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{serviceversion.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (svq *ServiceVersionQuery) FirstIDX(ctx context.Context) int {
	id, err := svq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ServiceVersion entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ServiceVersion entity is found.
// Returns a *NotFoundError when no ServiceVersion entities are found.
func (svq *ServiceVersionQuery) Only(ctx context.Context) (*ServiceVersion, error) {
	nodes, err := svq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{serviceversion.Label}
	default:
		return nil, &NotSingularError{serviceversion.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (svq *ServiceVersionQuery) OnlyX(ctx context.Context) *ServiceVersion {
	node, err := svq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ServiceVersion ID in the query.
// Returns a *NotSingularError when more than one ServiceVersion ID is found.
// Returns a *NotFoundError when no entities are found.
func (svq *ServiceVersionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = svq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{serviceversion.Label}
	default:
		err = &NotSingularError{serviceversion.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (svq *ServiceVersionQuery) OnlyIDX(ctx context.Context) int {
	id, err := svq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ServiceVersions.
func (svq *ServiceVersionQuery) All(ctx context.Context) ([]*ServiceVersion, error) {
	if err := svq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return svq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (svq *ServiceVersionQuery) AllX(ctx context.Context) []*ServiceVersion {
	nodes, err := svq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ServiceVersion IDs.
func (svq *ServiceVersionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := svq.Select(serviceversion.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (svq *ServiceVersionQuery) IDsX(ctx context.Context) []int {
	ids, err := svq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (svq *ServiceVersionQuery) Count(ctx context.Context) (int, error) {
	if err := svq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return svq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (svq *ServiceVersionQuery) CountX(ctx context.Context) int {
	count, err := svq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (svq *ServiceVersionQuery) Exist(ctx context.Context) (bool, error) {
	if err := svq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return svq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (svq *ServiceVersionQuery) ExistX(ctx context.Context) bool {
	exist, err := svq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ServiceVersionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (svq *ServiceVersionQuery) Clone() *ServiceVersionQuery {
	if svq == nil {
		return nil
	}
	return &ServiceVersionQuery{
		config:      svq.config,
		limit:       svq.limit,
		offset:      svq.offset,
		order:       append([]OrderFunc{}, svq.order...),
		predicates:  append([]predicate.ServiceVersion{}, svq.predicates...),
		withService: svq.withService.Clone(),
		// clone intermediate query.
		sql:    svq.sql.Clone(),
		path:   svq.path,
		unique: svq.unique,
	}
}

// WithService tells the query-builder to eager-load the nodes that are connected to
// the "service" edge. The optional arguments are used to configure the query builder of the edge.
func (svq *ServiceVersionQuery) WithService(opts ...func(*ServiceQuery)) *ServiceVersionQuery {
	query := &ServiceQuery{config: svq.config}
	for _, opt := range opts {
		opt(query)
	}
	svq.withService = query
	return svq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Version int `json:"version,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ServiceVersion.Query().
//		GroupBy(serviceversion.FieldVersion).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (svq *ServiceVersionQuery) GroupBy(field string, fields ...string) *ServiceVersionGroupBy {
	group := &ServiceVersionGroupBy{config: svq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := svq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return svq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Version int `json:"version,omitempty"`
//	}
//
//	client.ServiceVersion.Query().
//		Select(serviceversion.FieldVersion).
//		Scan(ctx, &v)
//
func (svq *ServiceVersionQuery) Select(fields ...string) *ServiceVersionSelect {
	svq.fields = append(svq.fields, fields...)
	return &ServiceVersionSelect{ServiceVersionQuery: svq}
}

func (svq *ServiceVersionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range svq.fields {
		if !serviceversion.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if svq.path != nil {
		prev, err := svq.path(ctx)
		if err != nil {
			return err
		}
		svq.sql = prev
	}
	return nil
}

func (svq *ServiceVersionQuery) sqlAll(ctx context.Context) ([]*ServiceVersion, error) {
	var (
		nodes       = []*ServiceVersion{}
		withFKs     = svq.withFKs
		_spec       = svq.querySpec()
		loadedTypes = [1]bool{
			svq.withService != nil,
		}
	)
	if svq.withService != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, serviceversion.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ServiceVersion{config: svq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, svq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := svq.withService; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*ServiceVersion)
		for i := range nodes {
			if nodes[i].service_service_versions == nil {
				continue
			}
			fk := *nodes[i].service_service_versions
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(service.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "service_service_versions" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Service = n
			}
		}
	}

	return nodes, nil
}

func (svq *ServiceVersionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := svq.querySpec()
	_spec.Node.Columns = svq.fields
	if len(svq.fields) > 0 {
		_spec.Unique = svq.unique != nil && *svq.unique
	}
	return sqlgraph.CountNodes(ctx, svq.driver, _spec)
}

func (svq *ServiceVersionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := svq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (svq *ServiceVersionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   serviceversion.Table,
			Columns: serviceversion.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: serviceversion.FieldID,
			},
		},
		From:   svq.sql,
		Unique: true,
	}
	if unique := svq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := svq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, serviceversion.FieldID)
		for i := range fields {
			if fields[i] != serviceversion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := svq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := svq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := svq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := svq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (svq *ServiceVersionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(svq.driver.Dialect())
	t1 := builder.Table(serviceversion.Table)
	columns := svq.fields
	if len(columns) == 0 {
		columns = serviceversion.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if svq.sql != nil {
		selector = svq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if svq.unique != nil && *svq.unique {
		selector.Distinct()
	}
	for _, p := range svq.predicates {
		p(selector)
	}
	for _, p := range svq.order {
		p(selector)
	}
	if offset := svq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := svq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ServiceVersionGroupBy is the group-by builder for ServiceVersion entities.
type ServiceVersionGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (svgb *ServiceVersionGroupBy) Aggregate(fns ...AggregateFunc) *ServiceVersionGroupBy {
	svgb.fns = append(svgb.fns, fns...)
	return svgb
}

// Scan applies the group-by query and scans the result into the given value.
func (svgb *ServiceVersionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := svgb.path(ctx)
	if err != nil {
		return err
	}
	svgb.sql = query
	return svgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (svgb *ServiceVersionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := svgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (svgb *ServiceVersionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(svgb.fields) > 1 {
		return nil, errors.New("ent: ServiceVersionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := svgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (svgb *ServiceVersionGroupBy) StringsX(ctx context.Context) []string {
	v, err := svgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (svgb *ServiceVersionGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = svgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serviceversion.Label}
	default:
		err = fmt.Errorf("ent: ServiceVersionGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (svgb *ServiceVersionGroupBy) StringX(ctx context.Context) string {
	v, err := svgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (svgb *ServiceVersionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(svgb.fields) > 1 {
		return nil, errors.New("ent: ServiceVersionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := svgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (svgb *ServiceVersionGroupBy) IntsX(ctx context.Context) []int {
	v, err := svgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (svgb *ServiceVersionGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = svgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serviceversion.Label}
	default:
		err = fmt.Errorf("ent: ServiceVersionGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (svgb *ServiceVersionGroupBy) IntX(ctx context.Context) int {
	v, err := svgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (svgb *ServiceVersionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(svgb.fields) > 1 {
		return nil, errors.New("ent: ServiceVersionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := svgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (svgb *ServiceVersionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := svgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (svgb *ServiceVersionGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = svgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serviceversion.Label}
	default:
		err = fmt.Errorf("ent: ServiceVersionGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (svgb *ServiceVersionGroupBy) Float64X(ctx context.Context) float64 {
	v, err := svgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (svgb *ServiceVersionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(svgb.fields) > 1 {
		return nil, errors.New("ent: ServiceVersionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := svgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (svgb *ServiceVersionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := svgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (svgb *ServiceVersionGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = svgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serviceversion.Label}
	default:
		err = fmt.Errorf("ent: ServiceVersionGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (svgb *ServiceVersionGroupBy) BoolX(ctx context.Context) bool {
	v, err := svgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (svgb *ServiceVersionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range svgb.fields {
		if !serviceversion.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := svgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := svgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (svgb *ServiceVersionGroupBy) sqlQuery() *sql.Selector {
	selector := svgb.sql.Select()
	aggregation := make([]string, 0, len(svgb.fns))
	for _, fn := range svgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(svgb.fields)+len(svgb.fns))
		for _, f := range svgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(svgb.fields...)...)
}

// ServiceVersionSelect is the builder for selecting fields of ServiceVersion entities.
type ServiceVersionSelect struct {
	*ServiceVersionQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (svs *ServiceVersionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := svs.prepareQuery(ctx); err != nil {
		return err
	}
	svs.sql = svs.ServiceVersionQuery.sqlQuery(ctx)
	return svs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (svs *ServiceVersionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := svs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (svs *ServiceVersionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(svs.fields) > 1 {
		return nil, errors.New("ent: ServiceVersionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := svs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (svs *ServiceVersionSelect) StringsX(ctx context.Context) []string {
	v, err := svs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (svs *ServiceVersionSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = svs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serviceversion.Label}
	default:
		err = fmt.Errorf("ent: ServiceVersionSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (svs *ServiceVersionSelect) StringX(ctx context.Context) string {
	v, err := svs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (svs *ServiceVersionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(svs.fields) > 1 {
		return nil, errors.New("ent: ServiceVersionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := svs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (svs *ServiceVersionSelect) IntsX(ctx context.Context) []int {
	v, err := svs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (svs *ServiceVersionSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = svs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serviceversion.Label}
	default:
		err = fmt.Errorf("ent: ServiceVersionSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (svs *ServiceVersionSelect) IntX(ctx context.Context) int {
	v, err := svs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (svs *ServiceVersionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(svs.fields) > 1 {
		return nil, errors.New("ent: ServiceVersionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := svs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (svs *ServiceVersionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := svs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (svs *ServiceVersionSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = svs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serviceversion.Label}
	default:
		err = fmt.Errorf("ent: ServiceVersionSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (svs *ServiceVersionSelect) Float64X(ctx context.Context) float64 {
	v, err := svs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (svs *ServiceVersionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(svs.fields) > 1 {
		return nil, errors.New("ent: ServiceVersionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := svs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (svs *ServiceVersionSelect) BoolsX(ctx context.Context) []bool {
	v, err := svs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (svs *ServiceVersionSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = svs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serviceversion.Label}
	default:
		err = fmt.Errorf("ent: ServiceVersionSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (svs *ServiceVersionSelect) BoolX(ctx context.Context) bool {
	v, err := svs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (svs *ServiceVersionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := svs.sql.Query()
	if err := svs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
