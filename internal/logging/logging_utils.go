package logging

import (
	"UrlShortener/internal/config"
	"log/slog"
	"os"
)

type LoggingConfig struct {
	Level string `env:"LOG_LEVEL" envDefault:"INFO"`
}

func Init() {
	conf := getLoggingConfig()

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: getLogLevelFromString(conf.Level),
	}))

	slog.SetDefault(logger)
}

func getLoggingConfig() LoggingConfig {
	loggingConfig := LoggingConfig{}
	err := config.LoadConfig(&loggingConfig)
	if err != nil {
		panic(err)
	}
	return loggingConfig
}

func getLogLevelFromString(level string) slog.Level {
	switch level {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
