package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mutex   *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		entries: make(map[string]cacheEntry),
		mutex:   &sync.Mutex{},
	}
	go c.reaploop(interval)
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
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reaploop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for {
			<-ticker.C
			c.mutex.Lock()

			now := time.Now()
			for key, entry := range c.entries {
				if now.Sub(entry.createdAt) > interval {
					delete(c.entries, key)
				}
			}
			c.mutex.Unlock()
		}
	}()
}
