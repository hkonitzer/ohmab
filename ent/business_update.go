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
	"hynie.de/ohmab/ent/business"
	"hynie.de/ohmab/ent/predicate"
	"hynie.de/ohmab/ent/tag"
	"hynie.de/ohmab/ent/user"
)

// BusinessUpdate is the builder for updating Business entities.
type BusinessUpdate struct {
	config
	hooks    []Hook
	mutation *BusinessMutation
}

// Where appends a list predicates to the BusinessUpdate builder.
func (bu *BusinessUpdate) Where(ps ...predicate.Business) *BusinessUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetUpdatedAt sets the "updated_at" field.
func (bu *BusinessUpdate) SetUpdatedAt(t time.Time) *BusinessUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// SetDeletedAt sets the "deleted_at" field.
func (bu *BusinessUpdate) SetDeletedAt(t time.Time) *BusinessUpdate {
	bu.mutation.SetDeletedAt(t)
	return bu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (bu *BusinessUpdate) SetNillableDeletedAt(t *time.Time) *BusinessUpdate {
	if t != nil {
		bu.SetDeletedAt(*t)
	}
	return bu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (bu *BusinessUpdate) ClearDeletedAt() *BusinessUpdate {
	bu.mutation.ClearDeletedAt()
	return bu
}

// SetName1 sets the "name1" field.
func (bu *BusinessUpdate) SetName1(s string) *BusinessUpdate {
	bu.mutation.SetName1(s)
	return bu
}

// SetName2 sets the "name2" field.
func (bu *BusinessUpdate) SetName2(s string) *BusinessUpdate {
	bu.mutation.SetName2(s)
	return bu
}

// SetNillableName2 sets the "name2" field if the given value is not nil.
func (bu *BusinessUpdate) SetNillableName2(s *string) *BusinessUpdate {
	if s != nil {
		bu.SetName2(*s)
	}
	return bu
}

// ClearName2 clears the value of the "name2" field.
func (bu *BusinessUpdate) ClearName2() *BusinessUpdate {
	bu.mutation.ClearName2()
	return bu
}

// SetTelephone sets the "telephone" field.
func (bu *BusinessUpdate) SetTelephone(s string) *BusinessUpdate {
	bu.mutation.SetTelephone(s)
	return bu
}

// SetNillableTelephone sets the "telephone" field if the given value is not nil.
func (bu *BusinessUpdate) SetNillableTelephone(s *string) *BusinessUpdate {
	if s != nil {
		bu.SetTelephone(*s)
	}
	return bu
}

// ClearTelephone clears the value of the "telephone" field.
func (bu *BusinessUpdate) ClearTelephone() *BusinessUpdate {
	bu.mutation.ClearTelephone()
	return bu
}

// SetEmail sets the "email" field.
func (bu *BusinessUpdate) SetEmail(s string) *BusinessUpdate {
	bu.mutation.SetEmail(s)
	return bu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (bu *BusinessUpdate) SetNillableEmail(s *string) *BusinessUpdate {
	if s != nil {
		bu.SetEmail(*s)
	}
	return bu
}

// ClearEmail clears the value of the "email" field.
func (bu *BusinessUpdate) ClearEmail() *BusinessUpdate {
	bu.mutation.ClearEmail()
	return bu
}

// SetWebsite sets the "website" field.
func (bu *BusinessUpdate) SetWebsite(s string) *BusinessUpdate {
	bu.mutation.SetWebsite(s)
	return bu
}

// SetNillableWebsite sets the "website" field if the given value is not nil.
func (bu *BusinessUpdate) SetNillableWebsite(s *string) *BusinessUpdate {
	if s != nil {
		bu.SetWebsite(*s)
	}
	return bu
}

// ClearWebsite clears the value of the "website" field.
func (bu *BusinessUpdate) ClearWebsite() *BusinessUpdate {
	bu.mutation.ClearWebsite()
	return bu
}

// SetComment sets the "comment" field.
func (bu *BusinessUpdate) SetComment(s string) *BusinessUpdate {
	bu.mutation.SetComment(s)
	return bu
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (bu *BusinessUpdate) SetNillableComment(s *string) *BusinessUpdate {
	if s != nil {
		bu.SetComment(*s)
	}
	return bu
}

// ClearComment clears the value of the "comment" field.
func (bu *BusinessUpdate) ClearComment() *BusinessUpdate {
	bu.mutation.ClearComment()
	return bu
}

// SetActive sets the "active" field.
func (bu *BusinessUpdate) SetActive(b bool) *BusinessUpdate {
	bu.mutation.SetActive(b)
	return bu
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (bu *BusinessUpdate) SetNillableActive(b *bool) *BusinessUpdate {
	if b != nil {
		bu.SetActive(*b)
	}
	return bu
}

// AddAddressIDs adds the "addresses" edge to the Address entity by IDs.
func (bu *BusinessUpdate) AddAddressIDs(ids ...uuid.UUID) *BusinessUpdate {
	bu.mutation.AddAddressIDs(ids...)
	return bu
}

// AddAddresses adds the "addresses" edges to the Address entity.
func (bu *BusinessUpdate) AddAddresses(a ...*Address) *BusinessUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return bu.AddAddressIDs(ids...)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (bu *BusinessUpdate) AddTagIDs(ids ...uuid.UUID) *BusinessUpdate {
	bu.mutation.AddTagIDs(ids...)
	return bu
}

// AddTags adds the "tags" edges to the Tag entity.
func (bu *BusinessUpdate) AddTags(t ...*Tag) *BusinessUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.AddTagIDs(ids...)
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (bu *BusinessUpdate) SetUsersID(id uuid.UUID) *BusinessUpdate {
	bu.mutation.SetUsersID(id)
	return bu
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (bu *BusinessUpdate) SetNillableUsersID(id *uuid.UUID) *BusinessUpdate {
	if id != nil {
		bu = bu.SetUsersID(*id)
	}
	return bu
}

// SetUsers sets the "users" edge to the User entity.
func (bu *BusinessUpdate) SetUsers(u *User) *BusinessUpdate {
	return bu.SetUsersID(u.ID)
}

// Mutation returns the BusinessMutation object of the builder.
func (bu *BusinessUpdate) Mutation() *BusinessMutation {
	return bu.mutation
}

// ClearAddresses clears all "addresses" edges to the Address entity.
func (bu *BusinessUpdate) ClearAddresses() *BusinessUpdate {
	bu.mutation.ClearAddresses()
	return bu
}

// RemoveAddressIDs removes the "addresses" edge to Address entities by IDs.
func (bu *BusinessUpdate) RemoveAddressIDs(ids ...uuid.UUID) *BusinessUpdate {
	bu.mutation.RemoveAddressIDs(ids...)
	return bu
}

// RemoveAddresses removes "addresses" edges to Address entities.
func (bu *BusinessUpdate) RemoveAddresses(a ...*Address) *BusinessUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return bu.RemoveAddressIDs(ids...)
}

// ClearTags clears all "tags" edges to the Tag entity.
func (bu *BusinessUpdate) ClearTags() *BusinessUpdate {
	bu.mutation.ClearTags()
	return bu
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (bu *BusinessUpdate) RemoveTagIDs(ids ...uuid.UUID) *BusinessUpdate {
	bu.mutation.RemoveTagIDs(ids...)
	return bu
}

// RemoveTags removes "tags" edges to Tag entities.
func (bu *BusinessUpdate) RemoveTags(t ...*Tag) *BusinessUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.RemoveTagIDs(ids...)
}

// ClearUsers clears the "users" edge to the User entity.
func (bu *BusinessUpdate) ClearUsers() *BusinessUpdate {
	bu.mutation.ClearUsers()
	return bu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BusinessUpdate) Save(ctx context.Context) (int, error) {
	if err := bu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BusinessUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BusinessUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BusinessUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bu *BusinessUpdate) defaults() error {
	if _, ok := bu.mutation.UpdatedAt(); !ok {
		if business.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized business.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := business.UpdateDefaultUpdatedAt()
		bu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (bu *BusinessUpdate) check() error {
	if v, ok := bu.mutation.Name1(); ok {
		if err := business.Name1Validator(v); err != nil {
			return &ValidationError{Name: "name1", err: fmt.Errorf(`ent: validator failed for field "Business.name1": %w`, err)}
		}
	}
	return nil
}

func (bu *BusinessUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(business.Table, business.Columns, sqlgraph.NewFieldSpec(business.FieldID, field.TypeUUID))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.SetField(business.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := bu.mutation.DeletedAt(); ok {
		_spec.SetField(business.FieldDeletedAt, field.TypeTime, value)
	}
	if bu.mutation.DeletedAtCleared() {
		_spec.ClearField(business.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := bu.mutation.Name1(); ok {
		_spec.SetField(business.FieldName1, field.TypeString, value)
	}
	if value, ok := bu.mutation.Name2(); ok {
		_spec.SetField(business.FieldName2, field.TypeString, value)
	}
	if bu.mutation.Name2Cleared() {
		_spec.ClearField(business.FieldName2, field.TypeString)
	}
	if value, ok := bu.mutation.Telephone(); ok {
		_spec.SetField(business.FieldTelephone, field.TypeString, value)
	}
	if bu.mutation.TelephoneCleared() {
		_spec.ClearField(business.FieldTelephone, field.TypeString)
	}
	if value, ok := bu.mutation.Email(); ok {
		_spec.SetField(business.FieldEmail, field.TypeString, value)
	}
	if bu.mutation.EmailCleared() {
		_spec.ClearField(business.FieldEmail, field.TypeString)
	}
	if value, ok := bu.mutation.Website(); ok {
		_spec.SetField(business.FieldWebsite, field.TypeString, value)
	}
	if bu.mutation.WebsiteCleared() {
		_spec.ClearField(business.FieldWebsite, field.TypeString)
	}
	if value, ok := bu.mutation.Comment(); ok {
		_spec.SetField(business.FieldComment, field.TypeString, value)
	}
	if bu.mutation.CommentCleared() {
		_spec.ClearField(business.FieldComment, field.TypeString)
	}
	if value, ok := bu.mutation.Active(); ok {
		_spec.SetField(business.FieldActive, field.TypeBool, value)
	}
	if bu.mutation.AddressesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   business.AddressesTable,
			Columns: business.AddressesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedAddressesIDs(); len(nodes) > 0 && !bu.mutation.AddressesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   business.AddressesTable,
			Columns: business.AddressesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.AddressesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   business.AddressesTable,
			Columns: business.AddressesPrimaryKey,
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
	if bu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   business.TagsTable,
			Columns: business.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedTagsIDs(); len(nodes) > 0 && !bu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   business.TagsTable,
			Columns: business.TagsPrimaryKey,
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
	if nodes := bu.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   business.TagsTable,
			Columns: business.TagsPrimaryKey,
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
	if bu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   business.UsersTable,
			Columns: []string{business.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   business.UsersTable,
			Columns: []string{business.UsersColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{business.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BusinessUpdateOne is the builder for updating a single Business entity.
type BusinessUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BusinessMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (buo *BusinessUpdateOne) SetUpdatedAt(t time.Time) *BusinessUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// SetDeletedAt sets the "deleted_at" field.
func (buo *BusinessUpdateOne) SetDeletedAt(t time.Time) *BusinessUpdateOne {
	buo.mutation.SetDeletedAt(t)
	return buo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (buo *BusinessUpdateOne) SetNillableDeletedAt(t *time.Time) *BusinessUpdateOne {
	if t != nil {
		buo.SetDeletedAt(*t)
	}
	return buo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (buo *BusinessUpdateOne) ClearDeletedAt() *BusinessUpdateOne {
	buo.mutation.ClearDeletedAt()
	return buo
}

// SetName1 sets the "name1" field.
func (buo *BusinessUpdateOne) SetName1(s string) *BusinessUpdateOne {
	buo.mutation.SetName1(s)
	return buo
}

// SetName2 sets the "name2" field.
func (buo *BusinessUpdateOne) SetName2(s string) *BusinessUpdateOne {
	buo.mutation.SetName2(s)
	return buo
}

// SetNillableName2 sets the "name2" field if the given value is not nil.
func (buo *BusinessUpdateOne) SetNillableName2(s *string) *BusinessUpdateOne {
	if s != nil {
		buo.SetName2(*s)
	}
	return buo
}

// ClearName2 clears the value of the "name2" field.
func (buo *BusinessUpdateOne) ClearName2() *BusinessUpdateOne {
	buo.mutation.ClearName2()
	return buo
}

// SetTelephone sets the "telephone" field.
func (buo *BusinessUpdateOne) SetTelephone(s string) *BusinessUpdateOne {
	buo.mutation.SetTelephone(s)
	return buo
}

// SetNillableTelephone sets the "telephone" field if the given value is not nil.
func (buo *BusinessUpdateOne) SetNillableTelephone(s *string) *BusinessUpdateOne {
	if s != nil {
		buo.SetTelephone(*s)
	}
	return buo
}

// ClearTelephone clears the value of the "telephone" field.
func (buo *BusinessUpdateOne) ClearTelephone() *BusinessUpdateOne {
	buo.mutation.ClearTelephone()
	return buo
}

// SetEmail sets the "email" field.
func (buo *BusinessUpdateOne) SetEmail(s string) *BusinessUpdateOne {
	buo.mutation.SetEmail(s)
	return buo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (buo *BusinessUpdateOne) SetNillableEmail(s *string) *BusinessUpdateOne {
	if s != nil {
		buo.SetEmail(*s)
	}
	return buo
}

// ClearEmail clears the value of the "email" field.
func (buo *BusinessUpdateOne) ClearEmail() *BusinessUpdateOne {
	buo.mutation.ClearEmail()
	return buo
}

// SetWebsite sets the "website" field.
func (buo *BusinessUpdateOne) SetWebsite(s string) *BusinessUpdateOne {
	buo.mutation.SetWebsite(s)
	return buo
}

// SetNillableWebsite sets the "website" field if the given value is not nil.
func (buo *BusinessUpdateOne) SetNillableWebsite(s *string) *BusinessUpdateOne {
	if s != nil {
		buo.SetWebsite(*s)
	}
	return buo
}

// ClearWebsite clears the value of the "website" field.
func (buo *BusinessUpdateOne) ClearWebsite() *BusinessUpdateOne {
	buo.mutation.ClearWebsite()
	return buo
}

// SetComment sets the "comment" field.
func (buo *BusinessUpdateOne) SetComment(s string) *BusinessUpdateOne {
	buo.mutation.SetComment(s)
	return buo
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (buo *BusinessUpdateOne) SetNillableComment(s *string) *BusinessUpdateOne {
	if s != nil {
		buo.SetComment(*s)
	}
	return buo
}

// ClearComment clears the value of the "comment" field.
func (buo *BusinessUpdateOne) ClearComment() *BusinessUpdateOne {
	buo.mutation.ClearComment()
	return buo
}

// SetActive sets the "active" field.
func (buo *BusinessUpdateOne) SetActive(b bool) *BusinessUpdateOne {
	buo.mutation.SetActive(b)
	return buo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (buo *BusinessUpdateOne) SetNillableActive(b *bool) *BusinessUpdateOne {
	if b != nil {
		buo.SetActive(*b)
	}
	return buo
}

// AddAddressIDs adds the "addresses" edge to the Address entity by IDs.
func (buo *BusinessUpdateOne) AddAddressIDs(ids ...uuid.UUID) *BusinessUpdateOne {
	buo.mutation.AddAddressIDs(ids...)
	return buo
}

// AddAddresses adds the "addresses" edges to the Address entity.
func (buo *BusinessUpdateOne) AddAddresses(a ...*Address) *BusinessUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return buo.AddAddressIDs(ids...)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (buo *BusinessUpdateOne) AddTagIDs(ids ...uuid.UUID) *BusinessUpdateOne {
	buo.mutation.AddTagIDs(ids...)
	return buo
}

// AddTags adds the "tags" edges to the Tag entity.
func (buo *BusinessUpdateOne) AddTags(t ...*Tag) *BusinessUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.AddTagIDs(ids...)
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (buo *BusinessUpdateOne) SetUsersID(id uuid.UUID) *BusinessUpdateOne {
	buo.mutation.SetUsersID(id)
	return buo
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (buo *BusinessUpdateOne) SetNillableUsersID(id *uuid.UUID) *BusinessUpdateOne {
	if id != nil {
		buo = buo.SetUsersID(*id)
	}
	return buo
}

// SetUsers sets the "users" edge to the User entity.
func (buo *BusinessUpdateOne) SetUsers(u *User) *BusinessUpdateOne {
	return buo.SetUsersID(u.ID)
}

// Mutation returns the BusinessMutation object of the builder.
func (buo *BusinessUpdateOne) Mutation() *BusinessMutation {
	return buo.mutation
}

// ClearAddresses clears all "addresses" edges to the Address entity.
func (buo *BusinessUpdateOne) ClearAddresses() *BusinessUpdateOne {
	buo.mutation.ClearAddresses()
	return buo
}

// RemoveAddressIDs removes the "addresses" edge to Address entities by IDs.
func (buo *BusinessUpdateOne) RemoveAddressIDs(ids ...uuid.UUID) *BusinessUpdateOne {
	buo.mutation.RemoveAddressIDs(ids...)
	return buo
}

// RemoveAddresses removes "addresses" edges to Address entities.
func (buo *BusinessUpdateOne) RemoveAddresses(a ...*Address) *BusinessUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return buo.RemoveAddressIDs(ids...)
}

// ClearTags clears all "tags" edges to the Tag entity.
func (buo *BusinessUpdateOne) ClearTags() *BusinessUpdateOne {
	buo.mutation.ClearTags()
	return buo
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (buo *BusinessUpdateOne) RemoveTagIDs(ids ...uuid.UUID) *BusinessUpdateOne {
	buo.mutation.RemoveTagIDs(ids...)
	return buo
}

// RemoveTags removes "tags" edges to Tag entities.
func (buo *BusinessUpdateOne) RemoveTags(t ...*Tag) *BusinessUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.RemoveTagIDs(ids...)
}

// ClearUsers clears the "users" edge to the User entity.
func (buo *BusinessUpdateOne) ClearUsers() *BusinessUpdateOne {
	buo.mutation.ClearUsers()
	return buo
}

// Where appends a list predicates to the BusinessUpdate builder.
func (buo *BusinessUpdateOne) Where(ps ...predicate.Business) *BusinessUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BusinessUpdateOne) Select(field string, fields ...string) *BusinessUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Business entity.
func (buo *BusinessUpdateOne) Save(ctx context.Context) (*Business, error) {
	if err := buo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BusinessUpdateOne) SaveX(ctx context.Context) *Business {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BusinessUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BusinessUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buo *BusinessUpdateOne) defaults() error {
	if _, ok := buo.mutation.UpdatedAt(); !ok {
		if business.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized business.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := business.UpdateDefaultUpdatedAt()
		buo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (buo *BusinessUpdateOne) check() error {
	if v, ok := buo.mutation.Name1(); ok {
		if err := business.Name1Validator(v); err != nil {
			return &ValidationError{Name: "name1", err: fmt.Errorf(`ent: validator failed for field "Business.name1": %w`, err)}
		}
	}
	return nil
}

func (buo *BusinessUpdateOne) sqlSave(ctx context.Context) (_node *Business, err error) {
	if err := buo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(business.Table, business.Columns, sqlgraph.NewFieldSpec(business.FieldID, field.TypeUUID))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Business.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, business.FieldID)
		for _, f := range fields {
			if !business.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != business.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.SetField(business.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := buo.mutation.DeletedAt(); ok {
		_spec.SetField(business.FieldDeletedAt, field.TypeTime, value)
	}
	if buo.mutation.DeletedAtCleared() {
		_spec.ClearField(business.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := buo.mutation.Name1(); ok {
		_spec.SetField(business.FieldName1, field.TypeString, value)
	}
	if value, ok := buo.mutation.Name2(); ok {
		_spec.SetField(business.FieldName2, field.TypeString, value)
	}
	if buo.mutation.Name2Cleared() {
		_spec.ClearField(business.FieldName2, field.TypeString)
	}
	if value, ok := buo.mutation.Telephone(); ok {
		_spec.SetField(business.FieldTelephone, field.TypeString, value)
	}
	if buo.mutation.TelephoneCleared() {
		_spec.ClearField(business.FieldTelephone, field.TypeString)
	}
	if value, ok := buo.mutation.Email(); ok {
		_spec.SetField(business.FieldEmail, field.TypeString, value)
	}
	if buo.mutation.EmailCleared() {
		_spec.ClearField(business.FieldEmail, field.TypeString)
	}
	if value, ok := buo.mutation.Website(); ok {
		_spec.SetField(business.FieldWebsite, field.TypeString, value)
	}
	if buo.mutation.WebsiteCleared() {
		_spec.ClearField(business.FieldWebsite, field.TypeString)
	}
	if value, ok := buo.mutation.Comment(); ok {
		_spec.SetField(business.FieldComment, field.TypeString, value)
	}
	if buo.mutation.CommentCleared() {
		_spec.ClearField(business.FieldComment, field.TypeString)
	}
	if value, ok := buo.mutation.Active(); ok {
		_spec.SetField(business.FieldActive, field.TypeBool, value)
	}
	if buo.mutation.AddressesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   business.AddressesTable,
			Columns: business.AddressesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedAddressesIDs(); len(nodes) > 0 && !buo.mutation.AddressesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   business.AddressesTable,
			Columns: business.AddressesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.AddressesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   business.AddressesTable,
			Columns: business.AddressesPrimaryKey,
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
	if buo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   business.TagsTable,
			Columns: business.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !buo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   business.TagsTable,
			Columns: business.TagsPrimaryKey,
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
	if nodes := buo.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   business.TagsTable,
			Columns: business.TagsPrimaryKey,
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
	if buo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   business.UsersTable,
			Columns: []string{business.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   business.UsersTable,
			Columns: []string{business.UsersColumn},
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
	_node = &Business{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{business.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
