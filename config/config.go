package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
)

type envConfig struct {
	Version string `env:"VERSION" envDefault:"0.0.1"`
}

var (
	// Env is the config
	Env = envConfig{}
)

// Setup setup config function
func Setup() {
	if err := env.Parse(&Env); err != nil {
		log.Fatalf("%+v\n", err)
	}

	fmt.Printf("%+v\n", Env)
}
