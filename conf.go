package conf

// Configure configures the configurable.
func Configure(config interface{}, configurable Configurable) error {
	return New(config).Configure(configurable)
}

// Load loads the config from readers.
func Load(config interface{}, readers ...Reader) (err error) {
	c := New(config)
	for _, r := range readers {
		if configurable, ok := r.(Configurable); ok {
			err = c.Configure(configurable)
			if err != nil {
				return
			}
		}
	}
	return c.Load(readers...)
}

// Template returns a config template with exporter.
func Template(config interface{}, exporter Exporter) (
	tmpl string, err error) {
	c := New(config)
	if configurable, ok := exporter.(Configurable); ok {
		err = c.Configure(configurable)
		if err != nil {
			return
		}
	}
	return c.Template(exporter)
}
