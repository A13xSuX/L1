package main

import (
	"fmt"
	"sync"
)

type counter struct {
	mu  sync.Mutex
	cnt int
}

func (c *counter) inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cnt++
}

func (c *counter) getCnt() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.cnt
}
func main() {
	var wg sync.WaitGroup
	count := &counter{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.inc()
		}()
	}

	wg.Wait()

	fmt.Println("count:", count.getCnt())
}
