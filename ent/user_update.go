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
	"hynie.de/ohmab/ent/business"
	"hynie.de/ohmab/ent/predicate"
	"hynie.de/ohmab/ent/tag"
	"hynie.de/ohmab/ent/timetable"
	"hynie.de/ohmab/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetDeletedAt sets the "deleted_at" field.
func (uu *UserUpdate) SetDeletedAt(t time.Time) *UserUpdate {
	uu.mutation.SetDeletedAt(t)
	return uu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDeletedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetDeletedAt(*t)
	}
	return uu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (uu *UserUpdate) ClearDeletedAt() *UserUpdate {
	uu.mutation.ClearDeletedAt()
	return uu
}

// SetSurname sets the "surname" field.
func (uu *UserUpdate) SetSurname(s string) *UserUpdate {
	uu.mutation.SetSurname(s)
	return uu
}

// SetFirstname sets the "firstname" field.
func (uu *UserUpdate) SetFirstname(s string) *UserUpdate {
	uu.mutation.SetFirstname(s)
	return uu
}

// SetTitle sets the "title" field.
func (uu *UserUpdate) SetTitle(s string) *UserUpdate {
	uu.mutation.SetTitle(s)
	return uu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (uu *UserUpdate) SetNillableTitle(s *string) *UserUpdate {
	if s != nil {
		uu.SetTitle(*s)
	}
	return uu
}

// ClearTitle clears the value of the "title" field.
func (uu *UserUpdate) ClearTitle() *UserUpdate {
	uu.mutation.ClearTitle()
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetPasswordhash sets the "passwordhash" field.
func (uu *UserUpdate) SetPasswordhash(s string) *UserUpdate {
	uu.mutation.SetPasswordhash(s)
	return uu
}

// SetNillablePasswordhash sets the "passwordhash" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePasswordhash(s *string) *UserUpdate {
	if s != nil {
		uu.SetPasswordhash(*s)
	}
	return uu
}

// ClearPasswordhash clears the value of the "passwordhash" field.
func (uu *UserUpdate) ClearPasswordhash() *UserUpdate {
	uu.mutation.ClearPasswordhash()
	return uu
}

// SetComment sets the "comment" field.
func (uu *UserUpdate) SetComment(s string) *UserUpdate {
	uu.mutation.SetComment(s)
	return uu
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (uu *UserUpdate) SetNillableComment(s *string) *UserUpdate {
	if s != nil {
		uu.SetComment(*s)
	}
	return uu
}

// ClearComment clears the value of the "comment" field.
func (uu *UserUpdate) ClearComment() *UserUpdate {
	uu.mutation.ClearComment()
	return uu
}

// SetActive sets the "active" field.
func (uu *UserUpdate) SetActive(b bool) *UserUpdate {
	uu.mutation.SetActive(b)
	return uu
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (uu *UserUpdate) SetNillableActive(b *bool) *UserUpdate {
	if b != nil {
		uu.SetActive(*b)
	}
	return uu
}

// SetRole sets the "role" field.
func (uu *UserUpdate) SetRole(i int) *UserUpdate {
	uu.mutation.ResetRole()
	uu.mutation.SetRole(i)
	return uu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRole(i *int) *UserUpdate {
	if i != nil {
		uu.SetRole(*i)
	}
	return uu
}

// AddRole adds i to the "role" field.
func (uu *UserUpdate) AddRole(i int) *UserUpdate {
	uu.mutation.AddRole(i)
	return uu
}

// AddBusinessIDs adds the "businesses" edge to the Business entity by IDs.
func (uu *UserUpdate) AddBusinessIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddBusinessIDs(ids...)
	return uu
}

// AddBusinesses adds the "businesses" edges to the Business entity.
func (uu *UserUpdate) AddBusinesses(b ...*Business) *UserUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uu.AddBusinessIDs(ids...)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (uu *UserUpdate) AddTagIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddTagIDs(ids...)
	return uu
}

// AddTags adds the "tags" edges to the Tag entity.
func (uu *UserUpdate) AddTags(t ...*Tag) *UserUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.AddTagIDs(ids...)
}

// AddTimetableIDs adds the "timetable" edge to the Timetable entity by IDs.
func (uu *UserUpdate) AddTimetableIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddTimetableIDs(ids...)
	return uu
}

