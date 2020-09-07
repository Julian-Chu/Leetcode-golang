package leetcode659

func isPossible(nums []int) bool {
	size := len(nums)
	count := make(map[int]int, size)
	next := make(map[int]int, size)

	for i := range nums {
		count[nums[i]]++
	}

	for _, n := range nums {
		if count[n] == 0 {
			continue
		}

		count[n]--

		if next[n] > 0 {
			next[n]--
			next[n+1]++
		} else if count[n+1] > 0 && count[n+2] > 0 {
			count[n+1]--
			count[n+2]--
			next[n+3]++
		} else {
			return false
		}
	}

	return true
}
