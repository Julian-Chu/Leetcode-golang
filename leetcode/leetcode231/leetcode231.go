package leetcode231

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}
	for n > 1 {
		if n%2 == 1 {
			return false
		}
		n /= 2
	}
	return true
}

//func isPowerOfTwo(n int) bool {
//	if n < 0 {
//		return false
//	}
//	cnt := 0
//	for i := 0; i < 32; i++ {
//		if n&1 == 1 {
//			cnt++
//		}
//		n = n >> 1
//	}
//	return cnt == 1
//}
