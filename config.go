package slogconfig

// SlogConfig contains common configuration items to the std slog logger of your app.
//
// It encapsulates methods to create NewHandler and NewLogger based on the config.
// It also supports directly SetSlogDefault.
type SlogConfig struct {
	Level  LogLevel  // one of "debug", "info" (default), "warn" or "error"
	Format LogFormat // one of "text" or "json" (default)
	Output LogOutput // one of "stderr", "stdout" (default) or "path/to/customFile.log"
}

type LogFormat = string

const (
	LogFormatText LogFormat = "text"
	LogFormatJson LogFormat = "json"
)

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

func (c SlogConfig) LevelOrDefault() LogLevel {
	if c.Level == "" {
		return Default.Level
	}
	return c.Level
}

func (c SlogConfig) FormatOrDefault() LogFormat {
	if c.Format == "" {
		return Default.Format
	}
	return c.Format
}

func (c SlogConfig) OutputOrDefault() LogOutput {
	if c.Output == "" {
		return Default.Output
	}
	return c.Output
}
