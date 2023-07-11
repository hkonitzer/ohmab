package admin

import (
	"github.com/spf13/cobra"
)

var CmdBusiness = &cobra.Command{
	Use:   "business",
	Short: "Manage Businesses",
	Long:  `Manage Businesses, e.g. View, Create and Update them`,
}

func init() {
	RootCmd.AddCommand(CmdBusiness)
	CmdBusiness.AddCommand(CmdListBusiness)

}
