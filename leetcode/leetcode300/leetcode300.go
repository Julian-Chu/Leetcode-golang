package leetcode300

import "fmt"

func lengthOfLIS(nums []int) int {
	res := make([]int, len(nums))
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				res[i] = max(res[j]+1, res[i])
			}
		}
		maxLen = max(maxLen, res[i]+1)
	}
	fmt.Printf("%v", res)
	return maxLen
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

//func lengthOfLIS(nums []int) int {
//	res := make([]int, 0, len(nums))
//	for _, v := range nums {
//		at := sort.SearchInts(res, v)
//		if at == len(res) {
//			res = append(res, v)
//		} else if res[at] > v {
//			res[at] = v
//		}
//	}
//	return len(res)
//}

//func lengthOfLIS(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	res := []int{nums[0]}
//OUTER:
//	for i := 1; i < len(nums); i++ {
//		for _, v := range res {
//			if v == nums[i] {
//				continue OUTER
//			}
//		}
//		if nums[i] < res[0] {
//			res[0] = nums[i]
//		} else if nums[i] > res[len(res)-1] {
//			res = append(res, nums[i])
//		} else {
//			l, r := 0, len(res)-1
//			mid := (l + r) / 2
//			for r-l > 1 {
//				if nums[i] > res[mid] {
//					l = mid
//				} else {
//					r = mid
//				}
//				mid = (l + r) / 2
//			}
//			res[r] = nums[i]
//		}
//	}
//	//fmt.Printf("%v", res)
//	return len(res)
//}
