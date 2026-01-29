package slogconfig

import (
	"fmt"
	"log"
	"log/slog"
	"os"
)

// Option defines a function type for customizing slog.HandlerOptions.
type Option func(slog.HandlerOptions) slog.HandlerOptions

// NewHandler creates a new slog.Handler based on the configuration.
//
// It accepts optional Option functions to override or extend the default slog.HandlerOptions.
func (c SlogConfig) NewHandler(opts ...Option) (slog.Handler, error) {
	handler := slog.Default().Handler()

	level := new(slog.LevelVar)

	err := level.UnmarshalText([]byte(c.ValidLevel()))
	if err != nil {
		return handler, fmt.Errorf("failed to unmarshal log level: %w", err)
	}

	out := os.Stdout

	switch c.ValidOutput() {
	case OutputStderr:
		out = os.Stderr
	case OutputStdout:
		out = os.Stdout
	default:
		log.Printf("output log to: %s", c.Output)
		f, err := os.OpenFile(c.ValidOutput(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return handler, fmt.Errorf("failed to open log file: %w", err)
		}
		// defer f.Close()
		out = f
	}

	opt := slog.HandlerOptions{
		Level:     level,
		AddSource: c.AddSource,
	}
	for _, o := range opts {
		opt = o(opt)
	}

	format := c.ValidFormat()
	newHandler, exists := logFormatsRegister[format]
	if !exists {
		return handler, fmt.Errorf("unknown log format: %s", c.Format)
	}

	handler = newHandler(out, &opt)

	return handler, nil
}

// NewLogger creates a new slog.Logger based on the configuration.
//
// Shorthand for:
//
//	handler, _ := c.NewHandler()
//	logger := slog.New(handler)
//
// It accepts optional Option functions to override or extend the default slog.HandlerOptions.
func (c SlogConfig) NewLogger(opts ...Option) (*slog.Logger, error) {
	handler, err := c.NewHandler(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create log handler: %w", err)
	}

	logger := slog.New(handler)
	return logger, nil
}

// OverrideSlogDefault sets the default logger of the slog package to the
// logger created based on the configuration.
//
// Shorthand for:
//
//	logger, _ := c.NewLogger()
//	slog.SetDefault(logger)
func (c SlogConfig) OverrideSlogDefault() error {
	logger, err := c.NewLogger()
	if err != nil {
		return fmt.Errorf("failed to create logger: %w", err)
	}
	slog.SetDefault(logger)
	return nil
}
