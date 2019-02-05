package leetcode897

import (
	"reflect"
	"testing"
)

func Test_increasingBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// {
		// 	"Case1",
		// 	args{
		// 		&TreeNode{
		// 			Val:   5,
		// 			Left:  &TreeNode{Val: 3},
		// 			Right: &TreeNode{Val: 6},
		// 		},
		// 	},
		// 	&TreeNode{
		// 		Val: 3,
		// 		Right: &TreeNode{
		// 			Val: 5,
		// 			Right: &TreeNode{
		// 				Val: 6,
		// 			},
		// 		},
		// 	},
		// },
		// {
		// 	"Case2",
		// 	args{
		// 		&TreeNode{
		// 			Val: 5,
		// 			Left: &TreeNode{
		// 				Val:   3,
		// 				Right: &TreeNode{Val: 4},
		// 			},
		// 			Right: &TreeNode{Val: 6},
		// 		},
		// 	},
		// 	&TreeNode{
		// 		Val: 3,
		// 		Right: &TreeNode{
		// 			Val: 4,
		// 			Right: &TreeNode{
		// 				Val: 5,
		// 				Right: &TreeNode{
		// 					Val: 6,
		// 				},
		// 			},
		// 		},
		// 	},
		// },

		{
			"Case3",
			args{
				&TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 2},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{Val: 6},
				},
			},
			&TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 5,
							Right: &TreeNode{
								Val: 6,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingBST(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				actual := got
				expect := tt.want
				for expect.Val == actual.Val {
					expect = expect.Right
					actual = actual.Right
				}
				t.Errorf("increasingBST() = %v, want %v", actual, expect)
			}
		})
	}
}
