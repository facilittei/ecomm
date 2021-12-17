package logging

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
)

// zeroLogger uses ZeroLog package for logging
type zeroLogger struct {
	logger zerolog.Logger
}

// NewZeroLogger creates an instance of zeroLogger
func NewZeroLogger() Logger {
	return &zeroLogger{
		logger: zerolog.New(os.Stdout).With().Timestamp().Logger(),
	}
}

// Fatal logs a msg with a fatal level
func (z *zeroLogger) Fatal(format string, content ...interface{}) {
	z.logger.Fatal().Msg(fmt.Sprintf(format, content...))
}

// Error logs a msg with an error level
func (z *zeroLogger) Error(format string, content ...interface{}) {
	z.logger.Error().Msg(fmt.Sprintf(format, content...))
}

// Warn logs a msg with a warning level
func (z *zeroLogger) Warn(format string, content ...interface{}) {
	z.logger.Warn().Msg(fmt.Sprintf(format, content...))
}

// Info logs a msg with an info level
func (z *zeroLogger) Info(format string, content ...interface{}) {
	z.logger.Info().Msg(fmt.Sprintf(format, content...))
}

// Debug logs a msg with a debug level
func (z *zeroLogger) Debug(format string, content ...interface{}) {
	z.logger.Debug().Msg(fmt.Sprintf(format, content...))
}
