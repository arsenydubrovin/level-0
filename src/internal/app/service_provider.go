package app

import (
	"log"

	"github.com/arsenydubrovin/level-0/src/internal/config"
	"github.com/arsenydubrovin/level-0/src/internal/controller/http"
	"github.com/arsenydubrovin/level-0/src/internal/controller/stan"
	"github.com/arsenydubrovin/level-0/src/internal/infrastructure/cache"
	"github.com/arsenydubrovin/level-0/src/internal/infrastructure/postgres"
	"github.com/arsenydubrovin/level-0/src/internal/service"
)

// Service provider is responsible for dependency injection.
type serviceProvider struct {
	applicationConfig config.ApplicationConfig
	httpConfig        config.HTTPConfig
	postgresConfig    config.PostgresConfig
	cacheConfig       config.CacheConfig
	stanConfig        config.StanConfig

	orderRepository service.OrderRepository
	orderService    orderService
	orderRouter     http.OrderRouter
	orderSubscriber stan.OrderSubscriber
}

// orderService combines requirements of http server and message broker.
type orderService interface {
	http.OrderService
	stan.OrderService
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) ApplicationConfig() config.ApplicationConfig {
	if s.applicationConfig == nil {
		cfg, err := config.NewApplicationConfig()
		if err != nil {
			log.Fatalf("failed to get application config: %s", err.Error())
		}

		s.applicationConfig = cfg
	}

	return s.applicationConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get HTTP config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) PostgresConfig() config.PostgresConfig {
	if s.postgresConfig == nil {
		cfg, err := config.NewPostgresConfig()
		if err != nil {
			log.Fatalf("failed to get postgres config: %s", err.Error())
		}

		s.postgresConfig = cfg
	}

	return s.postgresConfig
}

func (s *serviceProvider) CacheConfig() config.CacheConfig {
	if s.cacheConfig == nil {
		cfg, err := config.NewCacheConfig()
		if err != nil {
			log.Fatalf("failed to get cache config: %s", err.Error())
		}

		s.cacheConfig = cfg
	}

	return s.cacheConfig
}

func (s *serviceProvider) StanConfig() config.StanConfig {
	if s.stanConfig == nil {
		cfg, err := config.NewStanConfig()
		if err != nil {
			log.Fatalf("failed to get stan config: %s", err.Error())
		}

		s.stanConfig = cfg
	}

	return s.stanConfig
}

func (s *serviceProvider) OrderRepository() service.OrderRepository {
	if s.orderRepository == nil {
		pg, err := postgres.NewOrderRepository(
			s.PostgresConfig().DSN(),
			s.PostgresConfig().MaxOpenConns(),
			s.PostgresConfig().MaxIdleConns(),
			s.PostgresConfig().MaxIdleTime(),
		)
		if err != nil {
			log.Fatalf("failed to initialize order repository: %s", err.Error())
		}

		// decorate the database with a cache
		oc, err := cache.NewOrderCache(
			s.CacheConfig().DefaultExpiration(),
			s.CacheConfig().CleanupInterval(),
			pg,
		)
		if err != nil {
			log.Fatalf("failed to initialize order cache: %s", err.Error())
		}

		s.orderRepository = oc
	}

	return s.orderRepository
}

func (s *serviceProvider) OrderService() orderService {
	if s.orderService == nil {
		s.orderService = service.NewOrderService(s.OrderRepository())
	}

	return s.orderService
}

func (s *serviceProvider) OrderRouter() http.OrderRouter {
	if s.orderRouter == nil {
		s.orderRouter = http.NewOrderRouter(s.OrderService())
	}

	return s.orderRouter
}

func (s *serviceProvider) OrderSubscriber() stan.OrderSubscriber {
	if s.orderSubscriber == nil {
		s.orderSubscriber = stan.NewOrderSubscriber(s.OrderService())
	}

	return s.orderSubscriber
}
