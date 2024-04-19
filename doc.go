// Package slogconfig provides common configuration items to the std slog logger.
//
// It encapsulates methods to create NewHandler and NewLogger based on the config.
// It also supports directly SetupSlogDefaultLogger.
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
//	    cfg.SlogConfig.SetupSlogDefaultLogger()
//	    // other setup...
//	}
package slogconfig
