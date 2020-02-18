package conf

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
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
	return c.iterFields(
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
			switch field.Kind() {
			case reflect.String:
				finalVal = reflect.ValueOf(value)
			case reflect.Int:
				var i int64
				i, err = strconv.ParseInt(value, 0, 0)
				if err != nil {
					err = fmt.Errorf("failed to parse int for \"%s\" from \"%s\". %v", key, value, ErrInvalidValue)
					return
				}
				finalVal = reflect.ValueOf(int(i))
			case reflect.Float64:
				var f float64
				f, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("failed to parse float for \"%s\" from \"%s\". %v", key, value, ErrInvalidValue)
					return
				}
				finalVal = reflect.ValueOf(float64(f))
			case reflect.Bool:
				var b bool
				b, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("failed to parse bool for \"%s\" from \"%s\". %w", key, value, ErrInvalidValue)
					return
				}
				finalVal = reflect.ValueOf(b)
			default:
				err = ErrUnsupportedType
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
	err := c.iterFields(
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

func (c *Config) read(key string, readers []source.Reader) (val string, exists bool) {
	for _, reader := range readers {
		val, exists = reader.Read(key)
		if exists {
			return
		}
	}
	return
}

func (c *Config) iterFields(
	fn func(field reflect.Value, typeField reflect.StructField, key string) error,
) (err error) {
	return c.iterFieldsRecursively(reflect.ValueOf(c.ptr), "", make(map[string]struct{}), fn)
}

func (c *Config) iterFieldsRecursively(ptrValue reflect.Value, prefix string, loadedKeys map[string]struct{},
	fn func(field reflect.Value, typeField reflect.StructField, key string) error) (err error) {
	ptrType := ptrValue.Type()
	if ptrType == nil {
		return ErrNilConfig
	}

	configKind := ptrValue.Kind()
	if configKind != reflect.Ptr {
		return ErrConfigNotPtr
	}

	configValue := ptrValue.Elem()
	configType := configValue.Type()

	if configValue.Kind() != reflect.Struct {
		return ErrConfigNotStruct
	}

	for i := 0; i < configType.NumField(); i++ {
		v := configValue.Field(i)
		v.CanAddr()
		t := configType.Field(i)
		key := t.Name
		if prefix != "" {
			key = fmt.Sprintf("%s_%s", prefix, key)
		}
		key = strings.ToUpper(key)

		// recursive call for struct fields
		if v.Kind() == reflect.Struct {
			if !v.CanAddr() {
				return ErrCantAddr
			}
			err = c.iterFieldsRecursively(v.Addr(), key, loadedKeys, fn)
			if err != nil {
				return err
			}
			continue
		}

		if _, duplicate := loadedKeys[key]; duplicate {
			return fmt.Errorf("key=%s. %w", key, ErrDuplicateKey)
		}
		loadedKeys[key] = struct{}{}
		err = fn(v, t, key)
		if err != nil {
			return err
		}
	}
	return nil
}
