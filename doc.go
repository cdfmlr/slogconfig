// Package slogconfig provides common configuration items to the std slog logger.
//
// It encapsulates methods to create NewHandler and NewLogger based on the config.
// Call OverrideSlogDefault method to change the slog.Default() logger.
// based on the config values.
//
// Works with github.com/cdfmlr/configer:
//
//	type appConfig struct {
//	    SlogConfig slogconfig.SlogConfig
//	    // other fields...
//	}
//
//	func main() {
//	    var cfg appConfig
//	    configer.New(&cfg, configer.TOML).ReadFromFile("config.toml")
//
//	    cfg.SlogConfig.OverrideSlogDefault()
//	    // other setup...
//	}
//
// Example config.toml:
//
//	[Log]
//	Level = "info"    # one of "debug", "info" (default), "warn" or "error"
//	Format = "json"   # one of "text" or "json" (default)
//	Output = "stdout" # one of "stderr", "stdout" (default) or "path/to/customFile.log"
package slogconfig
