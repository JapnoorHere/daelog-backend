package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init() {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05",
	}

	log.Logger = zerolog.New(output).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}

func Info(msg string) *zerolog.Event {
	return log.Info().Str("msg", msg)
}

func Error(msg string, err error) {
	log.Error().Err(err).Msg(msg)
}

func Fatal(msg string, err error) {
	log.Fatal().Err(err).Msg(msg)
}

func Debug(msg string) *zerolog.Event {
	return log.Debug().Str("msg", msg)
}

func Warn(msg string) *zerolog.Event {
	return log.Warn().Str("msg", msg)
}
