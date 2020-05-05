package flag

import (
	"flag"
	"os"
	"reflect"
	"strings"

	"github.com/popodidi/conf"
)

func newGoFlagSys() *goflag {
	return &goflag{
		fs: flag.NewFlagSet("conf", flag.ExitOnError),
	}
}

type goflag struct {
	fs *flag.FlagSet
}

func (f *goflag) Flag(
	t reflect.Type, tag conf.FieldTag, name string) (
	ptr interface{}) {
	var (
		defaultValue interface{}
	)
	if tag.Default != nil {
		val, err := conf.ScanValue(t, *tag.Default)
		if err != nil {
			return f.fs.String(name, *tag.Default, tag.Usage)
		}
		if val.CanInterface() {
			defaultValue = val.Interface()
		}
	}
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		if defaultValue != nil {
			return f.fs.Int(name, defaultValue.(int), tag.Usage)
		}
		return f.fs.Int(name, 0, tag.Usage)
	case reflect.Int64:
		if defaultValue != nil {
			return f.fs.Int64(name, defaultValue.(int64), tag.Usage)
		}
		return f.fs.Int64(name, 0, tag.Usage)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		if defaultValue != nil {
			return f.fs.Uint(name, defaultValue.(uint), tag.Usage)
		}
		return f.fs.Uint(name, 0, tag.Usage)
	case reflect.Uint64:
		if defaultValue != nil {
			return f.fs.Uint64(name, defaultValue.(uint64), tag.Usage)
		}
		return f.fs.Uint64(name, 0, tag.Usage)
	case reflect.Float32, reflect.Float64:
		if defaultValue != nil {
			return f.fs.Float64(name, defaultValue.(float64), tag.Usage)
		}
		return f.fs.Float64(name, 0, tag.Usage)
	case reflect.Bool:
		if defaultValue != nil {
			return f.fs.Bool(name, defaultValue.(bool), tag.Usage)
		}
		return f.fs.Bool(name, false, tag.Usage)
	case reflect.String:
		if defaultValue != nil {
			return f.fs.String(name, defaultValue.(string), tag.Usage)
		}
		return f.fs.String(name, "", tag.Usage)
	default:
		if tag.Default != nil {
			return f.fs.String(name, *tag.Default, tag.Usage)
		}
		return f.fs.String(name, "", tag.Usage)
	}
}

func (f *goflag) Parse(args []string) error {
	return f.fs.Parse(args)
}

func (f *goflag) DefaultArgs() []string {
	return os.Args[1:]
}

func (f *goflag) IsFlagSet(name string) (found bool) {
	f.fs.Visit(func(f *flag.Flag) {
		if found {
			return
		}
		if f.Name == name {
			found = true
			return
		}
	})
	return
}

func (f *goflag) Usage() string {
	var str strings.Builder
	f.fs.SetOutput(&str)
	f.fs.PrintDefaults()
	return str.String()
}
