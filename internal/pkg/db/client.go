package db

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/hkonitzer/ohmab/ent"
	"github.com/hkonitzer/ohmab/internal/pkg/common/log"
	"github.com/spf13/viper"
)

// Get logger
var logger = log.GetLoggerInstance()

func ClientDebuglog(v ...any) {
	logger.Printf("%v", v...)
}

func CreateClient(ctx context.Context) (*ent.Client, error) {
	d := viper.GetString("database.dialect")
	if !viper.IsSet("database.dsn") {
		return nil, fmt.Errorf("DSN for database (dialect=%v) not set", d)
	}
	debug := viper.GetInt("DEBUG") > 1
	var client *ent.Client

	switch d {
	case dialect.MySQL:
		return nil, fmt.Errorf("MySQL not implemented yet")
	case dialect.Postgres:
		client = CreatePGClient(viper.GetString("database.dsn"), ctx, debug)
	case dialect.SQLite:
		client = CreateSQLiteClient(viper.GetString("database.dsn"), ctx, debug)
	default:
		return nil, fmt.Errorf("unknown dialect: %v", d)
	}
	if debug {
		client = client.Debug()
	}
	return client, nil
}
