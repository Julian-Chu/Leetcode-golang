package leetcode313

func nthSuperUglyNumber(n int, primes []int) int {
	l := len(primes)
	candidates := make([]int, l)
	pos := make([]int, l)
	res := make([]int, n)
	copy(candidates, primes)
	res[0] = 1

	for i := 1; i < n; i++ {
		res[i] = min(candidates)
		for j := 0; j < l; j++ {
			if res[i] == candidates[j] {
				pos[j]++
				candidates[j] = res[pos[j]] * primes[j]
			}
		}
	}
	return res[n-1]
}

func min(cs []int) int {
	min := cs[0]
	for _, v := range cs {
		if min > v {
			min = v
		}
	}
	return min
}
