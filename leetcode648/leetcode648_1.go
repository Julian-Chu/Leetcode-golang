package leetcode648

import "strings"

type Trie struct {
	children [26]*Trie
	end      bool
}

func (t *Trie) Add(word string) {
	for _, char := range word {
		if t.children[char-'a'] == nil {
			t.children[char-'a'] = &Trie{}
		}
		t = t.children[char-'a']
	}
	t.end = true
}

func (t *Trie) SearchRoot(word string) string {
	for index, char := range word {
		if t.children[char-'a'] == nil {
			return ""
		}
		t = t.children[char-'a']
		if t.end {
			return word[:index+1]
		}
	}
	return ""
}

func replaceWords_1(dictionary []string, sentence string) string {
	trie := &Trie{}

	for _, word := range dictionary {
		trie.Add(word)
	}

	words := strings.Split(sentence, " ")
	var res []string
	for _, word := range words {
		root := trie.SearchRoot(word)
		if root != "" {
			res = append(res, root)
		} else {
			res = append(res, word)
		}
	}
	return strings.Join(res, " ")
}
