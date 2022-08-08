package LeetCode_54_SpiralMatrix

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m := len(matrix)
	n := len(matrix[0])
	res := make([]int, m*n)

	top, bottom := 0, m-1
	l, r := 0, n-1
	x, y := 0, 0

	for i := range res {
		res[i] = matrix[x][y]

		if x == top && y < r {
			y++
			// additional check top+1 != bottom
			//  example edge case if only 2 rows
			if y == r && top+1 != bottom {
				top++
			}
			continue
		}

		if y == r && x < bottom {
			x++
			// additional check r-1 != l
			if x == bottom && r-1 != l {
				r--
			}
			continue
		}

		if x == bottom && y > l {
			y--
			if y == l {
				bottom--
			}
			continue
		}

		if y == l && x > top {
			x--
			if x == top {
				l++
			}
			continue
		}
	}
	return res
}
