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
	"github.com/hkonitzer/ohmab/ent/business"
)

// Address is the model entity for the Address schema.
type Address struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// The address addition
	Addition string `json:"addition,omitempty"`
	// Street holds the value of the "street" field.
	Street string `json:"street,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// Zip holds the value of the "zip" field.
	Zip string `json:"zip,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// The ICU locale identifier of the address, e.g. en_US, de_DE, ...
	Locale string `json:"locale,omitempty"`
	// Is this the primary address?
	Primary bool `json:"primary,omitempty"`
	// Telephone number
	Telephone *string `json:"telephone,omitempty"`
	// A comment for this address
	Comment string `json:"comment,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AddressQuery when eager-loading is set.
	Edges              AddressEdges `json:"edges"`
	business_addresses *uuid.UUID
	selectValues       sql.SelectValues
}

// AddressEdges holds the relations/edges for other nodes in the graph.
type AddressEdges struct {
	// Business holds the value of the business edge.
	Business *Business `json:"business,omitempty"`
	// Timetables holds the value of the timetables edge.
	Timetables []*Timetable `json:"timetables,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedTimetables map[string][]*Timetable
}

// BusinessOrErr returns the Business value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AddressEdges) BusinessOrErr() (*Business, error) {
	if e.loadedTypes[0] {
		if e.Business == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: business.Label}
		}
		return e.Business, nil
	}
	return nil, &NotLoadedError{edge: "business"}
}

// TimetablesOrErr returns the Timetables value or an error if the edge
// was not loaded in eager-loading.
func (e AddressEdges) TimetablesOrErr() ([]*Timetable, error) {
	if e.loadedTypes[1] {
		return e.Timetables, nil
	}
	return nil, &NotLoadedError{edge: "timetables"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Address) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case address.FieldPrimary:
			values[i] = new(sql.NullBool)
		case address.FieldAddition, address.FieldStreet, address.FieldCity, address.FieldZip, address.FieldState, address.FieldCountry, address.FieldLocale, address.FieldTelephone, address.FieldComment:
			values[i] = new(sql.NullString)
		case address.FieldCreatedAt, address.FieldUpdatedAt, address.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case address.FieldID:
			values[i] = new(uuid.UUID)
		case address.ForeignKeys[0]: // business_addresses
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Address fields.
func (a *Address) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case address.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case address.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case address.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case address.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = value.Time
			}
		case address.FieldAddition:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field addition", values[i])
			} else if value.Valid {
				a.Addition = value.String
			}
		case address.FieldStreet:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field street", values[i])
			} else if value.Valid {
				a.Street = value.String
			}
		case address.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				a.City = value.String
			}
		case address.FieldZip:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field zip", values[i])
			} else if value.Valid {
				a.Zip = value.String
			}
		case address.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				a.State = value.String
			}
		case address.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				a.Country = value.String
			}
		case address.FieldLocale:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field locale", values[i])
			} else if value.Valid {
				a.Locale = value.String
			}
		case address.FieldPrimary:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field primary", values[i])
			} else if value.Valid {
				a.Primary = value.Bool
			}
		case address.FieldTelephone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field telephone", values[i])
			} else if value.Valid {
				a.Telephone = new(string)
				*a.Telephone = value.String
			}
		case address.FieldComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comment", values[i])
			} else if value.Valid {
				a.Comment = value.String
			}
		case address.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field business_addresses", values[i])
			} else if value.Valid {
				a.business_addresses = new(uuid.UUID)
				*a.business_addresses = *value.S.(*uuid.UUID)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Address.
// This includes values selected through modifiers, order, etc.
func (a *Address) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryBusiness queries the "business" edge of the Address entity.
func (a *Address) QueryBusiness() *BusinessQuery {
	return NewAddressClient(a.config).QueryBusiness(a)
}

// QueryTimetables queries the "timetables" edge of the Address entity.
func (a *Address) QueryTimetables() *TimetableQuery {
	return NewAddressClient(a.config).QueryTimetables(a)
}

// Update returns a builder for updating this Address.
// Note that you need to call Address.Unwrap() before calling this method if this Address
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Address) Update() *AddressUpdateOne {
	return NewAddressClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Address entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Address) Unwrap() *Address {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Address is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Address) String() string {
	var builder strings.Builder
	builder.WriteString("Address(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(a.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("addition=")
	builder.WriteString(a.Addition)
	builder.WriteString(", ")
	builder.WriteString("street=")
	builder.WriteString(a.Street)
	builder.WriteString(", ")
	builder.WriteString("city=")
	builder.WriteString(a.City)
	builder.WriteString(", ")
	builder.WriteString("zip=")
	builder.WriteString(a.Zip)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(a.State)
	builder.WriteString(", ")
	builder.WriteString("country=")
	builder.WriteString(a.Country)
	builder.WriteString(", ")
	builder.WriteString("locale=")
	builder.WriteString(a.Locale)
	builder.WriteString(", ")
	builder.WriteString("primary=")
	builder.WriteString(fmt.Sprintf("%v", a.Primary))
	builder.WriteString(", ")
	if v := a.Telephone; v != nil {
		builder.WriteString("telephone=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("comment=")
	builder.WriteString(a.Comment)
	builder.WriteByte(')')
	return builder.String()
}

// NamedTimetables returns the Timetables named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Address) NamedTimetables(name string) ([]*Timetable, error) {
	if a.Edges.namedTimetables == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedTimetables[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Address) appendNamedTimetables(name string, edges ...*Timetable) {
	if a.Edges.namedTimetables == nil {
		a.Edges.namedTimetables = make(map[string][]*Timetable)
	}
	if len(edges) == 0 {
		a.Edges.namedTimetables[name] = []*Timetable{}
	} else {
		a.Edges.namedTimetables[name] = append(a.Edges.namedTimetables[name], edges...)
	}
}

// Addresses is a parsable slice of Address.
type Addresses []*Address
