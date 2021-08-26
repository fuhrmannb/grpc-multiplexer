package cmd

import (
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func LogLevelVar(cmd *cobra.Command, logLevelVar *string) {
	InitStringVar(cmd, logLevelVar, "log-level", "info",
		"Log level. Could be \"debug\", \"info\", \"warning\" or \"error\"")
}

func LogFormatVar(cmd *cobra.Command, logLevelVar *string) {
	InitStringVar(cmd, logLevelVar, "log-format", "pretty",
		"Log format. Could be \"json\" or \"pretty\"")
}

func logLevelFromVar(logLevel string) (zerolog.Level, error) {
	switch logLevel {
	case "debug":
		return zerolog.DebugLevel, nil
	case "info":
		return zerolog.InfoLevel, nil
	case "warning":
		return zerolog.WarnLevel, nil
	case "error":
		return zerolog.ErrorLevel, nil
	default:
		return zerolog.NoLevel, errors.New("wrong log level specified")
	}
}

func SetupLog(logLevel string, logFormat string) error {
	// Log format
	switch logFormat {
	case "json":
		// Keep default logger format already in JSON
	case "pretty":
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"})
	default:
		return errors.New("wrong log format specified")
	}

	// Log level
	level, err := logLevelFromVar(logLevel)
	if err != nil {
		return err
	}
	zerolog.SetGlobalLevel(level)

	log.Debug().Str("log_level", logLevel).Msg("log level set")
	log.Debug().Str("log_format", logFormat).Msg("log format set")

	return nil
}
