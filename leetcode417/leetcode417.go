package leetcode417

func pacificAtlantic(matrix [][]int) [][]int {
	m := len(matrix)
	if m == 0 {
		return [][]int{}
	}
	n := len(matrix[0])
	if n == 0 {
		return [][]int{}
	}
	const toPacific = 1
	const toAtlantic = 2
	flags := make([][]int, m)
	for i := range flags {
		flags[i] = make([]int, n)
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if flags[i][j] == 3 {
			return flags[i][j]
		}
		res := flags[i][j]
		if i == 0 || j == 0 {
			res |= toPacific
		}

		if i == m-1 || j == n-1 {
			res |= toAtlantic
		}

		if res == 3 {
			flags[i][j] = 3
			return 3
		}

		v := matrix[i][j]

		flags[i][j] = -1
		if i > 0 && matrix[i-1][j] <= v && flags[i-1][j] > -1 {
			res = res | dfs(i-1, j)
		}
		if j > 0 && matrix[i][j-1] <= v && flags[i][j-1] > -1 {
			res = res | dfs(i, j-1)
		}
		if i < m-1 && matrix[i+1][j] <= v && flags[i+1][j] > -1 {
			res = res | dfs(i+1, j)
		}
		if j < n-1 && matrix[i][j+1] <= v && flags[i][j+1] > -1 {
			res = res | dfs(i, j+1)
		}

		flags[i][j] = res
		return res
	}
	for i := range matrix {
		for j := range matrix[i] {
			dfs(i, j)
		}
	}

	for i := range matrix {
		for j := range matrix[i] {
			v := matrix[i][j]
			if i > 0 && matrix[i-1][j] >= v && flags[i-1][j] < flags[i][j] {
				flags[i-1][j] = flags[i][j]
			}
		}
	}

	var res [][]int
	for i := range flags {
		for j := range flags[i] {
			if flags[i][j] == 3 {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}
