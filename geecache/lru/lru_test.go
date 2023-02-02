package lru

import (
	"reflect"
	"testing"
)

type String string

func (s String) Len() int {
	return len(s)
}

// 测试添加和读取
func TestGet(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1", String("abc"))
	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "abc" {
		t.Fatalf("Get key1 fail")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("get key2 fail")
	}
}

// 测试自动移除
func TestRemoveoldest(t *testing.T) {
	k1, k2, k3 := "k1", "k2", "k3"
	v1, v2, v3 := String("v1"), String("v2"), String("v3")
	maxSize := int64(len(k1)) + int64(len(k2)) + int64(v1.Len()) + int64(v2.Len())
	lru := New(maxSize, nil)
	lru.Add(k1, v1)
	lru.Add(k2, v2)
	lru.Add(k3, v3)
	if _, ok := lru.Get("k1"); ok || lru.Len() != 2 {
		t.Fatalf("remove fail")
	}
}

func TestCallback(t *testing.T) {
	omitKeysList := []string{}
	callback := func(key string, value Value) {
		omitKeysList = append(omitKeysList, key)
	}
	lru := New(int64(10), callback)
	lru.Add("key1", String("123456"))
	lru.Add("k2", String("k2"))
	lru.Add("k3", String("k3"))
	lru.Add("k4", String("k4"))
	expect := []string{"key1", "k2"}
	if !reflect.DeepEqual(omitKeysList, expect) {
		t.Fatalf("callback fail")
	}
}
