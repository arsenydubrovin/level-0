package service

import (
	"context"

	"github.com/arsenydubrovin/level-0/src/internal/model"
	validator "github.com/go-playground/validator/v10"
)

type orderService struct {
	r OrderRepository
	v *validator.Validate
}

type OrderRepository interface {
	Get(ctx context.Context, uid string) (*model.Order, error)
	GetUIDs(ctx context.Context) (*[]string, error)
	Insert(ctx context.Context, order *model.Order) (string, error)
}

func NewOrderService(r OrderRepository) *orderService {
	return &orderService{
		r: r,
		v: validator.New(),
	}
}
