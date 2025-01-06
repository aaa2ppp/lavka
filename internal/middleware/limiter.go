package middleware

import (
	"net/http"
	"sync"
	"time"
)

func LimitRPS(rps int, h http.Handler) http.Handler {
	if rps <= 0 {
		return h
	}

	lim := newLimiter(rps)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !lim.pass() {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		h.ServeHTTP(w, r)
	})
}

type limiter struct {
	mu    sync.Mutex
	queue limiterQueue
}

func newLimiter(rps int) *limiter {
	return &limiter{
		queue: makeLimiterQueue(rps),
	}
}

func (lim *limiter) pass() bool {
	lim.mu.Lock()
	defer lim.mu.Unlock()

	now := time.Now()
	for !lim.queue.empty() && lim.queue.peek().Before(now) {
		lim.queue.pop()
	}

	if lim.queue.len() < lim.queue.cap() {
		lim.queue.push(now.Add(time.Second))
		return true
	}

	return false
}

// limiterQueue simple round queue with fixed capacity
type limiterQueue struct {
	items []time.Time
	first int
	_len  int
}

func makeLimiterQueue(size int) limiterQueue {
	return limiterQueue{
		items: make([]time.Time, size),
	}
}

func (q limiterQueue) len() int {
	return q._len
}

func (q limiterQueue) cap() int {
	return len(q.items)
}

func (q limiterQueue) empty() bool {
	return q._len == 0
}

func (q limiterQueue) peek() time.Time {
	if q._len == 0 {
		panic("queue is empty")
	}
	return q.items[q.first]
}

func (q *limiterQueue) pop() time.Time {
	v := q.peek()

	q.first++
	if q.first == len(q.items) {
		q.first = 0
	}

	q._len--
	return v
}

func (q *limiterQueue) push(v time.Time) {
	n := len(q.items)
	if q._len == n {
		panic("queue is full")
	}

	i := q.first + q._len
	if i >= n {
		i -= n
	}

	q.items[i] = v
	q._len++
}
