package leetcode96

func numTrees(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 || n == 2 {
		return n
	}
	if n == 3 {
		return 5
	}

	sum := 0
	for i := 1; i <= n/2; i++ {
		sum += numTrees(i-1) * numTrees(n-i)
	}
	sum *= 2
	if n%2 == 1 {
		temp := numTrees(n / 2)
		sum += temp * temp
	}
	return sum
}
