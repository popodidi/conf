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
func New() *Flag {
	return NewFlag(newGoFlagSys())
}

// NewFlag returns a flag reader with custom flag system.
func NewFlag(sys System) *Flag {
	return &Flag{
		sys:   sys,
		store: make(conf.Map),
	}
}

// System defines the flag parsing system.
type System interface {
	// Flag returns the ptr of the parsed value.
	Flag(t reflect.Type, tag conf.FieldTag, name string) (ptr interface{})

	// Parse parses arguments.
	Parse(args []string) error

	// DefaultArgs returns default arguments, e.g. os.Args[1:].
	DefaultArgs() []string

	// IsFlagSet returns if the flag was set.
	IsFlagSet(name string) (found bool)

	// Usage returns the usage string.
	Usage() string
}

// Flag defines the flag configurable reader.
type Flag struct {
	once  sync.Once
	sys   System
	store conf.Map
}

// Configure configures the flag with f.sys.
func (f *Flag) Configure(
	t reflect.Type, tag conf.FieldTag, key string, path ...string) error {
	if f.store.In(path...).Get(key) != nil {
		return nil
	}
	return f.store.MustIn(path...).Set(
		key, f.sys.Flag(t, tag, f.flagName(key, path...)))
}

// Parse parses the arguments with f.sys. Parse should be executed after
// configuration since it's only executed once. If not executed before Read,
// it will be executed with f.sys.DefaultArgs().
func (f *Flag) Parse(args []string) error {
	var err error
	f.once.Do(func() {
		err = f.sys.Parse(args)
	})
	return err
}

// Usage returns usage string.
func (f *Flag) Usage() string {
	return f.sys.Usage()
}

// Read reads the value and parses f.sys.DefaultArgs() if never parsed.
func (f *Flag) Read(key string, path ...string) (value string, exists bool) {
	err := f.Parse(f.sys.DefaultArgs())
	if err != nil {
		return
	}
	if !f.sys.IsFlagSet(f.flagName(key, path...)) {
		return
	}
	val := f.store.In(path...).Get(key)
	if val == nil {
		return
	}

	v := reflect.ValueOf(val)
	if v.Kind() != reflect.Ptr {
		return
	}
	if v.IsNil() {
		return
	}
	value = fmt.Sprintf("%v", v.Elem().Interface())
	exists = true
	return
}

func (f *Flag) flagName(key string, path ...string) string {
	key = strings.ToLower(key)
	if len(path) == 0 {
		return key
	}
	name := strings.Join(path, sep)
	name = strings.ToLower(name)
	return name + sep + key
}
