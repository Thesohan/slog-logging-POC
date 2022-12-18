package main

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	slog.Info("Go is best language!")
	slog.Debug("Go is best language!")
	slog.Warn("Go is best language!")
	slog.Error("Go is best language!", fmt.Errorf("error"))

	textHandler := slog.NewTextHandler(os.Stdout).
		WithAttrs([]slog.Attr{slog.String("app-version", "v0.0.1-beta")}) // 👈 add attributes to all logs

	logger := slog.New(textHandler)
	logger.Info("Go is best language")

	jsonTextHandler := slog.NewJSONHandler(os.Stdout) // 👈
	jsonLogger := slog.New(jsonTextHandler)

	jsonLogger.Info("Go is the best language!")

	logger.Info("Usage Statistics", slog.Int("current-memory", 50))

	logger.Info("Usage Statistics",
		slog.Group("memory",
			slog.Int("current", 50),
			slog.Int("min", 20),
			slog.Int("max", 80)),
		slog.Int("cpu", 10),
		slog.String("app-version", "v0.0.1-beta"),
	)

	ctx := slog.NewContext(context.Background(), logger)
	sendUsageStatus(ctx)
}

func sendUsageStatus(ctx context.Context) {
	logger := slog.FromContext(ctx) // 👈 grab logger from context

	logger.Info("Generating statistics")

	logger.Info("Usage Statistics",
		slog.Group("memory",
			slog.Int("current", 50),
			slog.Int("min", 20),
			slog.Int("max", 80)),
		slog.Int("cpu", 10),
	)
}
