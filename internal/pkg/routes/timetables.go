package routes

import (
	"context"
	"hynie.de/ohmab/ent"
	"hynie.de/ohmab/ent/business"
	"hynie.de/ohmab/ent/timetable"
	"hynie.de/ohmab/templates"
	"net/http"
	"strings"
	"time"
)

func (s *Server) Timetables(w http.ResponseWriter, r *http.Request) {
	reqttTypes := r.URL.Query().Get("type")
	ttTypes := strings.Split(reqttTypes, ",")

	data := DataTemplate{
		Title: "Timetables",
	}

	if len(ttTypes) == 0 {
		_, err := GetStandardTimetableData(&data, "", s.Client, r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		for _, ttType := range ttTypes {
			c, err := GetStandardTimetableData(&data, strings.TrimSpace(ttType), s.Client, r.Context())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			} else if c > 0 {
				data.Title += " " + ttType
			}
		}
	}

	if err := templates.FullTimeTableHTMLTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type DataTemplate struct {
	Title    string
	TypeData []TimetableData
}
type TimetableData struct {
	TimetableType string
	HeaderHTML    string
	FooterHTML    string
	Data          []*ent.Timetable
}

func GetStandardTimetableData(data *DataTemplate, timetabletype string, client *ent.Client, ctx context.Context) (int, error) {
	/* @TODO: This is not working, why? --> Order ByTimetables is not working
	a, err := s.Client.Address.Query().Where(address.HasTimetables()).
		WithTimetables(
			func(q *ent.TimetableQuery) {
				q.Where(timetable.DatetimeFromGTE(time.Now())).All(r.Context())
				//q.Order(ent.Asc(timetable.FieldDatetimeFrom))
			},
		).Order(address.ByTimetables(sql.OrderTerm(sql.OrderByField(timetable.FieldDatetimeFrom, sql.OrderDesc())))).
		WithBusiness().All(r.Context())
	*/
	// Create a map with address id as key and the business as value
	as, err := client.Address.Query().WithBusiness(
		func(q *ent.BusinessQuery) {
			q.Where(business.HasAddresses(), business.DeletedAtIsNil(), business.Active(true))
		},
	).All(ctx)
	if err != nil {
		logger.Err(err).Msgf("Error getting addresses with businesses")
		return 0, err
	}
	bsMap := make(map[string]*ent.Business)
	for _, a := range as {
		bsMap[a.ID.String()] = a.Edges.Business
	}
	// construct time for where clause: get all timetables valid from today 00:00:00
	t := time.Now()
	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)

	// get timetables for every requested type

	timetablesQuery := client.Timetable.Query().Where(timetable.DatetimeFromGTE(midnight)).WithAddress().
		Order(ent.Asc(timetable.FieldDatetimeFrom))

	if timetabletype != "" {
		timetablesQuery = timetablesQuery.Where(timetable.TypeEQ(timetable.Type(timetabletype)))
	}
	timetables, err := timetablesQuery.All(ctx)
	if err != nil {
		logger.Err(err).Msgf("Error getting timetables")
		return 0, err
	}
	if len(timetables) == 0 {
		logger.Debug().Msgf("No timetables found for type '%s'", timetabletype)
		return 0, nil
	}
	// fill timetable address business edge with map created above
	for _, tt := range timetables {
		tt.Edges.Address.Edges.Business = bsMap[tt.Edges.Address.ID.String()]
		tt.DatetimeTo = tt.DatetimeTo.In(time.Local)
		tt.DatetimeFrom = tt.DatetimeFrom.In(time.Local)
	}
	ttdata := TimetableData{
		TimetableType: timetabletype,
		Data:          timetables,
	}
	data.TypeData = append(data.TypeData, ttdata)
	return len(timetables), nil
}
