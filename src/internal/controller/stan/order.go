package stan

import (
	"context"
	"errors"
	"log/slog"

	"github.com/arsenydubrovin/level-0/src/internal/model"
	stan "github.com/nats-io/stan.go"
)

func (sb *orderSubscriber) CreateOrder(msg *stan.Msg) {
	slog.Debug("message recived", slog.Int("lenght", len(msg.Data)))

	uid, err := sb.s.Create(context.Background(), msg.Data)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrInvalidData):
			slog.Debug("received message is not valid")
		case errors.Is(err, model.ErrOrderExists):
			slog.Error("order with this already exists")
		default:
			slog.Error("error creating order", slog.String("err", err.Error()))
		}
		return
	}

	slog.Debug("order is created", slog.String("uid", uid))
}
