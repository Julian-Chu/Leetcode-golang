package leetcode338

func countBits(num int) []int {
	res := make([]int, num+1)

	for i := 0; i <= num; i++ {
		n := i
		for n > 0 {
			if n&1 == 1 {
				res[i]++
			}
			n = n >> 1
		}
	}
	return res
}
