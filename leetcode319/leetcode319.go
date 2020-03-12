package leetcode319

func bulbSwitch(n int) int {
	res := make([]bool, n+1)

	for i := 1; i < n+1; i++ {
		for j := i; j < n+1; j += i {
			res[j] = !res[j]
		}
	}

	cnt := 0
	for i := range res {
		if res[i] == true {
			cnt++
		}
	}
	return cnt
}
