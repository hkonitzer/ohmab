package hooks

import (
	"context"
	"errors"
	"hynie.de/ohmab/ent"
	"hynie.de/ohmab/ent/hook"
	"regexp"
	"strings"
)

func UpperCaseForBussinessFields() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.BusinessFunc(func(ctx context.Context, m *ent.BusinessMutation) (ent.Value, error) {
			if ent.OpCreate == m.Op() {
				alias, exists := m.Alias()
				if !exists {
					return nil, errors.New("alias field is missing")
				} else if len(alias) > 20 {
					return nil, errors.New("alias field is too long, max 20 chars")
				}
				// alias: Convert to uppercase,replace all spaces with nothing
				alias = strings.ToUpper(alias)
				alias = strings.ReplaceAll(alias, " ", "")
				// remove all special characters but - in string alias
				alias = regexp.MustCompile(`[^a-zA-Z0-9\-]+`).ReplaceAllString(alias, "")
				m.SetAlias(strings.ToUpper(alias))
			}
			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpCreate|ent.OpUpdate)
}
