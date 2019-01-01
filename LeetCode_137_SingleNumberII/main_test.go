package LeetCode_137_SingleNumberII

import "testing"

func singleNumber(nums []int) int {
	singleNumber := 0
	for i := 0; i < 32; i++ {

		bit := 1 << uint(i)
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[j]&bit != 0 {
				count++
			}
		}
		if count%3 != 0 {
			singleNumber |= bit
		}
	}
	return singleNumber
}

var testcases = []struct {
	input  []int
	output int
}{
	{input: []int{2, 2, 3, 2}, output: 3},
	{input: []int{0, 1, 0, 1, 0, 1, 99}, output: 99},
}

func Test_Cases(t *testing.T) {
	for _, c := range testcases {
		if singleNumber(c.input) != c.output {
			t.Error()
		}
	}

}

func Benchmark_SingleNumberII(b *testing.B) {
	for _, c := range testcases {
		singleNumber(c.input)
	}
}
