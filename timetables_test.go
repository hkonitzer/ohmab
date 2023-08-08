package OHMAB

import (
	"context"
	"fmt"
	"github.com/hkonitzer/ohmab/ent"
	"github.com/hkonitzer/ohmab/ent/enttest"
	"github.com/hkonitzer/ohmab/ent/schema/hooks"
	"github.com/hkonitzer/ohmab/internal/pkg/privacy"
	_ "github.com/mattn/go-sqlite3"
	"strings"
	"testing"
	"time"
)

func TestCreateTimeTable(t *testing.T) {

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		t.Fatal("failed creating schema resources: ", err)
	}

	// create admin viewer
	adminViewer := privacy.UserViewer{Role: privacy.Admin}
	adminViewer.SetUserID("TESTADMIN")
	// create admin context
	adminCtx := adminViewer.ToContext(ctx)
	// create owner viewer
	ownerViewer := privacy.UserViewer{Role: privacy.Owner}
	ownerViewer.SetUserID("TESTOWNER")
	// create owner context
	ownerCtx := ownerViewer.ToContext(ctx)

	// create users first
	// admin user
	adminUser, err := client.User.Create().
		SetComment("TESTDATA").
		SetRole(privacy.AdminRoleAsString()).
		SetLogin("TESTADMIN").
		SetSurname("Musterfrau").
		SetFirstname("Erika").
		SetEmail("admin@localhost").
		Save(adminCtx)
	if err != nil {
		t.Errorf("failed creating admin user: %v", err)
	}
	// owner user
	ownerUser, err := client.User.Create().
		SetComment("TESTDATA").
		SetRole(privacy.OwnerRoleAsString()).
		SetLogin("TESTOWNER").
		SetSurname("Doe").
		SetFirstname("Joey").
		SetEmail("owner@localhost").
		Save(adminCtx)
	if err != nil {
		t.Errorf("failed creating owner user: %v", err)
	}

	// update viewers
	adminViewer.SetUserID(adminUser.ID.String())
	ownerViewer.SetUserID(ownerUser.ID.String())

	// create a business for the address first
	business, err := client.Business.Create().
		SetComment("TESTDATA").
		SetName1("TEST").
		SetAlias("TEST").
		SetTelephone("123456789").
		AddUserIDs(adminUser.ID, ownerUser.ID).
		Save(adminCtx)
	if err != nil {
		t.Errorf("failed creating business: %v", err)
	}
	// set the scopes for the owner viewer, same as AddUserIDs above
	ownerViewer.Scopes = append(ownerViewer.Scopes, business.ID.String())
	ownerCtx = ownerViewer.ToContext(ownerCtx)
	// create a new address as Admin
	a1, err := client.Address.Create().
		SetComment("(TT) ADMIN created a address").
		SetBusiness(business).
		Save(adminCtx)
	if err != nil {
		t.Errorf("failed creating address: %v", err)
	}
	// create a new timetable for this address as Admin
	_, err = client.Timetable.Create().
		SetDatetimeFrom(time.Now()).
		SetDatetimeTo(time.Now().Add(24 * time.Hour)).
		SetComment("(TT) ADMIN created a timetable").
		SetAddress(a1).Save(adminCtx)
	if err != nil {
		t.Errorf("failed creating timetable as admin: %v", err)
	}
	// create a new address as Owner
	a2, err := client.Address.Create().
		SetComment("(TT) OWNER created a address").
		SetBusiness(business).
		Save(ownerCtx)
	if err != nil {
		t.Errorf("failed creating address: %v", err)
	}
	// create a new timetable for this address as Owner
	from := time.Date(2023, 01, 01, 00, 00, 00, 00, time.UTC)
	to := from.Add(4 * time.Hour)
	o, err := client.Timetable.Create().
		SetDatetimeFrom(from).
		SetDatetimeTo(to).
		SetComment("(TT) OWNER created a timetable").
		SetAddress(a2).Save(ownerCtx)
	if err != nil {
		t.Errorf("failed creating timetable as owner: %v", err)
	}

	// test validate datetime hook

	// creates a second entry with the same address and timeframe
	st2, err := client.Timetable.Create().
		SetDatetimeFrom(from).
		SetDatetimeTo(to).
		SetComment("(TT) OWNER created a second timetable").
		SetAddress(a2).Save(ownerCtx)
	if st2 != nil {
		t.Fatal("should not be able to create a second timetable for the same address and timeframe")
	} else if err != nil && !strings.HasPrefix(err.Error(), hooks.ErrorMessagePrefix) {
		t.Log("error creating a second timetable entry: ", err)
	}

	// creates a third entry with the same address and timeFrom + 1 hour
	st3, err := client.Timetable.Create().
		SetDatetimeFrom(from.Add(1 * time.Hour)).
		SetDatetimeTo(to).
		SetComment("(TT) OWNER created a third timetable").
		SetAddress(a2).Save(ownerCtx)
	if st3 != nil {
		t.Fatal("should not be able to create a third timetable for the same address and start time= +1 hour")
	} else if err != nil && !strings.HasPrefix(err.Error(), hooks.ErrorMessagePrefix) {
		t.Log("error creating a third timetable entry: ", err)
	}

	// creates a fourth entry with the same address and timeFrom +11 hour timeTo + 1 hour
	st4, err := client.Timetable.Create().
		SetDatetimeFrom(to.Add(-1 * time.Hour)).
		SetDatetimeTo(to.Add(1 * time.Hour)).
		SetComment("(TT) OWNER created a fourth timetable").
		SetAddress(a2).Save(ownerCtx)
	if st4 != nil {
		t.Logf("first entry      : %v\n", printTTEntry(o))
		t.Logf("created entry (4): %v\n", printTTEntry(st4))
		t.Fatal("should not be able to create a fourth timetable for the same address and startime= -1 hour startfrom and endtime= +1 hour")
	} else if err != nil && !strings.HasPrefix(err.Error(), hooks.ErrorMessagePrefix) {
		t.Log("error fourth a fourth timetable entry: ", err)
	}
	// creates a fifth entry one day later but with from bevor to
	st5, err := client.Timetable.Create().
		SetDatetimeFrom(to.Add(24 * time.Hour)).
		SetDatetimeTo(to.Add(4 * time.Hour)).
		SetComment("(TT) OWNER created a fifth timetable").
		SetAddress(a2).Save(ownerCtx)
	if st5 != nil {
		t.Fatal("should not be able to create a fifth timetable for the same address and startime= +24 hour first startfrom but with endtime 20 hours before starttime")
	} else if err != nil &&
		!strings.HasPrefix(err.Error(), hooks.ErrorMessagePrefix) &&
		!strings.HasPrefix(err.Error(), "ent:") { // this test should throw an ent validation error,
		// because of negative duration
		t.Log("error fourth a fifth timetable entry: ", err)
	}

	// creates a sixth entry one day later
	st6, err := client.Timetable.Create().
		SetDatetimeFrom(to.Add(24 * time.Hour)).
		SetDatetimeTo(to.Add(48 * time.Hour)).
		SetComment("(TT) OWNER created a sixth timetable").
		SetAddress(a2).Save(ownerCtx)
	if st6 == nil {
		t.Log(err)
		t.Fatal("should be able to create a sixth timetable for the same address and startime= +24 hour first startfrom and endtime= +24 hours")
	}
}

func printTTEntry(tt *ent.Timetable) string {
	return fmt.Sprintf("Timetable(id=%v, datetime_from=%v, datetime_to=%v, dur=%v type=%v", tt.ID, tt.DatetimeFrom.Format(time.RFC3339), tt.DatetimeTo.Format(time.RFC3339), tt.Duration, tt.TimetableType)
}
