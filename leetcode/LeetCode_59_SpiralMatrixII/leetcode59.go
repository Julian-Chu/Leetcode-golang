package LeetCode_59_SpiralMatrixII

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

//func generateMatrix(n int) [][]int {
//	matrix:= make([][]int, n)
//
//	for i := range matrix{
//		matrix[i] = make([]int,n)
//	}
//
//	i, j:= 0,0
//	left, right:= 0, n-1
//	top, bottom:= 0, n-1
//
//	for num:=1; num <= n*n; num++{
//		matrix[j][i] = num
//		if j  == top && i < right{
//			i++
//			if i == right{
//				top++
//			}
//			continue
//		}
//
//		if i == right && j < bottom{
//			j++
//			if j == bottom{
//				right--
//			}
//			continue
//		}
//
//		if j==bottom && i > left{
//			i--
//			if i==left{
//				bottom--
//				continue
//			}
//			continue
//		}
//
//		if i== left && j > top{
//			j--
//			if j==top{
//				left++
//			}
//			continue
//		}
//	}
//
//	return matrix
//}
