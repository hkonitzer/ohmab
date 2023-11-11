package hooks

import (
	"context"
	"fmt"
	"github.com/hkonitzer/ohmab/ent"
	"github.com/hkonitzer/ohmab/ent/address"
	"github.com/hkonitzer/ohmab/ent/hook"
	"github.com/hkonitzer/ohmab/ent/timetable"
	"time"
)

const ErrorMessagePrefix = "[VALIDATION ERROR]"

// EnsureOnlyOneTimetableEntry is a hook that ensures that only one timetable entry for an address exists
// for datetime_from in combination with timetable_type
func EnsureOnlyOneTimetableEntry() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.TimetableFunc(func(ctx context.Context, m *ent.TimetableMutation) (ent.Value, error) {
			ty, _ := m.TimetableType()
			f, _ := m.DatetimeFrom()
			to, _ := m.DatetimeTo()
			// query database
			aid, aide := m.AddressID()
			if !aide {
				return nil, fmt.Errorf("%s address_id field is missing", ErrorMessagePrefix)
			}
			tts, err := m.Client().Timetable.Query().Where(
				timetable.HasAddressWith(
					address.IDEQ(aid),
				),
				timetable.TimetableTypeEQ(ty),
				timetable.DatetimeToLTE(to),
				timetable.DeletedAtIsNil(),
			).All(ctx)
			if err != nil && !ent.IsNotFound(err) {
				return nil, err
			}
			// address_timetables is NIL? why?
			// cannot check against given address_id (aid)
			// TODO: get address_timetables from timetable entry and enable this check also for ent.OpUpdate|ent.OpUpdateOne
			// check if datetime_to is in range of another entry
			for _, tt := range tts {
				if to.After(tt.DatetimeTo) && !f.Before(tt.DatetimeTo) {
					continue
				}
				if tt.DatetimeFrom.Compare(f) == 0 {
					return nil, fmt.Errorf("%s there is already a timetable entry with the same datetime_from=%v and type: %v", ErrorMessagePrefix, tt.DatetimeFrom.Format(time.RFC3339), tt.ID)
				}
				if tt.DatetimeFrom.Compare(to) < 1 {
					return nil, fmt.Errorf("%s there is already a timetable entry with existing datetime_from=%v before new datetime_to=%v : %v", ErrorMessagePrefix, tt.DatetimeFrom.Format(time.RFC3339), to.Format(time.RFC3339), tt.ID)
				}
			}
			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpCreate)
}
