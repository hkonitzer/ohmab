package OHMAB

import (
	"context"
	"github.com/hkonitzer/ohmab/ent"
	"github.com/hkonitzer/ohmab/ent/business"
	"github.com/hkonitzer/ohmab/ent/enttest"
	"github.com/hkonitzer/ohmab/ent/operator"
	"github.com/hkonitzer/ohmab/ent/schema"
	"github.com/hkonitzer/ohmab/ent/user"
	"github.com/hkonitzer/ohmab/internal/pkg/privacy"
	"github.com/hkonitzer/ohmab/internal/pkg/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateUser(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	// client = client.Debug()
	defer client.Close()

	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal("failed creating schema resources: ", err)
	}

	// create admin viewer for context
	adminViewer := privacy.UserViewer{Role: privacy.Admin}
	adminViewer.SetUserID("TESTADMIN")
	// create admin context
	adminCtx := privacy.NewContext(context.Background(), &adminViewer)

	// create business first
	_, err := client.Business.Create().
		SetComment("TESTDATA").
		SetName1("TEST1").
		SetAlias("TEST").
		SetTelephone("123456789").
		Save(adminCtx)
	require.NoError(t, err)
	bus1, err := client.Business.Query().Where(business.Name1("TEST1")).
		Only(adminCtx)
	require.NoError(t, err)
	require.Equal(t, bus1.Name1, "TEST1", "name1 should be TEST1")

	// test user creation: Employee role
	// create password
	passwordHash, err := utils.HashPassword("TESTPASSWORD")
	require.NoError(t, err)
	_, err = client.User.Create().
		SetComment("TESTDATA").
		SetRole(privacy.EmployeeRoleAsString()).
		SetLogin("TESTUSER1").
		SetSurname("Musterfrau").
		SetFirstname("Erika").
		SetTitle("Dr.").
		SetEmail("emusterfrau@localhost").
		SetPasswordhash(passwordHash).
		AddBusinesses(bus1).
		SetUsePublicapi(schema.UsePublicApiValue).
		Save(adminCtx)
	require.NoError(t, err)
	quser1, err := client.User.Query().Where(user.Login("TESTUSER1")).WithBusinesses().Only(adminCtx)
	require.NoError(t, err)
	require.Contains(t, quser1.Comment, "TESTDATA", "comment should contain TESTDATA")
	require.Equal(t, quser1.Role, privacy.EmployeeRoleAsString(), "role should be Employee")
	require.Equal(t, quser1.Surname, "Musterfrau", "surname should be Musterfrau")
	require.Equal(t, quser1.Firstname, "Erika", "firstname should be Erika")
	require.Equal(t, quser1.Title, "Dr.", "title should be Dr.")
	require.Equal(t, quser1.Email, "emusterfrau@localhost", "email should be emusterfrau@localhost")
	require.Equal(t, quser1.Passwordhash, passwordHash, "passwordhash should be equal")
	require.Equal(t, utils.DoPasswordsMatch(quser1.Passwordhash, "TESTPASSWORD"), true, "passwordhash should match")
	require.Equal(t, len(quser1.Edges.Businesses), 1, "user should have exact one business")
	require.Equal(t, quser1.UsePublicapi, schema.UsePublicApiValue, "usePublicApi should be true")

	// query the business for the public user, created from the user above
	bus2, err := client.Business.Query().Where(business.Name1("TEST1")).WithOperators().Only(adminCtx)
	require.NoError(t, err)
	require.Len(t, bus2.Edges.Operators, 1, "business TEST1 should have one public user")
	pubus2 := bus2.Edges.Operators[0]
	require.Equal(t, "Erika", pubus2.Firstname, "surname should be Erika")
	require.Equal(t, "Musterfrau", pubus2.Surname, "surname should be Musterfrau")
	require.Equal(t, "Dr.", pubus2.Title, "title should be Dr.")
	require.Equal(t, "emusterfrau@localhost", pubus2.Email, "email should be emusterfrau@localhost")
	// create the same user on the bussiness again should fail
	_, err = client.User.Create().
		SetComment("TESTDATA").
		SetRole(privacy.EmployeeRoleAsString()).
		SetLogin("TESTUSER1").
		SetSurname("Musterfrau").
		SetFirstname("Erika").
		SetTitle("Dr.").
		SetEmail("emusterfrau@localhost").
		SetPasswordhash(passwordHash).
		AddBusinesses(bus2).
		SetUsePublicapi(schema.UsePublicApiValue).
		Save(adminCtx)
	require.Error(t, err, "should fail because user already exists: ", err)

	// update user (should update public user as well)
	_, err = client.User.UpdateOneID(quser1.ID).SetTitle("Prof.").Save(adminCtx)
	require.NoError(t, err)
	bus3, err := client.Business.Query().Where(business.Name1("TEST1")).WithOperators().Only(adminCtx)
	require.NoError(t, err)
	require.Len(t, bus3.Edges.Operators, 1, "business TEST1 should have one public user")
	pubus3 := bus3.Edges.Operators[0]
	require.Equal(t, "Erika", pubus3.Firstname, "surname should be Erika")
	require.Equal(t, "Musterfrau", pubus3.Surname, "surname should be Musterfrau")
	require.Equal(t, "Prof.", pubus3.Title, "title should be now Prof.")
	require.Equal(t, "emusterfrau@localhost", pubus3.Email, "email should be emusterfrau@localhost")
	// check the original user again
	quser2, err := client.User.Query().Where(user.Login("TESTUSER1")).WithBusinesses().Only(adminCtx)
	require.NoError(t, err)
	require.Contains(t, quser2.Comment, "TESTDATA", "comment should contain TESTDATA")
	require.Equal(t, privacy.EmployeeRoleAsString(), quser2.Role, "role should be Employee")
	require.Equal(t, "Musterfrau", quser2.Surname, "surname should be Musterfrau")
	require.Equal(t, "Erika", quser2.Firstname, "firstname should be Erika")
	require.Equal(t, "Prof.", quser2.Title, "title should be Dr.")
	require.Equal(t, "emusterfrau@localhost", quser2.Email, "email should be emusterfrau@localhost")
	require.Equal(t, passwordHash, quser2.Passwordhash, "passwordhash should be equal")
	require.Equal(t, true, utils.DoPasswordsMatch(quser1.Passwordhash, "TESTPASSWORD"), "passwordhash should match")
	require.Equal(t, 1, len(quser2.Edges.Businesses), "user should have exact one business")
	require.Equal(t, schema.UsePublicApiValue, quser2.UsePublicapi, "usePublicApi should be true")

	// crate a new user and then modify his public api usage
	quser3, err := client.User.Create().
		SetComment("TESTDATA").
		SetRole(privacy.EmployeeRoleAsString()).
		SetLogin("TESTUSER2").
		SetSurname("Doe").
		SetFirstname("Joe").
		SetTitle("Dr.").
		SetEmail("jdoe@localhost").
		SetPasswordhash(passwordHash).
		AddBusinesses(bus1).
		Save(adminCtx)
	require.NoError(t, err)
	require.Equal(t, "", quser3.UsePublicapi, "usePublicApi should be false")
	// switch to public api usage, after that a public user should be created
	_, err = client.User.UpdateOneID(quser3.ID).SetUsePublicapi(schema.UsePublicApiValue).Save(adminCtx)
	require.NoError(t, err)
	quser4, err := client.User.Query().Where(user.Login("TESTUSER2")).WithBusinesses().Only(adminCtx)
	require.NoError(t, err)
	require.Contains(t, "TESTDATA", quser4.Comment, "comment should contain TESTDATA")
	require.Equal(t, 1, len(quser4.Edges.Businesses), "user should have exact one business")
	require.Equal(t, quser4.UsePublicapi, schema.UsePublicApiValue, "usePublicApi should be true")
	// is there a public user for john doe? -should be since we set public_api to 1 above
	pquser5, err := client.Operator.Query().Where(operator.Email("jdoe@localhost")).Only(adminCtx)
	require.NoError(t, err)
	require.Equal(t, "Joe", pquser5.Firstname, "firstname should be Joe")
	require.Equal(t, "Doe", pquser5.Surname, "surname should be Doe")
	require.Equal(t, "Dr.", pquser5.Title, "title should be Dr.")
	require.Equal(t, "jdoe@localhost", pquser5.Email, "email should be jdoe@localhost")
	// switch back to normal api usage, after that the public user should be deleted
	_, err = client.User.UpdateOneID(quser4.ID).SetUsePublicapi("").Save(adminCtx)
	require.NoError(t, err)
	// check the result
	quser5, err := client.User.Query().Where(user.Login("TESTUSER2")).Only(adminCtx)
	require.NoError(t, err)
	require.Equal(t, "", quser5.UsePublicapi, "usePublicApi should be false")
	// the public user should be deleted
	_, err = client.Operator.Query().Where(operator.Email("jdoe@localhost")).Only(adminCtx)
	require.Error(t, err, "should fail because public user does not exist anymore")
	// delete the first user, should also delete the public user
	err = client.User.DeleteOneID(quser1.ID).Exec(adminCtx) // TESTUSER1
	require.NoError(t, err)
	_, err = client.User.Query().Where(user.Login("TESTUSER1")).OnlyID(adminCtx)
	require.Error(t, err, "should fail because user does not exist anymore")
	_, err = client.Operator.Query().Where(operator.Email("emusterfrau@localhost")).Only(adminCtx)
	require.Error(t, err, "should fail because user does not exist anymore")
}
