package leetcode473

import "sort"

func makesquare(nums []int) bool {
	if len(nums) < 4 {
		return false
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	sum := 0
	for _, num := range nums {
		sum += num
	}

	squareLen := sum / 4
	if sum%4 != 0 || nums[len(nums)-1] > squareLen || nums[0] > squareLen {
		return false
	}

	visited := make([]bool, len(nums))
	var dfs func(int, int) bool
	dfs = func(index int, rest int) bool {
		if rest == 0 {
			return true
		}
		if rest < 0 {
			return false
		}

		for index < len(nums) {
			if !visited[index] {
				visited[index] = true
				if dfs(index, rest-nums[index]) {
					return true
				}
				visited[index] = false
			}
			index++
		}
		return false
	}

	res := dfs(0, squareLen) && dfs(0, squareLen) && dfs(0, squareLen) && dfs(0, squareLen)
	for _, v := range visited {
		if v == false {
			return false
		}
	}
	return res
}
