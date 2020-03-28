package memory

import (
	"fmt"
	"io"
	"strings"

	"github.com/popodidi/conf"
)

// New returns a memory source.
func New(m conf.Map) conf.Source {
	return source(m)
}

type source conf.Map

func (s source) Read(key string, path ...string) (value string, exists bool) {
	val := conf.Map(s).In(path...).Get(key)
	if val == nil {
		return
	}
	value = fmt.Sprintf("%v", val)
	exists = true
	return
}

func (s source) Export(m conf.Map, writer io.Writer) error {
	var err error
	m.Iter(func(key string, val interface{}, path ...string) (next bool) {
		if len(path) == 0 {
			_, err = writer.Write([]byte(fmt.Sprintf("%s: %v", key, val)))
		} else {
			_, err = writer.Write([]byte(
				fmt.Sprintf("%s.%s: %v", strings.Join(path, "."), key, val)))
		}
		return err == nil
	})
	return err
}
