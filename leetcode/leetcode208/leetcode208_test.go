package leetcode208

import "testing"

func Test_Trie(t *testing.T) {
	trie := Constructor()

	trie.Insert("apple")
	if trie.Search("apple") != true {
		t.Error()
	}

	if trie.Search("app") != false {
		t.Error()
	}

	if trie.StartsWith("app") != true {
		t.Error()
	}

	trie.Insert("app")
	if trie.Search("app") != true {
		t.Error()
	}
}
