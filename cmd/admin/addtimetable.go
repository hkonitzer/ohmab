package admin

import (
	"context"
	"fmt"
	"github.com/hkonitzer/ohmab/ent"
	"github.com/hkonitzer/ohmab/ent/address"
	"github.com/hkonitzer/ohmab/ent/business"
	"github.com/hkonitzer/ohmab/ent/timetable"
	"github.com/hkonitzer/ohmab/internal/pkg/db"
	"github.com/hkonitzer/ohmab/internal/pkg/privacy"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

type newTimetableEntry struct {
	BusinessAlias string
	From          string
	Duration      int
	TTType        string
}

var NewTimetableEntry = newTimetableEntry{}

var CmdAddTimetableEntry = &cobra.Command{
	Use:   "addtimetableentry",
	Short: "Add a timetable entry",
	Long:  `Add a timetable entry for a business primary address`,
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
		b, err := client.Business.Query().Where(business.Alias(NewTimetableEntry.BusinessAlias)).WithAddresses(
			func(q *ent.AddressQuery) {
				q.Where(address.PrimaryEQ(true))
			},
		).
			First(ctx)
		if err != nil {
			fmt.Println(fmt.Errorf("failed get business for given alias: %v", err))
			return
		}
		a := b.Edges.Addresses[0]
		// check for existing entry
		from, err := time.Parse(time.DateTime, NewTimetableEntry.From)
		if err != nil {
			fmt.Println("Error parsing from date: ", err)
			return
		}
		if !from.After(time.Now()) {
			fmt.Println("Timetable entry must be in the future: ", from)
			return
		}
		tc := client.Timetable.Create().SetAddress(a).
			SetDatetimeFrom(from).
			SetDuration(NewTimetableEntry.Duration)
		var tttype timetable.TimetableType
		switch strings.ToUpper(NewTimetableEntry.TTType) {
		case timetable.TimetableTypeEMERGENCYSERVICE.String():
			tttype = timetable.TimetableTypeEMERGENCYSERVICE
		case timetable.TimetableTypeCLOSED.String():
			tttype = timetable.TimetableTypeCLOSED
		case timetable.TimetableTypeHOLIDAY.String():
			tttype = timetable.TimetableTypeHOLIDAY
		case timetable.TimetableTypeSPECIAL.String():
			tttype = timetable.TimetableTypeSPECIAL
		case timetable.TimetableTypeREGULAR.String():
			tttype = timetable.TimetableTypeREGULAR
		default:
			tttype = timetable.TimetableTypeDEFAULT
		}
		// check for existing entry
		c, err := client.Timetable.Query().Where(timetable.TimetableTypeEQ(tttype)).
			Where(timetable.HasAddressWith(address.IDEQ(a.ID))).
			Where(timetable.DatetimeFromEQ(from)).Count(ctx)
		if err != nil {
			fmt.Println(fmt.Errorf("failed to check for existing timetable entry: %v", err))
			return
		} else if c > 0 {
			fmt.Println(fmt.Errorf("timetable entry already exists"))
			return
		}
		tc.Mutation().SetTimetableType(tttype)
		err = tc.Exec(ctx)
		if err != nil {
			fmt.Println(fmt.Errorf("failed to create timetable entry: %v", err))
		}
		fmt.Println("Timetable entry created: ")

	},
}

func init() {
	CmdAddTimetableEntry.Flags().StringVarP(&NewTimetableEntry.BusinessAlias, "alias", "a", "", "Business alias for this entry")
	CmdAddTimetableEntry.MarkFlagRequired("alias")
	CmdAddTimetableEntry.Flags().StringVarP(&NewTimetableEntry.From, "from", "f", "", "Datetime-from which this entry is valid")
	CmdAddTimetableEntry.MarkFlagRequired("from")
	CmdAddTimetableEntry.Flags().StringVarP(&NewTimetableEntry.TTType, "type", "t", "EMERGENCYSERVICE", "Type of this Timetable entry, e.g. EMERGENCYSERVICE")
	CmdAddTimetableEntry.Flags().IntVarP(&NewTimetableEntry.Duration, "duration", "d", 24, "Duration in hours, default=24h")

}
