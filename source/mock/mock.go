package mock

import (
	"fmt"
	"io"

	"github.com/popodidi/conf/source"
)

// New returns a mock source with m.
func New(m map[string]string) source.Interface {
	return &mock{m}
}

type mock struct {
	m map[string]string
}

func (s *mock) Read(key string) (value string, exists bool) {
	value, exists = s.m[key]
	return
}

func (s *mock) Export(m map[string]interface{}, writer io.Writer) error {
	_, err := writer.Write([]byte(fmt.Sprintf("%v", m)))
	return err
}
