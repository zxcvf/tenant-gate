package redis

import (
	"crypto/tls"
	"time"
)

type Option func(*Redis)

func TLSConfig(tlsConfig *tls.Config) Option {
	return func(r *Redis) {
		r.TLSConfig = tlsConfig
	}
}

func PoolFIFO(fifo bool) Option {
	return func(r *Redis) {
		r.PoolFIFO = fifo
	}
}

func PoolSize(size int) Option {
	return func(r *Redis) {
		r.PoolSize = size
	}
}

func PoolTimeout(timeout time.Duration) Option {
	return func(r *Redis) {
		r.PoolTimeout = timeout
	}
}

func MinIdleConns(min int) Option {
	return func(r *Redis) {
		r.MinIdleConns = min
	}
}

func MaxIdleConns(max int) Option {
	return func(r *Redis) {
		r.MaxIdleConns = max
	}
}

func ConnMaxIdleTime(max time.Duration) Option {
	return func(r *Redis) {
		r.ConnMaxIdleTime = max
	}
}

func ConnMaxLifetime(max time.Duration) Option {
	return func(r *Redis) {
		r.ConnMaxLifetime = max
	}
}
