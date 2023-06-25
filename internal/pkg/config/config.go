package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

// Configurations exported
type Configurations struct {
	Server      ServerConfigurations
	Database    DatabaseConfigurations
	DEBUG       int
	ENVIRONMENT string
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port    int
	Headers map[string]string
}

type PostgresDatabaseConfigurations struct {
	SSLmode string // see https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-PARAMKEYWORDS
}

type Sqlite3DatabaseConfigurations struct {
	Cache string // see https://www.sqlite.org/uri.html#recognized_query_parameters
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int
	Dialect    string // possible values: postgres, mysql, sqlite3
	DSN        string // data source name
	Postgres   PostgresDatabaseConfigurations
	Sqlite3    Sqlite3DatabaseConfigurations
}

var configurations *Configurations
var once sync.Once

func GetConfiguration() (*Configurations, error) {
	var err error = nil
	once.Do(func() {
		configurations, err = ReadConfiguration()
	})
	return configurations, err
}

func ReadConfiguration() (*Configurations, error) {
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configurations Configurations

	// Set undefined variables (defaults)
	viper.SetDefault("server.port", 8081)
	viper.SetDefault("database.dialect", "sqlite3")
	viper.SetDefault("database.dbname", "ohmab")
	viper.SetDefault("database.dbuser", "ohmab")
	viper.SetDefault("DEBUG", 0)
	viper.SetDefault("ENVIRONMENT", "DEVELOPMENT")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return nil, err
	} else {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	if !viper.IsSet("database.dsn") {
		if viper.Get("database.dialect") == "postgres" {
			viper.SetDefault("database.postgres.sslmode", "disable")
			viper.Set("database.dsn",
				fmt.Sprintf("postgres://%s:%s@localhost:%d/%s?sslmode=%s&application_name=OHMAB",
					viper.Get("database.dbuser"),
					viper.Get("database.dbpassword"),
					viper.Get("database.dbport"),
					viper.Get("database.dbname"),
					viper.Get("database.postgres.sslmode"),
				))
		} else if viper.Get("database.dialect") == "sqlite3" {
			viper.SetDefault("database.sqlite3.cache", "shared")
			viper.SetDefault("database.sqlite3.mode", "memory")
			viper.Set("database.dsn",
				fmt.Sprintf("file:%s.db?mode=%s&cache=%s&_fk=1",
					viper.Get("database.dbname"),
					viper.Get("ddatabase.sqlite3.mode"),
					viper.Get("database.sqlite3.cache"),
				))
		}
	}

	// Unmarshal the configuration file into the Config variable.
	err := viper.Unmarshal(&configurations)
	if err != nil {
		fmt.Printf("Unable to decode config into struct, %v", err)
	}

	return &configurations, nil
}
