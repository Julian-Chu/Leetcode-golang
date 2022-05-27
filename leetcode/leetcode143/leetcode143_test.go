package leetcode143

import (
	"testing"

	"github.com/Julian-Chu/Leetcode-golang/utils"
)

func Test_reorderList(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1234->1423",
			args: args{
				ints: []int{1, 2, 3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(utils.IntsToListNode(tt.args.ints))
		})
	}
}
