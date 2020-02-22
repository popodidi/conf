# Conf

`conf` is a configuration parser that magically parses configurations into a
golang struct.

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
	err := conf.Load(&cfg, env.New(), yaml.New("basic.yaml"))
	if err != nil {
		log.Fatal(err)
	}

	// use values
	log.Println("cfg.Str", cfg.Str)
	log.Println("cfg.Int", cfg.Int)
	log.Println("cfg.Float", cfg.Float)
}
```

### Sources

The `Load` functions accept multiple sources. It falls back to the next one if
the value is not found.

### Config Struct

`conf` uses reflection to get config keys, which are identical in different
sources. Keys will all be UPPERCASE. Nested struct is also supported with a
separator `_`.

```go
type config struct {
	Yo  bool
	Hey struct {
		Hi int
	}
}
```

The corresponding configuration file will then be

```yaml
YO: true
HEY_HI: 1
```

### Default value

A configuration with the `default` tag is optional.

```go
type config struct {
	Name string `conf:"default:default_name"`
}
```

## Roadmap

The `conf` package is under actively development.

Source

- [x] Environment Variable
- [x] Yaml
- [ ] JSON
- [ ] Command line
  - [ ] Abstract Flag Parser
  - [ ] Prompt
