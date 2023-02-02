package geecache

import (
	"fmt"
	"log"
	"testing"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func TestGet(t *testing.T) {
	// 访问次数统计
	loadCounts := make(map[string]int, len(db))
	gee := NewGroup(
		"scores",
		2<<10,
		GetterFunc(
			func(key string) ([]byte, error) {
				log.Println("[SlowDB] search key", key)
				if v, ok := db[key]; ok {
					if _, ok := loadCounts[key]; !ok {
						// 计数器不存在此key
						loadCounts[key] = 0
					}
					loadCounts[key] += 1
					return []byte(v), nil
				}
				return nil, fmt.Errorf("%s not exist", key)
			},
		),
	)
	for k, v := range db {
		if view, err := gee.Get(k); err != nil || view.String() != v {
			t.Fatalf("failed to get value")
		}
		if _, err := gee.Get(k); err != nil || loadCounts[k] > 1 {
			t.Fatalf("cache %s miss", k)
		}
	}
	if view, err := gee.Get("unknow"); err == nil {
		t.Fatalf("the view of unknow should be empty,but %s got", view)
	}

}