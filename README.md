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
	flattened := true // false
	if flattened {
		err = conf.Load(&cfg, env.New(), yaml.New("basic.yaml"))       // flattened version
	} else {
		err = conf.LoadNested(&cfg, env.New(), yaml.New("basic.yaml")) // nested version
	}
	if err != nil {
		log.Fatal(err)
	}

	// use values
	log.Println("cfg.Str", cfg.Str)
	log.Println("cfg.Int", cfg.Int)
	log.Println("cfg.Float", cfg.Float)
}
```

The `Load`/`LoadNested` functions accept multiple sources. It falls back to the
next one if the value is not found. `conf` uses reflection to get config keys.

## Sources

The `conf` package currently supports the following sources.

- `env` - `env` reads from environment variables and always flattens struct
  field names into `UPPERCASE_WITH_UNDERCASE` as the variable names.
- `yaml` - `yaml` read config from a YAML. `yaml` supports both flattened and
  nested structures as flatted or nested functions are called.

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

The corresponding configuration YAML file will then be

```yaml
# Flattened
YO: true
HEY_HI: 1

# Nested
Yo: true
Hey:
  Hi: 1
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
- `int`
- `float64`
- `bool`

## Flattened / Nested

### Flattened Functions

The flattened functions of `conf` flattens every nested struct and formats keys
into `UPPERCASE_WITH_UNDERSCORE` to support nested structs.

> The very first version of `conf` flattens every nested struct by default. It
> was in this way because at the very beginning, `conf` was to parse only from
> environment variables. As more sources are added, the implicit flattening is no
> longer a good implementation. As a result, the source interface was updated to
> better support nested structs. At the same time, we still keep the original
> functions, which flatten config keys implicitly, for backward compatibility.

```go
func Load(config interface{}, readers ...Reader) (err error)
func Template(config interface{}, exporter Exporter)

func (c *Config) Map() (Map, error)
func (c *Config) Load(readers []Reader) error
func (c *Config) Template(exporter Exporter) (string, error)
```

### Nested Functions

With the nested functions, the flattening will be left to the implementation of
sources that do not support nesting, such `env`. For nested sources, such as
`yaml`, it is recommended to use nested functions instead.

```go
func LoadNested(config interface{}, readers ...Reader) (err error)
func NestedTemplate(config interface{}, exporter Exporter)

func (c *Config) NestedMap() (Map, error)
func (c *Config) LoadNested(readers []Reader) error
func (c *Config) NestedTemplate(exporter Exporter) (string, error)
```

> Since the flatten-formatting is implemented by the `env` source, it will just
> work to have nesting and not nesting sources together,
> `conf.LoadNested(&cfg, env.New(), yaml.New("my_conf.yaml"))`.
