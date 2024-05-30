package pokecache

import "time"

type cacheEntry struct {
	createAt time.Time
	val      []byte
}
