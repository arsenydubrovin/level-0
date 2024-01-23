package app

import (
	"log"

	"github.com/arsenydubrovin/level-0/src/internal/config"
	"github.com/arsenydubrovin/level-0/src/internal/controller/http"
	"github.com/arsenydubrovin/level-0/src/internal/infrastructure/postgres"
	"github.com/arsenydubrovin/level-0/src/internal/service"
)

// Service provider is responsible for dependency injection
type serviceProvider struct {
	applicationConfig config.ApplicationConfig
	httpConfig        config.HTTPConfig
	postgresConfig    config.PostgresConfig

	orderRepository service.OrderRepository
	orderService    http.OrderService
	orderRouter     http.OrderRouter
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

func (s *serviceProvider) OrderRepository() service.OrderRepository {
	if s.orderRepository == nil {
		repo, err := postgres.NewOrderRepository(
			s.PostgresConfig().DSN(),
			s.PostgresConfig().MaxOpenConns(),
			s.PostgresConfig().MaxIdleConns(),
			s.PostgresConfig().MaxIdleTime(),
		)
		if err != nil {
			log.Fatalf("failed to initialize order repository: %s", err.Error())
		}

		s.orderRepository = repo
	}

	return s.orderRepository
}

func (s *serviceProvider) OrderService() http.OrderService {
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
