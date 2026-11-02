package main

import (
	"sync"
	"time"
)

type Clock interface{ Now() time.Time }

type Limiter struct {
	mu     sync.Mutex
	clock  Clock
	rate   float64
	burst  int
	tokens float64
	last   time.Time
}

func NewLimiter(clock Clock, ratePerSec float64, burst int) *Limiter {
	limiter := &Limiter{
		clock:  clock,
		rate:   ratePerSec,
		burst:  burst,
		tokens: float64(burst),
	}
	if clock != nil {
		limiter.last = clock.Now()
	}
	return limiter
}

func (l *Limiter) AllowN(n int) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.clock == nil || l.burst <= 0 {
		return false
	}
	if n <= 0 || n > l.burst {
		return false
	}
	now := l.clock.Now()
	if now.After(l.last) {
		secondsPassed := now.Sub(l.last).Seconds()
		if l.rate > 0 {
			l.tokens = min(l.tokens+secondsPassed*l.rate, float64(l.burst))
		}
		l.last = now
	}
	tokensNeeded := float64(n)
	if l.tokens >= tokensNeeded {
		l.tokens -= tokensNeeded
		return true
	}
	return false
}
