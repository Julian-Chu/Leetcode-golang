package leetcode393

func validUtf8(data []int) bool {
	cnt := 0
	for _, d := range data {
		if cnt == 0 {
			switch {
			case d>>3 == 30:
				cnt = 3
			case d>>4 == 14:
				cnt = 2
			case d>>5 == 6:
				cnt = 1
			case d>>7 > 0:
				return false
			}
		} else {
			if d>>6 != 2 {
				return false
			}
			cnt--
		}
	}
	return 0 == cnt
}
