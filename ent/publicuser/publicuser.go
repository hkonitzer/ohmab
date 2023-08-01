// OHMAB
// Code generated by entc, DO NOT EDIT.

package publicuser

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the publicuser type in the database.
	Label = "public_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSurname holds the string denoting the surname field in the database.
	FieldSurname = "surname"
	// FieldFirstname holds the string denoting the firstname field in the database.
	FieldFirstname = "firstname"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// EdgeBusinesses holds the string denoting the businesses edge name in mutations.
	EdgeBusinesses = "businesses"
	// EdgeTimetable holds the string denoting the timetable edge name in mutations.
	EdgeTimetable = "timetable"
	// Table holds the table name of the publicuser in the database.
	Table = "public_users"
	// BusinessesTable is the table that holds the businesses relation/edge. The primary key declared below.
	BusinessesTable = "business_public_users"
	// BusinessesInverseTable is the table name for the Business entity.
	// It exists in this package in order to avoid circular dependency with the "business" package.
	BusinessesInverseTable = "businesses"
	// TimetableTable is the table that holds the timetable relation/edge. The primary key declared below.
	TimetableTable = "timetable_users_on_duty"
	// TimetableInverseTable is the table name for the Timetable entity.
	// It exists in this package in order to avoid circular dependency with the "timetable" package.
	TimetableInverseTable = "timetables"
)

// Columns holds all SQL columns for publicuser fields.
var Columns = []string{
	FieldID,
	FieldSurname,
	FieldFirstname,
	FieldTitle,
	FieldEmail,
}

var (
	// BusinessesPrimaryKey and BusinessesColumn2 are the table columns denoting the
	// primary key for the businesses relation (M2M).
	BusinessesPrimaryKey = []string{"business_id", "public_user_id"}
	// TimetablePrimaryKey and TimetableColumn2 are the table columns denoting the
	// primary key for the timetable relation (M2M).
	TimetablePrimaryKey = []string{"timetable_id", "public_user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/hkonitzer/ohmab/ent/runtime"
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// SurnameValidator is a validator for the "surname" field. It is called by the builders before save.
	SurnameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the PublicUser queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySurname orders the results by the surname field.
func BySurname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSurname, opts...).ToFunc()
}

// ByFirstname orders the results by the firstname field.
func ByFirstname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstname, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByBusinessesCount orders the results by businesses count.
func ByBusinessesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBusinessesStep(), opts...)
	}
}

// ByBusinesses orders the results by businesses terms.
func ByBusinesses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBusinessesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTimetableCount orders the results by timetable count.
func ByTimetableCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTimetableStep(), opts...)
	}
}

// ByTimetable orders the results by timetable terms.
func ByTimetable(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTimetableStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newBusinessesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BusinessesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, BusinessesTable, BusinessesPrimaryKey...),
	)
}
func newTimetableStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TimetableInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, TimetableTable, TimetablePrimaryKey...),
	)
}
