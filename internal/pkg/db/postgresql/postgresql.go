package postgresql

import (
	"context"
	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	"hynie.de/ohmab/ent"
	"hynie.de/ohmab/ent/migrate"
	"hynie.de/ohmab/internal/pkg/log"
)

func CreatePGClient(dsn string, ctx context.Context, debug bool) *ent.Client {
	// Get logger
	logger := log.GetLoggerInstance()
	// Create entutils.Client and run the schema migration.
	client, err := ent.Open(dialect.Postgres, dsn)
	if err != nil {
		logger.Fatal().Msgf("opening entutils pg client", err)
	}
	if err := client.Schema.Create(
		ctx,
		migrate.WithGlobalUniqueID(false),
	); err != nil {
		logger.Fatal().Msgf("creating schema", err)
	} else {
		logger.Debug().Msgf("entutils db client opened and schema created")
	}
	if debug {
		client = client.Debug()
	}
	return client
}
