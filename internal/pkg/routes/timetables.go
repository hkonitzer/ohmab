package routes

import (
	"hynie.de/ohmab/ent"
	"hynie.de/ohmab/ent/business"
	"hynie.de/ohmab/ent/timetable"
	"hynie.de/ohmab/templates"
	"net/http"
	"time"
)

func (s *Server) Timetables(w http.ResponseWriter, r *http.Request) {
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
	as, err := s.Client.Address.Query().WithBusiness(
		func(q *ent.BusinessQuery) {
			q.Where(business.HasAddresses(), business.DeletedAtIsNil(), business.Active(true))
		},
	).All(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bsMap := make(map[string]*ent.Business)
	for _, a := range as {
		bsMap[a.ID.String()] = a.Edges.Business
	}
	// get timetables
	timetablesQuery := s.Client.Timetable.Query().Where(timetable.DatetimeFromGTE(time.Now())).WithAddress().
		Order(ent.Asc(timetable.FieldDatetimeFrom))
	ttType := r.URL.Query().Get("type")
	if ttType != "" {
		timetablesQuery = timetablesQuery.Where(timetable.TypeEQ(timetable.Type(ttType)))
	}
	timetables, err := timetablesQuery.All(r.Context())
	// fill timetable address business edge with map created above
	for _, tt := range timetables {
		tt.Edges.Address.Edges.Business = bsMap[tt.Edges.Address.ID.String()]
		tt.DatetimeTo = tt.DatetimeTo.In(time.Local)
		tt.DatetimeFrom = tt.DatetimeFrom.In(time.Local)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := templates.Tmpl.Execute(w, timetables); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
