package leetcode146

import (
	"testing"
)

func TestLRUCache1(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	if cache.Get(1) != 1 {
		t.Error()
	}

	cache.Put(3, 3)
	if cache.Get(2) != -1 {
		t.Error()
	}

	cache.Put(4, 4)
	if cache.Get(1) != -1 {
		t.Error()
	}

	if cache.Get(3) != 3 {
		t.Error()
	}
	if cache.Get(4) != 4 {
		t.Error()
	}
}
func TestLRUCache2(t *testing.T) {
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(2, 2)
	if cache.Get(2) != 2 {
		t.Error()
	}

	cache.Put(1, 1)
	cache.Put(4, 1)
	if cache.Get(2) != -1 {
		t.Error()
	}

}
