package perfectNumber

func checkPerfectNumber(num int) bool {
	if num%2 != 0 || num == 0 {
		return false
	}

	sum := 0
	i := num / 2
	for i > 0 {
		if i == 1 {
			sum++
			i--
		} else if num%i == 0 {
			sum += i
			if i%2 == 0 {
				i = i / 2
			} else {
				i = (i + 1) / 2
			}
		} else {
			i--
		}
	}

	if sum == num {
		return true
	} else {
		return false
	}
}
