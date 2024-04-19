package slogconfig

// SlogConfig contains common configuration items to the std slog logger of your app.
//
// It encapsulates methods to create NewHandler and NewLogger based on the config.
// It also supports directly SetupSlogDefaultLogger.
type SlogConfig struct {
	Level  LogLevel  // one of "DEBUG", "INFO" (default), "WARN" or "ERROR"
	Format LogFormat // one of "text" (default) or "json"
	Output LogOutput // one of "stderr", "stdout" (default) or "path/to/customFile.log"
}

func (c SlogConfig) LevelOrDefault() LogLevel {
	if c.Level == "" {
		return LogLevelInfo
	}
	return c.Level
}

func (c SlogConfig) FormatOrDefault() LogFormat {
	if c.Format == "" {
		return LogFormatText
	}
	return c.Format
}

func (c SlogConfig) OutputOrDefault() LogOutput {
	if c.Output == "" {
		return LogOutputStdout
	}
	return c.Output
}
