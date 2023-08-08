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

func initFlags() (configFilePath *string, configFile *string) {
	configFilePath = flag.String("configpath", ".", "path to the config file")
	configFile = flag.String("config", "config", "name of the config file (does not inculde extension yml!)")
	flag.Parse()
	return configFilePath, configFile
}

func configFileExists(path string, filename string) bool {
	if filename == "" {
		filename = "config"
	}
	info, err := os.Stat(path + "/" + filename + ".yml")
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func read() {
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	// Set undefined variables (defaults)
	viper.SetDefault("server.port", 8081)
	viper.SetDefault("database.dbname", "ohmab")
	viper.SetDefault("database.dbuser", "ohmab")
	viper.SetDefault("database.dialect", "sqlite3")
	viper.SetDefault("DEBUG", 0)
	viper.SetDefault("ENVIRONMENT", DevelopmentEnvironment)
	viper.SetDefault("OAUTHSECRETKEY", "OHMAB-Secret-Key")

	// Set the file name of the config file
	configFilePath, configFile := initFlags()
	viper.SetConfigName(*configFile)
	viper.SetConfigType("yml")
	configFromFileUsed := false
	if configFileExists(*configFilePath, *configFile) { // search config in current directory
		viper.AddConfigPath(*configFilePath)
		configFromFileUsed = true
	} else { // use home directory
		home, err := os.UserHomeDir()
		if err != nil {
			panic(fmt.Sprintf("reading config file: error getting user home directory, %s", err))
		}
		if configFileExists(home, "") {
			viper.AddConfigPath(home)
			configFromFileUsed = true
		} else {
			fmt.Printf("WARNING: Could not read config file '%s' from %s or %s\n", *configFile, *configFilePath, home)
		}
	}
	if configFromFileUsed {
		if err := viper.ReadInConfig(); err != nil {
			if err != nil {
				panic(fmt.Sprintf("error reading config file: %s", err)) // TODO: continue if viper.ConfigFileNotFoundError == true, delete bool above then
			}
		}
	}
	if !viper.IsSet("database.dsn") {
		if viper.Get("database.dialect") == "postgres" {
			viper.SetDefault("database.postgres.sslmode", "disable")
			viper.Set("database.dsn",
				fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=%s&application_name=OHMAB",
					viper.Get("database.dbuser"),
					viper.Get("database.dbpassword"),
					viper.Get("database.dbhost"),
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
					viper.Get("database.sqlite3.mode"),
					viper.Get("database.sqlite3.cache"),
				))
		} else {
			panic("database.dialect not specified") // perhaps config file is missing
		}
	}
}
