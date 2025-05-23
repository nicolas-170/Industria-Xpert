package config

import (
	"os"
	"strings"
)

type Config struct {
	Port string
}

func InitConfig() *Config {
	return &Config{
		Port: asignarStrOValorPorDefecto(os.Getenv("PORT"), "7777"),
	}
}

func asignarStrOValorPorDefecto(str, porDefecto string) string {
	if strings.TrimSpace(str) != "" {
		return str
	}
	return porDefecto
}
