// OHMAB
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"hynie.de/ohmab/ent/address"
	"hynie.de/ohmab/ent/timetable"
	"hynie.de/ohmab/ent/user"
)

// TimetableCreate is the builder for creating a Timetable entity.
type TimetableCreate struct {
	config
	mutation *TimetableMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (tc *TimetableCreate) SetCreatedAt(t time.Time) *TimetableCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TimetableCreate) SetNillableCreatedAt(t *time.Time) *TimetableCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TimetableCreate) SetUpdatedAt(t time.Time) *TimetableCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TimetableCreate) SetNillableUpdatedAt(t *time.Time) *TimetableCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetDeletedAt sets the "deleted_at" field.
func (tc *TimetableCreate) SetDeletedAt(t time.Time) *TimetableCreate {
	tc.mutation.SetDeletedAt(t)
	return tc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tc *TimetableCreate) SetNillableDeletedAt(t *time.Time) *TimetableCreate {
	if t != nil {
		tc.SetDeletedAt(*t)
	}
	return tc
}

// SetType sets the "type" field.
func (tc *TimetableCreate) SetType(t timetable.Type) *TimetableCreate {
	tc.mutation.SetType(t)
	return tc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tc *TimetableCreate) SetNillableType(t *timetable.Type) *TimetableCreate {
	if t != nil {
		tc.SetType(*t)
	}
	return tc
}

// SetDatetimeFrom sets the "datetime_from" field.
func (tc *TimetableCreate) SetDatetimeFrom(t time.Time) *TimetableCreate {
	tc.mutation.SetDatetimeFrom(t)
	return tc
}

// SetNillableDatetimeFrom sets the "datetime_from" field if the given value is not nil.
func (tc *TimetableCreate) SetNillableDatetimeFrom(t *time.Time) *TimetableCreate {
	if t != nil {
		tc.SetDatetimeFrom(*t)
	}
	return tc
}

// SetDatetimeTo sets the "datetime_to" field.
func (tc *TimetableCreate) SetDatetimeTo(t time.Time) *TimetableCreate {
	tc.mutation.SetDatetimeTo(t)
	return tc
}

// SetNillableDatetimeTo sets the "datetime_to" field if the given value is not nil.
func (tc *TimetableCreate) SetNillableDatetimeTo(t *time.Time) *TimetableCreate {
	if t != nil {
		tc.SetDatetimeTo(*t)
	}
	return tc
}

// SetTimeWholeDay sets the "time_whole_day" field.
func (tc *TimetableCreate) SetTimeWholeDay(b bool) *TimetableCreate {
	tc.mutation.SetTimeWholeDay(b)
	return tc
}

// SetNillableTimeWholeDay sets the "time_whole_day" field if the given value is not nil.
func (tc *TimetableCreate) SetNillableTimeWholeDay(b *bool) *TimetableCreate {
	if b != nil {
		tc.SetTimeWholeDay(*b)
	}
	return tc
}

// SetComment sets the "comment" field.
func (tc *TimetableCreate) SetComment(s string) *TimetableCreate {
	tc.mutation.SetComment(s)
	return tc
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (tc *TimetableCreate) SetNillableComment(s *string) *TimetableCreate {
	if s != nil {
		tc.SetComment(*s)
	}
	return tc
}

// SetAvailabilityByPhone sets the "availability_by_phone" field.
func (tc *TimetableCreate) SetAvailabilityByPhone(s string) *TimetableCreate {
	tc.mutation.SetAvailabilityByPhone(s)
	return tc
}

// SetNillableAvailabilityByPhone sets the "availability_by_phone" field if the given value is not nil.
func (tc *TimetableCreate) SetNillableAvailabilityByPhone(s *string) *TimetableCreate {
	if s != nil {
		tc.SetAvailabilityByPhone(*s)
	}
	return tc
}

// SetAvailabilityByEmail sets the "availability_by_email" field.
func (tc *TimetableCreate) SetAvailabilityByEmail(s string) *TimetableCreate {
	tc.mutation.SetAvailabilityByEmail(s)
	return tc
}

// SetNillableAvailabilityByEmail sets the "availability_by_email" field if the given value is not nil.
func (tc *TimetableCreate) SetNillableAvailabilityByEmail(s *string) *TimetableCreate {
	if s != nil {
		tc.SetAvailabilityByEmail(*s)
	}
	return tc
}

// SetAvailabilityBySms sets the "availability_by_sms" field.
func (tc *TimetableCreate) SetAvailabilityBySms(s string) *TimetableCreate {
	tc.mutation.SetAvailabilityBySms(s)
	return tc
}

// SetNillableAvailabilityBySms sets the "availability_by_sms" field if the given value is not nil.
func (tc *TimetableCreate) SetNillableAvailabilityBySms(s *string) *TimetableCreate {
	if s != nil {
		tc.SetAvailabilityBySms(*s)
	}
	return tc
}

// SetAvailabilityByWhatsapp sets the "availability_by_whatsapp" field.
func (tc *TimetableCreate) SetAvailabilityByWhatsapp(s string) *TimetableCreate {
	tc.mutation.SetAvailabilityByWhatsapp(s)
	return tc
}

// SetNillableAvailabilityByWhatsapp sets the "availability_by_whatsapp" field if the given value is not nil.
func (tc *TimetableCreate) SetNillableAvailabilityByWhatsapp(s *string) *TimetableCreate {
	if s != nil {
		tc.SetAvailabilityByWhatsapp(*s)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TimetableCreate) SetID(u uuid.UUID) *TimetableCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TimetableCreate) SetNillableID(u *uuid.UUID) *TimetableCreate {
	if u != nil {
		tc.SetID(*u)
	}
	return tc
}

// SetAddressID sets the "address" edge to the Address entity by ID.
func (tc *TimetableCreate) SetAddressID(id uuid.UUID) *TimetableCreate {
	tc.mutation.SetAddressID(id)
	return tc
}

// SetAddress sets the "address" edge to the Address entity.
func (tc *TimetableCreate) SetAddress(a *Address) *TimetableCreate {
	return tc.SetAddressID(a.ID)
}

// AddUsersOnDutyIDs adds the "users_on_duty" edge to the User entity by IDs.
func (tc *TimetableCreate) AddUsersOnDutyIDs(ids ...uuid.UUID) *TimetableCreate {
	tc.mutation.AddUsersOnDutyIDs(ids...)
	return tc
}

// AddUsersOnDuty adds the "users_on_duty" edges to the User entity.
func (tc *TimetableCreate) AddUsersOnDuty(u ...*User) *TimetableCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tc.AddUsersOnDutyIDs(ids...)
}

