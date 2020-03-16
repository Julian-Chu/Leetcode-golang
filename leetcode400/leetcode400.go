package leetcode400

func findNthDigit(n int) int {
	for n < 10 {
		return n
	}
	digitsOfNumber := 1

	totalsize := 0
	numbersWithSameDigits := 9
	for n > totalsize {
		totalsize += numbersWithSameDigits * digitsOfNumber
		numbersWithSameDigits *= 10
		digitsOfNumber++
		if n == totalsize {
			break
		}
		if n < totalsize {
			numbersWithSameDigits /= 10
			digitsOfNumber--
			totalsize -= numbersWithSameDigits * digitsOfNumber
			break
		}
	}

	start := 1
	for i := 0; i < digitsOfNumber-1; i++ {
		start *= 10
	}
	start--
	nthDigitInNumber := (n - totalsize) % digitsOfNumber
	if nthDigitInNumber == 0 {
		nthDigitInNumber = digitsOfNumber
	}
	number := (n-totalsize+digitsOfNumber-nthDigitInNumber)/digitsOfNumber + start

	if nthDigitInNumber == digitsOfNumber {
		return number % 10
	}
	for i := 0; i < digitsOfNumber-nthDigitInNumber; i++ {
		number /= 10
	}
	return number % 10
}
