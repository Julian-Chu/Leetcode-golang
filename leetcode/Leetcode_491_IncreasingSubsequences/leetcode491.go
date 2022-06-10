package Leetcode_491_IncreasingSubsequences

func findSubsequences(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	var res [][]int
	tmp := make([]int, 0, len(nums))
	var dfs func(nums, tmp []int, start int)
	dfs = func(nums, tmp []int, start int) {
		if len(tmp) > 1 {
			res = append(res, append([]int{}, tmp...))
		}

		if start == len(nums) {
			return
		}

		set := make(map[int]bool)
		for i := start; i < len(nums); i++ {
			// if used on the same level, skip
			// this is not strict increasing sequence,
			// and can't be sorted because of subsequence
			if len(tmp) > 0 && nums[i] < tmp[len(tmp)-1] || set[nums[i]] {
				continue
			}

			tmp = append(tmp, nums[i])
			// nums[i] on the same level has been used
			set[nums[i]] = true
			dfs(nums, tmp, i+1)
			tmp = tmp[:len(tmp)-1]
		}
	}

	dfs(nums, tmp, 0)
	return res
}
