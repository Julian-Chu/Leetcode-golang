package LeetCode_344_ReverseString

func reverseString(s []byte) {
	n := len(s)

	l, r := 0, n-1

	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}

}

func reverseString1(s []byte) {
	if len(s) <= 1 {
		return
	}

	start, end := 0, len(s)-1
	// xor swap: need two different memory space, so start can't equal to end
	for start < end {
		s[start] ^= s[end]
		s[end] ^= s[start]
		s[start] ^= s[end]
		start++
		end--
	}

	return
}
