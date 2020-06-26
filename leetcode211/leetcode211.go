package leetcode211

type WordDictionary struct {
	items map[int][]string
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{items: map[int][]string{}}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	l := len(word)
	if list, ok := this.items[l]; ok {
		this.items[l] = append(list, word)
	} else {
		this.items[l] = []string{word}
	}
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	l := len(word)
	if list, ok := this.items[l]; ok {
		for _, w := range list {
			i := 0
			for i < l {
				if w[i] != word[i] && word[i] != '.' {
					break
				}
				i++
			}
			if i == l {
				return true
			}
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
