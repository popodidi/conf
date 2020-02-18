package env

import (
	"fmt"
	"io"
	"os"

	"github.com/popodidi/conf/source"
)

// New returns a env source.
func New() source.Interface {
	return &env{}
}

type env struct{}

func (e *env) Read(key string) (value string, exists bool) {
	return os.LookupEnv(key)
}

func (e *env) Export(m map[string]interface{}, writer io.Writer) error {
	for key, val := range m {
		_, err := writer.Write([]byte(fmt.Sprintf("%s=%v\n", key, val)))
		if err != nil {
			return err
		}
	}
	return nil
}
