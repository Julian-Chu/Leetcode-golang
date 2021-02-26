package leetcode532

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	m, res := make(map[int]int, len(nums)), 0

	for _, val := range nums {
		m[val]++
	}

	if k == 0 {
		for _, cnt := range m {
			if cnt > 1 {
				res++
			}
		}
	} else {
		for val := range m {
			if m[val+k] > 0 {
				res++
			}
		}
	}

	return res
}
