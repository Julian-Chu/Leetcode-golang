package leetcode228

import "fmt"

func summaryRanges(nums []int) []string {
	res := []string{}

	l := len(nums)
	if l == 0 {
		return res
	}

	begin := nums[0]
	str := ""

	for i := 0; i < l; i++ {
		if i == l-1 || nums[i]+1 != nums[i+1] {
			if nums[i] == begin {
				str = fmt.Sprintf("%d", begin)
			} else {
				str = fmt.Sprintf("%d->%d", begin, nums[i])
			}

			if i+1 < l {
				begin = nums[i+1]
			}

			res = append(res, str)
		}
	}
	return res
}

//func summaryRanges(nums []int) []string {
//	if len(nums) == 0 {
//		return []string{}
//	}
//
//	if len(nums) == 1 {
//		return []string{strconv.Itoa(nums[0])}
//	}
//
//	res := make([]string, 0)
//	start, end := nums[0], nums[0]
//	nums = append(nums, nums[len(nums)-1])
//	for i := 1; i < len(nums); i++ {
//		if nums[i]-nums[i-1] == 1 {
//			end = nums[i]
//			continue
//		}
//
//		if start == end {
//			res = append(res, strconv.Itoa(start))
//		} else {
//			res = append(res, fmt.Sprintf("%v->%v", start, end))
//		}
//		start, end = nums[i], nums[i]
//	}
//
//	return res
//}
