package hooks

import (
	"context"
	"errors"
	"github.com/hkonitzer/ohmab/ent"
	"github.com/hkonitzer/ohmab/ent/hook"
	"github.com/hkonitzer/ohmab/internal/pkg/common/log"
	"time"
)

func EnsureDurationIsSet() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		var logger = log.GetLoggerInstance()
		return hook.TimetableFunc(func(ctx context.Context, m *ent.TimetableMutation) (ent.Value, error) {
			fro, exists := m.DatetimeFrom()
			if !exists {
				return nil, errors.New("datetime_from field is missing")
			}
			// overwrite datetime_to if duration is set
			to, texists := m.DatetimeTo()
			dur, dexists := m.Duration()
			if texists && dexists {
				id, _ := m.ID()
				logger.Info().Msgf("both datetime_to and duration are set, ignoring datetime_to on id %v", id)
			}
			if dexists {
				tdur := time.Duration(dur) * time.Hour
				m.SetDatetimeTo(fro.Add(tdur))
			} else if texists { // duration is not set, but datetime_to is
				m.SetDuration(int(to.Sub(fro).Hours()))
			} else {
				return nil, errors.New("either datetime_to or duration field must be set")
			}
			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne)
}
