package leetcode228

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}

	res := make([]string, 0)
	start, end := nums[0], nums[0]
	nums = append(nums, nums[len(nums)-1])
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			end = nums[i]
			continue
		}

		if start == end {
			res = append(res, strconv.Itoa(start))
		} else {
			res = append(res, fmt.Sprintf("%v->%v", start, end))
		}
		start, end = nums[i], nums[i]
	}

	return res
}
