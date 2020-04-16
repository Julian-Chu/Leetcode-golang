package leetcode473

import "sort"

func makesquare(nums []int) bool {
	if len(nums) < 4 {
		return false
	}
	sort.Ints(nums)

	sum := 0
	m := make(map[int]int)
	for _, num := range nums {
		sum += num
		m[num]++
	}

	squareLen := sum / 4
	if sum%4 != 0 || nums[len(nums)-1] > squareLen {
		return false
	}

	keys := make([]int, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	sort.Ints(keys)
	l, r := 0, len(keys)-1
	for l < r {
		keys[l], keys[r] = keys[r], keys[l]
		l++
		r--
	}

	var dfs func(int, map[int]int) bool
	dfs = func(rest int, m map[int]int) bool {
		if rest == 0 {
			return true
		}
		if rest < 0 {
			return false
		}

		for _, key := range keys {
			if cnt, _ := m[key]; cnt > 0 {
				m[key]--
				if dfs(rest-key, m) {
					return true
				}
				m[key]++
			}
		}
		return false
	}

	return dfs(squareLen, m) && dfs(squareLen, m) && dfs(squareLen, m) && dfs(squareLen, m)
}
