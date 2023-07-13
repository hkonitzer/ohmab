// OHMAB
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/hkonitzer/ohmab/ent/address"
	"github.com/hkonitzer/ohmab/ent/timetable"
)

// Timetable is the model entity for the Timetable schema.
type Timetable struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// The type of the timetable entry
	TimetableType timetable.TimetableType `json:"timetable_type,omitempty"`
	// DatetimeFrom holds the value of the "datetime_from" field.
	DatetimeFrom time.Time `json:"datetime_from,omitempty"`
	// The duration of the timetable entry in hours, overwrites datetime_to
	Duration uint8 `json:"duration,omitempty"`
	// The end of the timetable entry, only used if duration is not set
	DatetimeTo time.Time `json:"datetime_to,omitempty"`
	// TimeWholeDay holds the value of the "time_whole_day" field.
	TimeWholeDay bool `json:"time_whole_day,omitempty"`
	// Comment holds the value of the "comment" field.
	Comment string `json:"comment,omitempty"`
	// AvailabilityByPhone holds the value of the "availability_by_phone" field.
	AvailabilityByPhone string `json:"availability_by_phone,omitempty"`
	// AvailabilityByEmail holds the value of the "availability_by_email" field.
	AvailabilityByEmail string `json:"availability_by_email,omitempty"`
	// AvailabilityBySms holds the value of the "availability_by_sms" field.
	AvailabilityBySms string `json:"availability_by_sms,omitempty"`
	// AvailabilityByWhatsapp holds the value of the "availability_by_whatsapp" field.
	AvailabilityByWhatsapp string `json:"availability_by_whatsapp,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TimetableQuery when eager-loading is set.
	Edges              TimetableEdges `json:"edges"`
	address_timetables *uuid.UUID
	selectValues       sql.SelectValues
}

// TimetableEdges holds the relations/edges for other nodes in the graph.
type TimetableEdges struct {
	// Address holds the value of the address edge.
	Address *Address `json:"address,omitempty"`
	// UsersOnDuty holds the value of the users_on_duty edge.
	UsersOnDuty []*User `json:"users_on_duty,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedUsersOnDuty map[string][]*User
}

// AddressOrErr returns the Address value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TimetableEdges) AddressOrErr() (*Address, error) {
	if e.loadedTypes[0] {
		if e.Address == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: address.Label}
		}
		return e.Address, nil
	}
	return nil, &NotLoadedError{edge: "address"}
}

