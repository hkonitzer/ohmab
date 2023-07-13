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
	"github.com/hkonitzer/ohmab/ent/address"
	"github.com/hkonitzer/ohmab/ent/business"
	"github.com/hkonitzer/ohmab/ent/timetable"
)

// AddressCreate is the builder for creating a Address entity.
type AddressCreate struct {
	config
	mutation *AddressMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ac *AddressCreate) SetCreatedAt(t time.Time) *AddressCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AddressCreate) SetNillableCreatedAt(t *time.Time) *AddressCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AddressCreate) SetUpdatedAt(t time.Time) *AddressCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AddressCreate) SetNillableUpdatedAt(t *time.Time) *AddressCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetDeletedAt sets the "deleted_at" field.
func (ac *AddressCreate) SetDeletedAt(t time.Time) *AddressCreate {
	ac.mutation.SetDeletedAt(t)
	return ac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ac *AddressCreate) SetNillableDeletedAt(t *time.Time) *AddressCreate {
	if t != nil {
		ac.SetDeletedAt(*t)
	}
	return ac
}

// SetAddition sets the "addition" field.
func (ac *AddressCreate) SetAddition(s string) *AddressCreate {
	ac.mutation.SetAddition(s)
	return ac
}

// SetNillableAddition sets the "addition" field if the given value is not nil.
func (ac *AddressCreate) SetNillableAddition(s *string) *AddressCreate {
	if s != nil {
		ac.SetAddition(*s)
	}
	return ac
}

// SetStreet sets the "street" field.
func (ac *AddressCreate) SetStreet(s string) *AddressCreate {
	ac.mutation.SetStreet(s)
	return ac
}

// SetNillableStreet sets the "street" field if the given value is not nil.
func (ac *AddressCreate) SetNillableStreet(s *string) *AddressCreate {
	if s != nil {
		ac.SetStreet(*s)
	}
	return ac
}

// SetCity sets the "city" field.
func (ac *AddressCreate) SetCity(s string) *AddressCreate {
	ac.mutation.SetCity(s)
	return ac
}

// SetNillableCity sets the "city" field if the given value is not nil.
func (ac *AddressCreate) SetNillableCity(s *string) *AddressCreate {
	if s != nil {
		ac.SetCity(*s)
	}
	return ac
}

// SetZip sets the "zip" field.
func (ac *AddressCreate) SetZip(s string) *AddressCreate {
	ac.mutation.SetZip(s)
	return ac
}

// SetNillableZip sets the "zip" field if the given value is not nil.
func (ac *AddressCreate) SetNillableZip(s *string) *AddressCreate {
	if s != nil {
		ac.SetZip(*s)
	}
	return ac
}

// SetState sets the "state" field.
func (ac *AddressCreate) SetState(s string) *AddressCreate {
	ac.mutation.SetState(s)
	return ac
}

// SetNillableState sets the "state" field if the given value is not nil.
func (ac *AddressCreate) SetNillableState(s *string) *AddressCreate {
	if s != nil {
		ac.SetState(*s)
	}
	return ac
}

