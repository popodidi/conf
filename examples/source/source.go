package main

import (
	"log"

	"github.com/popodidi/conf"
	"github.com/popodidi/conf/source/env"
	"github.com/popodidi/conf/source/mock"
	"github.com/popodidi/conf/source/yaml"
)

type config struct {
	Str   string
	Int   int
	Float float64
}

func main() {
	var cfg config
	err := conf.Load(&cfg,
		env.New(),
		yaml.New("config.yaml"),
		mock.New(map[string]string{
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
