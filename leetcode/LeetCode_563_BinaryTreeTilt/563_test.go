package leetcode563

import (
	"testing"
)

func Test_findTilt(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case1",
			args{
				&TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 3},
				},
			},
			1,
		},
		{"Case2",
			args{
				&TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}},
					Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}},
				},
			},
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTilt(tt.args.root); got != tt.want {
				t.Errorf("findTilt() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_sumOfNode(t *testing.T) {
// 	type args struct {
// 		root *TreeNode
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{"Case1", args{
// 			&TreeNode{Val: 1},
// 		}, 1},
// 		{"Case2",
// 			args{
// 				&TreeNode{
// 					Val:   1,
// 					Left:  &TreeNode{Val: 2},
// 					Right: &TreeNode{Val: 3},
// 				},
// 			},
// 			6,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := sumOfNode(tt.args.root); got != tt.want {
// 				t.Errorf("sumOfNode() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
