package flag

import (
	"flag"
	"os"
)

// NewSys returns a flag.System with go builtin flag package.
func NewSys() System {
	return &goSys{
		fs: flag.NewFlagSet("conf", flag.ExitOnError),
	}
}

type goSys struct {
	fs *flag.FlagSet
}

func (f *goSys) Parse(args []string) error {
	return f.fs.Parse(args)
}

func (f *goSys) DefaultArgs() []string {
	return os.Args[1:]
}

func (f *goSys) IsSet(name string) (found bool) {
	f.fs.Visit(func(flg *flag.Flag) {
		if found {
			return
		}
		if flg.Name == name {
			found = true
			return
		}
	})
	return
}

func (f *goSys) Read(name string) string {
	flg := f.fs.Lookup(name)
	if flg == nil {
		return ""
	}
	return flg.Value.String()
}

func (f *goSys) Int(name string, value int, usage string) Flag {
	_ = f.fs.Int(name, value, usage)
	return &goFlag{flag: f.fs.Lookup(name)}
}

func (f *goSys) Int64(name string, value int64, usage string) Flag {
	_ = f.fs.Int64(name, value, usage)
	return &goFlag{flag: f.fs.Lookup(name)}
}

func (f *goSys) Uint(name string, value uint, usage string) Flag {
	_ = f.fs.Uint(name, value, usage)
	return &goFlag{flag: f.fs.Lookup(name)}
}

func (f *goSys) Uint64(name string, value uint64, usage string) Flag {
	_ = f.fs.Uint64(name, value, usage)
	return &goFlag{flag: f.fs.Lookup(name)}
}

func (f *goSys) Float64(name string, value float64, usage string) Flag {
	_ = f.fs.Float64(name, value, usage)
	return &goFlag{flag: f.fs.Lookup(name)}
}

func (f *goSys) Bool(name string, value bool, usage string) Flag {
	_ = f.fs.Bool(name, value, usage)
	return &goFlag{flag: f.fs.Lookup(name)}
}

func (f *goSys) String(name, value, usage string) Flag {
	_ = f.fs.String(name, value, usage)
	return &goFlag{flag: f.fs.Lookup(name)}
}

type goFlag struct {
	flag *flag.Flag
}

func (f *goFlag) Name() string {
	return f.flag.Name
}
func (f *goFlag) Usage() string {
	return f.flag.Usage
}

func (f *goFlag) DefaultValue() string {
	return f.flag.DefValue
}
