package leetcode98

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "2 1 3",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: true,
		},
		{
			name: "[5,1,4,null,null,3,6",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 3,
						},
					},
				},
			},
			want: false,
		},
		{
			name: "1,1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 1,
					},
				},
			},
			want: false,
		},
		{
			name: "10,5,15,null,null,6,20",
			args: args{
				root: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val:   15,
						Left:  &TreeNode{Val: 6},
						Right: &TreeNode{Val: 20},
					},
				},
			},
			want: false,
		},
		{
			name: "[3,1,5,0,2,4,6",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
					Right: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "[3,1,5,0,2,4,6,null,null,null,3",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 2,
							Right: &TreeNode{
								Val: 3,
							},
						},
					},
					Right: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
			},
			want: false,
		},
		{
			name: "34,-6,null,-21",
			args: args{
				root: &TreeNode{
					Val: 34,
					Left: &TreeNode{
						Val: -6,
						Left: &TreeNode{
							Val: -21,
						},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
