package log

import (
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var requestLogger *zerolog.Logger
var once2 sync.Once

func GetRequestLoggerInstance() *zerolog.Logger {
	once2.Do(func() {
		requestLogger = createRequestLogger()
	})
	return requestLogger
}
func RequestLogger(next http.Handler) http.Handler {
	logger := GetRequestLoggerInstance()
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		time.Since(t1)
		rqid := r.Context().Value(middleware.RequestIDKey)
		if rqid == nil || rqid == "" {
			rqid = "unknown"
		}
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		defer func() {
			logger.Debug().
				Str(middleware.RequestIDHeader, fmt.Sprint(rqid)).
				Str("method", r.Method).
				Int("status", ww.Status()).
				Str("path", r.URL.Path).
				Str("remote_addr", r.RemoteAddr).
				Str("user_agent", r.UserAgent()).
				Int("length", ww.BytesWritten()).
				TimeDiff("duration", time.Now(), t1).
				Msg("request")
		}()
		w.Header().Add(middleware.RequestIDHeader, fmt.Sprint(rqid))
		next.ServeHTTP(ww, r)
	}

	return http.HandlerFunc(fn)
}
func createRequestLogger() *zerolog.Logger {
	// read debug enviroment variable
	debug_, err := strconv.ParseInt(os.Getenv("DEBUG"), 10, 64)
	if err == nil && debug_ == 1 {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else if debug_ == 2 {
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	hostname, _ := os.Hostname()
	var logger2 = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		With().
		Str("host", hostname).
		Timestamp().
		Logger()

	return &logger2
}
