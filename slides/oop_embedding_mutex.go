package main

import (
	"fmt"
	"sync"
)

// Count is a monotonically increasing counter
type Count struct {
	// embedded mutex has methods Lock() and Unlock()
	sync.Mutex
	val int64
}

// Value returns the current count
func (c *Count) Value() int64 {
	c.Lock()
	defer c.Unlock()
	return c.val
}

// Increment increases the counter by the given amount
func (c *Count) Increment(by int64) {
	c.Lock() // contrast with c.Mutex.Lock()
	c.val += by
	c.Unlock()
}

func main() {
	var c Count // zero value is useful
	for i := 0; i < 10; i++ {
		c.Increment(1)
	}
	fmt.Println("count is now:", c.Value())
	// count is now: 10
}
