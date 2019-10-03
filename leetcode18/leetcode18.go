package leetcode18

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}
	sort.Ints(nums)

	if nums[0]*4 > target || nums[len(nums)-1]*4 < target {
		return [][]int{}
	}
	res := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, len(nums)-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]

				switch {
				case sum > target:
					r--
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				case sum < target:
					l++
					for l < r && nums[l] == nums[l-1] {
						l++
					}
				default:
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					r--
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				}
			}
		}
	}
	return res
}
