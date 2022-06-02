package LeetCode_15_3Sum

import "sort"

// take big space by hashmap
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
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
	return res
}
