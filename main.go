package main

import (
	cacheGo "cache-go/cache"
	"fmt"
)

func main() {
	cache := cacheGo.NewCache()
	cache.Set("userId", "Hello")
	userId := cache.Get("userId")
	fmt.Println(userId)
	cache.Delete("userId")
	userId = cache.Get("userId")
	fmt.Println(userId)
}
