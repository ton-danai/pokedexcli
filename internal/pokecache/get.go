package pokecache

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	data, ok := c.cache[key]
	return data.val, ok
}
