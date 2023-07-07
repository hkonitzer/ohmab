package hooks

import (
	"context"
	"errors"
	"fmt"
	"hynie.de/ohmab/ent"
	"hynie.de/ohmab/ent/address"
	"hynie.de/ohmab/ent/hook"
	"hynie.de/ohmab/ent/schema/constants"
	"hynie.de/ohmab/internal/pkg/privacy"
)

type AuditLog struct {
	viewer privacy.Viewer
}

// AuditLogForUser is a hook that logs all operations on Users.
func (al *AuditLog) AuditLogForUser() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
			//@TODO: use Viewer from context
			fieldsAndValues := extractFieldsandherValues(m)
			userID, _ := m.ID()
			m.Client().AuditLog.Create().
				SetUser("admin").
				SetAction(m.Op().String()).
				SetEntitySchema(m.Type()).
				SetEntityUUID(userID.String()).
				SetEntityValues(fieldsAndValues).
				SaveX(ctx)
			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpCreate|ent.OpUpdate|ent.OpDelete|ent.OpDeleteOne|ent.OpUpdateOne)
}

func (al *AuditLog) AuditLogForTimetable() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.TimetableFunc(func(ctx context.Context, m *ent.TimetableMutation) (ent.Value, error) {
			// get authorize from context
			err := al.getAuth(ctx)
			if err != nil {
				return nil, err
			}
			// check if the viewer is the owner of the business or admin
			if !al.viewer.Admin() {
				auuid, _ := m.AddressID()
				business, err := m.Client().Business.Query().WithAddresses(
					func(q *ent.AddressQuery) {
						q.Where(address.IDEQ(auuid))
					}).First(ctx)
				if err != nil {
					logger.Err(err).Msgf("could not resolve business to dermine scope for address-id '%v'", auuid)
					return nil, errors.New("could not resolve scopes")
				}
				if !al.viewer.HasScope(business.ID.String()) {
					return nil, errors.New("not authorized (scope)")
				}
			}
			timetableID, _ := m.ID()
			err = al.createLogEntry(m.Client(), ctx, m, timetableID.String())
			if err != nil {
				return nil, errors.New("could not create audit log")
			}
			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpCreate|ent.OpUpdate|ent.OpDelete|ent.OpDeleteOne|ent.OpUpdateOne)
}

// AuditLogForBusiness is a hook that logs all operations on Businesses.
func (al *AuditLog) AuditLogForBusiness() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.BusinessFunc(func(ctx context.Context, m *ent.BusinessMutation) (ent.Value, error) {
			// validate name1 field is set
			_, exists := m.Name1()
			if !exists {
				return nil, errors.New("name1 field is missing")
			}
			// get authorize from context, request will be denied if the viewer has not the admin role
			err := al.getAuth(ctx)
			if err != nil || !al.viewer.Admin() {
				return nil, err
			}
			businessID, _ := m.ID()
			err = al.createLogEntry(m.Client(), ctx, m, businessID.String())
			if err != nil {
				return nil, errors.New("could not create audit log")
			}
			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpCreate|ent.OpUpdate|ent.OpDelete|ent.OpDeleteOne|ent.OpUpdateOne)
}

// AuditLogForAddress is a hook that logs all operations on Adresses.
func (al *AuditLog) AuditLogForAddress() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.AddressFunc(func(ctx context.Context, m *ent.AddressMutation) (ent.Value, error) {
			// get authorize from context
			err := al.getAuth(ctx)
			if err != nil {
				return nil, err
			}
			// check if the viewer is the owner of the business or admin
			if !al.viewer.Admin() {
				buuid, _ := m.BusinessID()
				if !al.viewer.HasScope(buuid.String()) {
					return nil, errors.New("not authorized (scope)")
				}
			}
			addressId, _ := m.ID()
			err = al.createLogEntry(m.Client(), ctx, m, addressId.String())
			if err != nil {
				return nil, errors.New("could not create audit log")
			}
			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpCreate|ent.OpUpdate|ent.OpDelete|ent.OpDeleteOne|ent.OpUpdateOne)
}

func (al *AuditLog) AuditLogForContent() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ContentFunc(func(ctx context.Context, m *ent.ContentMutation) (ent.Value, error) {
			// get authorize from context, request will be denied if the viewer has not the admin role
			err := al.getAuth(ctx)
			if err != nil || !al.viewer.Admin() {
				return nil, err
			}
			contentID, _ := m.ID()
			err = al.createLogEntry(m.Client(), ctx, m, contentID.String())
			if err != nil {
				return nil, errors.New("could not create audit log")
			}
			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpCreate|ent.OpUpdate|ent.OpDelete|ent.OpDeleteOne|ent.OpUpdateOne)
}

func extractFieldsandherValues(m ent.Mutation) map[string]string {
	fieldsAndValues := make(map[string]string)
	for _, field := range m.Fields() {
		if field == constants.TimeMixinUpdatedAtFieldName ||
			field == constants.TimeMixinCreatedAtFieldName { // skip this field in audit logs
			continue
		}
		value, _ := m.Field(field)
		fieldsAndValues[field] = fmt.Sprintf("%v", value)
	}
	return fieldsAndValues
}

func (al *AuditLog) getAuth(ctx context.Context) error {
	al.viewer = privacy.FromContext(ctx)
	if al.viewer == nil {
		return errors.New("not authorized")
	}
	return nil
}

func (al *AuditLog) createLogEntry(client *ent.Client, ctx context.Context, m ent.Mutation, eId string) error {
	_, err := client.AuditLog.Create().
		SetUser(al.viewer.GetUserID()).
		SetAction(m.Op().String()).
		SetEntitySchema(m.Type()).
		SetEntityUUID(eId).
		SetEntityValues(extractFieldsandherValues(m)).
		Save(ctx)
	if err != nil {
		logger.Err(err).Msgf("could not create audit log for entity-id %v", eId)
	}
	return err
}
