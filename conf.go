package conf

import (
	"github.com/popodidi/conf/source"
)

// Load loads the config from getters.
func Load(config interface{}, readers ...source.Reader) (err error) {
	return New(config).Load(readers)
}

// Template returns a config template with exporter.
func Template(config interface{}, exporter source.Exporter) (
	tmpl string, err error) {
	return New(config).Template(exporter)
}
