package leetcode208

type Trie struct {
	son [26]*Trie
	end bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	size := len(word)
	for i := 0; i < size; i++ {
		idx := word[i] - 'a'
		if node.son[idx] == nil {
			node.son[idx] = &Trie{}
		}
		node = node.son[idx]
	}
	node.end = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	size := len(word)
	for i := 0; i < size; i++ {
		idx := word[i] - 'a'
		if node.son[idx] == nil {
			return false
		}
		node = node.son[idx]
	}
	return node.end
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	size := len(prefix)
	for i := 0; i < size; i++ {
		idx := prefix[i] - 'a'
		if node.son[idx] == nil {
			return false
		}
		node = node.son[idx]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
