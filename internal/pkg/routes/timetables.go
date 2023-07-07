package routes

import (
	"context"
	"html/template"
	"hynie.de/ohmab/ent"
	"hynie.de/ohmab/ent/business"
	"hynie.de/ohmab/ent/content"
	"hynie.de/ohmab/ent/timetable"
	"hynie.de/ohmab/internal/pkg/utils"
	"hynie.de/ohmab/templates"
	"net/http"
	"strings"
	"time"
)

type DataTemplate struct {
	Title             string
	Locale            string
	TimetableTypeData []TimetableData
}
type TimetableData struct {
	TimetableType string
	CSSTop        template.CSS
	CSSBottom     template.CSS
	HTMLTop       template.HTML
	HTMLBottom    template.HTML
	Data          []*ent.Timetable
}

func (s *Server) Timetables(w http.ResponseWriter, r *http.Request) {
	// get request timetable types
	reqttTypes := r.URL.Query().Get("type")
	ttTypes := strings.Split(reqttTypes, ",")
	// parse requested locale @TODO: use query param also?
	langq := utils.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
	var locale string
	if len(langq) == 0 {
		locale = "en_US"
	} else {
		locale = langq[0].Locale
	}
	data := DataTemplate{
		Title:  "Timetables",
		Locale: locale,
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
		timetablesQuery = timetablesQuery.Where(timetable.TimetableTypeEQ(timetable.TimetableType(timetabletype)))
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
	// setup data for the template
	ttdata := TimetableData{
		TimetableType: timetabletype,
	}
	// fill timetable address business edge with map created above
	for _, tt := range timetables {
		tt.Edges.Address.Edges.Business = bsMap[tt.Edges.Address.ID.String()]
		tt.DatetimeTo = tt.DatetimeTo.In(time.Local)
		tt.DatetimeFrom = tt.DatetimeFrom.In(time.Local)
	}
	ttdata.Data = timetables
	// get content for the timetable type
	// only the latest one (published_at) will be fetched
	// location ignored for now @TODO add location handling below and in TimetableData struct
	contentHTMLQuery := client.Content.Query().
		Where(
			content.StatusEQ(content.StatusPUBLISHED),
			content.Or(
				content.TypeEQ(content.TypeHTML),
				content.TypeEQ(content.TypeCSS),
			),

			content.Locale(data.Locale),
			content.PublishedAtLTE(time.Now()),
		).
		Order(ent.Desc(content.FieldPublishedAt))
	if timetabletype != "" {
		contentHTMLQuery = contentHTMLQuery.Where(content.TimetableTypeEQ(content.TimetableType(timetabletype)))
	}
	co, err := contentHTMLQuery.All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		logger.Err(err).Msgf("Error getting HTML content for timetable type '%s'", timetabletype)
	} else if co == nil {
		logger.Debug().Msgf("No HTML content found for timetable type '%s' and locale '%s'", timetabletype, data.Locale)
	} else {
		for _, c := range co {
			if c.Type == content.TypeHTML {
				if c.Location == content.LocationTOP {
					ttdata.HTMLTop = template.HTML(c.Content)
				} else if c.Location == content.LocationBOTTOM {
					ttdata.HTMLBottom = template.HTML(c.Content)
				} else {
					logger.Debug().Msgf("Ignoring HTML content id '%s' with location '%s'", c.ID, c.Location)
				}
			} else if c.Type == content.TypeCSS {
				if c.Location == content.LocationTOP {
					ttdata.CSSTop = template.CSS(c.Content)
				} else if c.Location == content.LocationBOTTOM {
					ttdata.CSSBottom = template.CSS(c.Content)
				} else {
					logger.Debug().Msgf("Ignoring CSS content id '%s' with location '%s'", c.ID, c.Location)
				}
			}
		}
	}

	data.TimetableTypeData = append(data.TimetableTypeData, ttdata)
	return len(timetables), nil
}
