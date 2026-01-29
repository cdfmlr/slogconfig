package slogconfig

import (
	"fmt"
	"log"
	"log/slog"
	"os"
)

// NewHandler creates a new slog.Handler based on the configuration.
func (c SlogConfig) NewHandler() (slog.Handler, error) {
	handler := slog.Default().Handler()

	level := new(slog.LevelVar)

	if err := level.UnmarshalText([]byte(
		c.EffectiveLevel()),
	); err != nil {
		return handler, fmt.Errorf("failed to unmarshal log level: %w", err)
	}

	out := os.Stdout

	switch c.EffectiveOutput() {
	case LogOutputStderr:
		out = os.Stderr
	case LogOutputStdout:
		out = os.Stdout
	default:
		log.Printf("output log to: %s", c.Output)
		f, err := os.OpenFile(c.EffectiveOutput(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return handler, fmt.Errorf("failed to open log file: %w", err)
		}
		out = f
	}

	format := c.EffectiveFormat()
	if newHandler, exists := logFormatsRegister[format]; !exists {
		return handler, fmt.Errorf("unknown log format: %s", c.Format)
	} else {
		handler = newHandler(out, &slog.HandlerOptions{Level: level})
	}

	return handler, nil
}

// NewLogger creates a new slog.Logger based on the configuration.
//
// Shorthand for:
//
//	handler, _ := c.NewHandler()
//	logger := slog.New(handler)
func (c SlogConfig) NewLogger() (*slog.Logger, error) {
	handler, err := c.NewHandler()
	if err != nil {
		return nil, fmt.Errorf("failed to create log handler: %w", err)
	}

	logger := slog.New(handler)
	return logger, nil
}

// SetSlogDefault sets the default logger of the slog package to the
// logger created based on the configuration.
//
// Shorthand for:
//
//	logger, _ := c.NewLogger()
//	slog.SetDefault(logger)
func (c SlogConfig) SetSlogDefault() error {
	logger, err := c.NewLogger()
	if err != nil {
		return fmt.Errorf("failed to create logger: %w", err)
	}
	slog.SetDefault(logger)
	return nil
}
