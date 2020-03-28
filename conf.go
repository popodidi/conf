package conf

// Load loads the config from flattened readers.
func Load(config interface{}, readers ...Reader) (err error) {
	return New(config).Load(readers)
}

// LoadNested loads the config from nested readers.
func LoadNested(config interface{}, readers ...Reader) (err error) {
	return New(config).LoadNested(readers)
}

// Template returns a flattened config template with exporter.
func Template(config interface{}, exporter Exporter) (
	tmpl string, err error) {
	return New(config).Template(exporter)
}

// NestedTemplate returns a nested config template with exporter.
func NestedTemplate(config interface{}, exporter Exporter) (
	tmpl string, err error) {
	return New(config).NestedTemplate(exporter)
}
