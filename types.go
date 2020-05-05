package conf

import (
	"io"
	"reflect"
)

// Source defines the source interface.
type Source interface {
	Reader
	Exporter
}

// Reader reads value with key.
type Reader interface {
	Read(key string, path ...string) (value string, exists bool)
}

// Exporter exporters loaded config map to writer.
type Exporter interface {
	Export(m Map, writer io.Writer) error
}

// Scanner defines the Scanner interface for custom scannable types.
type Scanner interface {
	Scan(str string) error
}

// Configurable defines the source/reader/exported that should be configured
// beforehand, e.g. flag.
type Configurable interface {
	Configure(
		t reflect.Type, tag FieldTag, key string, path ...string) error
}

// ConfigurableSource defines the configurable source.
type ConfigurableSource interface {
	Configurable
	Source
}

// ConfigurableReader defines the configurable reader.
type ConfigurableReader interface {
	Configurable
	Reader
}

// ConfigurableExporter defines the configurable exporter.
type ConfigurableExporter interface {
	Configurable
	Exporter
}
