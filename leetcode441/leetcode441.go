package leetcode441

import "math"

func arrangeCoins(n int) int {
	return int(math.Sqrt(float64(1+8*n))-1) / 2
}
