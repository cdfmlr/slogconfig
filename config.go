package slogconfig

import (
	"io"
	"log/slog"
)

// SlogConfig contains common configuration items to the std slog logger of your app.
//
// It encapsulates methods to create NewHandler and NewLogger based on the config.
// It also supports directly OverrideSlogDefault.
type SlogConfig struct {
	Level     Level  // one of "debug", "info" (default), "warn" or "error"
	Format    Format // one of "text" or "json" (default), or using RegisterLogFormat to custom others.
	Output    Output // one of "stderr", "stdout" (default) or "path/to/customFile.log"
	AddSource bool   // whether to add source file and line number to log entries. Default false.
}

// Format represents a kind of slog.Handler.
//
// A LogFormat is bind to a NewHandlerFunc.
type Format = string

const (
	FormatText Format = "text"
	FormatJson Format = "json"
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
	FormatText: newTextHandler,
	FormatJson: newJsonHandler,
}

// RegisterLogFormat supports custom LogFormat.
//
// Param name is what you call this LogFormat in config files.
// Param builder is a NewHandlerFunc that makes slog.Handler objects with the custom format.
func RegisterLogFormat(name string, builder NewHandlerFunc) {
	logFormatsRegister[name] = builder
}

type Level = string

const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
)

type Output = string

const (
	OutputStderr Output = "stderr"
	OutputStdout Output = "stdout"
)

func (c SlogConfig) ValidLevel() Level {
	if c.Level == "" {
		return Default.Level
	}
	return c.Level
}

func (c SlogConfig) ValidFormat() Format {
	if c.Format == "" {
		return Default.Format
	}
	return c.Format
}

func (c SlogConfig) ValidOutput() Output {
	if c.Output == "" {
		return Default.Output
	}
	return c.Output
}
