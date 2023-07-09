// OHMAB
// Code generated by entc, DO NOT EDIT.

package address

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the address type in the database.
	Label = "address"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldAddition holds the string denoting the addition field in the database.
	FieldAddition = "addition"
	// FieldStreet holds the string denoting the street field in the database.
	FieldStreet = "street"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldZip holds the string denoting the zip field in the database.
	FieldZip = "zip"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldLocale holds the string denoting the locale field in the database.
	FieldLocale = "locale"
	// FieldPrimary holds the string denoting the primary field in the database.
	FieldPrimary = "primary"
	// FieldTelephone holds the string denoting the telephone field in the database.
	FieldTelephone = "telephone"
	// FieldComment holds the string denoting the comment field in the database.
	FieldComment = "comment"
	// EdgeBusiness holds the string denoting the business edge name in mutations.
	EdgeBusiness = "business"
	// EdgeTimetables holds the string denoting the timetables edge name in mutations.
	EdgeTimetables = "timetables"
	// Table holds the table name of the address in the database.
	Table = "addresses"
	// BusinessTable is the table that holds the business relation/edge.
	BusinessTable = "addresses"
	// BusinessInverseTable is the table name for the Business entity.
	// It exists in this package in order to avoid circular dependency with the "business" package.
	BusinessInverseTable = "businesses"
	// BusinessColumn is the table column denoting the business relation/edge.
	BusinessColumn = "business_addresses"
	// TimetablesTable is the table that holds the timetables relation/edge.
	TimetablesTable = "timetables"
	// TimetablesInverseTable is the table name for the Timetable entity.
	// It exists in this package in order to avoid circular dependency with the "timetable" package.
	TimetablesInverseTable = "timetables"
	// TimetablesColumn is the table column denoting the timetables relation/edge.
	TimetablesColumn = "address_timetables"
)

// Columns holds all SQL columns for address fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAddition,
	FieldStreet,
	FieldCity,
	FieldZip,
	FieldState,
	FieldCountry,
	FieldLocale,
	FieldPrimary,
	FieldTelephone,
	FieldComment,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "addresses"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"business_addresses",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "hynie.de/ohmab/ent/runtime"
var (
	Hooks [2]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultLocale holds the default value on creation for the "locale" field.
	DefaultLocale string
	// LocaleValidator is a validator for the "locale" field. It is called by the builders before save.
	LocaleValidator func(string) error
	// DefaultPrimary holds the default value on creation for the "primary" field.
	DefaultPrimary bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Address queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByAddition orders the results by the addition field.
func ByAddition(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddition, opts...).ToFunc()
}

// ByStreet orders the results by the street field.
func ByStreet(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStreet, opts...).ToFunc()
}

// ByCity orders the results by the city field.
func ByCity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCity, opts...).ToFunc()
}

// ByZip orders the results by the zip field.
func ByZip(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldZip, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// ByCountry orders the results by the country field.
func ByCountry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountry, opts...).ToFunc()
}

// ByLocale orders the results by the locale field.
func ByLocale(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocale, opts...).ToFunc()
}

// ByPrimary orders the results by the primary field.
func ByPrimary(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrimary, opts...).ToFunc()
}

// ByTelephone orders the results by the telephone field.
func ByTelephone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTelephone, opts...).ToFunc()
}

// ByComment orders the results by the comment field.
func ByComment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComment, opts...).ToFunc()
}

// ByBusinessField orders the results by business field.
func ByBusinessField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBusinessStep(), sql.OrderByField(field, opts...))
	}
}

// ByTimetablesCount orders the results by timetables count.
func ByTimetablesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTimetablesStep(), opts...)
	}
}

// ByTimetables orders the results by timetables terms.
func ByTimetables(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTimetablesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newBusinessStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BusinessInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, BusinessTable, BusinessColumn),
	)
}
func newTimetablesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TimetablesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TimetablesTable, TimetablesColumn),
	)
}
