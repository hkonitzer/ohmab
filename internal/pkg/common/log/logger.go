package log

import (
	"github.com/hkonitzer/ohmab/internal/pkg/common/config"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

var logger *zerolog.Logger
var once sync.Once

func GetLoggerInstance() *zerolog.Logger {
	once.Do(func() {
		logger = createLogger()
	})
	return logger

}

func createLogger() *zerolog.Logger {
	// Get the configuration
	config.Init()
	// read debug environment variable
	debug_ := viper.GetInt("DEBUG")
	if debug_ == 1 {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else if debug_ == 2 {
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	buildInfo, _ := debug.ReadBuildInfo()
	hostname, _ := os.Hostname()
	var logger1 zerolog.Logger
	if viper.GetString("ENVIRONMENT") == config.DevelopmentEnvironment {
		logger1 = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).
			With().
			Stack().
			Timestamp().
			Caller().
			Logger()
	} else {
		logger1 = zerolog.New(os.Stdout).
			With().
			Timestamp().
			Str("host", hostname).
			Caller().
			//Int("pid", os.Getpid()).
			//Str("go_version", buildInfo.GoVersion).
			Logger()
	}

	logger1.Debug().Msgf("Logger initialized, GO Version: %s", buildInfo.GoVersion)
	return &logger1
}