// UsersOnDutyOrErr returns the UsersOnDuty value or an error if the edge
// was not loaded in eager-loading.
func (e TimetableEdges) UsersOnDutyOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.UsersOnDuty, nil
	}
	return nil, &NotLoadedError{edge: "users_on_duty"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Timetable) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case timetable.FieldTimeWholeDay:
			values[i] = new(sql.NullBool)
		case timetable.FieldDuration:
			values[i] = new(sql.NullInt64)
		case timetable.FieldTimetableType, timetable.FieldComment, timetable.FieldAvailabilityByPhone, timetable.FieldAvailabilityByEmail, timetable.FieldAvailabilityBySms, timetable.FieldAvailabilityByWhatsapp:
			values[i] = new(sql.NullString)
		case timetable.FieldCreatedAt, timetable.FieldUpdatedAt, timetable.FieldDeletedAt, timetable.FieldDatetimeFrom, timetable.FieldDatetimeTo:
			values[i] = new(sql.NullTime)
		case timetable.FieldID:
			values[i] = new(uuid.UUID)
		case timetable.ForeignKeys[0]: // address_timetables
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Timetable fields.
func (t *Timetable) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case timetable.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case timetable.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case timetable.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case timetable.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				t.DeletedAt = value.Time
			}
		case timetable.FieldTimetableType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field timetable_type", values[i])
			} else if value.Valid {
				t.TimetableType = timetable.TimetableType(value.String)
			}
		case timetable.FieldDatetimeFrom:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field datetime_from", values[i])
			} else if value.Valid {
				t.DatetimeFrom = value.Time
			}
		case timetable.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				t.Duration = uint8(value.Int64)
			}
		case timetable.FieldDatetimeTo:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field datetime_to", values[i])
			} else if value.Valid {
				t.DatetimeTo = value.Time
			}
		case timetable.FieldTimeWholeDay:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field time_whole_day", values[i])
			} else if value.Valid {
				t.TimeWholeDay = value.Bool
			}
		case timetable.FieldComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comment", values[i])
			} else if value.Valid {
				t.Comment = value.String
			}
		case timetable.FieldAvailabilityByPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field availability_by_phone", values[i])
			} else if value.Valid {
				t.AvailabilityByPhone = value.String
			}
		case timetable.FieldAvailabilityByEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field availability_by_email", values[i])
			} else if value.Valid {
				t.AvailabilityByEmail = value.String
			}
		case timetable.FieldAvailabilityBySms:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field availability_by_sms", values[i])
			} else if value.Valid {
				t.AvailabilityBySms = value.String
			}
		case timetable.FieldAvailabilityByWhatsapp:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field availability_by_whatsapp", values[i])
			} else if value.Valid {
				t.AvailabilityByWhatsapp = value.String
			}
		case timetable.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field address_timetables", values[i])
			} else if value.Valid {
				t.address_timetables = new(uuid.UUID)
				*t.address_timetables = *value.S.(*uuid.UUID)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Timetable.
// This includes values selected through modifiers, order, etc.
func (t *Timetable) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryAddress queries the "address" edge of the Timetable entity.
func (t *Timetable) QueryAddress() *AddressQuery {
	return NewTimetableClient(t.config).QueryAddress(t)
}

// QueryUsersOnDuty queries the "users_on_duty" edge of the Timetable entity.
func (t *Timetable) QueryUsersOnDuty() *UserQuery {
	return NewTimetableClient(t.config).QueryUsersOnDuty(t)
}

// Update returns a builder for updating this Timetable.
// Note that you need to call Timetable.Unwrap() before calling this method if this Timetable
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Timetable) Update() *TimetableUpdateOne {
	return NewTimetableClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Timetable entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Timetable) Unwrap() *Timetable {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Timetable is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Timetable) String() string {
	var builder strings.Builder
	builder.WriteString("Timetable(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(t.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("timetable_type=")
	builder.WriteString(fmt.Sprintf("%v", t.TimetableType))
	builder.WriteString(", ")
	builder.WriteString("datetime_from=")
	builder.WriteString(t.DatetimeFrom.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("duration=")
	builder.WriteString(fmt.Sprintf("%v", t.Duration))
	builder.WriteString(", ")
	builder.WriteString("datetime_to=")
	builder.WriteString(t.DatetimeTo.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("time_whole_day=")
	builder.WriteString(fmt.Sprintf("%v", t.TimeWholeDay))
	builder.WriteString(", ")
	builder.WriteString("comment=")
	builder.WriteString(t.Comment)
	builder.WriteString(", ")
	builder.WriteString("availability_by_phone=")
	builder.WriteString(t.AvailabilityByPhone)
	builder.WriteString(", ")
	builder.WriteString("availability_by_email=")
	builder.WriteString(t.AvailabilityByEmail)
	builder.WriteString(", ")
	builder.WriteString("availability_by_sms=")
	builder.WriteString(t.AvailabilityBySms)
	builder.WriteString(", ")
	builder.WriteString("availability_by_whatsapp=")
	builder.WriteString(t.AvailabilityByWhatsapp)
	builder.WriteByte(')')
	return builder.String()
}

// NamedUsersOnDuty returns the UsersOnDuty named value or an error if the edge was not
// loaded in eager-loading with this name.
func (t *Timetable) NamedUsersOnDuty(name string) ([]*User, error) {
	if t.Edges.namedUsersOnDuty == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := t.Edges.namedUsersOnDuty[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (t *Timetable) appendNamedUsersOnDuty(name string, edges ...*User) {
	if t.Edges.namedUsersOnDuty == nil {
		t.Edges.namedUsersOnDuty = make(map[string][]*User)
	}
	if len(edges) == 0 {
		t.Edges.namedUsersOnDuty[name] = []*User{}
	} else {
		t.Edges.namedUsersOnDuty[name] = append(t.Edges.namedUsersOnDuty[name], edges...)
	}
}

// Timetables is a parsable slice of Timetable.
type Timetables []*Timetable
