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
	"github.com/hkonitzer/ohmab/ent/content"
)

// ContentCreate is the builder for creating a Content entity.
type ContentCreate struct {
	config
	mutation *ContentMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *ContentCreate) SetCreatedAt(t time.Time) *ContentCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *ContentCreate) SetNillableCreatedAt(t *time.Time) *ContentCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *ContentCreate) SetUpdatedAt(t time.Time) *ContentCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *ContentCreate) SetNillableUpdatedAt(t *time.Time) *ContentCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *ContentCreate) SetDeletedAt(t time.Time) *ContentCreate {
	cc.mutation.SetDeletedAt(t)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *ContentCreate) SetNillableDeletedAt(t *time.Time) *ContentCreate {
	if t != nil {
		cc.SetDeletedAt(*t)
	}
	return cc
}

// SetTimetableType sets the "timetable_type" field.
func (cc *ContentCreate) SetTimetableType(ct content.TimetableType) *ContentCreate {
	cc.mutation.SetTimetableType(ct)
	return cc
}

// SetNillableTimetableType sets the "timetable_type" field if the given value is not nil.
func (cc *ContentCreate) SetNillableTimetableType(ct *content.TimetableType) *ContentCreate {
	if ct != nil {
		cc.SetTimetableType(*ct)
	}
	return cc
}

// SetType sets the "type" field.
func (cc *ContentCreate) SetType(c content.Type) *ContentCreate {
	cc.mutation.SetType(c)
	return cc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cc *ContentCreate) SetNillableType(c *content.Type) *ContentCreate {
	if c != nil {
		cc.SetType(*c)
	}
	return cc
}

// SetLocale sets the "locale" field.
func (cc *ContentCreate) SetLocale(s string) *ContentCreate {
	cc.mutation.SetLocale(s)
	return cc
}

// SetNillableLocale sets the "locale" field if the given value is not nil.
func (cc *ContentCreate) SetNillableLocale(s *string) *ContentCreate {
	if s != nil {
		cc.SetLocale(*s)
	}
	return cc
}

