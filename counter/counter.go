package counter

import (
	"sync"
)

// Counter keep a count.
type Counter struct {
	sync.Mutex
	n int
}

// Increase increments the count.
func (c *Counter) Increase() {
	c.Lock()
	defer c.Unlock()
	c.n++
}

// Decrease decrements the count.
func (c *Counter) Decrease() {
	c.Lock()
	defer c.Unlock()
	c.n--
}

// Flush resets the count.
func (c *Counter) Flush() {
	c.Lock()
	defer c.Unlock()
	c.n = 0
}

// Get returns the count.
func (c Counter) Get() int {
	c.Lock()
	defer c.Unlock()
	return c.n
}
