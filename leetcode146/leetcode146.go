package leetcode146

type LRUCache struct {
	cache    map[int]int
	capacity int
	queue    []int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    make(map[int]int, 0),
		queue:    make([]int, 0, capacity),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		for i := range this.queue {
			if this.queue[i] == key {
				q := make([]int, 1, this.capacity)
				q[0] = key
				q = append(q, this.queue[:i]...)
				q = append(q, this.queue[i+1:]...)
				this.queue = q
			}
		}

		return v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok {
		for i := range this.queue {
			if this.queue[i] == key {
				q := make([]int, 1, this.capacity)
				q[0] = key
				q = append(q, this.queue[:i]...)
				q = append(q, this.queue[i+1:]...)
				this.queue = q
			}
		}
		this.cache[key] = value
		return
	}
	if len(this.cache) == this.capacity {
		size := len(this.queue)
		lastKey := this.queue[size-1]
		this.queue = this.queue[:size-1]
		delete(this.cache, lastKey)
	}
	this.cache[key] = value
	q := make([]int, 1+len(this.queue), this.capacity)
	q[0] = key
	copy(q[1:], this.queue)
	this.queue = q
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
