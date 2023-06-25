package log

import (
	"github.com/rs/zerolog"
	"os"
	"runtime/debug"
	"strconv"
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
	// read debug environment variable
	debug_, err := strconv.ParseInt(os.Getenv("DEBUG"), 10, 64)
	if err == nil && debug_ == 1 {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else if debug_ == 2 {
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	buildInfo, _ := debug.ReadBuildInfo()
	hostname, _ := os.Hostname()
	var logger1 = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		With().
		Timestamp().
		Str("host", hostname).
		Caller().
		//Int("pid", os.Getpid()).
		//Str("go_version", buildInfo.GoVersion).
		Logger()

	logger1.Debug().Msgf("Logger initialized, GO Version: %s", buildInfo.GoVersion)

	return &logger1
}
