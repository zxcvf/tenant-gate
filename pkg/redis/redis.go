package redis

import (
	"crypto/tls"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	_defaultTLSMinVersion   = tls.VersionTLS12
	_defaultPoolFIFO        = true
	_defaultPoolSize        = 10
	_defaultPoolTimeout     = 30 * time.Second
	_defaultMinIdleConns    = 0
	_defaultMaxIdleConns    = 10
	_defaultConnMaxIdleTime = 5 * time.Minute
	_defaultConnMaxLifetime = 0 // 0 means no limit
)

// Redis -.
type Redis struct {
	PoolFIFO        bool
	PoolSize        int
	PoolTimeout     time.Duration
	MinIdleConns    int
	MaxIdleConns    int
	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration
	TLSConfig       *tls.Config

	DB     int
	Client *redis.Client
}

func New(addr string, password string, db int, opts ...Option) (*Redis, error) {
	r := &Redis{
		PoolFIFO:        _defaultPoolFIFO,
		PoolSize:        _defaultPoolSize,
		PoolTimeout:     _defaultPoolTimeout,
		MinIdleConns:    _defaultMinIdleConns,
		MaxIdleConns:    _defaultMaxIdleConns,
		ConnMaxIdleTime: _defaultConnMaxIdleTime,
		ConnMaxLifetime: _defaultConnMaxLifetime,
		DB:              db,
	}

	for _, opt := range opts {
		opt(r)
	}

	r.Client = redis.NewClient(
		&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       r.DB,
		},
	)
	return r, nil
}
