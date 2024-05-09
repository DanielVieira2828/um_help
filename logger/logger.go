package logger

import (
	"os"

	"github.com/DanielVieirass/um_help/config"
	"github.com/rs/zerolog"
)

func New(cfg *config.Config) *zerolog.Logger {
	logger := zerolog.New(os.Stderr).With().Str("service", cfg.InternalConfig.ServiceName).Timestamp().Logger()

	return &logger
}
