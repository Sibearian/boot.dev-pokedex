package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
    createdAt time.Time
    val	[]byte
}

type Cache struct {
    cache map[string]cacheEntry
    mu    sync.Mutex
}

func NewCache(duration time.Duration) *Cache {
    cache := Cache{}

    cache.mu.Lock()
    cache.cache = make(map[string]cacheEntry)
    cache.mu.Unlock()
    go cache.reapLoop(duration)

    return &cache
}

func (c *Cache) Add(key string, val []byte) {
    c.mu.Lock()
    defer c.mu.Unlock()

    c.cache[key] = cacheEntry{
	createdAt: time.Now(),
	val: val,
    }
}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()

    val, ok := c.cache[key]
    if !ok {
	return nil, false
    }

    return val.val, true
}

func (c *Cache) reapLoop(duration time.Duration) {
    ticker := time.NewTicker(duration)

    for {
	select {
	case <-ticker.C:
	    for url, val := range c.cache {
		if val.canBeDeleted(duration) {
		    c.mu.Lock()
		    delete(c.cache, url)
		    c.mu.Unlock()
		}
	    }
	}
    }
}

func (c *cacheEntry) canBeDeleted(duration time.Duration) bool {
    if duration < time.Since(c.createdAt) {
	return false
    }
    return true
}
