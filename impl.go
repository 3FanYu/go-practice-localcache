package localcache

import (
	"sync"
	"time"
)

var (
	cacheInstance  *cache
	once           sync.Once
	expiredIn      = 30 * time.Second
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

func New() Cache {
	if cacheInstance == nil {
		once.Do(func() {
			cacheInstance = &cache{
				pool: make(map[string]value),
				nowFunc: func() time.Time {
					return time.Now().Add(expiredIn)
				},
			}
			go cacheInstance.spawnCacheChcker()
		})
		return cacheInstance
	}
	return cacheInstance
}

func (c *cache) Set(k string, v interface{}) (e error) {
	c.m.Lock()
	defer c.m.Unlock()
	expiredAt := c.nowFunc()
	c.pool[k] = value{data: v, expiredAt: expiredAt}
	e = nil
	return
}

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
