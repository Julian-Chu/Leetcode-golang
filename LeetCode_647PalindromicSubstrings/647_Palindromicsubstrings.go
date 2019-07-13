package LeetCode_647PalindromicSubstrings

func countSubstrings(s string) int {
	if len(s) < 2 {
		return 1
	}
	cnt := len(s) // one char

	for l := 2; l <= len(s); l++ {
		for i := 0; i+l-1 < len(s); i++ {
			slice := s[i : i+l]
			if isPalin(slice) {
				cnt++
			}
		}

	}
	return cnt
}

func isPalin(s string) bool {

	l := len(s)
	if l == 1 {
		return true
	}
	mid := l / 2
	b, e := mid, mid
	if l%2 == 0 {
		b, e = mid-1, mid
	}

	for e < l {
		if s[b] != s[e] {
			return false
		}
		b--
		e++
	}
	return true
}
