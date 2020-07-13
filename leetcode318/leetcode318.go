package leetcode318

func maxProduct(words []string) int {
	bits := make([]int, len(words))

	for i := range words {
		for _, ch := range words[i] {
			bits[i] |= 1 << uint32(ch-'a')
		}
	}

	cnt := 0

	n := len(words)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if bits[i]&bits[j] == 0 {
				product := len(words[i]) * len(words[j])
				if product > cnt {
					cnt = product
				}
			}
		}
	}
	return cnt
}
