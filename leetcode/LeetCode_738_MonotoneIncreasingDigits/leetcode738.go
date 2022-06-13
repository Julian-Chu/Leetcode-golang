package LeetCode_738_MonotoneIncreasingDigits

import "strconv"

func monotoneIncreasingDigits(n int) int {
	s := strconv.Itoa(n)
	bs := []byte(s)
	size := len(bs)

	for i := size - 1; i > 0; i-- {
		if bs[i-1] > bs[i] {
			bs[i-1]--
			for j := i; j < size; j++ {
				bs[j] = '9'
			}
		}
	}

	res, _ := strconv.Atoi(string(bs))
	return res
}
