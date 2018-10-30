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
	{[][]int{{1, 2}, {2, 2}}, [][]int{{1, 1}, {1, 1}}},
	{[][]int{{1, 3}, {1, 3}}, [][]int{{2, 2}, {2, 2}}},
	{[][]int{{6, 2}, {2, 2}}, [][]int{{3, 3}, {3, 3}}},
	{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
	{[][]int{{0, 0, 0}, {0, 4, 0}, {0, 0, 0}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
}

func TestImageSmoother(t *testing.T) {
	for _, testcase := range tests {

		matrix := testcase.input

		res := imageSmoother(matrix)
		target := testcase.expect
		if !reflect.DeepEqual(res, target) {
			fmt.Println("res", res)
			fmt.Println("target", target)
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
	target := [][]int{
		{0, 0},
		{0, 0},
	}
	if !reflect.DeepEqual(res, target) {
		t.Error("Failed")
	}
}

func BenchmarkImageSmoother(b *testing.B) {
	matrix := [][]int{
		{0, 1},
		{1, 1},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		imageSmoother(matrix)
	}
}

func BenchmarkImageSmootherWithoutSubfunc(b *testing.B) {
	matrix := [][]int{
		{0, 1},
		{1, 1},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		imageSmootherWithoutSubfunc(matrix)
	}
}
