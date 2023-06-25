package db

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"hynie.de/ohmab/ent"
	"hynie.de/ohmab/internal/pkg/config"
	"hynie.de/ohmab/internal/pkg/db/postgresql"
	"hynie.de/ohmab/internal/pkg/db/sqlite"
)

func CreateClient(ctx context.Context, configurations *config.Configurations) (*ent.Client, error) {
	if configurations.Database.DSN == "" {
		return nil, fmt.Errorf("DSN for database (dialect=%v) not set", configurations.Database.Dialect)
	}
	debug := false
	if configurations.DEBUG > 0 {
		debug = true
	}
	var client *ent.Client
	switch configurations.Database.Dialect {
	case dialect.MySQL:
		return nil, fmt.Errorf("MySQL not implemented yet")
	case dialect.Postgres:
		client = postgresql.CreatePGClient(configurations.Database.DSN, ctx, debug)
	case dialect.SQLite:
		client = sqlite.CreateSQLiteClient(configurations.Database.DSN, ctx, debug)
	default:
		return nil, fmt.Errorf("unknown dialect: %v", configurations.Database.Dialect)
	}
	return client, nil
}
