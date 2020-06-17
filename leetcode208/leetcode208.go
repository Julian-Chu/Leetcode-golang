package leetcode208

type Trie struct {
	fullStr map[string]bool
	prefix  map[string]bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		fullStr: make(map[string]bool),
		prefix:  make(map[string]bool),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	this.fullStr[word] = true

	for i := 0; i < len(word); i++ {
		this.prefix[word[:i+1]] = true
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	return this.fullStr[word]
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.prefix[prefix]
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
