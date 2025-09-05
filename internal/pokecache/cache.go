package pokecache

import (
	"sync"
	"time
)

type cacheField struct {
	createdAt  time.Time
	val        []byte
}

// TODO: Use RDMutex
type Cache struct {
	cache     map[string]cacheEntry
	mu       sync.RDMutex
}

func NewCache(interval time.Duration) Cache {
	return Cache{
		cache: map[string]cacheEntry{},
		mu: sync.Mutex,
	}
}

func (c Cache) Add(key string, val []byte) {
	cacheField = cacheField{
		createdAt: time.Now(),
		val: val,
	}
	c.cache[key] = cacheField
}

func (c Cache) Get(key string) []byte, bool {
	elem, exists := c.cache[key]

	if exists {
		return elem.val, true
	} else {
		return []byte{}, false
	}
}


func (c Cache) reapLoop(interval time.Duration) {
	for key, elem:= range c.cache {

	}
}
