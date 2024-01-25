package app

import (
	"context"
	"log"
	"log/slog"

	"github.com/arsenydubrovin/level-0/src/internal/config"
	"github.com/arsenydubrovin/level-0/src/internal/logger"
	echo "github.com/labstack/echo/v4"
	stan "github.com/nats-io/stan.go"
	slogecho "github.com/samber/slog-echo"
)

type App struct {
	sp               *serviceProvider
	httpServer       *echo.Echo
	stanSubscriber   stan.Conn
	stanSubscription stan.Subscription
	logger           *slog.Logger
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	deps := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initLogger,
		a.initHTTPServer,
		a.initStanSubscriber,
	}

	for _, f := range deps {
		err := f(ctx)
		if err != nil {
			return nil, err
		}
	}

	return a, nil
}

func (a *App) Run() error {
	err := a.runStanSubscriber()
	if err != nil {
		return err
	}

	err = a.runHTTPServer()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	err := a.stanSubscription.Unsubscribe()
	if err != nil {
		return err
	}

	err = a.stanSubscriber.Close()
	if err != nil {
		return err
	}

	if err := a.httpServer.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.sp = newServiceProvider()

	return nil
}

func (a *App) initLogger(_ context.Context) error {
	l := logger.Load(a.sp.ApplicationConfig().Env())

	slog.SetDefault(l)
	a.logger = l

	return nil
}

func (a *App) initHTTPServer(_ context.Context) error {
	e := echo.New()
	e.Use(slogecho.NewWithConfig(a.logger,
		slogecho.Config{
			WithRequestBody: true,
		}))

	a.sp.OrderRouter().RegisterRoutes(e)

	a.httpServer = e

	return nil
}

func (a *App) runHTTPServer() error {
	err := a.httpServer.Start(a.sp.HTTPConfig().Port())
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initStanSubscriber(_ context.Context) error {
	sc, err := stan.Connect(a.sp.StanConfig().StanClusterID(), "order-subscriber", stan.NatsURL(a.sp.StanConfig().NatsURL()))
	if err != nil {
		log.Fatalf("failed to connect to nats-streaming: %s", err.Error())
	}
	a.stanSubscriber = sc

	return nil
}

func (a *App) runStanSubscriber() error {
	sub, err := a.stanSubscriber.Subscribe("subject", a.sp.OrderSubscriber().CreateOrder)
	if err != nil {
		return err
	}
	a.stanSubscription = sub

	return nil
}
