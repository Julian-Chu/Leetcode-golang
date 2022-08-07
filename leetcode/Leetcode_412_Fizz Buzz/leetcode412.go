package Leetcode_412_Fizz_Buzz

import "strconv"

func fizzBuzz(n int) []string {
	if n == 0 {
		return nil
	}

	var res []string
	for i := 1; i < n+1; i++ {
		switch {
		case i%15 == 0:
			res = append(res, "FizzBuzz")
		case i%3 == 0:
			res = append(res, "Fizz")
		case i%5 == 0:
			res = append(res, "Buzz")
		default:
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}
