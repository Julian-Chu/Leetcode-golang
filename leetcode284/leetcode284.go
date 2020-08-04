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
	cache *int
	iter  *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
	next := iter.next()
	return &PeekingIterator{
		cache: &next,
		iter:  iter,
	}
}

func (this *PeekingIterator) hasNext() bool {
	return this.cache != nil
}

func (this *PeekingIterator) next() int {
	res := *this.cache
	var next *int
	if this.iter.hasNext() {
		tmp := this.iter.next()
		next = &tmp
	}
	this.cache = next
	return res
}

func (this *PeekingIterator) peek() int {
	return *this.cache
}
