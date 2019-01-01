package LeetCode_137_SingleNumberII

import "testing"

func singleNumber(nums []int) int {
	return 0
}

var testcases = []struct {
	input  []int
	output int
}{
	{input: []int{2, 2, 3, 2}, output: 3},
}

func Test_Cases(t *testing.T) {
	for _, c := range testcases {
		if singleNumber(c.input) != c.output {
			t.Error()
		}
	}

}
