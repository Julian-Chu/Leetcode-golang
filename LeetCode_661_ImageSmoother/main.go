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

	rowLength := len(M)
	colLength := len(M[0])
	res := make([][]int, rowLength)
	for index := range res {
		res[index] = make([]int, colLength)
	}

	for rowIndex, row := range M {
		for colIndex := range row {
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
			res[rowIndex][colIndex] = result
		}
	}
	return res
}
