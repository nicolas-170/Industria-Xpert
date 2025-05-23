package logger

import (
	"errors"
	"os"

	"github.com/rs/zerolog"
)

// Logger es una instancia global de zerolog.
// El cual nos va a dar los logs de cada proceso o informacion deseada
var log *zerolog.Logger
var ErrLoggerNotInitialized = errors.New("logger no inicialiazdo: Usa primero su NewLogger()")

// inicializa el logger con color
func NewLogger() {
	// Con ConsoleWriter tenemos -> INFO en verde, ERROR en rojo
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}
	logger := zerolog.New(output).With().Timestamp().Logger()
	log = &logger
}

func Info(msg string) error {
	if log == nil {
		return ErrLoggerNotInitialized
	}
	log.Info().Msg(msg)
	return nil
}

func Error(msg string) error {
	if log == nil {
		return ErrLoggerNotInitialized
	}
	log.Error().Msg(msg)
	return nil
}

func Warn(msg string) error {
	if log == nil {
		return ErrLoggerNotInitialized
	}
	log.Warn().Msg(msg)
	return nil
}

func Fatal(err error) error {
	if log == nil {
		return ErrLoggerNotInitialized
	}
	if err != nil {
		log.Fatal().Msg(err.Error())
		return nil
	}
	log.Fatal()
	return nil
}

func Debug(msg string) error {
	if log == nil {
		return ErrLoggerNotInitialized
	}
	log.Debug().Msg(msg)
	return nil
}
