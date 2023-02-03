package geecache

import (
	"fmt"
	"go_test_init/geecache/consistenthash"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

const (
	defaultBasePath = "/_geecache/"
	defaultReplicas = 50
)

type HTTPPool struct {
	self        string
	basePath    string
	mu          sync.Mutex
	peers       *consistenthash.Map
	httpGetters map[string]*httpGetter
}

func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self, // 自身的地址
		basePath: defaultBasePath,
	}
}

// 添加个日志方法

func (p *HTTPPool) Log(format string, v ...any) {
	log.Printf("[Server %s] %s", p.self, fmt.Sprintf(format, v...))
}

func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		// 没按照前缀来
		panic("HTTPPool serving unexpectd path: " + r.URL.Path)
	}
	p.Log("%s %s", r.Method, r.URL.Path)
	// <basepath>/<groupname>/<key>
	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	groupName := parts[0]
	key := parts[1]
	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group "+groupName, http.StatusNotFound)
		return
	}
	view, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/octet-steam")
	w.Write(view.ByteSlice())
}

// 设置节点列表
func (p *HTTPPool) Set(peers ...string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	// 创建一致性hash
	p.peers = consistenthash.New(defaultReplicas, nil)
	// 添加节点到一致性哈希中
	p.peers.Add(peers...)
	// 对节点创建http访问表
	p.httpGetters = make(map[string]*httpGetter, len(peers))
	for _, peer := range peers {
		p.httpGetters[peer] = &httpGetter{baseURL: peer + p.basePath}
	}
}

// 实现PeerPicker接口，提供方法从key获取应该请求的节点
func (p *HTTPPool) PickPeer(key string) (PeerGetter, bool) {
	if peer := p.peers.Get(key); peer != "" && peer != p.self {
		p.Log("Peek peer %s", peer)
		return p.httpGetters[peer], true
	}
	return nil, false
}

// 客户端类
type httpGetter struct {
	baseURL string
}

func (h *httpGetter) Get(group string, key string) ([]byte, error) {
	url := fmt.Sprintf("%v/%v/%v", h.baseURL, url.QueryEscape(group), url.QueryEscape(key))
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned: %v", res.Status)
	}
	bytes, err := ioutil.ReadAll(res.Body)
	return bytes, nil
}

// 确保类型实现了接口

var _ PeerGetter = (*httpGetter)(nil)
