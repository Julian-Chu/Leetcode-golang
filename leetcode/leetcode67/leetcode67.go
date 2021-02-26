package leetcode67

func addBinary(a string, b string) string {
	n := len(a)
	if n < len(b) {
		n = len(b)
	}
	n++
	res := make([]byte, n)

	idx1 := len(a) - 1
	idx2 := len(b) - 1
	for i := n - 1; i >= 0; i-- {
		if idx1 >= 0 {
			res[i] += a[idx1] - '0'
			idx1--
		}
		if idx2 >= 0 {
			res[i] += b[idx2] - '0'
			idx2--
		}
		if res[i] >= 2 {
			temp := res[i]
			res[i] = temp % 2
			res[i-1] = temp / 2
		}
	}
	if res[0] == 0 {
		res = res[1:]
	}

	for i := range res {
		res[i] += '0'
	}

	return string(res)
}
