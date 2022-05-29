package leetcode69

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	start, end := 0, x

	for start+1 < end {
		mid := start + (end-start)>>1

		if mid*mid == x {
			return mid
		}

		if mid*mid < x {
			start = mid
		} else {
			end = mid
		}
	}

	return start
}

//func mySqrt(x int) int {
//	res := x
//	for res*res > x {
//		res = (res + x/res) / 2
//	}
//	return res
//}
