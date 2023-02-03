package geecache

import (
	"fmt"
	"log"
	"sync"
)

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

type Group struct {
	name      string
	getter    Getter
	mainCache cache
	peers     PeerPicker
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

// 创建一个group，并存储到groups中
func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
	}
	groups[name] = g
	return g
}

func GetGroup(name string) *Group {
	mu.RLock()
	defer mu.RUnlock()
	g := groups[name]
	return g
}

func (g *Group) RegisterPeers(peers PeerPicker) {
	if g.peers != nil {
		panic("RegisterPeerPicker called more than once")
	}
	g.peers = peers
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
	if g.peers != nil {
		// 看看这个key从哪个节点上拿
		if peer, ok := g.peers.PickPeer(key); ok {
			var err error
			if value, err := g.getFromPeer(peer, key); err == nil {
				return value, nil
			}
			log.Println("[GeeCache] failed to get from peer", err)
		}
	}
	return g.getLocally(key)
}
func (g *Group) getFromPeer(peer PeerGetter, key string) (ByteView, error) {
	bytes, err := peer.Get(g.name, key)
	if err != nil {
		return ByteView{}, err
	}
	return ByteView{b: bytes}, nil
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
