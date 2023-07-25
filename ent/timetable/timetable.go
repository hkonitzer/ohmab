// OHMAB
// Code generated by entc, DO NOT EDIT.

package timetable

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the timetable type in the database.
	Label = "timetable"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldTimetableType holds the string denoting the timetable_type field in the database.
	FieldTimetableType = "timetable_type"
	// FieldDatetimeFrom holds the string denoting the datetime_from field in the database.
	FieldDatetimeFrom = "datetime_from"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldDatetimeTo holds the string denoting the datetime_to field in the database.
	FieldDatetimeTo = "datetime_to"
	// FieldTimeWholeDay holds the string denoting the time_whole_day field in the database.
	FieldTimeWholeDay = "time_whole_day"
	// FieldComment holds the string denoting the comment field in the database.
	FieldComment = "comment"
	// FieldAvailabilityByPhone holds the string denoting the availability_by_phone field in the database.
	FieldAvailabilityByPhone = "availability_by_phone"
	// FieldAvailabilityByEmail holds the string denoting the availability_by_email field in the database.
	FieldAvailabilityByEmail = "availability_by_email"
	// FieldAvailabilityBySms holds the string denoting the availability_by_sms field in the database.
	FieldAvailabilityBySms = "availability_by_sms"
	// FieldAvailabilityByWhatsapp holds the string denoting the availability_by_whatsapp field in the database.
	FieldAvailabilityByWhatsapp = "availability_by_whatsapp"
	// EdgeAddress holds the string denoting the address edge name in mutations.
	EdgeAddress = "address"
	// EdgeUsersOnDuty holds the string denoting the users_on_duty edge name in mutations.
	EdgeUsersOnDuty = "users_on_duty"
	// Table holds the table name of the timetable in the database.
	Table = "timetables"
	// AddressTable is the table that holds the address relation/edge.
	AddressTable = "timetables"
	// AddressInverseTable is the table name for the Address entity.
	// It exists in this package in order to avoid circular dependency with the "address" package.
	AddressInverseTable = "addresses"
	// AddressColumn is the table column denoting the address relation/edge.
	AddressColumn = "address_timetables"
	// UsersOnDutyTable is the table that holds the users_on_duty relation/edge. The primary key declared below.
	UsersOnDutyTable = "timetable_users_on_duty"
	// UsersOnDutyInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersOnDutyInverseTable = "users"
)

// Columns holds all SQL columns for timetable fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldTimetableType,
	FieldDatetimeFrom,
	FieldDuration,
	FieldDatetimeTo,
	FieldTimeWholeDay,
	FieldComment,
	FieldAvailabilityByPhone,
	FieldAvailabilityByEmail,
	FieldAvailabilityBySms,
	FieldAvailabilityByWhatsapp,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "timetables"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"address_timetables",
}

var (
	// UsersOnDutyPrimaryKey and UsersOnDutyColumn2 are the table columns denoting the
	// primary key for the users_on_duty relation (M2M).
	UsersOnDutyPrimaryKey = []string{"timetable_id", "user_id"}
)

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
//	import _ "github.com/hkonitzer/ohmab/ent/runtime"
var (
	Hooks [3]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DurationValidator is a validator for the "duration" field. It is called by the builders before save.
	DurationValidator func(uint8) error
	// DefaultTimeWholeDay holds the default value on creation for the "time_whole_day" field.
	DefaultTimeWholeDay bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// TimetableType defines the type for the "timetable_type" enum field.
type TimetableType string

// TimetableTypeDEFAULT is the default value of the TimetableType enum.
const DefaultTimetableType = TimetableTypeDEFAULT

// TimetableType values.
const (
	TimetableTypeDEFAULT          TimetableType = "DEFAULT"
	TimetableTypeREGULAR          TimetableType = "REGULAR"
	TimetableTypeCLOSED           TimetableType = "CLOSED"
	TimetableTypeEMERGENCYSERVICE TimetableType = "EMERGENCYSERVICE"
	TimetableTypeHOLIDAY          TimetableType = "HOLIDAY"
	TimetableTypeSPECIAL          TimetableType = "SPECIAL"
)

func (tt TimetableType) String() string {
	return string(tt)
}

// TimetableTypeValidator is a validator for the "timetable_type" field enum values. It is called by the builders before save.
func TimetableTypeValidator(tt TimetableType) error {
	switch tt {
	case TimetableTypeDEFAULT, TimetableTypeREGULAR, TimetableTypeCLOSED, TimetableTypeEMERGENCYSERVICE, TimetableTypeHOLIDAY, TimetableTypeSPECIAL:
		return nil
	default:
		return fmt.Errorf("timetable: invalid enum value for timetable_type field: %q", tt)
	}
}

// OrderOption defines the ordering options for the Timetable queries.
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

// ByTimetableType orders the results by the timetable_type field.
func ByTimetableType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimetableType, opts...).ToFunc()
}

// ByDatetimeFrom orders the results by the datetime_from field.
func ByDatetimeFrom(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDatetimeFrom, opts...).ToFunc()
}

// ByDuration orders the results by the duration field.
func ByDuration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDuration, opts...).ToFunc()
}

// ByDatetimeTo orders the results by the datetime_to field.
func ByDatetimeTo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDatetimeTo, opts...).ToFunc()
}

// ByTimeWholeDay orders the results by the time_whole_day field.
func ByTimeWholeDay(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimeWholeDay, opts...).ToFunc()
}

// ByComment orders the results by the comment field.
func ByComment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComment, opts...).ToFunc()
}

// ByAvailabilityByPhone orders the results by the availability_by_phone field.
func ByAvailabilityByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvailabilityByPhone, opts...).ToFunc()
}

// ByAvailabilityByEmail orders the results by the availability_by_email field.
func ByAvailabilityByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvailabilityByEmail, opts...).ToFunc()
}

// ByAvailabilityBySms orders the results by the availability_by_sms field.
func ByAvailabilityBySms(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvailabilityBySms, opts...).ToFunc()
}

// ByAvailabilityByWhatsapp orders the results by the availability_by_whatsapp field.
func ByAvailabilityByWhatsapp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvailabilityByWhatsapp, opts...).ToFunc()
}

// ByAddressField orders the results by address field.
func ByAddressField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAddressStep(), sql.OrderByField(field, opts...))
	}
}

// ByUsersOnDutyCount orders the results by users_on_duty count.
func ByUsersOnDutyCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUsersOnDutyStep(), opts...)
	}
}

// ByUsersOnDuty orders the results by users_on_duty terms.
func ByUsersOnDuty(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersOnDutyStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAddressStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AddressInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AddressTable, AddressColumn),
	)
}
func newUsersOnDutyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersOnDutyInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, UsersOnDutyTable, UsersOnDutyPrimaryKey...),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e TimetableType) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *TimetableType) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = TimetableType(str)
	if err := TimetableTypeValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid TimetableType", str)
	}
	return nil
}
