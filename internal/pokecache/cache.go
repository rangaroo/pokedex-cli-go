package pokecache

import (
	"sync"
	"time"
)

type cacheField struct {
	createdAt  time.Time
	val        []byte
}

type Cache struct {
	data          map[string]cacheField
	mu            sync.RWMutex
	interval      time.Duration
	done          chan struct{}
	ticker        *time.Ticker
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		data: make(map[string]cacheField),
		mu: sync.RWMutex{},
		interval: interval,
		done: make(chan struct{}),
		ticker: time.NewTicker(interval),
	}

	go cache.reapLoop()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	cacheField := cacheField{
		createdAt: time.Now(),
		val: val,
	}

	c.mu.Lock()
	c.data[key] = cacheField
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	elem, exists := c.data[key]

	if exists {
		return elem.val, true
	} else {
		return []byte{}, false
	}
}

func (c *Cache) reap() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for key, val := range c.data {
		if now.Sub(val.createdAt) > c.interval {
			delete(c.data, key)
		}
	}
}

func (c *Cache) reapLoop() {
	for {
		select {
			case <-c.done:
				return
			case <-c.ticker.C:
				c.reap()
		}
	}
}