// SetLocation sets the "location" field.
func (cc *ContentCreate) SetLocation(c content.Location) *ContentCreate {
	cc.mutation.SetLocation(c)
	return cc
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (cc *ContentCreate) SetNillableLocation(c *content.Location) *ContentCreate {
	if c != nil {
		cc.SetLocation(*c)
	}
	return cc
}

// SetContent sets the "content" field.
func (cc *ContentCreate) SetContent(s string) *ContentCreate {
	cc.mutation.SetContent(s)
	return cc
}

// SetStatus sets the "status" field.
func (cc *ContentCreate) SetStatus(c content.Status) *ContentCreate {
	cc.mutation.SetStatus(c)
	return cc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cc *ContentCreate) SetNillableStatus(c *content.Status) *ContentCreate {
	if c != nil {
		cc.SetStatus(*c)
	}
	return cc
}

// SetPublishedAt sets the "published_at" field.
func (cc *ContentCreate) SetPublishedAt(t time.Time) *ContentCreate {
	cc.mutation.SetPublishedAt(t)
	return cc
}

// SetNillablePublishedAt sets the "published_at" field if the given value is not nil.
func (cc *ContentCreate) SetNillablePublishedAt(t *time.Time) *ContentCreate {
	if t != nil {
		cc.SetPublishedAt(*t)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *ContentCreate) SetID(u uuid.UUID) *ContentCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *ContentCreate) SetNillableID(u *uuid.UUID) *ContentCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// Mutation returns the ContentMutation object of the builder.
func (cc *ContentCreate) Mutation() *ContentMutation {
	return cc.mutation
}

// Save creates the Content in the database.
func (cc *ContentCreate) Save(ctx context.Context) (*Content, error) {
	if err := cc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ContentCreate) SaveX(ctx context.Context) *Content {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ContentCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ContentCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ContentCreate) defaults() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		if content.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized content.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := content.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		if content.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized content.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := content.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.TimetableType(); !ok {
		v := content.DefaultTimetableType
		cc.mutation.SetTimetableType(v)
	}
	if _, ok := cc.mutation.GetType(); !ok {
		v := content.DefaultType
		cc.mutation.SetType(v)
	}
	if _, ok := cc.mutation.Locale(); !ok {
		v := content.DefaultLocale
		cc.mutation.SetLocale(v)
	}
	if _, ok := cc.mutation.Location(); !ok {
		v := content.DefaultLocation
		cc.mutation.SetLocation(v)
	}
	if _, ok := cc.mutation.Status(); !ok {
		v := content.DefaultStatus
		cc.mutation.SetStatus(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		if content.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized content.DefaultID (forgotten import ent/runtime?)")
		}
		v := content.DefaultID()
		cc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cc *ContentCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Content.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Content.updated_at"`)}
	}
	if _, ok := cc.mutation.TimetableType(); !ok {
		return &ValidationError{Name: "timetable_type", err: errors.New(`ent: missing required field "Content.timetable_type"`)}
	}
	if v, ok := cc.mutation.TimetableType(); ok {
		if err := content.TimetableTypeValidator(v); err != nil {
			return &ValidationError{Name: "timetable_type", err: fmt.Errorf(`ent: validator failed for field "Content.timetable_type": %w`, err)}
		}
	}
	if _, ok := cc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Content.type"`)}
	}
	if v, ok := cc.mutation.GetType(); ok {
		if err := content.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Content.type": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Locale(); !ok {
		return &ValidationError{Name: "locale", err: errors.New(`ent: missing required field "Content.locale"`)}
	}
	if v, ok := cc.mutation.Locale(); ok {
		if err := content.LocaleValidator(v); err != nil {
			return &ValidationError{Name: "locale", err: fmt.Errorf(`ent: validator failed for field "Content.locale": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Location(); !ok {
		return &ValidationError{Name: "location", err: errors.New(`ent: missing required field "Content.location"`)}
	}
	if v, ok := cc.mutation.Location(); ok {
		if err := content.LocationValidator(v); err != nil {
			return &ValidationError{Name: "location", err: fmt.Errorf(`ent: validator failed for field "Content.location": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Content.content"`)}
	}
	if _, ok := cc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Content.status"`)}
	}
	if v, ok := cc.mutation.Status(); ok {
		if err := content.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Content.status": %w`, err)}
		}
	}
	return nil
}

func (cc *ContentCreate) sqlSave(ctx context.Context) (*Content, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ContentCreate) createSpec() (*Content, *sqlgraph.CreateSpec) {
	var (
		_node = &Content{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(content.Table, sqlgraph.NewFieldSpec(content.FieldID, field.TypeUUID))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(content.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(content.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.SetField(content.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := cc.mutation.TimetableType(); ok {
		_spec.SetField(content.FieldTimetableType, field.TypeEnum, value)
		_node.TimetableType = value
	}
	if value, ok := cc.mutation.GetType(); ok {
		_spec.SetField(content.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := cc.mutation.Locale(); ok {
		_spec.SetField(content.FieldLocale, field.TypeString, value)
		_node.Locale = value
	}
	if value, ok := cc.mutation.Location(); ok {
		_spec.SetField(content.FieldLocation, field.TypeEnum, value)
		_node.Location = value
	}
	if value, ok := cc.mutation.Content(); ok {
		_spec.SetField(content.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := cc.mutation.Status(); ok {
		_spec.SetField(content.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := cc.mutation.PublishedAt(); ok {
		_spec.SetField(content.FieldPublishedAt, field.TypeTime, value)
		_node.PublishedAt = value
	}
	return _node, _spec
}

// ContentCreateBulk is the builder for creating many Content entities in bulk.
type ContentCreateBulk struct {
	config
	builders []*ContentCreate
}

// Save creates the Content entities in the database.
func (ccb *ContentCreateBulk) Save(ctx context.Context) ([]*Content, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Content, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ContentMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ContentCreateBulk) SaveX(ctx context.Context) []*Content {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ContentCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ContentCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
