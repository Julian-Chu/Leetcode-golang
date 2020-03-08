package leetcode204

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	isComposite := make([]bool, n)

	cnt := n / 2

	for i := 3; i*i < n; i += 2 {
		if isComposite[i] {
			continue
		}

		for j := i * i; j < n; j += 2 * i {
			if isComposite[j] {
				continue
			}
			cnt--
			isComposite[j] = true
		}
	}

	return cnt
}
