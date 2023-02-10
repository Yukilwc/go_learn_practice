package gee

import "net/http"

// 声明一个函数类型，此函数处理请求
type HandleFunc func(http.ResponseWriter, *http.Request)

// 实现ServeHTTP的引擎
type Engine struct {
	router map[string]HandleFunc
}

// 提供引擎的构造函数

func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}
