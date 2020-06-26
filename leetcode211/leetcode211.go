package leetcode211

type alphabet struct {
	alphabets [26]*alphabet
	end       bool
}

type WordDictionary struct {
	alphabets [26]*alphabet
	end       bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		alphabets: [26]*alphabet{},
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	nodes := &this.alphabets
	for i := 0; i < len(word); i++ {
		end := false
		if i == len(word)-1 {
			end = true
		}
		idx := word[i] - 'a'
		if nodes[idx] == nil {
			nodes[idx] = &alphabet{
				alphabets: [26]*alphabet{},
				end:       end,
			}
		} else if !nodes[idx].end {
			nodes[idx].end = end
		}
		nodes = &nodes[idx].alphabets
	}
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	queue := make([]*alphabet, 0)
	if word[0] == '.' {
		for _, alphabet := range this.alphabets {
			if alphabet == nil {
				continue
			}
			if len(word) == 1 && alphabet.end {
				return true
			}
			queue = append(queue, alphabet)
		}
	} else {
		idx := word[0] - 'a'
		if this.alphabets[idx] == nil {
			return false
		} else if len(word) == 1 && this.alphabets[idx].end {
			return true
		}
		queue = append(queue, this.alphabets[idx])
	}
	for i := 1; i < len(word); i++ {
		next_q := make([]*alphabet, 0)
		if word[i] == '.' {
			for _, prevAlphabet := range queue {
				if prevAlphabet == nil {
					continue
				}
				for _, alphabet := range prevAlphabet.alphabets {
					if alphabet == nil {
						continue
					}
					if i == len(word)-1 && alphabet.end {
						return true
					}
					next_q = append(next_q, alphabet)
				}
			}
		} else {
			idx := word[i] - 'a'
			for _, prevAlphabet := range queue {
				if prevAlphabet == nil || prevAlphabet.alphabets[idx] == nil {
					continue
				}
				alphabet := prevAlphabet.alphabets[idx]
				if i == len(word)-1 && alphabet.end == true {
					return true
				}
				next_q = append(next_q, alphabet)
			}
		}
		queue = next_q
	}

	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
