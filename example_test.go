package conf_test

import (
	"fmt"
	"log"

	"github.com/popodidi/conf"
	"github.com/popodidi/conf/source/env"
	"github.com/popodidi/conf/source/flag"
	"github.com/popodidi/conf/source/memory"
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
		flag.New(),
		env.New(),
		yaml.New("config.yaml"),
		memory.New(map[string]interface{}{
			"Str":   "??",
			"Int":   "12",
			"Float": "3.456",
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("cfg.Str", cfg.Str)
	log.Println("cfg.Int", cfg.Int)
	log.Println("cfg.Float", cfg.Float)
}

func Example_flagUsage() {
	type config struct {
		Str   string `conf:"usage:str is a string"`
		Int   int
		Float float64
	}
	flagSource := flag.New()
	var cfg config
	err := conf.Configure(&cfg, flagSource)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(flagSource.Usage())
}
