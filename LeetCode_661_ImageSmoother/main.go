package main

import (
	"fmt"
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
	for i := range res {
		res[i] = make([]int, colLength)
		for j := range res[i] {
			res[i][j] = calcAvg(M, i, j)
		}
	}

	return res
}

func calcAvg(M [][]int, i, j int) int {

	rowLength := len(M)
	colLength := len(M[0])
	count := 0
	sum := 0
	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if x >= 0 && x < rowLength && y >= 0 && y < colLength {
				count++
				sum += M[x][y]
			}
		}
	}

	return sum / count
}
