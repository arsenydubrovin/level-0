package stan

import (
	"context"
	"log/slog"

	stan "github.com/nats-io/stan.go"
)

func (sb *orderSubscriber) CreateOrder(msg *stan.Msg) {
	slog.Debug("message is recived", slog.String("msg", string(msg.Data[6:40])))

	uid, err := sb.s.Create(context.Background(), msg.Data)
	if err != nil {
		// TODO: handle error "already exists"
		slog.Error("error creating order", slog.String("err", err.Error()))
		return
	}

	slog.Debug("order is created", slog.String("uid", uid))
}
