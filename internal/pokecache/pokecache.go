// pokecache package implements a simple in-memory cache system

package pokecache

import (
	"sync"
	"time"
)

// Cache is a thread-safe map used to cache data.
type Cache struct {
	cache map[string]cacheEntry // cache  is a map which contains the cached data
	mux   *sync.Mutex           // mux is a Mutex which is used for handling access from multiple routines
}

// cacheEntry is a single entity/item stored in the cache
type cacheEntry struct {
	createdAt time.Time // Time when the cache entry was created
	val       []byte    // Value of the cache entry
}

// NewCache function creates a new Cache.
func NewCache(interval time.Duration) Cache {
	// A Cache struct is being initialized with an empty cache and a new mutex.
	c := Cache{
		cache: make(map[string]cacheEntry), // Creating an empty map of string keys to cacheEntry values.
		mux:   &sync.Mutex{},               // Initializing a new mutex for handling concurrent access to the cache.
	}

	// Invoking the reapLoop method in a separate goroutine, which continually clears expired entries from the cache at the provided interval.
	go c.reapLoop(interval)

	// Returning the newly created cache.
	return c
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

// reapLoop is a method on Cache struct that runs a loop at specific intervals.
// Each time the interval elapses, the method invokes the 'reap' method.
func (c *Cache) reapLoop(interval time.Duration) {
	// Create a new ticker that triggers at intervals based on the specified duration.
	ticker := time.NewTicker(interval)

	// The loop will continuously wait for the ticker's channel to send a signal
	// It triggers once the time interval specified when creating the ticker elapses
	for range ticker.C {
		// After each trigger, call the reap method on the Cache. Pass the current time and the interval.
		c.reap(time.Now().UTC(), interval)
	}
}

// `reap` is a method on `Cache` struct which clears items from the cache that were created
// before a specified time duration (`last`). The `now` parameter represents the current time.
func (c *Cache) reap(now time.Time, last time.Duration) {
	// Locking the mutex to prevent concurrent read/write on the map `c.cache`
	c.mux.Lock()

	// Unlocking the mutex once this function finishes its execution. This is done
	// using the `defer` statement which ensures that the lock is released
	// even if an error occurs during the execution of this function.
	defer c.mux.Unlock()

	// Iterating over each key-value pair in the cache
	for k, v := range c.cache {
		// Deleting all the cache entries that were created before the specified duration
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}
