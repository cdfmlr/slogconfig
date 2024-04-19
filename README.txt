package logconfig // import "github.com/cdfmlr/logconfig"

Package logconfig provides common configuration items to the std slog logger.

It encapsulates methods to create NewHandler and NewLogger based on the config.
It also supports directly SetupSlogDefaultLogger.

Works with github.com/cdfmlr/configer:

    type appConfig struct {
        LogConfig logconfig.LogConfig
        // other fields...
    }

    func main() {
        var cfg appConfig
        configer.New(&cfg, configer.TOML).ReadFromFile("config.toml")

        cfg.LogConfig.SetupSlogDefaultLogger()
        // other setup...
    }

TYPES

type LogConfig struct {
	Level  LogLevel  // one of "DEBUG", "INFO" (default), "WARN" or "ERROR"
	Format LogFormat // one of "text" (default) or "json"
	Output LogOutput // one of "stderr", "stdout" (default) or "path/to/customFile.log"
}
    LogConfig contains common configuration items to the std slog logger of your
    app.

    It encapsulates methods to create NewHandler and NewLogger based on the
    config. It also supports directly SetupSlogDefaultLogger.

func (c LogConfig) FormatOrDefault() LogFormat

func (c LogConfig) LevelOrDefault() LogLevel

func (c LogConfig) NewHandler() (slog.Handler, error)
    NewHandler creates a new slog.Handler based on the configuration.

func (c LogConfig) NewLogger() (*slog.Logger, error)
    NewLogger creates a new slog.Logger based on the configuration.

    Shorthand for:

        handler, _ := c.NewHandler()
        logger := slog.New(handler)

func (c LogConfig) OutputOrDefault() LogOutput

func (c LogConfig) SetupSlogDefaultLogger() error
    SetupSlogDefaultLogger sets the default logger of the slog package to the
    logger created based on the configuration.

    Shorthand for:

        logger, _ := c.NewLogger()
        slog.SetDefault(logger)

type LogFormat = string

const (
	LogFormatText LogFormat = "text"
	LogFormatJson LogFormat = "json"
)
type LogLevel = string

const (
	LogLevelDebug LogLevel = "DEBUG"
	LogLevelInfo  LogLevel = "INFO"
	LogLevelWarn  LogLevel = "WARN"
	LogLevelError LogLevel = "ERROR"
)
type LogOutput = string

const (
	LogOutputStderr LogOutput = "stderr"
	LogOutputStdout LogOutput = "stdout"
)
