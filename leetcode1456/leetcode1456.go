package leetcode1456

func maxVowels(s string, k int) int {
	n := len(s)

	l, max := 0, 0

	m := map[uint8]int{
		'a': 1,
		'e': 1,
		'i': 1,
		'o': 1,
		'u': 1,
	}

	for i := 0; i < n; i++ {
		l += m[s[i]]

		if i-k >= 0 {
			l -= m[s[i-k]]
		}

		if l > max {
			max = l
		}
	}
	return max
}
