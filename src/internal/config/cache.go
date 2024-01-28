package config

import "time"

const (
	defaultExpirationEnvName = "CACHE_DEFAULT_EXPIRATION_TIME"
	cleanupIntervalEnvName   = "CACHE_CLEANUP_INTERVAL"
)

type CacheConfig interface {
	DefaultExpiration() time.Duration
	CleanupInterval() time.Duration
}

type cacheConfig struct {
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
}

func NewCacheConfig() (CacheConfig, error) {
	defaultExpirationStr, err := getEnvVariable(defaultExpirationEnvName)
	if err != nil {
		return nil, err
	}
	defaultExpiration, err := time.ParseDuration(defaultExpirationStr)
	if err != nil {
		return nil, err
	}

	cleanupIntervalStr, err := getEnvVariable(cleanupIntervalEnvName)
	if err != nil {
		return nil, err
	}
	cleanupInterval, err := time.ParseDuration(cleanupIntervalStr)
	if err != nil {
		return nil, err
	}

	return &cacheConfig{
		defaultExpiration: defaultExpiration,
		cleanupInterval:   cleanupInterval,
	}, nil
}

func (cfg *cacheConfig) DefaultExpiration() time.Duration {
	return cfg.defaultExpiration
}

func (cfg *cacheConfig) CleanupInterval() time.Duration {
	return cfg.cleanupInterval
}
