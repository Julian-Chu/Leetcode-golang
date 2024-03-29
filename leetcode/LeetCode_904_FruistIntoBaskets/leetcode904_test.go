package LeetCode_904_FruistIntoBaskets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	tree []int
	ans  int
}{

	{
		[]int{0},
		1,
	},

	{
		[]int{1, 2, 1},
		3,
	},

	{
		[]int{0, 1, 2, 2},
		3,
	},

	{
		[]int{1, 2, 3, 2, 2},
		4,
	},

	{
		[]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4},
		5,
	},
}

func Test_totalFruit(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, totalFruit(tc.tree), "inout :%v", tc)
	}
}

func Benchmark_totalFruit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			totalFruit(tc.tree)
		}
	}
}
