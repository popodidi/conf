package conf_test

import (
	"log"

	"github.com/popodidi/conf"
	"github.com/popodidi/conf/source/env"
	"github.com/popodidi/conf/source/yaml"
)

func Example_basic() {
	type config struct {
		Str   string `conf:"default:hello"`
		Int   int
		Float float64
	}
	var err error
	var cfg config
	tmpl, _ := conf.Template(&cfg, yaml.NewExporter())
	log.Println(tmpl)
	err = conf.Load(&cfg, env.New(), yaml.New("example_basic.yaml"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("cfg.Str", cfg.Str)
	log.Println("cfg.Int", cfg.Int)
	log.Println("cfg.Float", cfg.Float)
}

func Example_source() {
	type config struct {
		Str   string
		Int   int
		Float float64
	}

	var cfg config
	err := conf.Load(&cfg,
		env.New(),
		yaml.New("config.yaml"),
		conf.NewMock(map[string]string{
			"STR":   "??",
			"INT":   "12",
			"FLOAT": "3.456",
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("cfg.Str", cfg.Str)
	log.Println("cfg.Int", cfg.Int)
	log.Println("cfg.Float", cfg.Float)
}
