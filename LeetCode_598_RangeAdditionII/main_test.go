package main

import (
	"testing"
)

func maxCount(m int, n int, ops [][]int) int {
	minRow := m
	minCol := n
	for _, row := range ops {
		if minRow > row[0] {
			minRow = row[0]
		}
		if minCol > row[1] {
			minCol = row[1]
		}
	}

	return minRow * minCol
}

func Test_testcase1(t *testing.T) {
	res := maxCount(2, 2, [][]int{{1, 1}})

	if res != 1 {
		t.Error("failed")
	}
}

func Test_testcase2(t *testing.T) {
	res := maxCount(2, 2, [][]int{{2, 2}})

	if res != 4 {
		t.Error("failed")
	}
}

func Test_testcase3(t *testing.T) {
	res := maxCount(3, 3, [][]int{{2, 2}, {3, 3}})

	if res != 4 {
		t.Error("failed")
	}
}
