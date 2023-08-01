package hooks

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/hkonitzer/ohmab/ent"
	"github.com/hkonitzer/ohmab/ent/business"
	"github.com/hkonitzer/ohmab/ent/hook"
	"github.com/hkonitzer/ohmab/ent/publicuser"
	"github.com/hkonitzer/ohmab/ent/user"
)

func UpdatePublicUser() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
			pu, ex := m.UsePublicapi()
			// get values from the mutation
			e, ee := m.Email()
			f, fe := m.Firstname()
			s, se := m.Surname()
			t, te := m.Title()
			// create a public user
			if m.Op() == ent.OpCreate {
				if !ex || pu != "1" { // no public usage wanted, skip
					return next.Mutate(ctx, m)
				}
				// We need active businesses on this user
				if len(m.BusinessesIDs()) == 0 {
					return nil, fmt.Errorf("no active businesses on user, but public usage wanted")
				}
				puc := m.Client().PublicUser.Create()
				if ee {
					puc.SetEmail(e)
				}
				if fe {
					puc.SetFirstname(f)
				}
				if se {
					puc.SetSurname(s)
				}
				if te {
					puc.SetTitle(t)
				}
				puc.AddBusinessIDs(m.BusinessesIDs()...)
				_, err := puc.Save(ctx)
				if err != nil {
					return nil, err
				}
				return next.Mutate(ctx, m)
			}
			// update the public user also
			opu, _ := m.OldUsePublicapi(ctx)
			npu, enpu := m.UsePublicapi()
			// old and new values for public_usage are false and we are not deleting the original user --> nothing to do
			deleteOp := m.Op() == ent.OpDeleteOne || m.Op() == ent.OpDelete
			if opu != "1" && npu != "1" && !deleteOp { // no public usage wanted, skip
				return next.Mutate(ctx, m)
			}
			// query the old entry first
			// get the user (for safety)
			uid, uidex := m.ID()
			if !uidex {
				return nil, fmt.Errorf("no user id given")
			}
			u, err := m.Client().User.Query().Where(user.ID(uid)).WithBusinesses().First(ctx)
			if err != nil {
				return nil, err
			}
			ue := u.Email
			uf := u.Firstname
			us := u.Surname
			ut := u.Title
			ub := u.Edges.Businesses
			var ubids []uuid.UUID
			for _, b := range ub {
				ubids = append(ubids, b.ID)
			}
			puq := m.Client().PublicUser.Query().
				Where(publicuser.Email(ue)).
				Where(publicuser.Firstname(uf)).
				Where(publicuser.Surname(us)).
				Where(publicuser.Title(ut))
			puq.WithBusinesses(func(query *ent.BusinessQuery) {
				//query.Where(business.IDIn(m.BusinessesIDs()...))
				query.Where(business.IDIn(ubids...))
			})
			puf, err := puq.First(ctx)
			if ent.IsNotFound(err) {
				// create a new public user (public_api changed to 1 on user)
				puc := m.Client().PublicUser.Create().
					SetEmail(ue).
					SetFirstname(uf).
					SetSurname(us)
				if ut != "" {
					puc.SetTitle(ut)
				}
				puc.AddBusinessIDs(ubids...)
				err = puc.Exec(ctx)
				if err != nil {
					return nil, err
				}
				return next.Mutate(ctx, m)
			} else if err != nil {
				return nil, err
			}
			// update or delete?
			if deleteOp || (enpu && npu != "1") {
				err = m.Client().PublicUser.DeleteOneID(puf.ID).Exec(ctx)
				if err != nil {
					return nil, err
				}
				return next.Mutate(ctx, m)
			} else {
				pufu := puf.Update()
				if ee {
					pufu.SetEmail(e)
				}
				if fe {
					pufu.SetFirstname(f)
				}
				if se {
					pufu.SetSurname(s)
				}
				if te {
					pufu.SetTitle(t)
				}
				err = pufu.Exec(ctx)
				if err != nil {
					return nil, err
				}
			}
			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne|ent.OpDeleteOne|ent.OpDelete)
}
