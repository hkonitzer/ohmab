package admin

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/hkonitzer/ohmab/ent/business"
	_ "github.com/hkonitzer/ohmab/ent/runtime"
	"github.com/hkonitzer/ohmab/internal/pkg/db"
	"github.com/hkonitzer/ohmab/internal/pkg/privacy"
	"github.com/hkonitzer/ohmab/internal/pkg/utils"
	"github.com/spf13/cobra"
	"os"
)

var CmdCreateUser = &cobra.Command{
	Use:   "create",
	Short: "Create a new user",
	Long:  `Create a new user`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.TODO()
		uv := privacy.UserViewer{Role: privacy.Admin}
		uv.SetUserID("cli")
		ctx = privacy.NewContext(ctx, &uv)
		client, clientError := db.CreateClient(ctx)
		if clientError != nil {
			panic(fmt.Sprintf("Error creating client: %v", clientError))
		}
		defer client.Close()
		pwh, err := utils.HashPassword(NewUser.Password)
		if err != nil {
			fmt.Println(fmt.Errorf("failed hashing password: %v", err))
			os.Exit(1)
		}
		u, err := client.User.Create().
			SetLogin(NewUser.Login).
			SetEmail(NewUser.EMail).
			SetFirstname(NewUser.Firstname).
			SetSurname(NewUser.Surname).
			SetPasswordhash(pwh).SetRole(fmt.Sprint(NewUser.Role)).
			Save(ctx)
		if err != nil {
			fmt.Println(fmt.Errorf("failed creating user: %v", err))
			os.Exit(1)
		}
		fmt.Println("User created with ID: ", u.ID)
		if NewUser.Scopes != nil && len(NewUser.Scopes) > 0 {

			for _, scope := range NewUser.Scopes {
				id := uuid.Must(uuid.Parse(scope))
				bq := client.Business.Query().Where(business.IDEQ(id)).WithUsers()
				b, err := bq.First(ctx)
				if err != nil {
					fmt.Println(fmt.Errorf("cannot get sccope: %v", err))
					os.Exit(1)
				}
				err = b.Update().AddUsers(u).Exec(ctx)
				if err != nil {
					fmt.Println(fmt.Errorf("cannot add user scope to business: %v", err))
					os.Exit(1)
				}
				fmt.Println("scope added: ", b.ID)
			}

		}
	},
}
