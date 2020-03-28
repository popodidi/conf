package json

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/popodidi/conf"
)

// New returns a JSON file source.
func New(path string) conf.Source {
	return &jsn{path: path}
}

// NewExporter returns an json exporter.
func NewExporter() conf.Exporter {
	return &exporter{}
}

type jsn struct {
	exporter
	path string
}

func (j *jsn) Read(key string, path ...string) (value string, exists bool) {
	b, err := ioutil.ReadFile(j.path)
	if err != nil {
		return
	}
	var m map[string]interface{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return
	}
	return j.read(m, key, path...)
}

func (j *jsn) read(m map[string]interface{}, key string, path ...string) (
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
	return j.read(subM, key, path[1:]...)
}

type exporter struct{}

func (e *exporter) Export(m conf.Map, writer io.Writer) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	_, err = writer.Write(b)
	return err
}
