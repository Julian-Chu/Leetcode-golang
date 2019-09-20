package leetcode77

func combine(n int, k int) [][]int {
	nums := make([]int, n)
	res := make([][]int, 0)

	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	recur(nums, []int{}, k, &res)
	return res

}
func recur(nums, solution []int, k int, res *[][]int) {
	if k == 0 {
		*res = append(*res, solution)
		return
	}

	if len(nums) < k {
		return
	}
	for i := range nums {
		solution = solution[:len(solution):len(solution)]
		recur(append([]int{}, nums[i+1:]...), append(solution, nums[i]), k-1, res)
	}

}
