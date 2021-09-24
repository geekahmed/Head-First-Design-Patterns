package main

type cache struct {
	storage map[interface{}]interface{}
	evictionAlgo ICache
	capacity int
	maxCapacity int
}

func initCache(e ICache) *cache{
	storage := make(map[interface{}]interface{})
	return &cache{
		storage: storage,
		evictionAlgo: e,
		capacity: 0,
		maxCapacity: 10,

	}
}

func (c *cache) setEvictionAlgo(e ICache) {
    c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
    if c.capacity == c.maxCapacity {
        c.evict()
    }
    c.capacity++
    c.storage[key] = value
}

func (c *cache) get(key string) {
    delete(c.storage, key)
}

func (c *cache) evict() {
    c.evictionAlgo.evict(c)
    c.capacity--
}
