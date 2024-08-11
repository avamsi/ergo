package group

import "sync"

type Collector[T any] struct {
	wg      sync.WaitGroup
	results chan T
}

func NewCollector[T any](results chan T) *Collector[T] {
	return &Collector[T]{results: results}
}

func (c *Collector[T]) Collect(v T) {
	c.results <- v
}

func (c *Collector[T]) Go(f func()) {
	c.wg.Add(1)
	go func() {
		f()
		c.wg.Done()
	}()
}

func (c *Collector[T]) Close() <-chan T {
	go func() {
		c.wg.Wait()
		close(c.results)
	}()
	return c.results
}
