package service

import (
	"context"

	"github.com/arsenydubrovin/level-0/src/internal/model"
)

func (s *orderService) Fetch(ctx context.Context, uid string) (*model.Order, error) {
	order, err := s.r.Get(ctx, uid)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *orderService) ListUIDs(ctx context.Context) (*[]string, error) {
	uids, err := s.r.GetUIDs(ctx)
	if err != nil {
		return nil, err
	}

	return uids, nil
}

func (s *orderService) Create(ctx context.Context, order *model.Order) (string, error) {
	// TODO: validate model

	uid, err := s.r.Insert(ctx, order)
	if err != nil {
		return "", err
	}

	return uid, nil
}
