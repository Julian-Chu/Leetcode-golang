package leetcode1456

func maxVowels(s string, k int) int {
	count, max := 0, 0
	v, vowel := []byte{'a', 'e', 'i', 'o', 'u'}, make([]int, 26)

	for _, c := range v {
		vowel[c-'a'] = 1
	}

	for i, v := range s {
		if vowel[v-'a'] == 1 {
			count++
		}
		if i >= k && vowel[s[i-k]-'a'] == 1 {
			count--
		}
		if count > max {
			max = count
		}
	}
	return max
}
