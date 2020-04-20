package leetcode441

func arrangeCoins(n int) int {
	calcSumFromZero := func(num int) int {
		return (num + 1) * num / 2
	}

	l, r := 0, n

	for l <= r {
		mid := l + (r-l)/2
		sum := calcSumFromZero(mid)

		if sum > n {
			r = mid - 1
		} else if sum < n {
			l = mid + 1
		} else {
			return mid
		}

		//if l+1 == r {
		//	break
		//}
	}
	return r
}
