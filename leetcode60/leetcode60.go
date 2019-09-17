package leetcode60

import (
	"strconv"
)

func getPermutation(n int, k int) string {
	res := make([]string, 0)
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}

	recur([]int{}, nums, &res)
	return res[k-1]
}

func recur(solution, nums []int, res *[]string) {
	if len(nums) == 0 {
		ret := ""
		for _, v := range solution {
			ret += strconv.Itoa(v)
		}
		*res = append(*res, ret)
		return
	}
	solution = solution[:len(solution):len(solution)]
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			recur(append(solution, nums[0]), append([]int{}, nums[1:]...), res)
			continue
		}

		recur(append(solution, nums[i]), append(append([]int{}, nums[0:i]...), nums[i+1:]...), res)
	}
}
