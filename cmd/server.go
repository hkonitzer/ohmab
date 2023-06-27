package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	ESDRA "hynie.de/ohmab"
	"hynie.de/ohmab/ent"
	_ "hynie.de/ohmab/ent/runtime"
	"hynie.de/ohmab/internal/pkg/config"
	"hynie.de/ohmab/internal/pkg/db"
	"hynie.de/ohmab/internal/pkg/log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func main() {
	// Get logger
	logger := log.GetLoggerInstance()

	// Get the configuration
	configurations, err := config.GetConfiguration()
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
	client, clientError := db.CreateClient(ctx, configurations)
	if clientError != nil {
		logger.Fatal().Msgf("Error creating client: %v", clientError)
	}
	defer client.Close()

	// Setup client
	// create test data if development mode is enabled
	if developmentMode {
		createTestData(client)
	}

	srv := newServer(client)
	r := newRouter(srv)

	logger.Info().Msgf("connect to http://localhost:%d/ for GraphQL playground", configurations.Server.Port)

	httpErr := http.ListenAndServe(fmt.Sprintf(":%d", configurations.Server.Port), r)
	if httpErr != nil {
		logger.Fatal().Msgf("http server terminated", httpErr)
	}

}

type server struct {
	client *ent.Client
}

func newServer(client *ent.Client) *server {
	return &server{client: client}
}

// newRouter creates a new router with the blog handlers mounted.
func newRouter(srv *server) chi.Router {
	r := chi.NewRouter()
	//r.Use(middleware.Logger)
	configs, _ := config.GetConfiguration()
	if configs.DEBUG > 0 {
		r.Use(log.RequestLogger)
	}
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(ESDRACtx)
	r.Use(SetHeadersHandler)
	r.Handle("/",
		playground.Handler("OHMAB", "/query"),
	)
	entServer := handler.NewDefaultServer(ESDRA.NewSchema(srv.client))
	r.Handle("/query",
		entServer,
	)

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
		configurations, _ := config.GetConfiguration()
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
		firstofdecember := time.Date(2023, 12, 1, 0, 0, 0, 0, time.Local)

		// create the address for the business
		address := client.Address.Create().
			SetStreet(fmt.Sprintf("B%v-street", i)).
			SetCity(fmt.Sprintf("B%v-city", i)).
			SetComment("TESTDATA").
			SetAddition(getRandomString(10)).
			SetZip(fmt.Sprintf("B%v-zip", i)).
			AddBusiness(tempBusiness).
			SaveX(ctx)
		// create the timetable entry
		client.Timetable.Create().
			SetAddress(address).
			SetComment("TESTDATA").
			SetDatetimeFrom(firstofdecember).
			SetDatetimeTo(firstofdecember.Add(time.Hour * 9)).
			SaveX(ctx)

	}
	knownBussiness, err := client.Business.Query().All(ctx)
	if err != nil {
		logger.Fatal().Msgf("failed querying businesses: %v", err)
	}
	client.User.Create().
		SetFirstname("unknown").
		SetSurname("developer").
		SetEmail("dev@localhost").
		SetComment("TESTDATA").
		SetRole(1).
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
