package leetcode387

func firstUniqChar(s string) int {
	m := make([]int, 26)
	for _, v := range s {
		m[v-'a']++
	}

	for k, v := range s {
		if m[v-'a'] == 1 {
			return k
		}
	}

	return -1
}
