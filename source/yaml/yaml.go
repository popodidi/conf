package yaml

import (
	"fmt"
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/popodidi/conf"
)

// New returns a Yaml file source.
func New(path string) conf.Source {
	return &yml{path: path}
}

// NewExporter returns an yaml exporter.
func NewExporter() conf.Exporter {
	return &exporter{}
}

type yml struct {
	exporter
	path string
}

func (y *yml) Read(key string, path ...string) (value string, exists bool) {
	b, err := ioutil.ReadFile(y.path)
	if err != nil {
		return
	}
	var m map[string]interface{}
	err = yaml.Unmarshal(b, &m)
	if err != nil {
		return
	}
	return y.read(m, key, path...)
}

func (y *yml) read(m map[string]interface{}, key string, path ...string) (
	value string, exists bool) {
	if len(path) == 0 {
		var val interface{}
		val, exists = m[key]
		if !exists {
			return
		}
		value = fmt.Sprintf("%v", val)
		return
	}
	sub, ok := m[path[0]]
	if !ok {
		return
	}
	subM, ok := sub.(map[string]interface{})
	if !ok {
		return
	}
	return y.read(subM, key, path[1:]...)
}

type exporter struct{}

func (e *exporter) Export(m conf.Map, writer io.Writer) error {
	b, err := yaml.Marshal(m)
	if err != nil {
		return err
	}
	_, err = writer.Write(b)
	return err
}
