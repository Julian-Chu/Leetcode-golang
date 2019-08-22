package factorialtrailingzeroes

func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}

	countOfFive := 0

	for n >= 5 {
		countOfFive += n / 5
		n = n / 5
	}

	return countOfFive
}
