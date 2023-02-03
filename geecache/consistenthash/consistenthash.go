package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type Map struct {
	hash Hash
	// 虚拟节点倍数
	replicas int
	// 哈希环
	keys []int
	// 虚拟/真实节点映射表 key是虚拟节点哈希值，value是真实节点名称
	hashMap map[int]string
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// 添加节点
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

// 拿到距离缓存最近的节点
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		// 没有缓存服务器
		return ""
	}
	hash := int(m.hash([]byte(key)))
	//sort.Search机制: Search returns the first true index. If there is no such index, Search returns n.
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	// 此处取余为了找不到时，idx为length，转为0
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
