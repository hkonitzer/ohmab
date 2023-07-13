package db

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/hkonitzer/ohmab/ent"
	"github.com/hkonitzer/ohmab/ent/migrate"
	_ "github.com/lib/pq"
)

func CreatePGClient(dsn string, ctx context.Context, debug bool) *ent.Client { // Create entutils.Client and run the schema migration.
	client, err := ent.Open(dialect.Postgres, dsn, ent.Log(ClientDebuglog))
	if err != nil {
		logger.Fatal().Msgf("opening entutils pg client: %v", err)
	}
	if err := client.Schema.Create(
		ctx,
		migrate.WithGlobalUniqueID(false),
	); err != nil {
		logger.Fatal().Msgf("creating schema: %v", err)
	} else {
		logger.Debug().Msgf("entutils db client opened and schema created")
	}
	if debug {
		client = client.Debug()
	}
	return client
}
