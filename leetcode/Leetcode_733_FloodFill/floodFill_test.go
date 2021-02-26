package floodFill

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	source := [][]int{
		{1, 1},
		{1, 0},
	}

	startingPixelRow := 1
	stsrtingPixelCol := 1
	newColor := 2
	expect := [][]int{
		{2, 2},
		{2, 0},
	}

	res := floodFill(source, startingPixelRow, stsrtingPixelCol, newColor)
	x := reflect.DeepEqual(res, expect)
	fmt.Println(x)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Not Equal")
	}
}
