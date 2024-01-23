package logger

import (
	"log/slog"
	"os"

	"github.com/arsenydubrovin/level-0/src/internal/config"
	console "github.com/phsym/console-slog"
)

func Load(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case config.EnvLocal:
		logger = slog.New(
			console.NewHandler(os.Stderr, &console.HandlerOptions{Level: slog.LevelDebug, AddSource: true}),
		)

	case config.EnvDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)

	case config.EnvProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)

	default: // prod logger by default increases security
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return logger
}
