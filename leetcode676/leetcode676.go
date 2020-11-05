package leetcode676

type MagicDictionary struct {
	lenDic map[int]map[string]bool
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{lenDic: make(map[int]map[string]bool)}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		if _, ok := this.lenDic[len(word)]; !ok {
			this.lenDic[len(word)] = make(map[string]bool)
		}
		this.lenDic[len(word)][word] = true
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	l := len(searchWord)
	wordMap, ok := this.lenDic[l]
	if !ok {
		return false
	}

	for word := range wordMap {
		diff := 0
		for i := range searchWord {
			if searchWord[i] != word[i] {
				diff++
			}
		}
		if diff == 1 {
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
