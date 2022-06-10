package permutation

func rmDup_sameLevel() {
	nums := []int{1, 1, 1}
	var res [][]int
	var dfs func(nums, tmp []int, used []bool)
	dfs = func(nums, tmp []int, used []bool) {
		if len(tmp) == len(nums) {
			res = append(res, append([]int{}, tmp...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			// can't use same nums twice on the same level
			if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
				continue
			}

			tmp = append(tmp, nums[i])
			used[i] = true
			dfs(nums, tmp, used)
			used[i] = false
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(nums, make([]int, 0, len(nums)), make([]bool, len(nums)))
}

func rmDup_sameBranch() {
	nums := []int{1, 1, 1}
	var res [][]int
	var dfs func(nums, tmp []int, used []bool)
	dfs = func(nums, tmp []int, used []bool) {
		if len(tmp) == len(nums) {
			res = append(res, append([]int{}, tmp...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			// can't use same nums twice through the same branch
			if i > 0 && nums[i] == nums[i-1] && used[i-1] == true {
				continue
			}

			tmp = append(tmp, nums[i])
			used[i] = true
			dfs(nums, tmp, used)
			used[i] = false
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(nums, make([]int, 0, len(nums)), make([]bool, len(nums)))
}
