package LeetCode_136_SingleNumber

import (
	"testing"
)

func singleNumber(nums []int) int {
	xor := 0
	for _, n := range nums {
		xor ^= n
	}
	return xor

	//ans := 0
	//var i uint

	//	count := 0
	//	bit := 1 << i
	//	for j := 0; j < len(nums); j++ {
	//		if nums[j]&bit != 0 {
	//			count++
	//		}
	//	}
	//	if count%2 != 0 {
	//		ans |= bit
	//	}
	//}
	//return ans
}

var cases = []struct {
	numbers []int
	num     int
}{
	{
		[]int{1},
		1,
	},
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

func BenchmarkSingleNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			singleNumber(c.numbers)
		}
	}
}
