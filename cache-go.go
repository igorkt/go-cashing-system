package cacheGo

import (
	"sync"
	"time"
)

const defaultTTL = 5 * time.Second

type Cache struct {
	pair       map[string]interface{}
	mu         *sync.RWMutex
	expireTime time.Time
}

func New() Cache {
	return Cache{
		pair:       make(map[string]interface{}),
		mu:         new(sync.RWMutex),
		expireTime: time.Now().Add(defaultTTL),
	}
}

// set only if cache is not null (memory allocated)
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	if c != nil {
		c.mu.Lock()
		c.pair[key] = value
		c.expireTime = time.Now().Add(ttl)
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
	c.mu.Lock()
	delete(c.pair, key)
	c.mu.Unlock()
}
