package hooks

import (
	"context"
	"errors"
	"github.com/hkonitzer/ohmab/ent"
	"github.com/hkonitzer/ohmab/ent/address"
	"github.com/hkonitzer/ohmab/ent/business"
	"github.com/hkonitzer/ohmab/ent/hook"
)

// EnsureOnlyOnePrimaryAddress is a hook that ensures that only one address on a business is marked as primary address
func EnsureOnlyOnePrimaryAddress() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.AddressFunc(func(ctx context.Context, m *ent.AddressMutation) (ent.Value, error) {
			primaryAddress, _ := m.Primary()
			addressId, _ := m.ID()
			if primaryAddress {
				businessId, _ := m.BusinessID()
				pc := m.Client().Business.Query().Where(business.ID(businessId)).QueryAddresses().Where(address.Primary(true), address.IDNEQ(addressId)).CountX(ctx)
				if pc > 0 {
					return nil, errors.New("there is already a primary address for this business, please change the existing")
				}
			}
			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne)
}
