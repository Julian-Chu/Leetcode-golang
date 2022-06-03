package LeetCode_541_ReverseStringII

func reverseStr(s string, k int) string {
	bs := []byte(s)

	// move 2k
	for i := 0; i < len(bs); i += 2 * k {
		// check the rest characters
		if i+k <= len(bs) {
			reverse(bs[i : i+k])
			continue
		}
		reverse(bs[i:])
	}

	return string(bs)
}

func reverse(bs []byte) {
	if len(bs) <= 1 {
		return
	}

	l, r := 0, len(bs)-1
	for l <= r {
		bs[l], bs[r] = bs[r], bs[l]
		l++
		r--
	}
}
