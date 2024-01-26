package http

import (
	"context"

	"github.com/arsenydubrovin/level-0/src/internal/model"
	echo "github.com/labstack/echo/v4"
)

type OrderRouter interface {
	fetchOrderHandler(ctx echo.Context) error
	listUIDsHandler(ctx echo.Context) error
	RegisterRoutes(e *echo.Echo)
}

type orderRouter struct {
	s OrderService
}

type OrderService interface {
	Fetch(ctx context.Context, uid string) (*model.Order, error)
	ListUIDs(ctx context.Context) (*[]string, error)
}

func NewOrderRouter(s OrderService) *orderRouter {
	return &orderRouter{
		s: s,
	}
}

func (r *orderRouter) RegisterRoutes(e *echo.Echo) {
	apiGroup := e.Group("/api")
	apiGroup.GET("/order/:uid", r.fetchOrderHandler)
	apiGroup.GET("/uids", r.listUIDsHandler)
}
