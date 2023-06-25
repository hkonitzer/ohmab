package sqlite

import (
	"context"
	"entgo.io/ent/dialect"
	"hynie.de/ohmab/ent"
	"hynie.de/ohmab/ent/migrate"
	"hynie.de/ohmab/internal/pkg/log"
)

func CreateSQLiteClient(dsn string, ctx context.Context, debug bool) *ent.Client {
	// Get logger
	logger := log.GetLoggerInstance()

	// Create entutils.Client and run the schema migration.
	client, err := ent.Open(dialect.SQLite, dsn)
	if err != nil {
		logger.Fatal().Msgf("opening entutils sqlite client", err)
	}
	if err := client.Schema.Create(
		ctx,
		migrate.WithGlobalUniqueID(false),
		//migrate.WithDropIndex(true),
		//migrate.WithDropColumn(true),
	); err != nil {
		logger.Fatal().Msgf("creating schema", err)
	} else {
		logger.Debug().Msgf("entutils db client created")
	}
	if debug {
		client = client.Debug()
	}

	return client
}
