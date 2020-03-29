package conf

import (
	"fmt"
	"io"
	"reflect"
	"strings"
)

// Config defines the config struct.
type Config struct {
	ptr interface{}
	m   Map
}

// New return a config instance.
func New(ptr interface{}) *Config {
	return &Config{
		ptr: ptr,
	}
}

// Load loads the config from flattened readers.
func (c *Config) Load(readers ...Reader) error {
	c.m = make(Map)
	return iterFields(reflect.ValueOf(c.ptr), nil,
		func(
			field reflect.Value, typeField reflect.StructField,
			key string, path ...string) (
			err error,
		) {
			// cache loaded value in the map
			var finalVal reflect.Value
			defer func() {
				if err != nil {
					return
				}
				err = c.m.MustIn(path...).Set(key, finalVal.Interface())
			}()

			tag, err := parseTag(typeField.Tag.Get("conf"))
			if err != nil {
				return
			}

			value, exists := c.read(readers, key, path...)

			// return error for required config not found
			if !exists && !tag.hasDefault {
				err = fmt.Errorf("key=%s. %w", key, ErrConfigNotFound)
				return
			}

			// set default value for not found config
			if !exists && tag.hasDefault {
				value = tag.defaultValue
			}

			// parse value
			finalVal, err = ParseValue(field.Kind(), value)
			if err != nil {
				return
			}

			// set value to ptr
			field.Set(finalVal)
			err = nil
			return
		},
	)
}

// Map returns a map of the loaded config.
func (c *Config) Map() (Map, error) {
	if c.m == nil {
		return nil, ErrConfigNotLoaded
	}
	return c.m.Clone()
}

// Export exports the loaded config c to writer with exporter.
func (c *Config) Export(exporter Exporter, writer io.Writer) error {
	m, err := c.Map()
	if err != nil {
		return err
	}
	return exporter.Export(m, writer)
}

//Template returns a template string with exporter format and nested map.
func (c *Config) Template(exporter Exporter) (string, error) {
	m := make(Map)
	err := iterFields(reflect.ValueOf(c.ptr), nil, mapRecorder(m))
	if err != nil {
		return "", err
	}
	var b strings.Builder
	err = exporter.Export(m, &b)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func (c *Config) read(readers []Reader, key string, path ...string) (
	val string, exists bool) {
	for _, reader := range readers {
		val, exists = reader.Read(key, path...)
		if exists {
			return
		}
	}
	return
}
