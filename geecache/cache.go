package geecache

import (
	"fmt"
	"go_test_init/geecache/lru"
	"sync"
)

type cache struct {
	mu         sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

// 增加了锁的add

func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}
	c.lru.Add(key, value)
}

// 增加了锁的get

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}
	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}
	return
}

func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}
	// 从缓存中取
	if v, ok := g.mainCache.get(key); ok {
		fmt.Println("hit cache")
		return v, nil
	}
	// 从数据源获取
	return g.load(key)
}

// 数据源载入方法
func (g *Group) load(key string) (ByteView, error) {
	return g.getLocally(key)
}

func (g *Group) getLocally(key string) (ByteView, error) {
	// 调用getter
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	// 拿到了数据，将其构造成ByteView
	value := ByteView{b: cloneBytes(bytes)}
	// 加入缓存
	g.populateCache(key, value)
	return value, nil
}

func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}
