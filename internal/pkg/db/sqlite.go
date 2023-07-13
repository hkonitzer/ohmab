package db

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/hkonitzer/ohmab/ent"
	"github.com/hkonitzer/ohmab/ent/migrate"
	"github.com/hkonitzer/ohmab/internal/pkg/common/log"
)

func CreateSQLiteClient(dsn string, ctx context.Context, debug bool) *ent.Client {
	// Get logger
	logger := log.GetLoggerInstance()

	// Create entutils.Client and run the schema migration.
	client, err := ent.Open(dialect.SQLite, dsn, ent.Log(ClientDebuglog))
	if err != nil {
		logger.Fatal().Msgf("opening entutils sqlite client: %v", err)
	}
	if err := client.Schema.Create(
		ctx,
		migrate.WithGlobalUniqueID(false),
		//migrate.WithDropIndex(true),
		//migrate.WithDropColumn(true),
	); err != nil {
		logger.Fatal().Msgf("creating schema: %v", err)
	} else {
		logger.Debug().Msgf("entutils db client created")
	}
	if debug {
		client = client.Debug()
	}

	return client
}
