package hooks

import (
	"context"
	"errors"
	"github.com/hkonitzer/ohmab/ent"
	"github.com/hkonitzer/ohmab/ent/hook"
	"github.com/hkonitzer/ohmab/internal/pkg/privacy"
)

func VerifyUserRole() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
			r, exists := m.Role()
			if !exists {
				return nil, errors.New("user role is missing")
			}
			if r == privacy.AdminRoleAsString() ||
				r == privacy.OwnerRoleAsString() ||
				r == privacy.EmployeeRoleAsString() ||
				r == privacy.ViewerRoleAsString() {
				return next.Mutate(ctx, m)
			} else {
				return nil, errors.New("user role '" + r + "' is invalid")
			}
		})
	}
	return hook.On(hk, ent.OpCreate)
}
