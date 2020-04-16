package leetcode473

import "sort"

func makesquare(nums []int) bool {
	if len(nums) < 4 {
		return false
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	avg := sum / 4
	if sum%4 != 0 || nums[len(nums)-1] > avg || nums[0] > avg {
		return false
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	edges := make([]int, 4)
	var dfs func(int, int) bool
	dfs = func(index int, target int) bool {
		if index == len(nums) {
			if edges[0] == target && edges[1] == target && edges[2] == target {
				return true
			}
			return false
		}

		for i := 0; i < 4; i++ {
			if edges[i]+nums[index] > target {
				continue
			}
			edges[i] += nums[index]
			if dfs(index+1, target) {
				return true
			}
			edges[i] -= nums[index]
		}
		return false
	}

	return dfs(0, avg)
}
