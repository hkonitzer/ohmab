package admin

import (
	"context"
	"fmt"
	"github.com/hkonitzer/ohmab/internal/pkg/db"
	"github.com/hkonitzer/ohmab/internal/pkg/privacy"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"time"
)

var CmdListBusiness = &cobra.Command{
	Use:   "list",
	Short: "List all businesses",
	Long:  `List all businesses`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.TODO()
		uv := privacy.UserViewer{Role: privacy.Admin}
		uv.SetUserID("cli")
		ctx = uv.ToContext(ctx)
		client, clientError := db.CreateClient(ctx)
		if clientError != nil {
			panic(fmt.Sprintf("Error creating client: %v", clientError))
		}
		defer client.Close()
		b, err := client.Business.Query().All(ctx)
		if err != nil {
			fmt.Println(fmt.Errorf("failed listing businesses: %v", err))
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Name1", "Name2", "Alias", "Telephone", "Active", "Updated"})

		for _, v := range b {
			table.Append([]string{
				v.ID.String(),
				v.Name1,
				*v.Name2,
				v.Alias,
				*v.Telephone,
				strconv.FormatBool(v.Active),
				v.UpdatedAt.Format(time.DateTime),
			})
		}
		table.Render()
	},
}
