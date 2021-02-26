package leetcode166

import (
	"fmt"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	if numerator*denominator < 0 {
		return "-" + fractionToDecimal(-numerator, denominator)
	}
	v := numerator / denominator
	r := numerator % denominator
	if r == 0 {
		return fmt.Sprintf("%d", v)
	}

	rec := make(map[int]int)
	idx := 0
	digits := make([]byte, 0)
	for {
		if i, ok := rec[r]; ok {
			return fmt.Sprintf("%d.%s(%s)", v, digits[:i], digits[i:])
		}
		rec[r] = idx
		r *= 10
		digits = append(digits, byte(r/denominator)+'0')
		idx++
		r %= denominator

		if r == 0 {
			str := fmt.Sprintf("%d.%s", v, string(digits))
			return str
		}
	}
}
