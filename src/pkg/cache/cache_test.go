package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	testKey   = "testKey"
	testValue = "testValue"
)

func TestCacheSetAndGet(t *testing.T) {
	c := New(time.Second, 0)

	c.Set(testKey, testValue)
	result, ok := c.Get(testKey)

	assert.True(t, ok, "expected value to be present in cache")
	assert.Equal(t, testValue, result, "expected value to match the set value")
}

func TestCacheExpiration(t *testing.T) {
	c := New(time.Millisecond*500, time.Second)

	c.Set(testKey, testValue)
	time.Sleep(time.Millisecond * 600)
	result, ok := c.Get(testKey)

	assert.False(t, ok, "expected value to be expired")
	assert.Nil(t, result, "expected value to be nil")
}
