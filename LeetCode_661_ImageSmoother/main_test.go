package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestImageSmoother_First(t *testing.T) {

	matrix := [][]int{
		{0, 1},
		{1, 1},
	}

	res := imageSmoother(matrix)
	fmt.Println("res",res)
	target := [][]int{
		{0, 0},
		{0, 0},
	}
	if !reflect.DeepEqual(res, target) {
		t.Error("Failed")
	}
}
