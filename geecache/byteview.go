package geecache

// 此文件构造一个高兼容的值类型

// 缓存值的类型，定义为此格式能支持各种类别得缓存
type ByteView struct {
	b []byte
}

// 让其实现Value接口
func (vp ByteView) Len() int {
	return len(vp.b)
}

// 返回一份b的拷贝 是为了让b只读
func (v ByteView) ByteSlice() []byte {
	c := make([]byte, len(v.b))
	copy(c, v.b)
	return c
}

// 返回字节切片的字符串

func (v ByteView) String() string {
	return string(v.b)
}

// 工具函数
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
