package leetcode300

import (
	"fmt"
	"sort"
)

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

func lengthOfLIS_1(nums []int) int {
	res := make([]int, 0, len(nums))
	for _, v := range nums {
		at := sort.SearchInts(res, v)
		if at == len(res) {
			res = append(res, v)
		} else if res[at] > v {
			res[at] = v
		}
	}
	return len(res)
}

func lengthOfLIS_2(nums []int) int {
	sub := []int{}
	for _, num := range nums {
		if len(sub) == 0 || sub[len(sub)-1] < num {
			sub = append(sub, num)
		} else {
			idx := bisect(sub, num)
			sub[idx] = num
		}
	}
	return len(sub)

}

func bisect(sub []int, num int) int {
	if len(sub) == 0 {
		return 0
	}

	start, end := 0, len(sub)-1

	for start+1 < end {
		mid := start + (end-start)>>1
		if sub[mid] >= num {
			end = mid
		} else {
			start = mid
		}
	}

	if sub[start] >= num {
		return start
	}
	return end
}
