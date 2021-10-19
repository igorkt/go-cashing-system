package cacheGo

type Cache struct {
	pair map[string]interface{}
}

func NewCache() Cache {
	return Cache{
		pair: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.pair[key] = value
}

func (c *Cache) Get(key string) interface{} {
	value, ok := c.pair[key]
	if ok {
		return value
	}
	return nil
}

func (c *Cache) Delete(key string) {
	_, ok := c.pair[key]
	if ok {
		delete(c.pair, key)
	}
}
