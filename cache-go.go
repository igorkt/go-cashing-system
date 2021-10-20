package cacheGo

type Cache struct {
	pair map[string]interface{}
}

func NewCache() Cache {
	return Cache{
		pair: make(map[string]interface{}),
	}
}

// set only if cache is not null (memory allocated)
func (c *Cache) Set(key string, value interface{}) {
	if c != nil {
		c.pair[key] = value
	}
}

// return nil if no such value
func (c *Cache) Get(key string) interface{} {
	value, ok := c.pair[key]
	if ok {
		return value
	}
	return nil
}

// delete only if key exists or nothing happens
func (c *Cache) Delete(key string) {
	delete(c.pair, key)
}
