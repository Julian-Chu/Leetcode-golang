package leetcode397

func integerReplacement(n int) int {
	var cnt int
	return helper(n, cnt)
}

func helper(n int, cnt int) int {
	if n == 1 {
		return cnt
	}

	cnt++
	if n%2 == 0 {
		return helper(n/2, cnt)
	} else if n%4 == 1 || n == 3 {
		return helper(n-1, cnt)
	}

	return helper(n+1, cnt)

}
