package LeetCode_216_CombinationSumIII

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	temp := make([]int, k+1)
	used := make([]bool, 10)

	var dfs func(int, int)

	dfs = func(idx int, remain int) {
		if idx == k {
			if remain < 10 && !used[remain] {
				temp[idx] = remain
				t := make([]int, k)
				copy(t, temp[1:])
				res = append(res, t)
			}
			return
		}

		for i := temp[idx-1] + 1; i < 10; i++ {
			if remain-i < i {
				return
			}
			used[i] = true
			temp[idx] = i
			dfs(idx+1, remain-i)
			used[i] = false
		}
	}
	dfs(1, n)
	return res
}

func combinationSum3_1(k int, n int) [][]int {
	var res [][]int
	var dfs func(k, n, start int, comb []int)
	dfs = func(k, sum, start int, comb []int) {
		if len(comb) == k && sum == 0 {
			res = append(res, append([]int{}, comb...))
			return
		}

		if len(comb) >= k || sum <= 0 {
			return
		}

		for i := start; i < 10; i++ {
			comb = append(comb, i)
			dfs(k, sum-i, i+1, comb)
			comb = comb[:len(comb)-1]
		}
	}

	dfs(k, n, 1, nil)
	return res
}

func combinationSum3_2(k int, n int) [][]int {
	var res [][]int
	comb := make([]int, k)
	var dfs func(combIdx, sum, start int)
	dfs = func(combIdx, sum, start int) {
		if sum == n && combIdx == k {
			res = append(res, append([]int{}, comb...))
			return
		}

		if combIdx >= k || sum >= n {
			return
		}

		for i := start; i < 10; i++ {
			comb[combIdx] = i
			dfs(combIdx+1, sum+i, i+1)
		}
	}

	dfs(0, 0, 1)
	return res
}
