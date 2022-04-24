package localcache

import (
	"sync"
	"time"
)

var (
	expiredIn = 30 * time.Second
)

type cache struct {
	pool    map[string]value
	nowFunc func() time.Time
	m       sync.Mutex
}

type value struct {
	data      interface{}
	expiredAt time.Time
}

// returns a new instance of Cache
func New() Cache {
	c := &cache{
		pool: make(map[string]value),
		nowFunc: func() time.Time {
			return time.Now().Add(expiredIn)
		},
	}
	go c.spawnCacheChcker()
	return c
}

// takes a key and a value and sets new cache
func (c *cache) Set(k string, v interface{}) (e error) {
	c.m.Lock()
	defer c.m.Unlock()
	expiredAt := c.nowFunc()
	c.pool[k] = value{data: v, expiredAt: expiredAt}
	e = nil
	return
}

// finds and returns a value from the pool using given key
func (c *cache) Get(k string) (v interface{}, e error) {
	c.m.Lock()
	defer c.m.Unlock()
	v = c.pool[k].data
	e = nil
	return
}

func (c *cache) spawnCacheChcker() {
	for {
		for k, v := range c.pool {
			if !(v.expiredAt.Before(time.Now())) {
				continue
			}
			c.evictCache(k)
		}
	}
}

func (c *cache) evictCache(k string) {
	c.m.Lock()
	defer c.m.Unlock()
	delete(c.pool, k)
}
