package leetcode284

//   Below is the interface for Iterator, which is already defined for you.

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return false
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	return -1
}

//

type PeekingIterator struct {
	cache        int
	iter         *Iterator
	hasNextCache bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	hasNextCache := iter.hasNext()
	cache := 0
	if hasNextCache {
		cache = iter.next()
	}
	return &PeekingIterator{
		cache:        cache,
		iter:         iter,
		hasNextCache: hasNextCache,
	}
}

func (this *PeekingIterator) hasNext() bool {
	return this.hasNextCache
}

func (this *PeekingIterator) next() int {
	res := this.cache
	if this.iter.hasNext() {
		this.cache = this.iter.next()
	} else {
		this.hasNextCache = false
	}
	return res
}

func (this *PeekingIterator) peek() int {
	return this.cache
}
