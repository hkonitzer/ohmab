// OHMAB
// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldLogin holds the string denoting the login field in the database.
	FieldLogin = "login"
	// FieldSurname holds the string denoting the surname field in the database.
	FieldSurname = "surname"
	// FieldFirstname holds the string denoting the firstname field in the database.
	FieldFirstname = "firstname"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPasswordhash holds the string denoting the passwordhash field in the database.
	FieldPasswordhash = "passwordhash"
	// FieldComment holds the string denoting the comment field in the database.
	FieldComment = "comment"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// EdgeBusinesses holds the string denoting the businesses edge name in mutations.
	EdgeBusinesses = "businesses"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeTimetable holds the string denoting the timetable edge name in mutations.
	EdgeTimetable = "timetable"
	// Table holds the table name of the user in the database.
	Table = "users"
	// BusinessesTable is the table that holds the businesses relation/edge. The primary key declared below.
	BusinessesTable = "business_users"
	// BusinessesInverseTable is the table name for the Business entity.
	// It exists in this package in order to avoid circular dependency with the "business" package.
	BusinessesInverseTable = "businesses"
	// TagsTable is the table that holds the tags relation/edge. The primary key declared below.
	TagsTable = "user_tags"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// TimetableTable is the table that holds the timetable relation/edge. The primary key declared below.
	TimetableTable = "timetable_users_on_duty"
	// TimetableInverseTable is the table name for the Timetable entity.
	// It exists in this package in order to avoid circular dependency with the "timetable" package.
	TimetableInverseTable = "timetables"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldLogin,
	FieldSurname,
	FieldFirstname,
	FieldTitle,
	FieldEmail,
	FieldPasswordhash,
	FieldComment,
	FieldActive,
	FieldRole,
}

var (
	// BusinessesPrimaryKey and BusinessesColumn2 are the table columns denoting the
	// primary key for the businesses relation (M2M).
	BusinessesPrimaryKey = []string{"business_id", "user_id"}
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"user_id", "tag_id"}
	// TimetablePrimaryKey and TimetableColumn2 are the table columns denoting the
	// primary key for the timetable relation (M2M).
	TimetablePrimaryKey = []string{"timetable_id", "user_id"}
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
//	import _ "hynie.de/ohmab/ent/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// LoginValidator is a validator for the "login" field. It is called by the builders before save.
	LoginValidator func(string) error
	// SurnameValidator is a validator for the "surname" field. It is called by the builders before save.
	SurnameValidator func(string) error
	// DefaultActive holds the default value on creation for the "active" field.
	DefaultActive bool
	// DefaultRole holds the default value on creation for the "role" field.
	DefaultRole int
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the User queries.
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

// ByLogin orders the results by the login field.
func ByLogin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLogin, opts...).ToFunc()
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

// ByPasswordhash orders the results by the passwordhash field.
func ByPasswordhash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPasswordhash, opts...).ToFunc()
}

// ByComment orders the results by the comment field.
func ByComment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComment, opts...).ToFunc()
}

// ByActive orders the results by the active field.
func ByActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActive, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
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

// ByTagsCount orders the results by tags count.
func ByTagsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTagsStep(), opts...)
	}
}

// ByTags orders the results by tags terms.
func ByTags(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTagsStep(), append([]sql.OrderTerm{term}, terms...)...)
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
func newTagsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TagsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
	)
}
func newTimetableStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TimetableInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, TimetableTable, TimetablePrimaryKey...),
	)
}
