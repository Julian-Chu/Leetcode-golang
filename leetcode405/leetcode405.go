package leetcode405

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	n := uint32(num)
	idx := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	res := make([]byte, 0, 8)
	for n != 0 {
		res = append([]byte{idx[n&15]}, res...)
		n >>= 4
	}

	return string(res)
}
