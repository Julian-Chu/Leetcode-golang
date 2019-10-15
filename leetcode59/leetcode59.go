package leetcode59

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	top, bottom := 0, n-1
	l, r := 0, n-1
	i, j := 0, 0
	for k := 1; k <= n*n; k++ {
		res[i][j] = k

		if i == top && j < r {
			j++
			if j == r {
				top++
			}
			continue
		}

		if j == r && i < bottom {
			i++
			if i == bottom {
				r--
			}
			continue
		}

		if i == bottom && j > l {
			j--
			if j == l {
				bottom--
			}
			continue
		}

		if j == l && i > top {
			i--
			if i == top {
				l++
			}
			continue
		}

	}
	return res
}
