package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"sync"
)

const DevelopmentEnvironment = "DEVELOPMENT"

// Configurations exported
type Configurations struct {
	Server         ServerConfigurations
	Database       DatabaseConfigurations
	Ghost          GhostConfigurations
	DEBUG          int
	ENVIRONMENT    string
	OAUTHSECRETKEY string
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

type GhostConfigurations struct {
	BaseURL string
	Key     string
	Locale  string
	Content []map[string]GhostContentConfigurations `mapstructure:"content"`
}

type GhostContentConfigurations struct {
	PostID    string `mapstructure:"post"`
	PostTitle string `mapstructure:"posttitle"`
	PageID    string `mapstructure:"page"`
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
var initOnce, readOnce sync.Once

var Version string

func Get() (*Configurations, error) {
	var err error = nil
	readOnce.Do(func() {
		Init()
		// Unmarshal the configuration file into the Config variable.
		err = viper.Unmarshal(&configurations)
	})
	return configurations, err
}
func GetX() *Configurations {
	if configurations == nil {
		readOnce.Do(func() {
			Init()
			// Unmarshal the configuration file into the Config variable.
			err := viper.Unmarshal(&configurations)
			if err != nil {
				panic(fmt.Sprintf("Unable to decode config into struct, %v", err))
			}
		})
	}
	return configurations
}

func Init() {
	initOnce.Do(read)
}

func configFileExists(path string) bool {
	info, err := os.Stat(path + "/config.yml")
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func read() {
	// Set the file name of the config file
	viper.SetConfigName("config")

	if configFileExists(".") { // search config in current directory
		viper.AddConfigPath(".")
	} else { // use home directory
		home, err := os.UserHomeDir()
		if err != nil {
			panic(fmt.Sprintf("Error getting user home directory, %s", err))
		}
		if configFileExists(home) {
			viper.AddConfigPath(home)
		} else {
			// use config provided by flag
			// Set the path to look for the configurations' file
			configFile := flag.String("config", "./config.yml", "path to the config file")
			flag.Parse()
			viper.AddConfigPath(*configFile)
		}
	}

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	// Set undefined variables (defaults)
	viper.SetDefault("server.port", 8081)
	viper.SetDefault("database.dialect", "sqlite3")
	viper.SetDefault("database.dbname", "ohmab")
	viper.SetDefault("database.dbuser", "ohmab")
	viper.SetDefault("DEBUG", 0)
	viper.SetDefault("ENVIRONMENT", DevelopmentEnvironment)
	viper.SetDefault("OAUTHSECRETKEY", "OHMAB-Secret-Key")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Error reading config file, %s", err))
	}

	if !viper.IsSet("database.dsn") {
		if viper.Get("database.dialect") == "postgres" {
			viper.SetDefault("database.postgres.sslmode", "disable")
			viper.Set("database.dsn",
				fmt.Sprintf("postgres://%s:%s@l%s:%v/%s?sslmode=%s&application_name=OHMAB",
					viper.Get("database.dbuser"),
					viper.Get("database.dbpassword"),
					viper.Get("database.host"),
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
}
