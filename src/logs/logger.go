package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func Logger() *zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}

	logger := zerolog.New(file).With().Timestamp().Str("AppName", "Audit System").Logger()
	return &logger

}
