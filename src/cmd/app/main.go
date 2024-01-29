package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/arsenydubrovin/level-0/src/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		err = a.Run()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("shutting down the app: %s", err.Error())
		}
	}()

	// Graceful Shutdown
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	slog.Debug("completing background tasks...")

	if err := a.Stop(ctx); err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	slog.Debug("shutting down the app...")
}
