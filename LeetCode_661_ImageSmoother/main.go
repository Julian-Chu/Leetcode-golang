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

	fmt.Println(imageSmoother(m))
}

func imageSmoother(M [][]int) [][]int {

	res := make([][]int, 0)

	for rowIndex, row := range M {
		avgs := make([]int, 0)
		rowMax := len(M)
		colMax := len(M[0])
		for colIndex, _ := range row {
			rowBegin := rowIndex - 1
			rowEnd := rowIndex + 1
			colBegin := colIndex - 1
			colEnd := colIndex + 1
			if rowIndex == 0 {
				rowBegin = 0
			}
			if rowIndex == rowMax {
				rowEnd = rowMax
			}
			if colIndex == 0 {
				colBegin = 0
			}
			if colIndex == colMax {
				colEnd = colMax
			}
			subSlice := M[rowBegin:rowEnd][colBegin:colEnd]
			count := 0
			sum := 0
			for _, row := range subSlice {
				for _, e := range row {
					sum += e
					count++
				}
			}
			result := int(math.Floor(float64(sum / count)))

			avgs = append(avgs, result)
		}
		res = append(res, avgs)
	}
	return res
}
