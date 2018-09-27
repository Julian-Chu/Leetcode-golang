package main

import (
	"fmt"
	"reflect"
	"testing"
)

var tests = []struct {
	input  [][]int
	expect [][]int
}{
	{[][]int{{0, 1}, {1, 1}}, [][]int{{0, 0}, {0, 0}}},
	{[][]int{{1, 1}, {1, 1}}, [][]int{{1, 1}, {1, 1}}},
}

func TestImageSmoother(t *testing.T) {
	for _, testcase := range tests {

		matrix := testcase.input

		res := imageSmoother(matrix)
		target := testcase.expect
		if !reflect.DeepEqual(res, target) {
			t.Error("Failed")
		}
	}
}
func TestImageSmoother_First(t *testing.T) {

	matrix := [][]int{
		{0, 1},
		{1, 1},
	}

	res := imageSmoother(matrix)
	fmt.Println("res", res)
	target := [][]int{
		{0, 0},
		{0, 0},
	}
	if !reflect.DeepEqual(res, target) {
		t.Error("Failed")
	}
}
