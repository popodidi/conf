package source

import "io"

// Interface defines the source interface.
type Interface interface {
	Reader
	Exporter
}

// Reader reads value with key.
type Reader interface {
	Read(key string) (value string, exists bool)
}

// Exporter exporters loaded config map to writer.
type Exporter interface {
	Export(m map[string]interface{}, writer io.Writer) error
}
