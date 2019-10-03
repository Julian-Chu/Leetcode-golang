package leetcode15

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			switch {
			case sum > 0:
				r--
			case sum < 0:
				l++
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				if nums[l] != 0 {
					l++
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					break
				}

				for l < r && nums[l] == 0 {
					l++
				}
			}
		}
	}
	return res
}
