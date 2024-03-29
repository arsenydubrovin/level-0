package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/arsenydubrovin/level-0/src/internal/model"
	validator "github.com/go-playground/validator/v10"
)

func (s *orderService) Fetch(ctx context.Context, uid string) (*model.Order, error) {
	order, err := s.r.Get(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch order: %w", err)
	}
	return order, nil
}

func (s *orderService) ListUIDs(ctx context.Context) (*[]string, error) {
	uids, err := s.r.GetUIDs(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list order UIDs: %w", err)
	}
	return uids, nil
}

func (s *orderService) Create(ctx context.Context, orderJSON []byte) (string, error) {
	var order model.Order

	if err := json.Unmarshal(orderJSON, &order); err != nil {
		return "", model.ErrInvalidData
	}

	err := s.v.Struct(order)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			slog.Debug("failed to validate JSON",
				slog.String("field", e.Field()),
				slog.Any("value", e.Value()),
			)
		}
		return "", model.ErrInvalidData
	}

	uid, err := s.r.Insert(ctx, &order)
	if err != nil {
		return "", fmt.Errorf("failed to create order: %w", err)
	}
	return uid, nil
}
