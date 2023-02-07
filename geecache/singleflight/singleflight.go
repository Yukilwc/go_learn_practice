package singleflight

import (
	"sync"
)

// 代表正在进行中，或已经结束的请求。使用 sync.WaitGroup 锁避免重入。
type call struct {
	wg  sync.WaitGroup
	val any
	err error
}

// 是 singleflight 的主数据结构，管理不同 key 的请求(call)。
type Group struct {
	mu sync.Mutex // 对m操作加锁
	m  map[string]*call
}

func (g *Group) Do(key string, fn func() (any, error)) (any, error) {
	g.mu.Lock()
	// 给m分配空间
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	// 看看进来的这个key，是否正在获取
	if c, ok := g.m[key]; ok {
		// 正在获取，那就不动了，解锁就行，然后等待
		g.mu.Unlock()
		c.wg.Wait()
		// 好了，返回等待的这个call的结果
		return c.val, c.err
	}
	// 这个key还没在获取，那就要自己去获取下
	// 先建立一个call
	c := new(call)
	// 加一次
	c.wg.Add(1)
	// 注册进来这个key对应的call，表示正在获取了
	g.m[key] = c
	g.mu.Unlock()
	// 获取值
	c.val, c.err = fn()
	// 解除其它等待
	c.wg.Done()
	g.mu.Lock()
	// 删除表示离开时调用状态
	delete(g.m, key)
	g.mu.Unlock()
	return c.val, c.err
}
