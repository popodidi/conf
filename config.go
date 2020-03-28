package conf

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/popodidi/conf/source"
)

// Config defines the config struct.
type Config struct {
	ptr interface{}
	m   map[string]interface{}
}

// New return a config instance.
func New(ptr interface{}) *Config {
	return &Config{
		ptr: ptr,
	}
}

// Load loads the config from readers.
func (c *Config) Load(readers []source.Reader) error {
	c.m = make(map[string]interface{})
	return iterFields(reflect.ValueOf(c.ptr), "",
		func(field reflect.Value, typeField reflect.StructField, key string) (
			err error) {
			// cache loaded value in the map
			var finalVal reflect.Value
			defer func() {
				if err != nil {
					return
				}
				c.m[key] = finalVal.Interface()
			}()

			tag, err := parseTag(typeField.Tag.Get("conf"))
			if err != nil {
				return
			}
			value, exists := c.read(key, readers)

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
			valuer, ok := valuers[field.Kind()]
			if !ok {
				err = ErrUnsupportedType
				return
			}

			finalVal, err = valuer(value)
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

// Map returns the loaded config c as golang map type.
func (c *Config) Map() (map[string]interface{}, error) {
	if c.m == nil {
		return nil, ErrConfigNotLoaded
	}
	m := make(map[string]interface{})
	for k, v := range c.m {
		m[k] = v
	}
	return m, nil
}

// Export exports the loaded config c to writer with exporter.
func (c *Config) Export(exporter source.Exporter, writer io.Writer) error {
	m, err := c.Map()
	if err != nil {
		return err
	}
	return exporter.Export(m, writer)
}

// Template returns a template string with exporter format.
func (c *Config) Template(exporter source.Exporter) (string, error) {
	m := make(map[string]interface{})
	err := iterFields(reflect.ValueOf(c.ptr), "",
		func(field reflect.Value, typeField reflect.StructField, key string) error {
			m[key] = typeField.Type.Name()
			return nil
		})
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

func (c *Config) read(key string, readers []source.Reader) (
	val string, exists bool) {
	for _, reader := range readers {
		val, exists = reader.Read(key)
		if exists {
			return
		}
	}
	return
}
