package leetcode146

import "container/list"

type LRUCache struct {
	capacity int
	lru      *list.List
	values   map[int]int
	keys     map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		lru:      list.New(),
		values:   make(map[int]int, capacity),
		keys:     make(map[int]*list.Element, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.values[key]; ok {
		this.lru.MoveToFront(this.keys[key])
		return v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.values[key]; ok {
		this.values[key] = value
		this.lru.MoveToFront(this.keys[key])
		return
	}

	if this.lru.Len() >= this.capacity {
		last := this.lru.Back()
		this.lru.Remove(last)
		delete(this.values, last.Value.(int))
		delete(this.keys, last.Value.(int))
	}

	e := this.lru.PushFront(key)
	this.values[key] = value
	this.keys[key] = e
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
