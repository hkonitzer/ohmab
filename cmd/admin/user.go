package admin

import (
	"github.com/hkonitzer/ohmab/internal/pkg/privacy"
	"github.com/spf13/cobra"
)

type newUser struct {
	Login     string
	EMail     string
	Password  string
	Firstname string
	Surname   string
	Title     string
	Role      int
	PublicAPI string
	Scopes    []string
}

var NewUser = newUser{}
var CmdUser = &cobra.Command{
	Use:   "user",
	Short: "Manage users",
	Long:  `Manage users`,
}

func init() {
	RootCmd.AddCommand(CmdUser)
	CmdUser.AddCommand(CmdCreateUser)
	CmdCreateUser.Flags().StringVarP(&NewUser.Login, "login", "l", "", "Login name")
	CmdCreateUser.MarkFlagRequired("login")
	CmdCreateUser.Flags().StringVarP(&NewUser.EMail, "email", "e", "", "E-Mail")
	CmdCreateUser.MarkFlagRequired("email")
	CmdCreateUser.Flags().StringVarP(&NewUser.Password, "password", "p", "", "Desired Password")
	CmdCreateUser.MarkFlagRequired("password")
	CmdCreateUser.Flags().StringVarP(&NewUser.Firstname, "firstname", "f", "", "Firstname")
	CmdCreateUser.MarkFlagRequired("firstname")
	CmdCreateUser.Flags().StringVarP(&NewUser.Surname, "surname", "s", "", "Surname")
	CmdCreateUser.MarkFlagRequired("surname")
	CmdCreateUser.Flags().StringVarP(&NewUser.Title, "title", "t", "", "Title")
	CmdCreateUser.Flags().IntVarP(&NewUser.Role, "role", "r", privacy.ViewerRole(), "Role")
	CmdCreateUser.Flags().StringSliceVarP(&NewUser.Scopes, "scopes", "", nil, "Scopes")
	CmdCreateUser.Flags().StringVarP(&NewUser.PublicAPI, "publicapi", "a", "", "User will be shown in the public api as operator")
}
