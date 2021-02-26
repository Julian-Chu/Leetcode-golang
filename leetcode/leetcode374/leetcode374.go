package leetcode374

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */
//func guessNumber(n int) int {
//	for i, j := 0, n; ; {
//		m := i + (j-i)/2
//		if ret := guess(m); ret == 1 {
//			i = m + 1
//		} else if ret == -1 {
//			j = m - 1
//		} else {
//			return m
//		}
//	}
//}

func guessNumber(n int) int {
	l, r := 0, n
	for l <= r {
		mid := l + (r-l)/2
		switch guess(mid) {
		case 1:
			l = mid + 1
		case -1:
			r = mid - 1
		default:
			return mid
		}
	}
	return 0
}

var pickup = 0

func guess(num int) int {
	switch {
	case num > pickup:
		return -1
	case num < pickup:
		return 1
	default:
		return 0
	}
}
