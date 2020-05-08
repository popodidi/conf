package flag

// System defines the flag parsing system.
type System interface {
	FlagSet

	// Parse parses arguments.
	Parse(args []string) error

	// DefaultArgs returns default arguments, e.g. os.Args[1:].
	DefaultArgs() []string
}

// nolint: golint
// FlagSet defines the interface flag set to manage flags.
type FlagSet interface {
	IsSet(name string) bool
	Read(name string) string

	Int(name string, value int, usage string) Flag
	Int64(name string, value int64, usage string) Flag
	Uint(name string, value uint, usage string) Flag
	Uint64(name string, value uint64, usage string) Flag
	Float64(name string, value float64, usage string) Flag
	Bool(name string, value bool, usage string) Flag
	String(name, value, usage string) Flag
}

// Flag defines the flag interface.
type Flag interface {
	Name() string
	Usage() string
	DefaultValue() string
}
