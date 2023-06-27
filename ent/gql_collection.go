// OHMAB
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"hynie.de/ohmab/ent/address"
	"hynie.de/ohmab/ent/auditlog"
	"hynie.de/ohmab/ent/business"
	"hynie.de/ohmab/ent/tag"
	"hynie.de/ohmab/ent/timetable"
	"hynie.de/ohmab/ent/user"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (a *AddressQuery) CollectFields(ctx context.Context, satisfies ...string) (*AddressQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return a, nil
	}
	if err := a.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *AddressQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(address.Columns))
		selectedFields = []string{address.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "business":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&BusinessClient{config: a.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			a.withBusiness = query
		case "timetables":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TimetableClient{config: a.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			a.WithNamedTimetables(alias, func(wq *TimetableQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[address.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, address.FieldCreatedAt)
				fieldSeen[address.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[address.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, address.FieldUpdatedAt)
				fieldSeen[address.FieldUpdatedAt] = struct{}{}
			}
		case "deletedAt":
			if _, ok := fieldSeen[address.FieldDeletedAt]; !ok {
				selectedFields = append(selectedFields, address.FieldDeletedAt)
				fieldSeen[address.FieldDeletedAt] = struct{}{}
			}
		case "addition":
			if _, ok := fieldSeen[address.FieldAddition]; !ok {
				selectedFields = append(selectedFields, address.FieldAddition)
				fieldSeen[address.FieldAddition] = struct{}{}
			}
		case "street":
			if _, ok := fieldSeen[address.FieldStreet]; !ok {
				selectedFields = append(selectedFields, address.FieldStreet)
				fieldSeen[address.FieldStreet] = struct{}{}
			}
		case "city":
			if _, ok := fieldSeen[address.FieldCity]; !ok {
				selectedFields = append(selectedFields, address.FieldCity)
				fieldSeen[address.FieldCity] = struct{}{}
			}
		case "zip":
			if _, ok := fieldSeen[address.FieldZip]; !ok {
				selectedFields = append(selectedFields, address.FieldZip)
				fieldSeen[address.FieldZip] = struct{}{}
			}
		case "state":
			if _, ok := fieldSeen[address.FieldState]; !ok {
				selectedFields = append(selectedFields, address.FieldState)
				fieldSeen[address.FieldState] = struct{}{}
			}
		case "country":
			if _, ok := fieldSeen[address.FieldCountry]; !ok {
				selectedFields = append(selectedFields, address.FieldCountry)
				fieldSeen[address.FieldCountry] = struct{}{}
			}
		case "primary":
			if _, ok := fieldSeen[address.FieldPrimary]; !ok {
				selectedFields = append(selectedFields, address.FieldPrimary)
				fieldSeen[address.FieldPrimary] = struct{}{}
			}
		case "telephone":
			if _, ok := fieldSeen[address.FieldTelephone]; !ok {
				selectedFields = append(selectedFields, address.FieldTelephone)
				fieldSeen[address.FieldTelephone] = struct{}{}
			}
		case "comment":
			if _, ok := fieldSeen[address.FieldComment]; !ok {
				selectedFields = append(selectedFields, address.FieldComment)
				fieldSeen[address.FieldComment] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		a.Select(selectedFields...)
	}
	return nil
}

type addressPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AddressPaginateOption
}

func newAddressPaginateArgs(rv map[string]any) *addressPaginateArgs {
	args := &addressPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*AddressWhereInput); ok {
		args.opts = append(args.opts, WithAddressFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (al *AuditLogQuery) CollectFields(ctx context.Context, satisfies ...string) (*AuditLogQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return al, nil
	}
	if err := al.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return al, nil
}

func (al *AuditLogQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(auditlog.Columns))
		selectedFields = []string{auditlog.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "user":
			if _, ok := fieldSeen[auditlog.FieldUser]; !ok {
				selectedFields = append(selectedFields, auditlog.FieldUser)
				fieldSeen[auditlog.FieldUser] = struct{}{}
			}
		case "action":
			if _, ok := fieldSeen[auditlog.FieldAction]; !ok {
				selectedFields = append(selectedFields, auditlog.FieldAction)
				fieldSeen[auditlog.FieldAction] = struct{}{}
			}
		case "entitySchema":
			if _, ok := fieldSeen[auditlog.FieldEntitySchema]; !ok {
				selectedFields = append(selectedFields, auditlog.FieldEntitySchema)
				fieldSeen[auditlog.FieldEntitySchema] = struct{}{}
			}
		case "entityValues":
			if _, ok := fieldSeen[auditlog.FieldEntityValues]; !ok {
				selectedFields = append(selectedFields, auditlog.FieldEntityValues)
				fieldSeen[auditlog.FieldEntityValues] = struct{}{}
			}
		case "entityUUID":
			if _, ok := fieldSeen[auditlog.FieldEntityUUID]; !ok {
				selectedFields = append(selectedFields, auditlog.FieldEntityUUID)
				fieldSeen[auditlog.FieldEntityUUID] = struct{}{}
			}
		case "timestamp":
			if _, ok := fieldSeen[auditlog.FieldTimestamp]; !ok {
				selectedFields = append(selectedFields, auditlog.FieldTimestamp)
				fieldSeen[auditlog.FieldTimestamp] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		al.Select(selectedFields...)
	}
	return nil
}

type auditlogPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AuditLogPaginateOption
}

func newAuditLogPaginateArgs(rv map[string]any) *auditlogPaginateArgs {
	args := &auditlogPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*AuditLogWhereInput); ok {
		args.opts = append(args.opts, WithAuditLogFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (b *BusinessQuery) CollectFields(ctx context.Context, satisfies ...string) (*BusinessQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return b, nil
	}
	if err := b.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return b, nil
}

func (b *BusinessQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(business.Columns))
		selectedFields = []string{business.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "addresses":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&AddressClient{config: b.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			b.WithNamedAddresses(alias, func(wq *AddressQuery) {
				*wq = *query
			})
		case "tags":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TagClient{config: b.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			b.WithNamedTags(alias, func(wq *TagQuery) {
				*wq = *query
			})
		case "users":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: b.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			b.withUsers = query
		case "createdAt":
			if _, ok := fieldSeen[business.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, business.FieldCreatedAt)
				fieldSeen[business.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[business.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, business.FieldUpdatedAt)
				fieldSeen[business.FieldUpdatedAt] = struct{}{}
			}
		case "deletedAt":
			if _, ok := fieldSeen[business.FieldDeletedAt]; !ok {
				selectedFields = append(selectedFields, business.FieldDeletedAt)
				fieldSeen[business.FieldDeletedAt] = struct{}{}
			}
		case "name1":
			if _, ok := fieldSeen[business.FieldName1]; !ok {
				selectedFields = append(selectedFields, business.FieldName1)
				fieldSeen[business.FieldName1] = struct{}{}
			}
		case "name2":
			if _, ok := fieldSeen[business.FieldName2]; !ok {
				selectedFields = append(selectedFields, business.FieldName2)
				fieldSeen[business.FieldName2] = struct{}{}
			}
		case "alias":
			if _, ok := fieldSeen[business.FieldAlias]; !ok {
				selectedFields = append(selectedFields, business.FieldAlias)
				fieldSeen[business.FieldAlias] = struct{}{}
			}
		case "telephone":
			if _, ok := fieldSeen[business.FieldTelephone]; !ok {
				selectedFields = append(selectedFields, business.FieldTelephone)
				fieldSeen[business.FieldTelephone] = struct{}{}
			}
		case "email":
			if _, ok := fieldSeen[business.FieldEmail]; !ok {
				selectedFields = append(selectedFields, business.FieldEmail)
				fieldSeen[business.FieldEmail] = struct{}{}
			}
		case "website":
			if _, ok := fieldSeen[business.FieldWebsite]; !ok {
				selectedFields = append(selectedFields, business.FieldWebsite)
				fieldSeen[business.FieldWebsite] = struct{}{}
			}
		case "comment":
			if _, ok := fieldSeen[business.FieldComment]; !ok {
				selectedFields = append(selectedFields, business.FieldComment)
				fieldSeen[business.FieldComment] = struct{}{}
			}
		case "active":
			if _, ok := fieldSeen[business.FieldActive]; !ok {
				selectedFields = append(selectedFields, business.FieldActive)
				fieldSeen[business.FieldActive] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		b.Select(selectedFields...)
	}
	return nil
}

type businessPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []BusinessPaginateOption
}

func newBusinessPaginateArgs(rv map[string]any) *businessPaginateArgs {
	args := &businessPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case []*BusinessOrder:
			args.opts = append(args.opts, WithBusinessOrder(v))
		case []any:
			var orders []*BusinessOrder
			for i := range v {
				mv, ok := v[i].(map[string]any)
				if !ok {
					continue
				}
				var (
					err1, err2 error
					order      = &BusinessOrder{Field: &BusinessOrderField{}, Direction: entgql.OrderDirectionAsc}
				)
				if d, ok := mv[directionField]; ok {
					err1 = order.Direction.UnmarshalGQL(d)
				}
				if f, ok := mv[fieldField]; ok {
					err2 = order.Field.UnmarshalGQL(f)
				}
				if err1 == nil && err2 == nil {
					orders = append(orders, order)
				}
			}
			args.opts = append(args.opts, WithBusinessOrder(orders))
		}
	}
	if v, ok := rv[whereField].(*BusinessWhereInput); ok {
		args.opts = append(args.opts, WithBusinessFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TagQuery) CollectFields(ctx context.Context, satisfies ...string) (*TagQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return t, nil
	}
	if err := t.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *TagQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(tag.Columns))
		selectedFields = []string{tag.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "business":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&BusinessClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			t.WithNamedBusiness(alias, func(wq *BusinessQuery) {
				*wq = *query
			})
		case "user":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			t.WithNamedUser(alias, func(wq *UserQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[tag.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, tag.FieldCreatedAt)
				fieldSeen[tag.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[tag.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, tag.FieldUpdatedAt)
				fieldSeen[tag.FieldUpdatedAt] = struct{}{}
			}
		case "deletedAt":
			if _, ok := fieldSeen[tag.FieldDeletedAt]; !ok {
				selectedFields = append(selectedFields, tag.FieldDeletedAt)
				fieldSeen[tag.FieldDeletedAt] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[tag.FieldName]; !ok {
				selectedFields = append(selectedFields, tag.FieldName)
				fieldSeen[tag.FieldName] = struct{}{}
			}
		case "comment":
			if _, ok := fieldSeen[tag.FieldComment]; !ok {
				selectedFields = append(selectedFields, tag.FieldComment)
				fieldSeen[tag.FieldComment] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		t.Select(selectedFields...)
	}
	return nil
}

type tagPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TagPaginateOption
}

func newTagPaginateArgs(rv map[string]any) *tagPaginateArgs {
	args := &tagPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &TagOrder{Field: &TagOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithTagOrder(order))
			}
		case *TagOrder:
			if v != nil {
				args.opts = append(args.opts, WithTagOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*TagWhereInput); ok {
		args.opts = append(args.opts, WithTagFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TimetableQuery) CollectFields(ctx context.Context, satisfies ...string) (*TimetableQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return t, nil
	}
	if err := t.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *TimetableQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(timetable.Columns))
		selectedFields = []string{timetable.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "address":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&AddressClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			t.withAddress = query
		case "usersOnDuty":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			t.WithNamedUsersOnDuty(alias, func(wq *UserQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[timetable.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, timetable.FieldCreatedAt)
				fieldSeen[timetable.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[timetable.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, timetable.FieldUpdatedAt)
				fieldSeen[timetable.FieldUpdatedAt] = struct{}{}
			}
		case "deletedAt":
			if _, ok := fieldSeen[timetable.FieldDeletedAt]; !ok {
				selectedFields = append(selectedFields, timetable.FieldDeletedAt)
				fieldSeen[timetable.FieldDeletedAt] = struct{}{}
			}
		case "type":
			if _, ok := fieldSeen[timetable.FieldType]; !ok {
				selectedFields = append(selectedFields, timetable.FieldType)
				fieldSeen[timetable.FieldType] = struct{}{}
			}
		case "datetimeFrom":
			if _, ok := fieldSeen[timetable.FieldDatetimeFrom]; !ok {
				selectedFields = append(selectedFields, timetable.FieldDatetimeFrom)
				fieldSeen[timetable.FieldDatetimeFrom] = struct{}{}
			}
		case "datetimeTo":
			if _, ok := fieldSeen[timetable.FieldDatetimeTo]; !ok {
				selectedFields = append(selectedFields, timetable.FieldDatetimeTo)
				fieldSeen[timetable.FieldDatetimeTo] = struct{}{}
			}
		case "timeWholeDay":
			if _, ok := fieldSeen[timetable.FieldTimeWholeDay]; !ok {
				selectedFields = append(selectedFields, timetable.FieldTimeWholeDay)
				fieldSeen[timetable.FieldTimeWholeDay] = struct{}{}
			}
		case "comment":
			if _, ok := fieldSeen[timetable.FieldComment]; !ok {
				selectedFields = append(selectedFields, timetable.FieldComment)
				fieldSeen[timetable.FieldComment] = struct{}{}
			}
		case "availabilityByPhone":
			if _, ok := fieldSeen[timetable.FieldAvailabilityByPhone]; !ok {
				selectedFields = append(selectedFields, timetable.FieldAvailabilityByPhone)
				fieldSeen[timetable.FieldAvailabilityByPhone] = struct{}{}
			}
		case "availabilityByEmail":
			if _, ok := fieldSeen[timetable.FieldAvailabilityByEmail]; !ok {
				selectedFields = append(selectedFields, timetable.FieldAvailabilityByEmail)
				fieldSeen[timetable.FieldAvailabilityByEmail] = struct{}{}
			}
		case "availabilityBySms":
			if _, ok := fieldSeen[timetable.FieldAvailabilityBySms]; !ok {
				selectedFields = append(selectedFields, timetable.FieldAvailabilityBySms)
				fieldSeen[timetable.FieldAvailabilityBySms] = struct{}{}
			}
		case "availabilityByWhatsapp":
			if _, ok := fieldSeen[timetable.FieldAvailabilityByWhatsapp]; !ok {
				selectedFields = append(selectedFields, timetable.FieldAvailabilityByWhatsapp)
				fieldSeen[timetable.FieldAvailabilityByWhatsapp] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		t.Select(selectedFields...)
	}
	return nil
}

type timetablePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TimetablePaginateOption
}

func newTimetablePaginateArgs(rv map[string]any) *timetablePaginateArgs {
	args := &timetablePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*TimetableWhereInput); ok {
		args.opts = append(args.opts, WithTimetableFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(user.Columns))
		selectedFields = []string{user.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "businesses":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&BusinessClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedBusinesses(alias, func(wq *BusinessQuery) {
				*wq = *query
			})
		case "tags":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TagClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedTags(alias, func(wq *TagQuery) {
				*wq = *query
			})
		case "timetable":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TimetableClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedTimetable(alias, func(wq *TimetableQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[user.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, user.FieldCreatedAt)
				fieldSeen[user.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[user.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, user.FieldUpdatedAt)
				fieldSeen[user.FieldUpdatedAt] = struct{}{}
			}
		case "deletedAt":
			if _, ok := fieldSeen[user.FieldDeletedAt]; !ok {
				selectedFields = append(selectedFields, user.FieldDeletedAt)
				fieldSeen[user.FieldDeletedAt] = struct{}{}
			}
		case "surname":
			if _, ok := fieldSeen[user.FieldSurname]; !ok {
				selectedFields = append(selectedFields, user.FieldSurname)
				fieldSeen[user.FieldSurname] = struct{}{}
			}
		case "firstname":
			if _, ok := fieldSeen[user.FieldFirstname]; !ok {
				selectedFields = append(selectedFields, user.FieldFirstname)
				fieldSeen[user.FieldFirstname] = struct{}{}
			}
		case "title":
			if _, ok := fieldSeen[user.FieldTitle]; !ok {
				selectedFields = append(selectedFields, user.FieldTitle)
				fieldSeen[user.FieldTitle] = struct{}{}
			}
		case "email":
			if _, ok := fieldSeen[user.FieldEmail]; !ok {
				selectedFields = append(selectedFields, user.FieldEmail)
				fieldSeen[user.FieldEmail] = struct{}{}
			}
		case "comment":
			if _, ok := fieldSeen[user.FieldComment]; !ok {
				selectedFields = append(selectedFields, user.FieldComment)
				fieldSeen[user.FieldComment] = struct{}{}
			}
		case "active":
			if _, ok := fieldSeen[user.FieldActive]; !ok {
				selectedFields = append(selectedFields, user.FieldActive)
				fieldSeen[user.FieldActive] = struct{}{}
			}
		case "role":
			if _, ok := fieldSeen[user.FieldRole]; !ok {
				selectedFields = append(selectedFields, user.FieldRole)
				fieldSeen[user.FieldRole] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		u.Select(selectedFields...)
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]any) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &UserOrder{Field: &UserOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithUserOrder(order))
			}
		case *UserOrder:
			if v != nil {
				args.opts = append(args.opts, WithUserOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*UserWhereInput); ok {
		args.opts = append(args.opts, WithUserFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if condition is enabled (Node/Nodes) and it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond string) []string {
	if len(satisfies) == 0 {
		return satisfies
	}
	for _, s := range satisfies {
		if typeCond == s {
			return satisfies
		}
	}
	return append(satisfies, typeCond)
}
