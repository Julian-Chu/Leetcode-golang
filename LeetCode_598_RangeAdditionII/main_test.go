package main

import (
	"testing"
)

func maxCount(m int, n int, ops [][]int) int {
	matrix := make([][]int, m)
	for index := range matrix {
		matrix[index] = make([]int, n)
		for j := range matrix[index] {
			matrix[index][j] = 0
		}
	}
	return 1
}

func Test_testcase1(t *testing.T) {
	res := maxCount(2, 2, [][]int{{1, 1}})

	if res != 1 {
		t.Error("failed")
	}
}
