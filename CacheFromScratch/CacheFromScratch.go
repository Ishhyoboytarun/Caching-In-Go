package CacheFromScratch

import (
	"fmt"
	"sync"
	"time"
)

// CacheItem represents an item in the cache.
type CacheItem struct {
	value      interface{}
	expiration time.Time
}

// Cache represents an in-memory cache.
type Cache struct {
	items map[string]CacheItem
	lock  sync.RWMutex
}

// NewCache creates a new instance of Cache.
func NewCache() *Cache {
	return &Cache{items: make(map[string]CacheItem)}
}

// Set sets a value in the cache with the given key and expiration time.
func (c *Cache) Set(key string, value interface{}, expiration time.Time) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.items[key] = CacheItem{value: value, expiration: expiration}
}

// Get gets a value from the cache with the given key.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	item, found := c.items[key]
	if !found {
		return nil, false
	}
	if item.expiration.Before(time.Now()) {
		delete(c.items, key)
		return nil, false
	}
	return item.value, true
}

// Delete deletes a value from the cache with the given key.
func (c *Cache) Delete(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.items, key)
}

func main() {
	cache := NewCache()

	// Set a value in the cache with a 1 second expiration time.
	cache.Set("key", "value", time.Now().Add(1*time.Second))

	// Get the value from the cache.
	value, found := cache.Get("key")
	if found {
		fmt.Println("Value found in cache:", value)
	} else {
		fmt.Println("Value not found in cache")
	}

	// Wait for the expiration time to pass.
	time.Sleep(2 * time.Second)

	// Get the value from the cache again.
	value, found = cache.Get("key")
	if found {
		fmt.Println("Value found in cache:", value)
	} else {
		fmt.Println("Value not found in cache")
	}
}
