package hooks

import (
	"context"
	"github.com/hkonitzer/ohmab/ent"
	"github.com/hkonitzer/ohmab/ent/hook"
	"github.com/hkonitzer/ohmab/internal/pkg/common/log"
)

func RemoveZeroTimeValues() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		var logger = log.GetLoggerInstance()
		return hook.TimetableFunc(func(ctx context.Context, m *ent.TimetableMutation) (ent.Value, error) {
			to, exists := m.DatetimeTo()
			if exists && to.IsZero() {
				m.ClearDatetimeTo()
				id, _ := m.ID()
				logger.Debug().Msgf("datetime_to is zero, clearing value on id %v", id)
			}
			del, exists := m.DeletedAt()
			if exists && del.IsZero() {
				m.ClearDeletedAt()
				id, _ := m.ID()
				logger.Debug().Msgf("deleted_at is zero, clearing value on id %v", id)
			}
			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne)
}
