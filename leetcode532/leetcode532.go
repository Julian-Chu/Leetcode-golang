package leetcode532

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	m := make(map[int]int)

	for i := range nums {
		m[nums[i]]++
	}

	cnt := 0
	for key := range m {
		if k == 0 {
			if m[key] > 1 {
				cnt++
			}
		} else {
			cnt += min(min(m[key], m[key+k]), 1)
		}
	}

	return cnt

}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
