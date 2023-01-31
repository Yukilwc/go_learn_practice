package lru

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	// t.Fatalf("failed")
	gap := int64(0)
	lru := New(gap, nil)
	fmt.Println(lru, gap)
}
