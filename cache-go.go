package cacheGo

import (
	"sync"
	"time"
)

type CacheItem struct {
	value      interface{}
	expireTime time.Time
}

type Cache struct {
	items map[string]CacheItem
	mu    *sync.RWMutex
}

func New() *Cache {
	newCache := Cache{
		items: make(map[string]CacheItem),
		mu:    new(sync.RWMutex),
	}

	newCache.StartBackgroundTasks()

	return &newCache
}

// set only if cache is not null (memory allocated)
func (c *Cache) Set(key string, val interface{}, ttl time.Duration) {
	if c != nil {
		c.mu.Lock()
		c.items[key] = CacheItem{
			value:      val,
			expireTime: time.Now().Add(ttl),
		}
		c.mu.Unlock()
	}
}

// return nil if no such value
func (c *Cache) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.items[key]
	if ok && c.items[key].expireTime.UnixNano() < time.Now().UnixNano() {
		return item.value
	}
	return nil
}

// delete only if key exists or nothing happens
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	delete(c.items, key)
	c.mu.Unlock()
}

func (c *Cache) StartBackgroundTasks() {
	go c.DeleteExpired()
}

func (c *Cache) DeleteExpired() {
	for {
		for key, item := range c.items {
			if item.expireTime.UnixNano() > time.Now().UnixNano() {
				c.Delete(key)
			}
		}
	}
}