// Mutation returns the TimetableMutation object of the builder.
func (tc *TimetableCreate) Mutation() *TimetableMutation {
	return tc.mutation
}

// Save creates the Timetable in the database.
func (tc *TimetableCreate) Save(ctx context.Context) (*Timetable, error) {
	if err := tc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TimetableCreate) SaveX(ctx context.Context) *Timetable {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TimetableCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TimetableCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TimetableCreate) defaults() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		if timetable.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized timetable.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := timetable.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		if timetable.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized timetable.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := timetable.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tc.mutation.GetType(); !ok {
		v := timetable.DefaultType
		tc.mutation.SetType(v)
	}
	if _, ok := tc.mutation.TimeWholeDay(); !ok {
		v := timetable.DefaultTimeWholeDay
		tc.mutation.SetTimeWholeDay(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		if timetable.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized timetable.DefaultID (forgotten import ent/runtime?)")
		}
		v := timetable.DefaultID()
		tc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tc *TimetableCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Timetable.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Timetable.updated_at"`)}
	}
	if _, ok := tc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Timetable.type"`)}
	}
	if v, ok := tc.mutation.GetType(); ok {
		if err := timetable.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Timetable.type": %w`, err)}
		}
	}
	if _, ok := tc.mutation.TimeWholeDay(); !ok {
		return &ValidationError{Name: "time_whole_day", err: errors.New(`ent: missing required field "Timetable.time_whole_day"`)}
	}
	if _, ok := tc.mutation.AddressID(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required edge "Timetable.address"`)}
	}
	return nil
}

func (tc *TimetableCreate) sqlSave(ctx context.Context) (*Timetable, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TimetableCreate) createSpec() (*Timetable, *sqlgraph.CreateSpec) {
	var (
		_node = &Timetable{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(timetable.Table, sqlgraph.NewFieldSpec(timetable.FieldID, field.TypeUUID))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(timetable.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(timetable.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tc.mutation.DeletedAt(); ok {
		_spec.SetField(timetable.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := tc.mutation.GetType(); ok {
		_spec.SetField(timetable.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := tc.mutation.DatetimeFrom(); ok {
		_spec.SetField(timetable.FieldDatetimeFrom, field.TypeTime, value)
		_node.DatetimeFrom = value
	}
	if value, ok := tc.mutation.DatetimeTo(); ok {
		_spec.SetField(timetable.FieldDatetimeTo, field.TypeTime, value)
		_node.DatetimeTo = value
	}
	if value, ok := tc.mutation.TimeWholeDay(); ok {
		_spec.SetField(timetable.FieldTimeWholeDay, field.TypeBool, value)
		_node.TimeWholeDay = value
	}
	if value, ok := tc.mutation.Comment(); ok {
		_spec.SetField(timetable.FieldComment, field.TypeString, value)
		_node.Comment = value
	}
	if value, ok := tc.mutation.AvailabilityByPhone(); ok {
		_spec.SetField(timetable.FieldAvailabilityByPhone, field.TypeString, value)
		_node.AvailabilityByPhone = value
	}
	if value, ok := tc.mutation.AvailabilityByEmail(); ok {
		_spec.SetField(timetable.FieldAvailabilityByEmail, field.TypeString, value)
		_node.AvailabilityByEmail = value
	}
	if value, ok := tc.mutation.AvailabilityBySms(); ok {
		_spec.SetField(timetable.FieldAvailabilityBySms, field.TypeString, value)
		_node.AvailabilityBySms = value
	}
	if value, ok := tc.mutation.AvailabilityByWhatsapp(); ok {
		_spec.SetField(timetable.FieldAvailabilityByWhatsapp, field.TypeString, value)
		_node.AvailabilityByWhatsapp = value
	}
	if nodes := tc.mutation.AddressIDs(); len(nodes) > 0 {
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
		_node.address_timetables = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.UsersOnDutyIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TimetableCreateBulk is the builder for creating many Timetable entities in bulk.
type TimetableCreateBulk struct {
	config
	builders []*TimetableCreate
}

// Save creates the Timetable entities in the database.
func (tcb *TimetableCreateBulk) Save(ctx context.Context) ([]*Timetable, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Timetable, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TimetableMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TimetableCreateBulk) SaveX(ctx context.Context) []*Timetable {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TimetableCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TimetableCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
