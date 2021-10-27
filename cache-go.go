package cacheGo

import "sync"

type Cache struct {
	pair map[string]interface{}
	mu   *sync.RWMutex
}

func New() Cache {
	return Cache{
		pair: make(map[string]interface{}),
		mu:   new(sync.RWMutex),
	}
}

// set only if cache is not null (memory allocated)
func (c *Cache) Set(key string, value interface{}) {
	if c != nil {
		c.mu.Lock()
		c.pair[key] = value
		c.mu.Unlock()
	}
}

// return nil if no such value
func (c *Cache) Get(key string) interface{} {
	c.mu.RLock()
	value, ok := c.pair[key]
	c.mu.RUnlock()
	if ok {
		return value
	}
	return nil
}

// delete only if key exists or nothing happens
func (c *Cache) Delete(key string) {
	delete(c.pair, key)
}
