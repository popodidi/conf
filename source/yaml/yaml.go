package yaml

import (
	"fmt"
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/popodidi/conf/source"
)

// New returns a Yaml file source.
func New(path string) source.Interface {
	return &yml{path: path}
}

// NewExporter returns an yaml exporter.
func NewExporter() source.Exporter {
	return &exporter{}
}

type yml struct {
	exporter
	path string
}

func (y *yml) Read(key string) (value string, exists bool) {
	b, err := ioutil.ReadFile(y.path)
	if err != nil {
		return
	}
	var m map[string]interface{}
	err = yaml.Unmarshal(b, &m)
	if err != nil {
		return
	}
	var val interface{}
	val, exists = m[key]
	if !exists {
		return
	}
	value = fmt.Sprintf("%v", val)
	return
}

type exporter struct{}

func (e *exporter) Export(m map[string]interface{}, writer io.Writer) error {
	b, err := yaml.Marshal(m)
	if err != nil {
		return err
	}
	_, err = writer.Write(b)
	return err
}
