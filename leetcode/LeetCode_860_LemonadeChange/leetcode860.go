package LeetCode_860_LemonadeChange

func lemonadeChange(bills []int) bool {
	fiveDollars := 0
	tenDollars := 0

	for _, b := range bills {
		if b == 5 {
			fiveDollars++
			continue
		}

		if b == 10 {
			if fiveDollars > 0 {
				tenDollars++
				fiveDollars--
				continue
			}
			return false
		}

		if b == 20 {
			if tenDollars > 0 && fiveDollars > 0 {
				tenDollars--
				fiveDollars--
				continue
			}

			if fiveDollars >= 3 {
				fiveDollars -= 3
				continue
			}
			return false
		}
	}
	return true
}
