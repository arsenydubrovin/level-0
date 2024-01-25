package main

import (
	"context"
	"log"
	"net/http"

	"github.com/arsenydubrovin/level-0/src/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to run app: %s", err.Error())
	}

	// TODO: gracefull shutdown
}
