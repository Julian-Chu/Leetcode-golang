package leetcode676

type MagicDictionary struct {
	keys []string
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{keys: make([]string, 0, 100)}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, w := range dictionary {
		this.keys = append(this.keys, w)
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	n := len(searchWord)

	for _, w := range this.keys {
		if len(w) != n || w == searchWord {
			continue
		}

		res := 0
		for i := range w {
			if w[i] != searchWord[i] {
				res++
			}
		}
		if res == 1 {
			return true
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
