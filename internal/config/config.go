package config

import (
	logora "github.com/Edward-Alphonse/logora/writers"

	"github.com/Edward-Alphonse/pkg/config"
)

type Configuration struct {
	Logs *logora.FileConfig
}

var cfg *Configuration

func GetConfig() *Configuration {
	return cfg
}

func Init(path string) *Configuration {
	cfg = config.Init[Configuration](path)
	return cfg
}
