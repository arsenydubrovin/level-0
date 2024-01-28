package cache

import (
	"sync"
	"time"
)

type Cache struct {
	mu                sync.RWMutex
	data              map[string]item
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
}

type item struct {
	value          any
	expirationTime int64
}

// New initializes a new cache.
// defaultExpiration set to zero means no cache expiration, and cleanupInterval set to zero means no cache cleaning.
func New(defaultExpiration, cleanupInterval time.Duration) *Cache {
	cache := Cache{
		data:              make(map[string]item),
		cleanupInterval:   cleanupInterval,
		defaultExpiration: defaultExpiration,
	}

	if cleanupInterval > 0 {
		go cache.collectGarbage()
	}

	return &cache
}

// Set adds a value to the cache by key.
func (c *Cache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = item{
		value:          value,
		expirationTime: time.Now().Add(c.defaultExpiration).UnixNano(),
	}
}

// Get gets a value from the cache by key.
func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.data[key]
	if !ok {
		return nil, false
	}

	if item.expirationTime > 0 {
		// if cache expired, but still exists
		if time.Now().UnixNano() > item.expirationTime {
			return nil, false
		}
	}

	return item.value, true
}

// collectGarbage removes expired items from cache.
func (c *Cache) collectGarbage() {
	for {

		<-time.After(c.cleanupInterval)

		if c.data == nil {
			return
		}

		if keys := c.expiredKeys(); len(keys) != 0 {
			c.clearItems(keys)
		}
	}
}

// expiredKeys returns key list which are expired.
func (c *Cache) expiredKeys() (keys []string) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for k, i := range c.data {
		if time.Now().UnixNano() > i.expirationTime && i.expirationTime > 0 {
			keys = append(keys, k)
		}
	}

	return
}

// clearItems removes all the items which key in keys.
func (c *Cache) clearItems(keys []string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, k := range keys {
		delete(c.data, k)
	}
}
