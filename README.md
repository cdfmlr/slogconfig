# slogconfig

[![Go Reference](https://pkg.go.dev/badge/github.com/cdfmlr/slogconfig.svg)](https://pkg.go.dev/github.com/cdfmlr/slogconfig)

Package slogconfig provides common configuration items for "log/slog".

## install

```sh
go get github.com/cdfmlr/slogconfig
```

## usage

Works with [`github.com/cdfmlr/configer`](https://github.com/cdfmlr/configer):

```go
package main

import (
	"github.com/cdfmlr/configer"
	"github.com/cdfmlr/slogconfig"
)

type appConfig struct {
	SlogConfig slogconfig.SlogConfig
	// other fields...
}

func main() {
	var cfg appConfig
	configer.New(&cfg, configer.TOML).ReadFromFile("config.toml")

	cfg.SlogConfig.OverrideSlogDefault()
	// now slog.Default() is set to a new Logger based on the config.
}
```

example `config.toml`:

```toml
[Log]
Level = "info"    # one of "debug", "info" (default), "warn" or "error"
Format = "json"   # one of "text" or "json" (default)
Output = "stdout" # one of "stderr", "stdout" (default) or "path/to/customFile.log"
```
