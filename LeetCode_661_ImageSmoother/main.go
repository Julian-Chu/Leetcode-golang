package main

import (
	"fmt"
	"math"
)

func main() {
	m := [][]int{
		{0, 1},
		{1, 1},
	}
	fmt.Print(len(m))
	fmt.Println(imageSmoother(m))
}

func imageSmoother(M [][]int) [][]int {

	res := make([][]int, 0)

	for rowIndex, row := range M {
		avgs := make([]int, 0)
		rowLength := len(M)
		colLength := len(M[0])
		for colIndex := range row {
			// fmt.Println(rowIndex, colIndex)
			rowBegin := rowIndex - 1
			if rowIndex == 0 {
				rowBegin = 0
			}
			rowEnd := rowIndex + 1
			if rowIndex == rowLength-1 {
				rowEnd = rowIndex
			}
			colBegin := colIndex - 1
			if colIndex == 0 {
				colBegin = 0
			}
			colEnd := colIndex + 1
			if colIndex == colLength-1 {
				colEnd = colIndex
			}
			count := 0
			sum := 0
			for i := rowBegin; i <= rowEnd; i++ {
				for j := colBegin; j <= colEnd; j++ {
					count++
					sum += M[i][j]
				}
			}
			result := int(math.Floor(float64(sum / count)))
			avgs = append(avgs, result)
		}
		res = append(res, avgs)
	}
	return res
}
