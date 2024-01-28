package cache

import (
	"context"
	"time"

	"github.com/arsenydubrovin/level-0/src/internal/model"
	"github.com/arsenydubrovin/level-0/src/pkg/cache"
)

var _ OrderRepository = (*orderCache)(nil)

// orderCache struct decorates an OrderRepository.
type orderCache struct {
	cache *cache.Cache
	db    OrderRepository
}

type OrderRepository interface {
	Get(ctx context.Context, uid string) (*model.Order, error)
	GetUIDs(ctx context.Context) (*[]string, error)
	Insert(ctx context.Context, order *model.Order) (string, error)
}

// NewOrderCache creates a new orderCache instance.
func NewOrderCache(defaultExpiration, cleanupInterval time.Duration, r OrderRepository) (*orderCache, error) {
	oc := &orderCache{
		cache: cache.New(defaultExpiration, cleanupInterval),
		db:    r,
	}

	err := oc.warmUp()
	if err != nil {
		return nil, err
	}

	return oc, nil
}

// warmUp preloads orders into the cache from the underlying database.
func (c *orderCache) warmUp() error {
	ctx := context.Background()

	uids, err := c.db.GetUIDs(ctx)
	if err != nil {
		return err
	}

	for _, uid := range *uids {
		if _, ok := c.cache.Get(uid); ok {
			continue
		}

		order, err := c.db.Get(ctx, uid)
		if err != nil {
			return err
		}

		c.cache.Set(uid, order)
	}

	return nil
}

// Get retrieves an order by UID, caching from the database if not found in the cache.
func (c *orderCache) Get(ctx context.Context, uid string) (*model.Order, error) {
	value, ok := c.cache.Get(uid)
	if !ok {
		order, err := c.db.Get(ctx, uid)
		if err != nil {
			return nil, err
		}

		c.cache.Set(uid, order)

		return order, nil
	}

	return value.(*model.Order), nil
}

// GetUIDs wraps OrderRepository.GetUIDs.
func (c *orderCache) GetUIDs(ctx context.Context) (*[]string, error) {
	uids, err := c.db.GetUIDs(ctx)
	if err != nil {
		return nil, err
	}

	return uids, nil
}

// Insert adds a new order to the database and updates the cache.
func (c *orderCache) Insert(ctx context.Context, order *model.Order) (string, error) {
	uid, err := c.db.Insert(ctx, order)
	if err != nil {
		return "", err
	}

	c.cache.Set(uid, order)

	return uid, nil
}
