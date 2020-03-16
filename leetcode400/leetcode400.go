package leetcode400

func findNthDigit(n int) int {
	count, digits := 9, 1
	num := 1
	for n-count*digits > 0 {
		n -= count * digits
		count *= 10
		digits++
		num *= 10
	}

	index := n % digits
	if index == 0 {
		index = digits
	}

	num += n / digits
	if index == digits {
		num--
	}

	for i := index; i < digits; i++ {
		num /= 10
	}
	return num % 10
}
