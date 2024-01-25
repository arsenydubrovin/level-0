package stan

import (
	"context"

	stan "github.com/nats-io/stan.go"
)

type OrderSubscriber interface {
	CreateOrder(msg *stan.Msg)
}

type orderSubscriber struct {
	s OrderService
}

type OrderService interface {
	Create(ctx context.Context, orderJSON []byte) (string, error)
}

func NewOrderSubscriber(s OrderService) *orderSubscriber {
	return &orderSubscriber{
		s: s,
	}
}
