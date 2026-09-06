package tokenbucket_test

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"tenant-gate/pkg/tokenbucket"

	"github.com/stretchr/testify/assert"
)

func TestTokenBucket(t *testing.T) {
	t.Parallel()

	// Create a new token bucket with a capacity of 5 tokens and a refill rate of 1 token per second
	bucket := tokenbucket.NewTokenBucket(1, 5)

	// Test consuming tokens
	for i := 0; i < 5; i++ {
		if !bucket.Allow() {
			t.Errorf("Expected to allow token consumption, but it was denied at iteration %d", i)
		}
	}

	// Test exceeding the token limit
	if bucket.Allow() {
		t.Errorf("Expected to deny token consumption, but it was allowed")
	}
}

func TestTokenBucketRefill(t *testing.T) {
	t.Parallel()
	bucket := tokenbucket.NewTokenBucket(1, 5)

	// Consume all tokens
	for i := 0; i < 5; i++ {
		if !bucket.Allow() {
			t.Errorf("Expected to allow token consumption, but it was denied at iteration %d", i)
		}
	}

	// Wait for 2 seconds to allow tokens to refill
	// Test exceeding the token limit
	if !bucket.AllowN(time.Now().Add(2*time.Second), 2) {
		t.Errorf("Expected to allow token consumption, but it was denied after waiting for refill")
	}
}

func TestTokenBucketConcurrency(t *testing.T) {
	t.Parallel()
	refill := 0.0
	capacity := 60000.0

	bucket := tokenbucket.NewTokenBucket(refill, capacity)

	// Add concurrency tests here
	wg := sync.WaitGroup{}
	var allowedCount, deniedCount int32
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if bucket.Allow() {
				atomic.AddInt32(&allowedCount, 1)
			} else {
				atomic.AddInt32(&deniedCount, 1)
			}
		}()
	}
	wg.Wait()

	t.Logf("Allowed: %d, Denied: %d", allowedCount, deniedCount)
	assert.EqualValues(t, allowedCount, 60000)
	assert.EqualValues(t, deniedCount, 40000)
}

func TestTokenBucketSetParam(t *testing.T) {
	t.Parallel()
	bucket := tokenbucket.NewTokenBucket(2, 10)
	bucket.SetParam(time.Now(), 1, 10)

	for i := 0; i < 10; i++ {
		if !bucket.Allow() {
			t.Errorf("Expected to allow token consumption, but it was denied at iteration %d", i)
		}
	}

	bucket.SetParam(time.Now(), 1, 20)
	assert.EqualValues(t, 0, bucket.GetTokens())
}
