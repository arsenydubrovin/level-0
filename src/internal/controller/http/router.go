package http

import (
	"context"
	"text/template"

	"github.com/arsenydubrovin/level-0/src/internal/model"
	echo "github.com/labstack/echo/v4"
)

type OrderRouter interface {
	fetchOrderHandler(ctx echo.Context) error
	listUIDsHandler(ctx echo.Context) error
	RegisterRoutes(e *echo.Echo)
}

type orderRouter struct {
	service  OrderService
	renderer *renderer
}

type OrderService interface {
	Fetch(ctx context.Context, uid string) (*model.Order, error)
	ListUIDs(ctx context.Context) (*[]string, error)
}

func NewOrderRouter(s OrderService) *orderRouter {
	return &orderRouter{
		service: s,
		renderer: &renderer{
			t: template.Must(template.ParseGlob("src/ui/html/*.html")),
		},
	}
}

func (r *orderRouter) RegisterRoutes(e *echo.Echo) {
	e.Renderer = r.renderer
	e.GET("/", r.mainPageHandler)

	apiGroup := e.Group("/api")
	apiGroup.GET("/order", r.fetchOrderHandler)
	apiGroup.GET("/uids", r.listUIDsHandler)
}
