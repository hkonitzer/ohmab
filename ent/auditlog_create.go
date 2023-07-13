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
	"github.com/hkonitzer/ohmab/ent/auditlog"
)

// AuditLogCreate is the builder for creating a AuditLog entity.
type AuditLogCreate struct {
	config
	mutation *AuditLogMutation
	hooks    []Hook
}

// SetUser sets the "user" field.
func (alc *AuditLogCreate) SetUser(s string) *AuditLogCreate {
	alc.mutation.SetUser(s)
	return alc
}

// SetAction sets the "action" field.
func (alc *AuditLogCreate) SetAction(s string) *AuditLogCreate {
	alc.mutation.SetAction(s)
	return alc
}

// SetEntitySchema sets the "entity_schema" field.
func (alc *AuditLogCreate) SetEntitySchema(s string) *AuditLogCreate {
	alc.mutation.SetEntitySchema(s)
	return alc
}

// SetEntityValues sets the "entity_values" field.
func (alc *AuditLogCreate) SetEntityValues(m map[string]string) *AuditLogCreate {
	alc.mutation.SetEntityValues(m)
	return alc
}

// SetEntityUUID sets the "entity_uuid" field.
func (alc *AuditLogCreate) SetEntityUUID(s string) *AuditLogCreate {
	alc.mutation.SetEntityUUID(s)
	return alc
}

// SetNillableEntityUUID sets the "entity_uuid" field if the given value is not nil.
func (alc *AuditLogCreate) SetNillableEntityUUID(s *string) *AuditLogCreate {
	if s != nil {
		alc.SetEntityUUID(*s)
	}
	return alc
}

// SetTimestamp sets the "timestamp" field.
func (alc *AuditLogCreate) SetTimestamp(t time.Time) *AuditLogCreate {
	alc.mutation.SetTimestamp(t)
	return alc
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (alc *AuditLogCreate) SetNillableTimestamp(t *time.Time) *AuditLogCreate {
	if t != nil {
		alc.SetTimestamp(*t)
	}
	return alc
}

// SetID sets the "id" field.
func (alc *AuditLogCreate) SetID(u uuid.UUID) *AuditLogCreate {
	alc.mutation.SetID(u)
	return alc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (alc *AuditLogCreate) SetNillableID(u *uuid.UUID) *AuditLogCreate {
	if u != nil {
		alc.SetID(*u)
	}
	return alc
}

// Mutation returns the AuditLogMutation object of the builder.
func (alc *AuditLogCreate) Mutation() *AuditLogMutation {
	return alc.mutation
}

// Save creates the AuditLog in the database.
func (alc *AuditLogCreate) Save(ctx context.Context) (*AuditLog, error) {
	alc.defaults()
	return withHooks(ctx, alc.sqlSave, alc.mutation, alc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (alc *AuditLogCreate) SaveX(ctx context.Context) *AuditLog {
	v, err := alc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (alc *AuditLogCreate) Exec(ctx context.Context) error {
	_, err := alc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alc *AuditLogCreate) ExecX(ctx context.Context) {
	if err := alc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (alc *AuditLogCreate) defaults() {
	if _, ok := alc.mutation.Timestamp(); !ok {
		v := auditlog.DefaultTimestamp()
		alc.mutation.SetTimestamp(v)
	}
	if _, ok := alc.mutation.ID(); !ok {
		v := auditlog.DefaultID()
		alc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (alc *AuditLogCreate) check() error {
	if _, ok := alc.mutation.User(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required field "AuditLog.user"`)}
	}
	if v, ok := alc.mutation.User(); ok {
		if err := auditlog.UserValidator(v); err != nil {
			return &ValidationError{Name: "user", err: fmt.Errorf(`ent: validator failed for field "AuditLog.user": %w`, err)}
		}
	}
	if _, ok := alc.mutation.Action(); !ok {
		return &ValidationError{Name: "action", err: errors.New(`ent: missing required field "AuditLog.action"`)}
	}
	if v, ok := alc.mutation.Action(); ok {
		if err := auditlog.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf(`ent: validator failed for field "AuditLog.action": %w`, err)}
		}
	}
	if _, ok := alc.mutation.EntitySchema(); !ok {
		return &ValidationError{Name: "entity_schema", err: errors.New(`ent: missing required field "AuditLog.entity_schema"`)}
	}
	if v, ok := alc.mutation.EntitySchema(); ok {
		if err := auditlog.EntitySchemaValidator(v); err != nil {
			return &ValidationError{Name: "entity_schema", err: fmt.Errorf(`ent: validator failed for field "AuditLog.entity_schema": %w`, err)}
		}
	}
	if _, ok := alc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "timestamp", err: errors.New(`ent: missing required field "AuditLog.timestamp"`)}
	}
	return nil
}

func (alc *AuditLogCreate) sqlSave(ctx context.Context) (*AuditLog, error) {
	if err := alc.check(); err != nil {
		return nil, err
	}
	_node, _spec := alc.createSpec()
	if err := sqlgraph.CreateNode(ctx, alc.driver, _spec); err != nil {
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
	alc.mutation.id = &_node.ID
	alc.mutation.done = true
	return _node, nil
}

func (alc *AuditLogCreate) createSpec() (*AuditLog, *sqlgraph.CreateSpec) {
	var (
		_node = &AuditLog{config: alc.config}
		_spec = sqlgraph.NewCreateSpec(auditlog.Table, sqlgraph.NewFieldSpec(auditlog.FieldID, field.TypeUUID))
	)
	if id, ok := alc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := alc.mutation.User(); ok {
		_spec.SetField(auditlog.FieldUser, field.TypeString, value)
		_node.User = value
	}
	if value, ok := alc.mutation.Action(); ok {
		_spec.SetField(auditlog.FieldAction, field.TypeString, value)
		_node.Action = value
	}
	if value, ok := alc.mutation.EntitySchema(); ok {
		_spec.SetField(auditlog.FieldEntitySchema, field.TypeString, value)
		_node.EntitySchema = value
	}
	if value, ok := alc.mutation.EntityValues(); ok {
		_spec.SetField(auditlog.FieldEntityValues, field.TypeJSON, value)
		_node.EntityValues = value
	}
	if value, ok := alc.mutation.EntityUUID(); ok {
		_spec.SetField(auditlog.FieldEntityUUID, field.TypeString, value)
		_node.EntityUUID = value
	}
	if value, ok := alc.mutation.Timestamp(); ok {
		_spec.SetField(auditlog.FieldTimestamp, field.TypeTime, value)
		_node.Timestamp = value
	}
	return _node, _spec
}

// AuditLogCreateBulk is the builder for creating many AuditLog entities in bulk.
type AuditLogCreateBulk struct {
	config
	builders []*AuditLogCreate
}

// Save creates the AuditLog entities in the database.
func (alcb *AuditLogCreateBulk) Save(ctx context.Context) ([]*AuditLog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(alcb.builders))
	nodes := make([]*AuditLog, len(alcb.builders))
	mutators := make([]Mutator, len(alcb.builders))
	for i := range alcb.builders {
		func(i int, root context.Context) {
			builder := alcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuditLogMutation)
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
					_, err = mutators[i+1].Mutate(root, alcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, alcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, alcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (alcb *AuditLogCreateBulk) SaveX(ctx context.Context) []*AuditLog {
	v, err := alcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (alcb *AuditLogCreateBulk) Exec(ctx context.Context) error {
	_, err := alcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alcb *AuditLogCreateBulk) ExecX(ctx context.Context) {
	if err := alcb.Exec(ctx); err != nil {
		panic(err)
	}
}
