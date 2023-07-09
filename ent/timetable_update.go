// OHMAB
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"hynie.de/ohmab/ent/address"
	"hynie.de/ohmab/ent/predicate"
	"hynie.de/ohmab/ent/timetable"
	"hynie.de/ohmab/ent/user"
)

// TimetableUpdate is the builder for updating Timetable entities.
type TimetableUpdate struct {
	config
	hooks    []Hook
	mutation *TimetableMutation
}

// Where appends a list predicates to the TimetableUpdate builder.
func (tu *TimetableUpdate) Where(ps ...predicate.Timetable) *TimetableUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TimetableUpdate) SetUpdatedAt(t time.Time) *TimetableUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetDeletedAt sets the "deleted_at" field.
func (tu *TimetableUpdate) SetDeletedAt(t time.Time) *TimetableUpdate {
	tu.mutation.SetDeletedAt(t)
	return tu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tu *TimetableUpdate) SetNillableDeletedAt(t *time.Time) *TimetableUpdate {
	if t != nil {
		tu.SetDeletedAt(*t)
	}
	return tu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (tu *TimetableUpdate) ClearDeletedAt() *TimetableUpdate {
	tu.mutation.ClearDeletedAt()
	return tu
}

// SetTimetableType sets the "timetable_type" field.
func (tu *TimetableUpdate) SetTimetableType(tt timetable.TimetableType) *TimetableUpdate {
	tu.mutation.SetTimetableType(tt)
	return tu
}

// SetNillableTimetableType sets the "timetable_type" field if the given value is not nil.
func (tu *TimetableUpdate) SetNillableTimetableType(tt *timetable.TimetableType) *TimetableUpdate {
	if tt != nil {
		tu.SetTimetableType(*tt)
	}
	return tu
}

// SetDatetimeFrom sets the "datetime_from" field.
func (tu *TimetableUpdate) SetDatetimeFrom(t time.Time) *TimetableUpdate {
	tu.mutation.SetDatetimeFrom(t)
	return tu
}

// SetNillableDatetimeFrom sets the "datetime_from" field if the given value is not nil.
func (tu *TimetableUpdate) SetNillableDatetimeFrom(t *time.Time) *TimetableUpdate {
	if t != nil {
		tu.SetDatetimeFrom(*t)
	}
	return tu
}

// ClearDatetimeFrom clears the value of the "datetime_from" field.
func (tu *TimetableUpdate) ClearDatetimeFrom() *TimetableUpdate {
	tu.mutation.ClearDatetimeFrom()
	return tu
}

// SetDuration sets the "duration" field.
func (tu *TimetableUpdate) SetDuration(u uint8) *TimetableUpdate {
	tu.mutation.ResetDuration()
	tu.mutation.SetDuration(u)
	return tu
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (tu *TimetableUpdate) SetNillableDuration(u *uint8) *TimetableUpdate {
	if u != nil {
		tu.SetDuration(*u)
	}
	return tu
}

// AddDuration adds u to the "duration" field.
func (tu *TimetableUpdate) AddDuration(u int8) *TimetableUpdate {
	tu.mutation.AddDuration(u)
	return tu
}

// ClearDuration clears the value of the "duration" field.
func (tu *TimetableUpdate) ClearDuration() *TimetableUpdate {
	tu.mutation.ClearDuration()
	return tu
}

// SetDatetimeTo sets the "datetime_to" field.
func (tu *TimetableUpdate) SetDatetimeTo(t time.Time) *TimetableUpdate {
	tu.mutation.SetDatetimeTo(t)
	return tu
}

// SetTimeWholeDay sets the "time_whole_day" field.
func (tu *TimetableUpdate) SetTimeWholeDay(b bool) *TimetableUpdate {
	tu.mutation.SetTimeWholeDay(b)
	return tu
}

// SetNillableTimeWholeDay sets the "time_whole_day" field if the given value is not nil.
func (tu *TimetableUpdate) SetNillableTimeWholeDay(b *bool) *TimetableUpdate {
	if b != nil {
		tu.SetTimeWholeDay(*b)
	}
	return tu
}

// SetComment sets the "comment" field.
func (tu *TimetableUpdate) SetComment(s string) *TimetableUpdate {
	tu.mutation.SetComment(s)
	return tu
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (tu *TimetableUpdate) SetNillableComment(s *string) *TimetableUpdate {
	if s != nil {
		tu.SetComment(*s)
	}
	return tu
}

// ClearComment clears the value of the "comment" field.
func (tu *TimetableUpdate) ClearComment() *TimetableUpdate {
	tu.mutation.ClearComment()
	return tu
}

// SetAvailabilityByPhone sets the "availability_by_phone" field.
func (tu *TimetableUpdate) SetAvailabilityByPhone(s string) *TimetableUpdate {
	tu.mutation.SetAvailabilityByPhone(s)
	return tu
}

// SetNillableAvailabilityByPhone sets the "availability_by_phone" field if the given value is not nil.
func (tu *TimetableUpdate) SetNillableAvailabilityByPhone(s *string) *TimetableUpdate {
	if s != nil {
		tu.SetAvailabilityByPhone(*s)
	}
	return tu
}

// ClearAvailabilityByPhone clears the value of the "availability_by_phone" field.
func (tu *TimetableUpdate) ClearAvailabilityByPhone() *TimetableUpdate {
	tu.mutation.ClearAvailabilityByPhone()
	return tu
}

// SetAvailabilityByEmail sets the "availability_by_email" field.
func (tu *TimetableUpdate) SetAvailabilityByEmail(s string) *TimetableUpdate {
	tu.mutation.SetAvailabilityByEmail(s)
	return tu
}

// SetNillableAvailabilityByEmail sets the "availability_by_email" field if the given value is not nil.
func (tu *TimetableUpdate) SetNillableAvailabilityByEmail(s *string) *TimetableUpdate {
	if s != nil {
		tu.SetAvailabilityByEmail(*s)
	}
	return tu
}

// ClearAvailabilityByEmail clears the value of the "availability_by_email" field.
func (tu *TimetableUpdate) ClearAvailabilityByEmail() *TimetableUpdate {
	tu.mutation.ClearAvailabilityByEmail()
	return tu
}

// SetAvailabilityBySms sets the "availability_by_sms" field.
func (tu *TimetableUpdate) SetAvailabilityBySms(s string) *TimetableUpdate {
	tu.mutation.SetAvailabilityBySms(s)
	return tu
}

// SetNillableAvailabilityBySms sets the "availability_by_sms" field if the given value is not nil.
func (tu *TimetableUpdate) SetNillableAvailabilityBySms(s *string) *TimetableUpdate {
	if s != nil {
		tu.SetAvailabilityBySms(*s)
	}
	return tu
}

// ClearAvailabilityBySms clears the value of the "availability_by_sms" field.
func (tu *TimetableUpdate) ClearAvailabilityBySms() *TimetableUpdate {
	tu.mutation.ClearAvailabilityBySms()
	return tu
}

// SetAvailabilityByWhatsapp sets the "availability_by_whatsapp" field.
func (tu *TimetableUpdate) SetAvailabilityByWhatsapp(s string) *TimetableUpdate {
	tu.mutation.SetAvailabilityByWhatsapp(s)
	return tu
}

// SetNillableAvailabilityByWhatsapp sets the "availability_by_whatsapp" field if the given value is not nil.
func (tu *TimetableUpdate) SetNillableAvailabilityByWhatsapp(s *string) *TimetableUpdate {
	if s != nil {
		tu.SetAvailabilityByWhatsapp(*s)
	}
	return tu
}

// ClearAvailabilityByWhatsapp clears the value of the "availability_by_whatsapp" field.
func (tu *TimetableUpdate) ClearAvailabilityByWhatsapp() *TimetableUpdate {
	tu.mutation.ClearAvailabilityByWhatsapp()
	return tu
}

// SetAddressID sets the "address" edge to the Address entity by ID.
func (tu *TimetableUpdate) SetAddressID(id uuid.UUID) *TimetableUpdate {
	tu.mutation.SetAddressID(id)
	return tu
}

// SetAddress sets the "address" edge to the Address entity.
func (tu *TimetableUpdate) SetAddress(a *Address) *TimetableUpdate {
	return tu.SetAddressID(a.ID)
}

// AddUsersOnDutyIDs adds the "users_on_duty" edge to the User entity by IDs.
func (tu *TimetableUpdate) AddUsersOnDutyIDs(ids ...uuid.UUID) *TimetableUpdate {
	tu.mutation.AddUsersOnDutyIDs(ids...)
	return tu
}

// AddUsersOnDuty adds the "users_on_duty" edges to the User entity.
func (tu *TimetableUpdate) AddUsersOnDuty(u ...*User) *TimetableUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.AddUsersOnDutyIDs(ids...)
}

// Mutation returns the TimetableMutation object of the builder.
func (tu *TimetableUpdate) Mutation() *TimetableMutation {
	return tu.mutation
}

// ClearAddress clears the "address" edge to the Address entity.
func (tu *TimetableUpdate) ClearAddress() *TimetableUpdate {
	tu.mutation.ClearAddress()
	return tu
}

// ClearUsersOnDuty clears all "users_on_duty" edges to the User entity.
func (tu *TimetableUpdate) ClearUsersOnDuty() *TimetableUpdate {
	tu.mutation.ClearUsersOnDuty()
	return tu
}

// RemoveUsersOnDutyIDs removes the "users_on_duty" edge to User entities by IDs.
func (tu *TimetableUpdate) RemoveUsersOnDutyIDs(ids ...uuid.UUID) *TimetableUpdate {
	tu.mutation.RemoveUsersOnDutyIDs(ids...)
	return tu
}

// RemoveUsersOnDuty removes "users_on_duty" edges to User entities.
func (tu *TimetableUpdate) RemoveUsersOnDuty(u ...*User) *TimetableUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.RemoveUsersOnDutyIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TimetableUpdate) Save(ctx context.Context) (int, error) {
	if err := tu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TimetableUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TimetableUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TimetableUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TimetableUpdate) defaults() error {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		if timetable.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized timetable.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := timetable.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tu *TimetableUpdate) check() error {
	if v, ok := tu.mutation.TimetableType(); ok {
		if err := timetable.TimetableTypeValidator(v); err != nil {
			return &ValidationError{Name: "timetable_type", err: fmt.Errorf(`ent: validator failed for field "Timetable.timetable_type": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Duration(); ok {
		if err := timetable.DurationValidator(v); err != nil {
			return &ValidationError{Name: "duration", err: fmt.Errorf(`ent: validator failed for field "Timetable.duration": %w`, err)}
		}
	}
	if _, ok := tu.mutation.AddressID(); tu.mutation.AddressCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Timetable.address"`)
	}
	return nil
}

func (tu *TimetableUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(timetable.Table, timetable.Columns, sqlgraph.NewFieldSpec(timetable.FieldID, field.TypeUUID))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(timetable.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.DeletedAt(); ok {
		_spec.SetField(timetable.FieldDeletedAt, field.TypeTime, value)
	}
	if tu.mutation.DeletedAtCleared() {
		_spec.ClearField(timetable.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := tu.mutation.TimetableType(); ok {
		_spec.SetField(timetable.FieldTimetableType, field.TypeEnum, value)
	}
	if value, ok := tu.mutation.DatetimeFrom(); ok {
		_spec.SetField(timetable.FieldDatetimeFrom, field.TypeTime, value)
	}
	if tu.mutation.DatetimeFromCleared() {
		_spec.ClearField(timetable.FieldDatetimeFrom, field.TypeTime)
	}
	if value, ok := tu.mutation.Duration(); ok {
		_spec.SetField(timetable.FieldDuration, field.TypeUint8, value)
	}
	if value, ok := tu.mutation.AddedDuration(); ok {
		_spec.AddField(timetable.FieldDuration, field.TypeUint8, value)
	}
	if tu.mutation.DurationCleared() {
		_spec.ClearField(timetable.FieldDuration, field.TypeUint8)
	}
	if value, ok := tu.mutation.DatetimeTo(); ok {
		_spec.SetField(timetable.FieldDatetimeTo, field.TypeTime, value)
	}
	if value, ok := tu.mutation.TimeWholeDay(); ok {
		_spec.SetField(timetable.FieldTimeWholeDay, field.TypeBool, value)
	}
	if value, ok := tu.mutation.Comment(); ok {
		_spec.SetField(timetable.FieldComment, field.TypeString, value)
	}
	if tu.mutation.CommentCleared() {
		_spec.ClearField(timetable.FieldComment, field.TypeString)
	}
	if value, ok := tu.mutation.AvailabilityByPhone(); ok {
		_spec.SetField(timetable.FieldAvailabilityByPhone, field.TypeString, value)
	}
	if tu.mutation.AvailabilityByPhoneCleared() {
		_spec.ClearField(timetable.FieldAvailabilityByPhone, field.TypeString)
	}
	if value, ok := tu.mutation.AvailabilityByEmail(); ok {
		_spec.SetField(timetable.FieldAvailabilityByEmail, field.TypeString, value)
	}
	if tu.mutation.AvailabilityByEmailCleared() {
		_spec.ClearField(timetable.FieldAvailabilityByEmail, field.TypeString)
	}
	if value, ok := tu.mutation.AvailabilityBySms(); ok {
		_spec.SetField(timetable.FieldAvailabilityBySms, field.TypeString, value)
	}
	if tu.mutation.AvailabilityBySmsCleared() {
		_spec.ClearField(timetable.FieldAvailabilityBySms, field.TypeString)
	}
	if value, ok := tu.mutation.AvailabilityByWhatsapp(); ok {
		_spec.SetField(timetable.FieldAvailabilityByWhatsapp, field.TypeString, value)
	}
	if tu.mutation.AvailabilityByWhatsappCleared() {
		_spec.ClearField(timetable.FieldAvailabilityByWhatsapp, field.TypeString)
	}
	if tu.mutation.AddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   timetable.AddressTable,
			Columns: []string{timetable.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.AddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   timetable.AddressTable,
			Columns: []string{timetable.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.UsersOnDutyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   timetable.UsersOnDutyTable,
			Columns: timetable.UsersOnDutyPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedUsersOnDutyIDs(); len(nodes) > 0 && !tu.mutation.UsersOnDutyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   timetable.UsersOnDutyTable,
			Columns: timetable.UsersOnDutyPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.UsersOnDutyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   timetable.UsersOnDutyTable,
			Columns: timetable.UsersOnDutyPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{timetable.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TimetableUpdateOne is the builder for updating a single Timetable entity.
type TimetableUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TimetableMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TimetableUpdateOne) SetUpdatedAt(t time.Time) *TimetableUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetDeletedAt sets the "deleted_at" field.
func (tuo *TimetableUpdateOne) SetDeletedAt(t time.Time) *TimetableUpdateOne {
	tuo.mutation.SetDeletedAt(t)
	return tuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tuo *TimetableUpdateOne) SetNillableDeletedAt(t *time.Time) *TimetableUpdateOne {
	if t != nil {
		tuo.SetDeletedAt(*t)
	}
	return tuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (tuo *TimetableUpdateOne) ClearDeletedAt() *TimetableUpdateOne {
	tuo.mutation.ClearDeletedAt()
	return tuo
}

// SetTimetableType sets the "timetable_type" field.
func (tuo *TimetableUpdateOne) SetTimetableType(tt timetable.TimetableType) *TimetableUpdateOne {
	tuo.mutation.SetTimetableType(tt)
	return tuo
}

// SetNillableTimetableType sets the "timetable_type" field if the given value is not nil.
func (tuo *TimetableUpdateOne) SetNillableTimetableType(tt *timetable.TimetableType) *TimetableUpdateOne {
	if tt != nil {
		tuo.SetTimetableType(*tt)
	}
	return tuo
}

// SetDatetimeFrom sets the "datetime_from" field.
func (tuo *TimetableUpdateOne) SetDatetimeFrom(t time.Time) *TimetableUpdateOne {
	tuo.mutation.SetDatetimeFrom(t)
	return tuo
}

// SetNillableDatetimeFrom sets the "datetime_from" field if the given value is not nil.
func (tuo *TimetableUpdateOne) SetNillableDatetimeFrom(t *time.Time) *TimetableUpdateOne {
	if t != nil {
		tuo.SetDatetimeFrom(*t)
	}
	return tuo
}

// ClearDatetimeFrom clears the value of the "datetime_from" field.
func (tuo *TimetableUpdateOne) ClearDatetimeFrom() *TimetableUpdateOne {
	tuo.mutation.ClearDatetimeFrom()
	return tuo
}

// SetDuration sets the "duration" field.
func (tuo *TimetableUpdateOne) SetDuration(u uint8) *TimetableUpdateOne {
	tuo.mutation.ResetDuration()
	tuo.mutation.SetDuration(u)
	return tuo
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (tuo *TimetableUpdateOne) SetNillableDuration(u *uint8) *TimetableUpdateOne {
	if u != nil {
		tuo.SetDuration(*u)
	}
	return tuo
}

// AddDuration adds u to the "duration" field.
func (tuo *TimetableUpdateOne) AddDuration(u int8) *TimetableUpdateOne {
	tuo.mutation.AddDuration(u)
	return tuo
}

// ClearDuration clears the value of the "duration" field.
func (tuo *TimetableUpdateOne) ClearDuration() *TimetableUpdateOne {
	tuo.mutation.ClearDuration()
	return tuo
}

// SetDatetimeTo sets the "datetime_to" field.
func (tuo *TimetableUpdateOne) SetDatetimeTo(t time.Time) *TimetableUpdateOne {
	tuo.mutation.SetDatetimeTo(t)
	return tuo
}

// SetTimeWholeDay sets the "time_whole_day" field.
func (tuo *TimetableUpdateOne) SetTimeWholeDay(b bool) *TimetableUpdateOne {
	tuo.mutation.SetTimeWholeDay(b)
	return tuo
}

// SetNillableTimeWholeDay sets the "time_whole_day" field if the given value is not nil.
func (tuo *TimetableUpdateOne) SetNillableTimeWholeDay(b *bool) *TimetableUpdateOne {
	if b != nil {
		tuo.SetTimeWholeDay(*b)
	}
	return tuo
}

// SetComment sets the "comment" field.
func (tuo *TimetableUpdateOne) SetComment(s string) *TimetableUpdateOne {
	tuo.mutation.SetComment(s)
	return tuo
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (tuo *TimetableUpdateOne) SetNillableComment(s *string) *TimetableUpdateOne {
	if s != nil {
		tuo.SetComment(*s)
	}
	return tuo
}

// ClearComment clears the value of the "comment" field.
func (tuo *TimetableUpdateOne) ClearComment() *TimetableUpdateOne {
	tuo.mutation.ClearComment()
	return tuo
}

// SetAvailabilityByPhone sets the "availability_by_phone" field.
func (tuo *TimetableUpdateOne) SetAvailabilityByPhone(s string) *TimetableUpdateOne {
	tuo.mutation.SetAvailabilityByPhone(s)
	return tuo
}

// SetNillableAvailabilityByPhone sets the "availability_by_phone" field if the given value is not nil.
func (tuo *TimetableUpdateOne) SetNillableAvailabilityByPhone(s *string) *TimetableUpdateOne {
	if s != nil {
		tuo.SetAvailabilityByPhone(*s)
	}
	return tuo
}

// ClearAvailabilityByPhone clears the value of the "availability_by_phone" field.
func (tuo *TimetableUpdateOne) ClearAvailabilityByPhone() *TimetableUpdateOne {
	tuo.mutation.ClearAvailabilityByPhone()
	return tuo
}

// SetAvailabilityByEmail sets the "availability_by_email" field.
func (tuo *TimetableUpdateOne) SetAvailabilityByEmail(s string) *TimetableUpdateOne {
	tuo.mutation.SetAvailabilityByEmail(s)
	return tuo
}

// SetNillableAvailabilityByEmail sets the "availability_by_email" field if the given value is not nil.
func (tuo *TimetableUpdateOne) SetNillableAvailabilityByEmail(s *string) *TimetableUpdateOne {
	if s != nil {
		tuo.SetAvailabilityByEmail(*s)
	}
	return tuo
}

// ClearAvailabilityByEmail clears the value of the "availability_by_email" field.
func (tuo *TimetableUpdateOne) ClearAvailabilityByEmail() *TimetableUpdateOne {
	tuo.mutation.ClearAvailabilityByEmail()
	return tuo
}

// SetAvailabilityBySms sets the "availability_by_sms" field.
func (tuo *TimetableUpdateOne) SetAvailabilityBySms(s string) *TimetableUpdateOne {
	tuo.mutation.SetAvailabilityBySms(s)
	return tuo
}

// SetNillableAvailabilityBySms sets the "availability_by_sms" field if the given value is not nil.
func (tuo *TimetableUpdateOne) SetNillableAvailabilityBySms(s *string) *TimetableUpdateOne {
	if s != nil {
		tuo.SetAvailabilityBySms(*s)
	}
	return tuo
}

// ClearAvailabilityBySms clears the value of the "availability_by_sms" field.
func (tuo *TimetableUpdateOne) ClearAvailabilityBySms() *TimetableUpdateOne {
	tuo.mutation.ClearAvailabilityBySms()
	return tuo
}

// SetAvailabilityByWhatsapp sets the "availability_by_whatsapp" field.
func (tuo *TimetableUpdateOne) SetAvailabilityByWhatsapp(s string) *TimetableUpdateOne {
	tuo.mutation.SetAvailabilityByWhatsapp(s)
	return tuo
}

// SetNillableAvailabilityByWhatsapp sets the "availability_by_whatsapp" field if the given value is not nil.
func (tuo *TimetableUpdateOne) SetNillableAvailabilityByWhatsapp(s *string) *TimetableUpdateOne {
	if s != nil {
		tuo.SetAvailabilityByWhatsapp(*s)
	}
	return tuo
}

// ClearAvailabilityByWhatsapp clears the value of the "availability_by_whatsapp" field.
func (tuo *TimetableUpdateOne) ClearAvailabilityByWhatsapp() *TimetableUpdateOne {
	tuo.mutation.ClearAvailabilityByWhatsapp()
	return tuo
}

// SetAddressID sets the "address" edge to the Address entity by ID.
func (tuo *TimetableUpdateOne) SetAddressID(id uuid.UUID) *TimetableUpdateOne {
	tuo.mutation.SetAddressID(id)
	return tuo
}

// SetAddress sets the "address" edge to the Address entity.
func (tuo *TimetableUpdateOne) SetAddress(a *Address) *TimetableUpdateOne {
	return tuo.SetAddressID(a.ID)
}

// AddUsersOnDutyIDs adds the "users_on_duty" edge to the User entity by IDs.
func (tuo *TimetableUpdateOne) AddUsersOnDutyIDs(ids ...uuid.UUID) *TimetableUpdateOne {
	tuo.mutation.AddUsersOnDutyIDs(ids...)
	return tuo
}

// AddUsersOnDuty adds the "users_on_duty" edges to the User entity.
func (tuo *TimetableUpdateOne) AddUsersOnDuty(u ...*User) *TimetableUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.AddUsersOnDutyIDs(ids...)
}

// Mutation returns the TimetableMutation object of the builder.
func (tuo *TimetableUpdateOne) Mutation() *TimetableMutation {
	return tuo.mutation
}

// ClearAddress clears the "address" edge to the Address entity.
func (tuo *TimetableUpdateOne) ClearAddress() *TimetableUpdateOne {
	tuo.mutation.ClearAddress()
	return tuo
}

// ClearUsersOnDuty clears all "users_on_duty" edges to the User entity.
func (tuo *TimetableUpdateOne) ClearUsersOnDuty() *TimetableUpdateOne {
	tuo.mutation.ClearUsersOnDuty()
	return tuo
}

// RemoveUsersOnDutyIDs removes the "users_on_duty" edge to User entities by IDs.
func (tuo *TimetableUpdateOne) RemoveUsersOnDutyIDs(ids ...uuid.UUID) *TimetableUpdateOne {
	tuo.mutation.RemoveUsersOnDutyIDs(ids...)
	return tuo
}

// RemoveUsersOnDuty removes "users_on_duty" edges to User entities.
func (tuo *TimetableUpdateOne) RemoveUsersOnDuty(u ...*User) *TimetableUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.RemoveUsersOnDutyIDs(ids...)
}

// Where appends a list predicates to the TimetableUpdate builder.
func (tuo *TimetableUpdateOne) Where(ps ...predicate.Timetable) *TimetableUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TimetableUpdateOne) Select(field string, fields ...string) *TimetableUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Timetable entity.
func (tuo *TimetableUpdateOne) Save(ctx context.Context) (*Timetable, error) {
	if err := tuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TimetableUpdateOne) SaveX(ctx context.Context) *Timetable {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TimetableUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TimetableUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TimetableUpdateOne) defaults() error {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		if timetable.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized timetable.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := timetable.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TimetableUpdateOne) check() error {
	if v, ok := tuo.mutation.TimetableType(); ok {
		if err := timetable.TimetableTypeValidator(v); err != nil {
			return &ValidationError{Name: "timetable_type", err: fmt.Errorf(`ent: validator failed for field "Timetable.timetable_type": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Duration(); ok {
		if err := timetable.DurationValidator(v); err != nil {
			return &ValidationError{Name: "duration", err: fmt.Errorf(`ent: validator failed for field "Timetable.duration": %w`, err)}
		}
	}
	if _, ok := tuo.mutation.AddressID(); tuo.mutation.AddressCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Timetable.address"`)
	}
	return nil
}

func (tuo *TimetableUpdateOne) sqlSave(ctx context.Context) (_node *Timetable, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(timetable.Table, timetable.Columns, sqlgraph.NewFieldSpec(timetable.FieldID, field.TypeUUID))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Timetable.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, timetable.FieldID)
		for _, f := range fields {
			if !timetable.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != timetable.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(timetable.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.DeletedAt(); ok {
		_spec.SetField(timetable.FieldDeletedAt, field.TypeTime, value)
	}
	if tuo.mutation.DeletedAtCleared() {
		_spec.ClearField(timetable.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.TimetableType(); ok {
		_spec.SetField(timetable.FieldTimetableType, field.TypeEnum, value)
	}
	if value, ok := tuo.mutation.DatetimeFrom(); ok {
		_spec.SetField(timetable.FieldDatetimeFrom, field.TypeTime, value)
	}
	if tuo.mutation.DatetimeFromCleared() {
		_spec.ClearField(timetable.FieldDatetimeFrom, field.TypeTime)
	}
	if value, ok := tuo.mutation.Duration(); ok {
		_spec.SetField(timetable.FieldDuration, field.TypeUint8, value)
	}
	if value, ok := tuo.mutation.AddedDuration(); ok {
		_spec.AddField(timetable.FieldDuration, field.TypeUint8, value)
	}
	if tuo.mutation.DurationCleared() {
		_spec.ClearField(timetable.FieldDuration, field.TypeUint8)
	}
	if value, ok := tuo.mutation.DatetimeTo(); ok {
		_spec.SetField(timetable.FieldDatetimeTo, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.TimeWholeDay(); ok {
		_spec.SetField(timetable.FieldTimeWholeDay, field.TypeBool, value)
	}
	if value, ok := tuo.mutation.Comment(); ok {
		_spec.SetField(timetable.FieldComment, field.TypeString, value)
	}
	if tuo.mutation.CommentCleared() {
		_spec.ClearField(timetable.FieldComment, field.TypeString)
	}
	if value, ok := tuo.mutation.AvailabilityByPhone(); ok {
		_spec.SetField(timetable.FieldAvailabilityByPhone, field.TypeString, value)
	}
	if tuo.mutation.AvailabilityByPhoneCleared() {
		_spec.ClearField(timetable.FieldAvailabilityByPhone, field.TypeString)
	}
	if value, ok := tuo.mutation.AvailabilityByEmail(); ok {
		_spec.SetField(timetable.FieldAvailabilityByEmail, field.TypeString, value)
	}
	if tuo.mutation.AvailabilityByEmailCleared() {
		_spec.ClearField(timetable.FieldAvailabilityByEmail, field.TypeString)
	}
	if value, ok := tuo.mutation.AvailabilityBySms(); ok {
		_spec.SetField(timetable.FieldAvailabilityBySms, field.TypeString, value)
	}
	if tuo.mutation.AvailabilityBySmsCleared() {
		_spec.ClearField(timetable.FieldAvailabilityBySms, field.TypeString)
	}
	if value, ok := tuo.mutation.AvailabilityByWhatsapp(); ok {
		_spec.SetField(timetable.FieldAvailabilityByWhatsapp, field.TypeString, value)
	}
	if tuo.mutation.AvailabilityByWhatsappCleared() {
		_spec.ClearField(timetable.FieldAvailabilityByWhatsapp, field.TypeString)
	}
	if tuo.mutation.AddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   timetable.AddressTable,
			Columns: []string{timetable.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.AddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   timetable.AddressTable,
			Columns: []string{timetable.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.UsersOnDutyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   timetable.UsersOnDutyTable,
			Columns: timetable.UsersOnDutyPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedUsersOnDutyIDs(); len(nodes) > 0 && !tuo.mutation.UsersOnDutyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   timetable.UsersOnDutyTable,
			Columns: timetable.UsersOnDutyPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.UsersOnDutyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   timetable.UsersOnDutyTable,
			Columns: timetable.UsersOnDutyPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Timetable{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{timetable.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
