package LeetCode_268_MissingNumber

import (
	"sort"
	"testing"
)

func missingNumber(nums []int) int {
	sort.Ints(nums)
	n := nums[len(nums)-1]
	for i := 0; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}
	return n + 1
}

var cases = []struct {
	numbers       []int
	missingNumber int
}{
	{[]int{3, 0, 1},
		2,
	},
	{
		[]int{9, 6, 4, 2, 3, 5, 7, 0, 1},
		8,
	},
	{
		[]int{0},
		1,
	},
}

func Test_MissingNumber(t *testing.T) {
	for _, c := range cases {
		if missingNumber(c.numbers) != c.missingNumber {
			t.Error("Failed")
		}
	}
}
