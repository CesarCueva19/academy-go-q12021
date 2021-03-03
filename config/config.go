package config

import (
	"time"
)

// Configuration contains application configuration
type Configuration struct {
	AppName  string `mapstructure:"app_name" validate:"required"`
	HTTPPort string `mapstructure:"http_port" validate:"required"`
	LogLevel string `mapstructure:"log_level" validate:"required"`
	Services *struct {
		PokemonAPI *ServConf `mapstructure:"pokemonAPI" validate:"required"`
	} `mapstructure:"services" validate:"required"`
	Pokedex string `mapstructure:"pokedex" validate:"required"`
}

// ServConf contains Service configurations
type ServConf struct {
	Host    string        `mapstructure:"host" validate:"required,url"`
	Timeout time.Duration `mapstructure:"timeout" validate:"required"`
}
