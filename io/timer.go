package io

import "time"

// Timer type
type Timer struct {
	start time.Time
	Total time.Duration
}

// Start timer
func (t *Timer) Start() {
	t.start = time.Now()
}

// Stop timer and add to total
func (t *Timer) Stop() {
	t.Total += time.Now().Sub(t.start)
}
