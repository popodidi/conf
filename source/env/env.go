package env

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/popodidi/conf"
)

const sep = "_"

// New returns a env source.
func New() conf.Source {
	return &env{}
}

type env struct{}

func (e *env) Read(key string, path ...string) (value string, exists bool) {
	varName := e.varName(key, path...)
	return os.LookupEnv(varName)
}

func (e *env) Export(m conf.Map, writer io.Writer) error {
	var err error
	m.Iter(func(key string, val interface{}, path ...string) bool {
		name := e.varName(key, path...)
		_, err = writer.Write([]byte(fmt.Sprintf("%s=%v\n", name, val)))
		return err == nil
	})
	return err
}

func (e *env) varName(key string, path ...string) string {
	name := strings.Join(path, sep)
	name = strings.ToUpper(name)
	return name + sep + strings.ToUpper(key)
}