// SetCountry sets the "country" field.
func (ac *AddressCreate) SetCountry(s string) *AddressCreate {
	ac.mutation.SetCountry(s)
	return ac
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (ac *AddressCreate) SetNillableCountry(s *string) *AddressCreate {
	if s != nil {
		ac.SetCountry(*s)
	}
	return ac
}

// SetLocale sets the "locale" field.
func (ac *AddressCreate) SetLocale(s string) *AddressCreate {
	ac.mutation.SetLocale(s)
	return ac
}

// SetNillableLocale sets the "locale" field if the given value is not nil.
func (ac *AddressCreate) SetNillableLocale(s *string) *AddressCreate {
	if s != nil {
		ac.SetLocale(*s)
	}
	return ac
}

// SetPrimary sets the "primary" field.
func (ac *AddressCreate) SetPrimary(b bool) *AddressCreate {
	ac.mutation.SetPrimary(b)
	return ac
}

// SetNillablePrimary sets the "primary" field if the given value is not nil.
func (ac *AddressCreate) SetNillablePrimary(b *bool) *AddressCreate {
	if b != nil {
		ac.SetPrimary(*b)
	}
	return ac
}

// SetTelephone sets the "telephone" field.
func (ac *AddressCreate) SetTelephone(s string) *AddressCreate {
	ac.mutation.SetTelephone(s)
	return ac
}

// SetNillableTelephone sets the "telephone" field if the given value is not nil.
func (ac *AddressCreate) SetNillableTelephone(s *string) *AddressCreate {
	if s != nil {
		ac.SetTelephone(*s)
	}
	return ac
}

// SetComment sets the "comment" field.
func (ac *AddressCreate) SetComment(s string) *AddressCreate {
	ac.mutation.SetComment(s)
	return ac
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (ac *AddressCreate) SetNillableComment(s *string) *AddressCreate {
	if s != nil {
		ac.SetComment(*s)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AddressCreate) SetID(u uuid.UUID) *AddressCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AddressCreate) SetNillableID(u *uuid.UUID) *AddressCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// SetBusinessID sets the "business" edge to the Business entity by ID.
func (ac *AddressCreate) SetBusinessID(id uuid.UUID) *AddressCreate {
	ac.mutation.SetBusinessID(id)
	return ac
}

// SetNillableBusinessID sets the "business" edge to the Business entity by ID if the given value is not nil.
func (ac *AddressCreate) SetNillableBusinessID(id *uuid.UUID) *AddressCreate {
	if id != nil {
		ac = ac.SetBusinessID(*id)
	}
	return ac
}

// SetBusiness sets the "business" edge to the Business entity.
func (ac *AddressCreate) SetBusiness(b *Business) *AddressCreate {
	return ac.SetBusinessID(b.ID)
}

// AddTimetableIDs adds the "timetables" edge to the Timetable entity by IDs.
func (ac *AddressCreate) AddTimetableIDs(ids ...uuid.UUID) *AddressCreate {
	ac.mutation.AddTimetableIDs(ids...)
	return ac
}

// AddTimetables adds the "timetables" edges to the Timetable entity.
func (ac *AddressCreate) AddTimetables(t ...*Timetable) *AddressCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ac.AddTimetableIDs(ids...)
}

// Mutation returns the AddressMutation object of the builder.
func (ac *AddressCreate) Mutation() *AddressMutation {
	return ac.mutation
}

// Save creates the Address in the database.
func (ac *AddressCreate) Save(ctx context.Context) (*Address, error) {
	if err := ac.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AddressCreate) SaveX(ctx context.Context) *Address {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AddressCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AddressCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AddressCreate) defaults() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		if address.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized address.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := address.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		if address.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized address.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := address.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.Locale(); !ok {
		v := address.DefaultLocale
		ac.mutation.SetLocale(v)
	}
	if _, ok := ac.mutation.Primary(); !ok {
		v := address.DefaultPrimary
		ac.mutation.SetPrimary(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		if address.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized address.DefaultID (forgotten import ent/runtime?)")
		}
		v := address.DefaultID()
		ac.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ac *AddressCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Address.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Address.updated_at"`)}
	}
	if _, ok := ac.mutation.Locale(); !ok {
		return &ValidationError{Name: "locale", err: errors.New(`ent: missing required field "Address.locale"`)}
	}
	if v, ok := ac.mutation.Locale(); ok {
		if err := address.LocaleValidator(v); err != nil {
			return &ValidationError{Name: "locale", err: fmt.Errorf(`ent: validator failed for field "Address.locale": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Primary(); !ok {
		return &ValidationError{Name: "primary", err: errors.New(`ent: missing required field "Address.primary"`)}
	}
	return nil
}

func (ac *AddressCreate) sqlSave(ctx context.Context) (*Address, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
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
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AddressCreate) createSpec() (*Address, *sqlgraph.CreateSpec) {
	var (
		_node = &Address{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(address.Table, sqlgraph.NewFieldSpec(address.FieldID, field.TypeUUID))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(address.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(address.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.DeletedAt(); ok {
		_spec.SetField(address.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := ac.mutation.Addition(); ok {
		_spec.SetField(address.FieldAddition, field.TypeString, value)
		_node.Addition = value
	}
	if value, ok := ac.mutation.Street(); ok {
		_spec.SetField(address.FieldStreet, field.TypeString, value)
		_node.Street = value
	}
	if value, ok := ac.mutation.City(); ok {
		_spec.SetField(address.FieldCity, field.TypeString, value)
		_node.City = value
	}
	if value, ok := ac.mutation.Zip(); ok {
		_spec.SetField(address.FieldZip, field.TypeString, value)
		_node.Zip = value
	}
	if value, ok := ac.mutation.State(); ok {
		_spec.SetField(address.FieldState, field.TypeString, value)
		_node.State = value
	}
	if value, ok := ac.mutation.Country(); ok {
		_spec.SetField(address.FieldCountry, field.TypeString, value)
		_node.Country = value
	}
	if value, ok := ac.mutation.Locale(); ok {
		_spec.SetField(address.FieldLocale, field.TypeString, value)
		_node.Locale = value
	}
	if value, ok := ac.mutation.Primary(); ok {
		_spec.SetField(address.FieldPrimary, field.TypeBool, value)
		_node.Primary = value
	}
	if value, ok := ac.mutation.Telephone(); ok {
		_spec.SetField(address.FieldTelephone, field.TypeString, value)
		_node.Telephone = value
	}
	if value, ok := ac.mutation.Comment(); ok {
		_spec.SetField(address.FieldComment, field.TypeString, value)
		_node.Comment = value
	}
	if nodes := ac.mutation.BusinessIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   address.BusinessTable,
			Columns: []string{address.BusinessColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.business_addresses = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.TimetablesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   address.TimetablesTable,
			Columns: []string{address.TimetablesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timetable.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AddressCreateBulk is the builder for creating many Address entities in bulk.
type AddressCreateBulk struct {
	config
	builders []*AddressCreate
}

// Save creates the Address entities in the database.
func (acb *AddressCreateBulk) Save(ctx context.Context) ([]*Address, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Address, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AddressMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AddressCreateBulk) SaveX(ctx context.Context) []*Address {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AddressCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AddressCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
