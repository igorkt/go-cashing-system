package cacheGo

type MemoryCache interface {
	Set(key string, value interface{}) bool
	Get(key string)
	Delete(key string)
}

func (memCache MemoryCache) Set(key string, value interface{}) bool {

}
