package leetcode698

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	size := len(nums)
	sum := 0
	max := nums[0]
	for _, n := range nums {
		sum += n
		if max < n {
			max = n
		}
	}

	if sum%k != 0 || sum/k < max {
		return false
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	isUsed := make([]bool, size)
	return dfs(nums, isUsed, k, 0, 0, sum/k)
}

func dfs(nums []int, isUsed []bool, k, start, sum, target int) bool {
	if k == 1 {
		return true
	}

	if sum == target {
		return dfs(nums, isUsed, k-1, 0, 0, target)
	}

	for i := start; i < len(nums); i++ {
		if !isUsed[i] && sum+nums[i] <= target {
			isUsed[i] = true
			if dfs(nums, isUsed, k, i+1, sum+nums[i], target) {
				return true
			}
			isUsed[i] = false
		}
	}
	return false
}
