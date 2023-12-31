package OHMAB

import (
	"context"
	"github.com/hkonitzer/ohmab/ent/enttest"
	"github.com/hkonitzer/ohmab/internal/pkg/privacy"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestCreateAddress(t *testing.T) {

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
		SetSurname("Erika").
		SetFirstname("Musterfrau").
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
	adminCtx = adminViewer.ToContext(adminCtx)
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
	ownerCtx = ownerViewer.ToContext(ownerCtx) // update the context with the new scopes
	// create a new address as Admin
	_, err = client.Address.Create().
		SetComment("(AT) ADMIN created a address").
		SetBusiness(business).
		Save(adminCtx)
	if err != nil {
		t.Errorf("failed creating address: %v", err)
	}
	// Create a new address as Owner
	_, err = client.Address.Create().
		SetComment("(AT) OWNER created a address").
		SetBusiness(business).
		Save(ownerCtx)
	if err != nil {
		t.Errorf("failed creating address: %v", err)
	}
}
