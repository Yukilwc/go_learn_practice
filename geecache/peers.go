package geecache

// 节点选择接口

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// 缓存值获取接口
type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
