package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `yaml:"env" env-default="local"`
	StoragePath string        `yaml:"storage_path" env-required=true`
	TokenTTL    time.Duration `yaml:"token_ttl`
	GRPC        GRPCConfig    `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	var cfg Config

	if err := cleanenv.ReadConfig("../../config/local.yaml", &cfg); err != nil {
		panic("failded to read config: " + err.Error())
	}

	return &cfg
}
