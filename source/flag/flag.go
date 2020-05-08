package flag

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/popodidi/conf"
)

const sep = "-"

// New returns a flag reader with go builtin flag package.
func New() *Reader {
	return NewReader(NewSys())
}

// NewReader returns a flag reader with custom flag system.
func NewReader(sys System) *Reader {
	return &Reader{
		sys:   sys,
		store: make(conf.Map),
	}
}

// Reader defines the flag configurable reader.
type Reader struct {
	once  sync.Once
	sys   System
	store conf.Map
}

// Configure configures the flag with r.sys.
func (r *Reader) Configure(
	t reflect.Type, tag conf.FieldTag, key string, path ...string) error {
	if r.store.In(path...).Get(key) != nil {
		return nil
	}
	return r.store.MustIn(path...).Set(
		key, r.makeFlag(t, tag, r.flagName(key, path...)))
}

func (r *Reader) makeFlag(
	t reflect.Type, tag conf.FieldTag, name string) Flag {
	var (
		defaultValue interface{}
	)
	if tag.Default != nil {
		val, err := conf.ScanValue(t, *tag.Default)
		if err != nil {
			return r.sys.String(name, *tag.Default, tag.Usage)
		}
		if val.CanInterface() {
			defaultValue = val.Interface()
		}
	}
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		if defaultValue != nil {
			return r.sys.Int(name, defaultValue.(int), tag.Usage)
		}
		return r.sys.Int(name, 0, tag.Usage)
	case reflect.Int64:
		if defaultValue != nil {
			return r.sys.Int64(name, defaultValue.(int64), tag.Usage)
		}
		return r.sys.Int64(name, 0, tag.Usage)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		if defaultValue != nil {
			return r.sys.Uint(name, defaultValue.(uint), tag.Usage)
		}
		return r.sys.Uint(name, 0, tag.Usage)
	case reflect.Uint64:
		if defaultValue != nil {
			return r.sys.Uint64(name, defaultValue.(uint64), tag.Usage)
		}
		return r.sys.Uint64(name, 0, tag.Usage)
	case reflect.Float32, reflect.Float64:
		if defaultValue != nil {
			return r.sys.Float64(name, defaultValue.(float64), tag.Usage)
		}
		return r.sys.Float64(name, 0, tag.Usage)
	case reflect.Bool:
		if defaultValue != nil {
			return r.sys.Bool(name, defaultValue.(bool), tag.Usage)
		}
		return r.sys.Bool(name, false, tag.Usage)
	case reflect.String:
		if defaultValue != nil {
			return r.sys.String(name, defaultValue.(string), tag.Usage)
		}
		return r.sys.String(name, "", tag.Usage)
	default:
		if tag.Default != nil {
			return r.sys.String(name, *tag.Default, tag.Usage)
		}
		return r.sys.String(name, "", tag.Usage)
	}
}

// Parse parses the arguments with r.sys. Parse should be executed after
// configuration since it's only executed once. If not executed before Read,
// it will be executed with r.sys.DefaultArgs().
func (r *Reader) Parse(args []string) error {
	var err error
	r.once.Do(func() {
		err = r.sys.Parse(args)
	})
	return err
}

// Usage returns usage string.
func (r *Reader) Usage() string {
	var str strings.Builder
	r.store.Iter(
		func(key string, val interface{}, path ...string) (next bool) {
			flg := val.(Flag)
			_, _ = str.WriteString(fmt.Sprintf("  --%-20s", flg.Name()))
			if flg.Usage() != "" {
				_, _ = str.WriteString(fmt.Sprintf(" %s", flg.Usage()))
			}
			_, _ = str.WriteString("\n")
			return true
		},
	)
	return str.String()
}

// Read reads the value and parses r.sys.DefaultArgs() if never parsed.
func (r *Reader) Read(key string, path ...string) (value string, exists bool) {
	err := r.Parse(r.sys.DefaultArgs())
	if err != nil {
		return
	}
	name := r.flagName(key, path...)
	if !r.sys.IsSet(name) {
		return
	}
	value = r.sys.Read(name)
	exists = true
	return
}

func (r *Reader) flagName(key string, path ...string) string {
	key = strings.ToLower(key)
	if len(path) == 0 {
		return key
	}
	name := strings.Join(path, sep)
	name = strings.ToLower(name)
	return name + sep + key
}
