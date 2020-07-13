package leetcode318

func maxProduct(words []string) int {
	n := len(words)
	if n <= 0 {
		return 0
	}

	count := 0

	for i := 0; i < n-1; i++ {
		m := make(map[int32]bool)
		len1 := len(words[i])
		for _, b := range words[i] {
			m[b] = true
		}

	loop:
		for j := i + 1; j < n; j++ {
			len2 := len(words[j])
			if len1*len2 <= count {
				continue
			}

			for _, b := range words[j] {
				if m[b] {
					continue loop
				}
			}
			count = len1 * len2
		}
	}
	return count
}
