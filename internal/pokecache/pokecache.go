package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	CreatedAt time.Time
	val       []byte
}

type Cache struct {
	mu   *sync.Mutex
	data map[string]cacheEntry
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		mu:   &sync.Mutex{},
		data: make(map[string]cacheEntry),
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{
		CreatedAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, ok := c.data[key]
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.data {
		if v.CreatedAt.Before(now.Add(-last)) {
			delete(c.data, k)
		}
	}
}