// AddTimetable adds the "timetable" edges to the Timetable entity.
func (uu *UserUpdate) AddTimetable(t ...*Timetable) *UserUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.AddTimetableIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearBusinesses clears all "businesses" edges to the Business entity.
func (uu *UserUpdate) ClearBusinesses() *UserUpdate {
	uu.mutation.ClearBusinesses()
	return uu
}

// RemoveBusinessIDs removes the "businesses" edge to Business entities by IDs.
func (uu *UserUpdate) RemoveBusinessIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveBusinessIDs(ids...)
	return uu
}

// RemoveBusinesses removes "businesses" edges to Business entities.
func (uu *UserUpdate) RemoveBusinesses(b ...*Business) *UserUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uu.RemoveBusinessIDs(ids...)
}

// ClearTags clears all "tags" edges to the Tag entity.
func (uu *UserUpdate) ClearTags() *UserUpdate {
	uu.mutation.ClearTags()
	return uu
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (uu *UserUpdate) RemoveTagIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveTagIDs(ids...)
	return uu
}

// RemoveTags removes "tags" edges to Tag entities.
func (uu *UserUpdate) RemoveTags(t ...*Tag) *UserUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.RemoveTagIDs(ids...)
}

// ClearTimetable clears all "timetable" edges to the Timetable entity.
func (uu *UserUpdate) ClearTimetable() *UserUpdate {
	uu.mutation.ClearTimetable()
	return uu
}

// RemoveTimetableIDs removes the "timetable" edge to Timetable entities by IDs.
func (uu *UserUpdate) RemoveTimetableIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveTimetableIDs(ids...)
	return uu
}

