package leetcode393

func validUtf8(data []int) bool {
	var run int
	for _, d := range data {
		if run == 0 {
			switch {
			case d&0b10000000 == 0b0:
			case d&0b11111000 == 0b11110000:
				run = 3
			case d&0b11110000 == 0b11100000:
				run = 2
			case d&0b11100000 == 0b11000000:
				run = 1
			default:
				return false
			}
		} else {
			if d&0b11000000 == 0b10000000 {
				run--
			} else {
				return false
			}
		}
	}
	return run == 0
}
