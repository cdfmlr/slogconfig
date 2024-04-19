package logconfig

// LogConfig contains common configuration items to the std slog logger of your app.
//
// It encapsulates methods to create NewHandler and NewLogger based on the config.
// It also supports directly SetupSlogDefaultLogger.
type LogConfig struct {
	Level  LogLevel  // one of "DEBUG", "INFO" (default), "WARN" or "ERROR"
	Format LogFormat // one of "text" (default) or "json"
	Output LogOutput // one of "stderr", "stdout" (default) or "path/to/customFile.log"
}

func (c LogConfig) LevelOrDefault() LogLevel {
	if c.Level == "" {
		return LogLevelInfo
	}
	return c.Level
}

func (c LogConfig) FormatOrDefault() LogFormat {
	if c.Format == "" {
		return LogFormatText
	}
	return c.Format
}

func (c LogConfig) OutputOrDefault() LogOutput {
	if c.Output == "" {
		return LogOutputStdout
	}
	return c.Output
}
