package LeetCode_136_SingleNumber

import "testing"

func singleNumber(nums []int) int {
	xor := 0
	for _, n := range nums {
		xor ^= n
	}
	return xor
}

var cases = []struct {
	numbers []int
	num     int
}{
	{
		[]int{2, 2, 1},
		1,
	},
	{
		[]int{4, 1, 2, 1, 2},
		4,
	},
}

func TestCases(t *testing.T) {
	for _, c := range cases {
		if singleNumber(c.numbers) != c.num {
			t.Error("Failed")
		}
	}

}
