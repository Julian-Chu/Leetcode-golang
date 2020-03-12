package leetcode319

func bulbSwitch(n int) int {
	res := make([]int, n+1)

	for i := 1; i < n+1; i++ {
		for j := i; j < n+1; j += i {
			res[j] = res[j] ^ 1
		}
	}

	cnt := 0
	for i := range res {
		cnt += res[i]
	}
	return cnt
}
