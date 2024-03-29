package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"

	"github.com/arsenydubrovin/level-0/src/internal/model"
	"github.com/lib/pq"
)

func (r *orderRepository) Insert(ctx context.Context, order *model.Order) (string, error) {
	orderJSON, err := json.Marshal(order)
	if err != nil {
		slog.Error("orderRepository.Insert", slog.String("err", err.Error()))
		return "", fmt.Errorf("failed to marshal order to JSON: %w", err)
	}

	stmt := "INSERT INTO orders (data, uid) VALUES ($1, $2) RETURNING uid"

	var uid string

	// order.OrderUID is duplicated as a separate value for quick retrieval
	err = r.db.QueryRowContext(ctx, stmt, orderJSON, order.OrderUID).Scan(&uid)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			// if the error is a unique key violation
			if pqErr.Code == "23505" {
				return "", model.ErrOrderExists
			}
		}
		slog.Error("orderRepository.Insert", slog.String("err", err.Error()))
		return "", fmt.Errorf("failed to insert order: %w", err)
	}

	return uid, nil
}

func (r *orderRepository) Get(ctx context.Context, uid string) (*model.Order, error) {
	stmt := "SELECT data FROM orders WHERE uid = $1"

	var orderData []byte

	err := r.db.QueryRowContext(ctx, stmt, uid).Scan(&orderData)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			slog.Error("orderRepository.Get", slog.String("err", err.Error()))
			return nil, model.ErrOrderNotFound
		default:
			slog.Error("orderRepository.Get", slog.String("err", err.Error()))
			return nil, fmt.Errorf("failed to get order: %w", err)
		}
	}

	var order model.Order

	err = json.Unmarshal(orderData, &order)
	if err != nil {
		slog.Error("orderRepository.Get", slog.String("err", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal order data: %w", err)
	}

	return &order, nil
}

func (r *orderRepository) GetUIDs(ctx context.Context) (*[]string, error) {
	stmt := "SELECT uid FROM orders"

	rows, err := r.db.QueryContext(ctx, stmt)
	if err != nil {
		slog.Error("orderRepository.GetUIDs", slog.String("err", err.Error()))
		return nil, fmt.Errorf("failed to get uids: %w", err)
	}
	defer rows.Close()

	var uids []string

	for rows.Next() {
		var uid string
		err := rows.Scan(&uid)
		if err != nil {
			slog.Error("orderRepository.GetUIDs", slog.String("err", err.Error()))
			return nil, fmt.Errorf("failed to scan uid: %w", err)
		}
		uids = append(uids, uid)
	}

	if err := rows.Err(); err != nil {
		slog.Error("orderRepository.GetUIDs", slog.String("err", err.Error()))
		return nil, fmt.Errorf("error reading rows: %w", err)
	}

	return &uids, nil
}
