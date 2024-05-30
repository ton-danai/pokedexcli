package pokecache

import "time"

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		createAt: time.Now().UTC(),
		val:      val,
	}
}
