package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mutex   *sync.Mutex
	entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		mutex:   &sync.Mutex{},
		entries: map[string]cacheEntry{},
	}

	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, found := c.entries[key]
	return entry.val, found
}

func (c *Cache) reapLoop(ttl time.Duration) {
	ticker := time.NewTicker(ttl)
	for range ticker.C {
		for k, v := range c.entries {
			if v.createdAt.Add(ttl).Before(time.Now()) {
				c.mutex.Lock()
				delete(c.entries, k)
				c.mutex.Unlock()
			}
		}
	}

}
