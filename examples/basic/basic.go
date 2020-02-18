package main

import (
	"log"

	"github.com/popodidi/conf"
	"github.com/popodidi/conf/source/env"
	"github.com/popodidi/conf/source/yaml"
)

func main() {
	type config struct {
		Str   string `conf:"default:hello"`
		Int   int
		Float float64
	}
	var err error
	var cfg config
	tmpl, _ := conf.Template(&cfg, yaml.NewExporter())
	log.Println(tmpl)
	err = conf.Load(&cfg, env.New(), yaml.New("basic.yaml"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("cfg.Str", cfg.Str)
	log.Println("cfg.Int", cfg.Int)
	log.Println("cfg.Float", cfg.Float)
}
