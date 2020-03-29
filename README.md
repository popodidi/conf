# Conf

![Go Test](https://github.com/popodidi/conf/workflows/Go%20Test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/popodidi/conf)](https://goreportcard.com/report/github.com/popodidi/conf)
[![Documentation](https://godoc.org/github.com/popodidi/conf?status.svg)](http://godoc.org/github.com/popodidi/conf)

`conf` is a configuration parser that parses configurations into a golang struct.

## Usage

### Basic

```go
package main

import (
	"log"

	"github.com/popodidi/conf"
	"github.com/popodidi/conf/source/env"
	"github.com/popodidi/conf/source/yaml"
)

func main() {
	// define the config struct
	type config struct {
		Str   string `conf:"default:hello"`
		Int   int
		Float float64
	}

	// load from sources
	var cfg config
	var err error
	err = conf.Load(&cfg, env.New(), yaml.New("basic.yaml"))
	if err != nil {
		log.Fatal(err)
	}

	// use values
	log.Println("cfg.Str", cfg.Str)
	log.Println("cfg.Int", cfg.Int)
	log.Println("cfg.Float", cfg.Float)
}
```

The `Load` functions accept multiple sources. It falls back to the next one if
the value is not found. `conf` uses reflection to get config keys.

## Sources

The `conf` package currently supports the following sources.

- `env` - `env` reads from environment variables and always flattens struct
  field names into `UPPERCASE_WITH_UNDERCASE` as the variable names.
- `yaml` - `yaml` reads config from a YAML file.
- `json` - `json` reads config from a JSON file.

## Config Struct

`conf` uses reflection to get config keys. Nested structs are also supported.

```go
type config struct {
	Yo  bool
	Hey struct {
		Hi int
	}
}
```

### ENV

```
YO: true
HEY_HI: 1
```

### YAML

```yaml
Yo: true
Hey:
  Hi: 1
```

### JSON

```json
{
  "Yo": true,
  "Hey": {
    "Hi": 1
  }
}
```

### Default value

A configuration with the `default` tag is optional.

```go
type config struct {
	Name string `conf:"default:default_name"`
}
```

### Supported Types

The `conf` package currently supports the following field types.

- `string`
- `int(8/16/32/64)`
- `uint(8/16/32/64)`
- `float(32/64)64`
- `bool`
