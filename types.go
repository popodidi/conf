package conf

import (
	"io"
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
