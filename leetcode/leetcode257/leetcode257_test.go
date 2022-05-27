package leetcode257

import (
	"testing"

	"github.com/Julian-Chu/Leetcode-golang/utils"
)

func Test_binaryTreePaths(t *testing.T) {
	type args struct {
		node []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "[1,2,3,null,5]",
			args: args{
				node: []int{1, 2, 3, utils.NULL, 5},
			},
			want: []string{"1->2->5", "1->3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compareStringSlice := func(s1, s2 []string) bool {
				if len(s1) != len(s2) {
					return false
				}

				for i := range s1 {
					flag := false
					for j := range s2 {
						if s1[i] == s2[j] {
							flag = true
							break
						}
					}
					if !flag {
						return false
					}
				}
				return true
			}
			if got := binaryTreePaths(utils.Ints2Tree(tt.args.node)); !compareStringSlice(got, tt.want) {
				t.Errorf("binaryTreePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
