package slogconfig

var Default = SlogConfig{
	Level:  LogLevelInfo,
	Format: LogFormatJson,
	Output: LogOutputStdout,
}
