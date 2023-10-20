// pokecache package implements a simple in-memory cache system

package pokecache

import (
	"sync"
	"time"
)

// Cache is a thread-safe map used to cache data.
type Cache struct {
	cache map[string]cacheEntry // Cache variable is a map which contains the cached data
	mux   *sync.Mutex           // Mutex is used for handling access from multiple routines
}

// cacheEntry is a single entity/item stored in the cache
type cacheEntry struct {
	createdAt time.Time // Time when the cache entry was created
	val       []byte    // Value of the cache entry
}

// NewCache function creates a new cache.
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry), // initialize cache map
		mux:   &sync.Mutex{},               // initialize mutex lock
	}

	return c // return new cache
}

// Add method adds a new entry to the cache. It is thread-safe, i.e.,
// it permits concurrent access to the cache.
func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()         // Lock before writing to the cache
	defer c.mux.Unlock() // Unlock after writing to the cache is complete
	c.cache[key] = cacheEntry{
		createdAt: time.Now(), // set creation time
		val:       value,      // store value
	}
}

// Get method fetches an entry from the cache. It is thread-safe, i.e.,
// it permits concurrent access to the cache.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()         // Lock before reading the cache
	defer c.mux.Unlock() // Unlock after reading the cache is complete
	val, ok := c.cache[key]
	return val.val, ok // returns value and a boolean indicating if a value was found
}

// reapLoop is responsible for periodically checking and deleting expired cache entries
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval) // call reap method to remove expired entries
	}
}

// reap method checks for expired cache entries and deletes them. This method
// must be safely used in a concurrent context.
func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()         // Lock before modifying the cache
	defer c.mux.Unlock() // Unlock after modification is complete
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k) // delete expired entries
		}
	}
}
