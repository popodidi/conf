package conf

// Load loads the config from getters.
func Load(config interface{}, readers ...Reader) (err error) {
	return New(config).Load(readers)
}

// Template returns a config template with exporter.
func Template(config interface{}, exporter Exporter) (
	tmpl string, err error) {
	return New(config).Template(exporter)
}
