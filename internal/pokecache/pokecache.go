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
	mx      sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	entry, ok := c.entries[key]
	if !ok {
		return []byte{}, false
	}
	val := entry.val
	if val == nil {
		return []byte{}, false
	}
	return val, true
}

func (c *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		t := <-ticker.C
		c.mx.Lock()
		reapTime := t.Add(-interval)
		for k := range c.entries {
			if reapTime.After(c.entries[k].createdAt) {
				delete(c.entries, k)
			}
		}
		c.mx.Unlock()
	}

}

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{entries: make(map[string]cacheEntry),}
	// possible wrong implementation of the reapLoop
	go newCache.reapLoop(interval)
	return newCache

}
