package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rotisserie/eris"
)

const AppName = "Ocram"

type Config struct {
	Valkey
	Google
	App
}

func Load() (Config, error) {
	var valkey Valkey
	if err := envconfig.Process("VALKEY", &valkey); err != nil {
		return Config{}, eris.Wrap(err, "error loading VALKEY env")
	}

	var google Google
	if err := envconfig.Process("GOOGLE", &google); err != nil {
		return Config{}, eris.Wrap(err, "error loading GOOGLE env")
	}

	var app App
	if err := envconfig.Process("APP", &app); err != nil {
		return Config{}, eris.Wrap(err, "error loading APP env")
	}

	return Config{
		Valkey: valkey,
		Google: google,
		App:    app,
	}, nil
}