// RemoveTimetable removes "timetable" edges to Timetable entities.
func (uu *UserUpdate) RemoveTimetable(t ...*Timetable) *UserUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.RemoveTimetableIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if err := uu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() error {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		if user.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized user.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Surname(); ok {
		if err := user.SurnameValidator(v); err != nil {
			return &ValidationError{Name: "surname", err: fmt.Errorf(`ent: validator failed for field "User.surname": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeTime, value)
	}
	if uu.mutation.DeletedAtCleared() {
		_spec.ClearField(user.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := uu.mutation.Surname(); ok {
		_spec.SetField(user.FieldSurname, field.TypeString, value)
	}
	if value, ok := uu.mutation.Firstname(); ok {
		_spec.SetField(user.FieldFirstname, field.TypeString, value)
	}
	if value, ok := uu.mutation.Title(); ok {
		_spec.SetField(user.FieldTitle, field.TypeString, value)
	}
	if uu.mutation.TitleCleared() {
		_spec.ClearField(user.FieldTitle, field.TypeString)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Passwordhash(); ok {
		_spec.SetField(user.FieldPasswordhash, field.TypeString, value)
	}
	if uu.mutation.PasswordhashCleared() {
		_spec.ClearField(user.FieldPasswordhash, field.TypeString)
	}
	if value, ok := uu.mutation.Comment(); ok {
		_spec.SetField(user.FieldComment, field.TypeString, value)
	}
	if uu.mutation.CommentCleared() {
		_spec.ClearField(user.FieldComment, field.TypeString)
	}
	if value, ok := uu.mutation.Active(); ok {
		_spec.SetField(user.FieldActive, field.TypeBool, value)
	}
	if value, ok := uu.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedRole(); ok {
		_spec.AddField(user.FieldRole, field.TypeInt, value)
	}
	if uu.mutation.BusinessesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.BusinessesTable,
			Columns: []string{user.BusinessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedBusinessesIDs(); len(nodes) > 0 && !uu.mutation.BusinessesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.BusinessesTable,
			Columns: []string{user.BusinessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.BusinessesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.BusinessesTable,
			Columns: []string{user.BusinessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TagsTable,
			Columns: user.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedTagsIDs(); len(nodes) > 0 && !uu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TagsTable,
			Columns: user.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TagsTable,
			Columns: user.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.TimetableCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.TimetableTable,
			Columns: user.TimetablePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timetable.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedTimetableIDs(); len(nodes) > 0 && !uu.mutation.TimetableCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.TimetableTable,
			Columns: user.TimetablePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timetable.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.TimetableIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.TimetableTable,
			Columns: user.TimetablePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timetable.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetDeletedAt sets the "deleted_at" field.
func (uuo *UserUpdateOne) SetDeletedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetDeletedAt(t)
	return uuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDeletedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetDeletedAt(*t)
	}
	return uuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (uuo *UserUpdateOne) ClearDeletedAt() *UserUpdateOne {
	uuo.mutation.ClearDeletedAt()
	return uuo
}

// SetSurname sets the "surname" field.
func (uuo *UserUpdateOne) SetSurname(s string) *UserUpdateOne {
	uuo.mutation.SetSurname(s)
	return uuo
}

// SetFirstname sets the "firstname" field.
func (uuo *UserUpdateOne) SetFirstname(s string) *UserUpdateOne {
	uuo.mutation.SetFirstname(s)
	return uuo
}

// SetTitle sets the "title" field.
func (uuo *UserUpdateOne) SetTitle(s string) *UserUpdateOne {
	uuo.mutation.SetTitle(s)
	return uuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTitle(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetTitle(*s)
	}
	return uuo
}

// ClearTitle clears the value of the "title" field.
func (uuo *UserUpdateOne) ClearTitle() *UserUpdateOne {
	uuo.mutation.ClearTitle()
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetPasswordhash sets the "passwordhash" field.
func (uuo *UserUpdateOne) SetPasswordhash(s string) *UserUpdateOne {
	uuo.mutation.SetPasswordhash(s)
	return uuo
}

// SetNillablePasswordhash sets the "passwordhash" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePasswordhash(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPasswordhash(*s)
	}
	return uuo
}

// ClearPasswordhash clears the value of the "passwordhash" field.
func (uuo *UserUpdateOne) ClearPasswordhash() *UserUpdateOne {
	uuo.mutation.ClearPasswordhash()
	return uuo
}

// SetComment sets the "comment" field.
func (uuo *UserUpdateOne) SetComment(s string) *UserUpdateOne {
	uuo.mutation.SetComment(s)
	return uuo
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableComment(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetComment(*s)
	}
	return uuo
}

// ClearComment clears the value of the "comment" field.
func (uuo *UserUpdateOne) ClearComment() *UserUpdateOne {
	uuo.mutation.ClearComment()
	return uuo
}

// SetActive sets the "active" field.
func (uuo *UserUpdateOne) SetActive(b bool) *UserUpdateOne {
	uuo.mutation.SetActive(b)
	return uuo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableActive(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetActive(*b)
	}
	return uuo
}

// SetRole sets the "role" field.
func (uuo *UserUpdateOne) SetRole(i int) *UserUpdateOne {
	uuo.mutation.ResetRole()
	uuo.mutation.SetRole(i)
	return uuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRole(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetRole(*i)
	}
	return uuo
}

// AddRole adds i to the "role" field.
func (uuo *UserUpdateOne) AddRole(i int) *UserUpdateOne {
	uuo.mutation.AddRole(i)
	return uuo
}

// AddBusinessIDs adds the "businesses" edge to the Business entity by IDs.
func (uuo *UserUpdateOne) AddBusinessIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddBusinessIDs(ids...)
	return uuo
}

// AddBusinesses adds the "businesses" edges to the Business entity.
func (uuo *UserUpdateOne) AddBusinesses(b ...*Business) *UserUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uuo.AddBusinessIDs(ids...)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (uuo *UserUpdateOne) AddTagIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddTagIDs(ids...)
	return uuo
}

// AddTags adds the "tags" edges to the Tag entity.
func (uuo *UserUpdateOne) AddTags(t ...*Tag) *UserUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.AddTagIDs(ids...)
}

// AddTimetableIDs adds the "timetable" edge to the Timetable entity by IDs.
func (uuo *UserUpdateOne) AddTimetableIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddTimetableIDs(ids...)
	return uuo
}

// AddTimetable adds the "timetable" edges to the Timetable entity.
func (uuo *UserUpdateOne) AddTimetable(t ...*Timetable) *UserUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.AddTimetableIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearBusinesses clears all "businesses" edges to the Business entity.
func (uuo *UserUpdateOne) ClearBusinesses() *UserUpdateOne {
	uuo.mutation.ClearBusinesses()
	return uuo
}

// RemoveBusinessIDs removes the "businesses" edge to Business entities by IDs.
func (uuo *UserUpdateOne) RemoveBusinessIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveBusinessIDs(ids...)
	return uuo
}

// RemoveBusinesses removes "businesses" edges to Business entities.
func (uuo *UserUpdateOne) RemoveBusinesses(b ...*Business) *UserUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uuo.RemoveBusinessIDs(ids...)
}

// ClearTags clears all "tags" edges to the Tag entity.
func (uuo *UserUpdateOne) ClearTags() *UserUpdateOne {
	uuo.mutation.ClearTags()
	return uuo
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (uuo *UserUpdateOne) RemoveTagIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveTagIDs(ids...)
	return uuo
}

// RemoveTags removes "tags" edges to Tag entities.
func (uuo *UserUpdateOne) RemoveTags(t ...*Tag) *UserUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.RemoveTagIDs(ids...)
}

// ClearTimetable clears all "timetable" edges to the Timetable entity.
func (uuo *UserUpdateOne) ClearTimetable() *UserUpdateOne {
	uuo.mutation.ClearTimetable()
	return uuo
}

// RemoveTimetableIDs removes the "timetable" edge to Timetable entities by IDs.
func (uuo *UserUpdateOne) RemoveTimetableIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveTimetableIDs(ids...)
	return uuo
}

// RemoveTimetable removes "timetable" edges to Timetable entities.
func (uuo *UserUpdateOne) RemoveTimetable(t ...*Timetable) *UserUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.RemoveTimetableIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if err := uuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() error {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		if user.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized user.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Surname(); ok {
		if err := user.SurnameValidator(v); err != nil {
			return &ValidationError{Name: "surname", err: fmt.Errorf(`ent: validator failed for field "User.surname": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeTime, value)
	}
	if uuo.mutation.DeletedAtCleared() {
		_spec.ClearField(user.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := uuo.mutation.Surname(); ok {
		_spec.SetField(user.FieldSurname, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Firstname(); ok {
		_spec.SetField(user.FieldFirstname, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Title(); ok {
		_spec.SetField(user.FieldTitle, field.TypeString, value)
	}
	if uuo.mutation.TitleCleared() {
		_spec.ClearField(user.FieldTitle, field.TypeString)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Passwordhash(); ok {
		_spec.SetField(user.FieldPasswordhash, field.TypeString, value)
	}
	if uuo.mutation.PasswordhashCleared() {
		_spec.ClearField(user.FieldPasswordhash, field.TypeString)
	}
	if value, ok := uuo.mutation.Comment(); ok {
		_spec.SetField(user.FieldComment, field.TypeString, value)
	}
	if uuo.mutation.CommentCleared() {
		_spec.ClearField(user.FieldComment, field.TypeString)
	}
	if value, ok := uuo.mutation.Active(); ok {
		_spec.SetField(user.FieldActive, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedRole(); ok {
		_spec.AddField(user.FieldRole, field.TypeInt, value)
	}
	if uuo.mutation.BusinessesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.BusinessesTable,
			Columns: []string{user.BusinessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedBusinessesIDs(); len(nodes) > 0 && !uuo.mutation.BusinessesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.BusinessesTable,
			Columns: []string{user.BusinessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.BusinessesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.BusinessesTable,
			Columns: []string{user.BusinessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TagsTable,
			Columns: user.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !uuo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TagsTable,
			Columns: user.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TagsTable,
			Columns: user.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.TimetableCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.TimetableTable,
			Columns: user.TimetablePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timetable.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedTimetableIDs(); len(nodes) > 0 && !uuo.mutation.TimetableCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.TimetableTable,
			Columns: user.TimetablePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timetable.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.TimetableIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.TimetableTable,
			Columns: user.TimetablePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timetable.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
