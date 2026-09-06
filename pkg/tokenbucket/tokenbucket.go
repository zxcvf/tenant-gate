package tokenbucket

import (
	"sync"
	"time"
)

// token bucket for rate limiting
const (
	_defaultGlobalRefillRate = 10 // tokens per second
	_defaultGlobalCapacity   = 100
)

type TokenBucket struct {
	mu sync.Mutex

	refillRate float64   // refill tokens per second
	capacity   float64   // maximum number of tokens in the bucket
	tokens     float64   // current number of tokens in the bucket
	lastRefill time.Time // timestamp in nanoseconds
}

func NewTokenBucket(refillRate, capacity float64) *TokenBucket {
	if refillRate <= 0 {
		refillRate = _defaultGlobalRefillRate
	}
	if capacity <= 0 {
		capacity = _defaultGlobalCapacity
	}
	return &TokenBucket{
		refillRate: refillRate,
		capacity:   capacity,
		tokens:     capacity,
		lastRefill: time.Now(),
	}
}

func (tb *TokenBucket) refill(now time.Time) {
	if now.Before(tb.lastRefill) {
		tb.lastRefill = now
		return
	}

	elapsed := now.Sub(tb.lastRefill).Seconds()
	tb.tokens += elapsed * tb.refillRate
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
	tb.lastRefill = now
}

func (tb *TokenBucket) Allow() bool {
	return tb.AllowN(time.Now(), 1)
}

func (tb *TokenBucket) AllowN(now time.Time, n int) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill(now)
	if tb.tokens >= float64(n) {
		tb.tokens -= float64(n)
		return true
	}
	return false
}

func (tb *TokenBucket) SetParam(now time.Time, newRate float64, newCapacity float64) {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill(now)

	if newRate > 0 {
		tb.refillRate = newRate
	}
	if newCapacity > 0 {
		tb.capacity = newCapacity
		if tb.tokens > tb.capacity {
			tb.tokens = tb.capacity
		}
	}
}

func (tb *TokenBucket) GetTokens() float64 {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	return tb.tokens
}
