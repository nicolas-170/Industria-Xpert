package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName  string
	Port     string
	ConfigDb ConfigDb
	Email    Email
}

type Email struct {
	EmailSend     string
	EmailPassword string
}

type ConfigDb struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbName     string
}

func InitConfig() *Config {
	// Cargamos las variables de entorno o valores del archivo .dev.env
	if err := godotenv.Load(".dev.env"); err != nil {
		log.Fatal("Error cargando el archivo .dev.env: ", err)
	}
	return &Config{
		AppName: asignarStrOValorPorDefecto(os.Getenv("APP_NAME"), "Industria-Xpert"),
		Port:    asignarStrOValorPorDefecto(os.Getenv("PORT"), "7777"),
		ConfigDb: ConfigDb{
			DbUser:     asignarStrOValorPorDefecto(os.Getenv("DB_USER"), "root"),
			DbPassword: os.Getenv("DB_PASSWORD"),
			DbHost:     asignarStrOValorPorDefecto(os.Getenv("DB_HOST"), "localhost:3306"),
			DbName:     os.Getenv("DB_NAME"),
		},
		Email: Email{
			EmailSend: os.Getenv("EMAIL_SEND"), 
			EmailPassword: os.Getenv("EMAIL_PASSWORD"),
		},
	}
}

func asignarStrOValorPorDefecto(str, porDefecto string) string {
	if strings.TrimSpace(str) != "" {
		return str
	}
	return porDefecto
}

//
