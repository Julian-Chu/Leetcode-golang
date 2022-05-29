package LeetCode_367_ValidPerfectSquare

//func isPerfectSquare(num int) bool {
//	if num == 0 {
//		return true
//	}
//	for i := 1; i*i <= num; i++ {
//		if num%i == 0 && num/i == i {
//			return true
//		}
//	}
//
//	return false
//}

func isPerfectSquare(num int) bool {
	if num <= 1 {
		return true
	}

	start, end := 0, num

	for start+1 < end {
		mid := start + (end-start)/2

		if mid*mid == num {
			return true
		}

		if mid*mid < num {
			start = mid
		} else {
			end = mid
		}
	}

	if start*start == num || end*end == num {
		return true
	}
	return false
}
