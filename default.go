package slogconfig

var Default = SlogConfig{
	Level:     LevelInfo,
	Format:    FormatJson,
	Output:    OutputStdout,
	AddSource: false,
}
