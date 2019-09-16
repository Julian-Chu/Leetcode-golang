package leetcode47

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums)
	fmt.Println(nums)
	res := make([][]int, 0)

	rec(nums, []int{}, &res)
	return res
}

func rec(nums, solution []int, res *[][]int) {
	if len(nums) == 0 {
		*res = append(*res, solution)
		return
	}

	mapper := make(map[int]bool)

	for i := range nums {
		if v := mapper[nums[i]]; v {
			continue
		}
		mapper[nums[i]] = true
		newNums := append([]int{}, nums[1:]...)
		if i != 0 {
			newNums = append(append([]int{}, nums[0:i]...), nums[i+1:]...)
		}

		rec(newNums, append(solution, nums[i]), res)
	}
}
