package slogconfig

import (
	"io"
	"log/slog"
)

// SlogConfig contains common configuration items to the std slog logger of your app.
//
// It encapsulates methods to create NewHandler and NewLogger based on the config.
// It also supports directly SetSlogDefault.
type SlogConfig struct {
	Level  LogLevel  // one of "debug", "info" (default), "warn" or "error"
	Format LogFormat // one of "text" or "json" (default), or using RegisterLogFormat to custom others.
	Output LogOutput // one of "stderr", "stdout" (default) or "path/to/customFile.log"
}

// LogFormat represents a kind of slog.Handler.
//
// A LogFormat is bind to a NewHandlerFunc.
type LogFormat = string

const (
	LogFormatText LogFormat = "text"
	LogFormatJson LogFormat = "json"
)

// NewHandlerFunc builds out a slog.Handler.
type NewHandlerFunc func(w io.Writer, opts *slog.HandlerOptions) slog.Handler

func newTextHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	return slog.NewTextHandler(w, opts)
}

func newJsonHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	return slog.NewJSONHandler(w, opts)
}

var logFormatsRegister = map[string]NewHandlerFunc{
	LogFormatText: newTextHandler,
	LogFormatJson: newJsonHandler,
}

// RegisterLogFormat supports custom LogFormat.
//
// Param name is what you call this LogFormat in config files.
// Param builder is a NewHandlerFunc that makes slog.Handler objects with the custom format.
func RegisterLogFormat(name string, builder NewHandlerFunc) {
	logFormatsRegister[name] = builder
}

type LogLevel = string

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
)

type LogOutput = string

const (
	LogOutputStderr LogOutput = "stderr"
	LogOutputStdout LogOutput = "stdout"
)

func (c SlogConfig) EffectiveLevel() LogLevel {
	if c.Level == "" {
		return Default.Level
	}
	return c.Level
}

func (c SlogConfig) EffectiveFormat() LogFormat {
	if c.Format == "" {
		return Default.Format
	}
	return c.Format
}

func (c SlogConfig) EffectiveOutput() LogOutput {
	if c.Output == "" {
		return Default.Output
	}
	return c.Output
}
