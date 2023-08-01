// OHMAB
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hkonitzer/ohmab/ent/business"
	"github.com/hkonitzer/ohmab/ent/publicuser"
	"github.com/hkonitzer/ohmab/ent/timetable"
)

// PublicUserCreate is the builder for creating a PublicUser entity.
type PublicUserCreate struct {
	config
	mutation *PublicUserMutation
	hooks    []Hook
}

// SetSurname sets the "surname" field.
func (puc *PublicUserCreate) SetSurname(s string) *PublicUserCreate {
	puc.mutation.SetSurname(s)
	return puc
}

// SetFirstname sets the "firstname" field.
func (puc *PublicUserCreate) SetFirstname(s string) *PublicUserCreate {
	puc.mutation.SetFirstname(s)
	return puc
}

// SetTitle sets the "title" field.
func (puc *PublicUserCreate) SetTitle(s string) *PublicUserCreate {
	puc.mutation.SetTitle(s)
	return puc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (puc *PublicUserCreate) SetNillableTitle(s *string) *PublicUserCreate {
	if s != nil {
		puc.SetTitle(*s)
	}
	return puc
}

// SetEmail sets the "email" field.
func (puc *PublicUserCreate) SetEmail(s string) *PublicUserCreate {
	puc.mutation.SetEmail(s)
	return puc
}

// SetID sets the "id" field.
func (puc *PublicUserCreate) SetID(u uuid.UUID) *PublicUserCreate {
	puc.mutation.SetID(u)
	return puc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (puc *PublicUserCreate) SetNillableID(u *uuid.UUID) *PublicUserCreate {
	if u != nil {
		puc.SetID(*u)
	}
	return puc
}

// AddBusinessIDs adds the "businesses" edge to the Business entity by IDs.
func (puc *PublicUserCreate) AddBusinessIDs(ids ...uuid.UUID) *PublicUserCreate {
	puc.mutation.AddBusinessIDs(ids...)
	return puc
}

// AddBusinesses adds the "businesses" edges to the Business entity.
func (puc *PublicUserCreate) AddBusinesses(b ...*Business) *PublicUserCreate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return puc.AddBusinessIDs(ids...)
}

// AddTimetableIDs adds the "timetable" edge to the Timetable entity by IDs.
func (puc *PublicUserCreate) AddTimetableIDs(ids ...uuid.UUID) *PublicUserCreate {
	puc.mutation.AddTimetableIDs(ids...)
	return puc
}

// AddTimetable adds the "timetable" edges to the Timetable entity.
func (puc *PublicUserCreate) AddTimetable(t ...*Timetable) *PublicUserCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puc.AddTimetableIDs(ids...)
}

// Mutation returns the PublicUserMutation object of the builder.
func (puc *PublicUserCreate) Mutation() *PublicUserMutation {
	return puc.mutation
}

// Save creates the PublicUser in the database.
func (puc *PublicUserCreate) Save(ctx context.Context) (*PublicUser, error) {
	if err := puc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, puc.sqlSave, puc.mutation, puc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (puc *PublicUserCreate) SaveX(ctx context.Context) *PublicUser {
	v, err := puc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (puc *PublicUserCreate) Exec(ctx context.Context) error {
	_, err := puc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puc *PublicUserCreate) ExecX(ctx context.Context) {
	if err := puc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puc *PublicUserCreate) defaults() error {
	if _, ok := puc.mutation.ID(); !ok {
		if publicuser.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized publicuser.DefaultID (forgotten import ent/runtime?)")
		}
		v := publicuser.DefaultID()
		puc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (puc *PublicUserCreate) check() error {
	if _, ok := puc.mutation.Surname(); !ok {
		return &ValidationError{Name: "surname", err: errors.New(`ent: missing required field "PublicUser.surname"`)}
	}
	if v, ok := puc.mutation.Surname(); ok {
		if err := publicuser.SurnameValidator(v); err != nil {
			return &ValidationError{Name: "surname", err: fmt.Errorf(`ent: validator failed for field "PublicUser.surname": %w`, err)}
		}
	}
	if _, ok := puc.mutation.Firstname(); !ok {
		return &ValidationError{Name: "firstname", err: errors.New(`ent: missing required field "PublicUser.firstname"`)}
	}
	if _, ok := puc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "PublicUser.email"`)}
	}
	return nil
}

func (puc *PublicUserCreate) sqlSave(ctx context.Context) (*PublicUser, error) {
	if err := puc.check(); err != nil {
		return nil, err
	}
	_node, _spec := puc.createSpec()
	if err := sqlgraph.CreateNode(ctx, puc.driver, _spec); err != nil {
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
	puc.mutation.id = &_node.ID
	puc.mutation.done = true
	return _node, nil
}

func (puc *PublicUserCreate) createSpec() (*PublicUser, *sqlgraph.CreateSpec) {
	var (
		_node = &PublicUser{config: puc.config}
		_spec = sqlgraph.NewCreateSpec(publicuser.Table, sqlgraph.NewFieldSpec(publicuser.FieldID, field.TypeUUID))
	)
	if id, ok := puc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := puc.mutation.Surname(); ok {
		_spec.SetField(publicuser.FieldSurname, field.TypeString, value)
		_node.Surname = value
	}
	if value, ok := puc.mutation.Firstname(); ok {
		_spec.SetField(publicuser.FieldFirstname, field.TypeString, value)
		_node.Firstname = value
	}
	if value, ok := puc.mutation.Title(); ok {
		_spec.SetField(publicuser.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := puc.mutation.Email(); ok {
		_spec.SetField(publicuser.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if nodes := puc.mutation.BusinessesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   publicuser.BusinessesTable,
			Columns: publicuser.BusinessesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := puc.mutation.TimetableIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   publicuser.TimetableTable,
			Columns: publicuser.TimetablePrimaryKey,
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

// PublicUserCreateBulk is the builder for creating many PublicUser entities in bulk.
type PublicUserCreateBulk struct {
	config
	builders []*PublicUserCreate
}

// Save creates the PublicUser entities in the database.
func (pucb *PublicUserCreateBulk) Save(ctx context.Context) ([]*PublicUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pucb.builders))
	nodes := make([]*PublicUser, len(pucb.builders))
	mutators := make([]Mutator, len(pucb.builders))
	for i := range pucb.builders {
		func(i int, root context.Context) {
			builder := pucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PublicUserMutation)
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
					_, err = mutators[i+1].Mutate(root, pucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pucb *PublicUserCreateBulk) SaveX(ctx context.Context) []*PublicUser {
	v, err := pucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pucb *PublicUserCreateBulk) Exec(ctx context.Context) error {
	_, err := pucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pucb *PublicUserCreateBulk) ExecX(ctx context.Context) {
	if err := pucb.Exec(ctx); err != nil {
		panic(err)
	}
}