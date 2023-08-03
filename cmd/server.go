package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/oauth"
	"github.com/hkonitzer/ohmab"
	"github.com/hkonitzer/ohmab/ent"
	"github.com/hkonitzer/ohmab/ent/intercept"
	_ "github.com/hkonitzer/ohmab/ent/runtime"
	"github.com/hkonitzer/ohmab/ent/schema"
	"github.com/hkonitzer/ohmab/ent/user"
	"github.com/hkonitzer/ohmab/internal/pkg/common/config"
	"github.com/hkonitzer/ohmab/internal/pkg/common/log"
	"github.com/hkonitzer/ohmab/internal/pkg/db"
	"github.com/hkonitzer/ohmab/internal/pkg/privacy"
	"github.com/hkonitzer/ohmab/internal/pkg/routes"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var srv *routes.Server

func main() {
	// Get logger
	logger := log.GetLoggerInstance()
	logger.Info().Msg("Starting OHMAB V" + config.Version)
	// Get the configuration
	configurations, err := config.Get()
	if err != nil {
		logger.Fatal().Msgf("Error reading configurations: %v", err)
	}

	// Get environment variable
	environment := configurations.ENVIRONMENT
	developmentMode := false
	if environment != "" && strings.ToLower(environment) == "development" {
		developmentMode = true
		logger.Debug().Msgf("DEVELOPMENT mode enabled, DEBUG LEVEL: %v", configurations.DEBUG)
	}

	ctx := context.TODO()
	client, clientError := db.CreateClient(ctx)
	if clientError != nil {
		logger.Fatal().Msgf("Error creating client: %v", clientError)
	}
	// Do not expose users on public api routes unless other stated in database
	// See Policy also in ent/schema/user.go
	client.User.Intercept(
		intercept.TraverseUser(func(ctx context.Context, q *ent.UserQuery) error {
			uv := privacy.FromContext(ctx)
			if uv != nil { // auth in context found, skip
				return nil
			}
			rc := chi.RouteContext(ctx)
			if !routes.HasPublicRoute(rc.RoutePatterns) {
				return nil
			}
			q.Where(user.UsePublicapiEQ(schema.UsePublicApiValue), user.Active(true)).
				Select(user.FieldTitle, user.FieldFirstname, user.FieldSurname)
			return nil
		}),
	)
	defer client.Close()

	// Setup client
	// create test data if development mode is enabled
	if developmentMode {
		createTestData(client)
	}

	srv = routes.NewServer(client)
	r := newRouter(srv)

	logger.Info().Msgf("connect to http://localhost:%d/ for GraphQL playground", configurations.Server.Port)

	httpErr := http.ListenAndServe(fmt.Sprintf(":%d", configurations.Server.Port), r)
	if httpErr != nil {
		logger.Fatal().Msgf("http Server terminated: %v", httpErr)
	}

}

// newRouter creates a new router with the blog handlers mounted.
func newRouter(srv *routes.Server) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	configs, _ := config.Get()
	if configs.DEBUG > 0 {
		r.Use(log.RequestLogger)
	}
	r.Use(middleware.Recoverer)
	r.Use(ESDRACtx)
	r.Use(SetHeadersHandler)
	r.Use(middleware.ThrottleBacklog(10, 100, 10*time.Second)) // @TODO: make this configurable

	routes.RegisterOAuthAPI(r, srv)

	// standard api routes, protected with OAuth
	r.Route("/", func(r chi.Router) {
		r.Use(oauth.Authorize(configs.OAUTHSECRETKEY, nil))
		r.Use(srv.CheckUser)
		if configs.ENVIRONMENT == config.DevelopmentEnvironment {
			r.Handle("/",
				playground.Handler("OHMAB", "/query"),
			)
		}
		entServer := handler.NewDefaultServer(OHMAB.NewSchema(srv.Client))
		r.Handle("/query", entServer)
	})

	// anonymous routes, including playground
	gqlPublicServer := handler.NewDefaultServer(OHMAB.NewSchema(srv.Client))
	r.Route(routes.PublicAPIRoute, func(r chi.Router) {
		r.Handle("/query", gqlPublicServer)
		r.Handle("/",
			playground.Handler("OHMAB", "/query"),
		)
		// html exports
		r.Get("/exports/timetables", srv.Timetables)
	})

	return r
}

func ESDRACtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "OHMAB", "OHMAB")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func SetHeadersHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers for the preflight request
		configurations, _ := config.Get()
		for key, val := range configurations.Server.Headers {
			w.Header().Set(key, val)
		}
		/*
			if r.Method == "OPTIONS" {
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			}
		*/
		next.ServeHTTP(w, r)
	},
	)
}

func createTestData(client *ent.Client) {
	// Get logger
	logger := log.GetLoggerInstance()

	var ctx = context.Background()
	ctx = privacy.NewContext(ctx, &privacy.UserViewer{Role: privacy.Admin})
	// Check if database is empty
	count, err := client.Business.Query().Count(ctx)
	if err != nil {
		logger.Fatal().Msgf("failed querying count of businesses: %v", err)

	}
	if count > 0 {
		logger.Debug().Msgf("Database not empty, found businesses. Skipping test data creation")
		return
	}

	logger.Debug().Msgf("creating test data")

	// create a TESTDATA tag first
	tempTagTestdata, err := client.Tag.Create().SetName("TESTDATA").SetComment("TESTDATA").Save(ctx)
	if err != nil {
		logger.Fatal().Msgf("failed creating a tag: %v", err)
	}
	for i := 1; i < 30; i++ {
		// create a random tag for each business
		tempTag := client.Tag.Create().SetName(getRandomString(20)).SetComment("TESTDATA").SaveX(ctx)
		// create a business with random telephonnumber
		name1 := fmt.Sprintf("B%v-name1 %v%v%v", i, i, i, i)
		name2 := fmt.Sprintf("B%v-name2", i)
		tempBusiness, err := client.Business.Create().
			SetName1(name1).
			SetName2(name2).
			SetAlias(name1).
			SetTelephone(getRandomIntAsString(10000, 90000)).
			SetComment("TESTDATA").
			AddTags(tempTagTestdata, tempTag).
			Save(ctx)
		if err != nil {
			logger.Fatal().Msgf("failed creating a business: %v", err)
		}
		// create a timetable entry

		// create two addresses for the business
		for ii := 1; ii < 3; ii++ {
			addressCreate := client.Address.Create().
				SetStreet(fmt.Sprintf("B%v-street %v", i, ii)).
				SetCity(fmt.Sprintf("B%v-city %v", i, ii)).
				SetComment("TESTDATA").
				SetAddition(getRandomString(10)).
				SetZip(fmt.Sprintf("B%v-zip %v", i, ii)).
				SetBusiness(tempBusiness)
			if ii == 1 {
				addressCreate.SetPrimary(true)
			}
			address := addressCreate.
				SaveX(ctx)

			daysindecember := time.Date(2023, 12, ii, 0, 0, 0, 0, time.Local)
			// create the timetable entry
			client.Timetable.Create().
				SetAddress(address).
				SetComment("TESTDATA").
				SetDatetimeFrom(daysindecember).
				SetDatetimeTo(daysindecember.Add(time.Hour * 9)).
				SaveX(ctx)
		}

	}
	knownBussiness, err := client.Business.Query().All(ctx)
	if err != nil {
		logger.Fatal().Msgf("failed querying businesses: %v", err)
	}
	client.User.Create().
		SetFirstname("unknown").
		SetSurname("developer").
		SetLogin("developer").
		SetEmail("dev@localhost").
		SetComment("TESTDATA").
		SetRole(privacy.AdminRoleAsString()).
		AddBusinesses(knownBussiness...).
		SaveX(ctx)
	logger.Debug().Msgf("test data created")
}

func getRandomInt(min, max int) int {
	if min == 0 {
		min = 1000
	}
	if max == 0 {
		max = 9999
	}
	return rand.Intn(max-min) + min
}

func getRandomIntAsString(min, max int) string {
	return fmt.Sprint(getRandomInt(min, max))
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func getRandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
