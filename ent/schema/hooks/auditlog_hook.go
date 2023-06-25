package hooks

import (
	"context"
	"errors"
	"fmt"
	"hynie.de/ohmab/ent"
	"hynie.de/ohmab/ent/hook"
	"hynie.de/ohmab/ent/schema/constants"
)

// AuditLogForUser is a hook that logs all operations on Users.
func AuditLogForUser() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
			fieldsAndValues := extractFieldsandherValues(m)
			userId, _ := m.ID()
			m.Client().AuditLog.Create().
				SetUser("admin").
				SetAction(m.Op().String()).
				SetEntitySchema(m.Type()).
				SetEntityUUID(userId.String()).
				SetEntityValues(fieldsAndValues).
				SaveX(ctx)
			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpCreate|ent.OpUpdate|ent.OpDelete|ent.OpDeleteOne|ent.OpUpdateOne)
}

func AuditLogForTimetable() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.TimetableFunc(func(ctx context.Context, m *ent.TimetableMutation) (ent.Value, error) {

			timetableId, _ := m.ID()
			m.Client().AuditLog.Create().
				SetUser("admin").
				SetAction(m.Op().String()).
				SetEntitySchema(m.Type()).
				SetEntityUUID(timetableId.String()).
				SetEntityValues(extractFieldsandherValues(m)).
				SaveX(ctx)
			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpCreate|ent.OpUpdate|ent.OpDelete|ent.OpDeleteOne|ent.OpUpdateOne)
}

// AuditLogForBusiness is a hook that logs all operations on Businesses.
func AuditLogForBusiness() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.BusinessFunc(func(ctx context.Context, m *ent.BusinessMutation) (ent.Value, error) {
			if ent.OpCreate == m.Op() {
				_, exists := m.Name1()
				if !exists {
					return nil, errors.New("name1 field is missing")
				}
			}
			//start := time.Now()
			defer func() {
				//fmt.Printf("Op=%s\tType=%s\tTime=%s\tConcreteType=%TFields=%v\n", m.Op(), m.Type(), time.Since(start), m, m.Fields())
				businessId, _ := m.ID()
				m.Client().AuditLog.Create().
					SetUser("admin").
					SetAction(m.Op().String()).
					SetEntitySchema(m.Type()).
					SetEntityUUID(businessId.String()).
					SetEntityValues(extractFieldsandherValues(m)).
					SaveX(ctx)
			}()
			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpCreate|ent.OpUpdate|ent.OpDelete|ent.OpDeleteOne|ent.OpUpdateOne)
}

func extractFieldsandherValues(m ent.Mutation) map[string]string {
	fieldsAndValues := make(map[string]string)
	for _, field := range m.Fields() {
		if field == constants.TimeMixinUpdatedAtFieldName || field == constants.TimeMixinCreatedAtFieldName { // skip this field in audit logs
			continue
		}
		value, _ := m.Field(field)
		fieldsAndValues[field] = fmt.Sprintf("%v", value)
	}
	return fieldsAndValues
}
