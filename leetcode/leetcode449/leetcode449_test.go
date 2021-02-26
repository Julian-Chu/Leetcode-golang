package leetcode449

import (
	"Leetcode-golang/utils"
	"reflect"
	"testing"
)

func TestCodec_serializeAndDeserialize(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "case2",
			args: args{
				root: utils.Ints2Tree([]int{
					41, 37, 44, 24, 39, 42, 48, 1, 35, 38, 40, utils.NULL, 43, 46, 49, 0, 2, 30, 36, utils.NULL, utils.NULL, utils.NULL, utils.NULL, utils.NULL, utils.NULL, 45, 47, utils.NULL, utils.NULL, utils.NULL, utils.NULL, utils.NULL, 4, 29, 32, utils.NULL, utils.NULL, utils.NULL, utils.NULL, utils.NULL, utils.NULL, 3, 9, 26, utils.NULL, 31, 34, utils.NULL, utils.NULL, 7, 11, 25, 27, utils.NULL, utils.NULL, 33, utils.NULL, 6, 8, 10, 16, utils.NULL, utils.NULL, utils.NULL, 28, utils.NULL, utils.NULL, 5, utils.NULL, utils.NULL, utils.NULL, utils.NULL, utils.NULL, 15, 19, utils.NULL, utils.NULL, utils.NULL, utils.NULL, 12, utils.NULL, 18, 20, utils.NULL, 13, 17, utils.NULL, utils.NULL, 22, utils.NULL, 14, utils.NULL, utils.NULL, 21, 23,
				}),
			},
			want: utils.Ints2Tree([]int{
				41, 37, 44, 24, 39, 42, 48, 1, 35, 38, 40, utils.NULL, 43, 46, 49, 0, 2, 30, 36, utils.NULL, utils.NULL, utils.NULL, utils.NULL, utils.NULL, utils.NULL, 45, 47, utils.NULL, utils.NULL, utils.NULL, utils.NULL, utils.NULL, 4, 29, 32, utils.NULL, utils.NULL, utils.NULL, utils.NULL, utils.NULL, utils.NULL, 3, 9, 26, utils.NULL, 31, 34, utils.NULL, utils.NULL, 7, 11, 25, 27, utils.NULL, utils.NULL, 33, utils.NULL, 6, 8, 10, 16, utils.NULL, utils.NULL, utils.NULL, 28, utils.NULL, utils.NULL, 5, utils.NULL, utils.NULL, utils.NULL, utils.NULL, utils.NULL, 15, 19, utils.NULL, utils.NULL, utils.NULL, utils.NULL, 12, utils.NULL, 18, 20, utils.NULL, 13, 17, utils.NULL, utils.NULL, 22, utils.NULL, 14, utils.NULL, utils.NULL, 21, 23,
			}),
		},
		{
			name: "41 37 44 24 39 42 48 1",
			args: args{
				root: &TreeNode{
					Val:   41,
					Left:  &utils.TreeNode{Val: 37, Left: &utils.TreeNode{Val: 24, Left: &utils.TreeNode{Val: 1}}, Right: &utils.TreeNode{Val: 39}},
					Right: &utils.TreeNode{Val: 44, Left: &utils.TreeNode{Val: 42}, Right: &TreeNode{Val: 48}},
				},
			},
			want: &TreeNode{
				Val:   41,
				Left:  &utils.TreeNode{Val: 37, Left: &utils.TreeNode{Val: 24, Left: &utils.TreeNode{Val: 1}}, Right: &utils.TreeNode{Val: 39}},
				Right: &utils.TreeNode{Val: 44, Left: &utils.TreeNode{Val: 42}, Right: &TreeNode{Val: 48}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			str := this.serialize(tt.args.root)
			if got := this.deserialize(str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestCodec_serialize(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "41 37 44 24 39 42",
			args: args{
				root: &TreeNode{
					Val:   41,
					Left:  &utils.TreeNode{Val: 37, Left: &utils.TreeNode{Val: 24}, Right: &utils.TreeNode{Val: 39}},
					Right: &utils.TreeNode{Val: 44, Left: &utils.TreeNode{Val: 42}},
				},
			},
			want: "41,37,24,39,44,42,",
		},
		{
			name: "[1,2]",
			args: args{
				root: &TreeNode{
					Val:   1,
					Right: &utils.TreeNode{Val: 2},
				},
			},
			want: "1,2,",
		},
		{
			name: "[1,2,3]",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &utils.TreeNode{Val: 2,
						Right: &utils.TreeNode{Val: 3}},
				},
			},
			want: "1,2,3,",
		},
		{
			name: "[5,2,6]",
			args: args{
				root: &TreeNode{
					Val:   5,
					Left:  &utils.TreeNode{Val: 2},
					Right: &utils.TreeNode{Val: 6},
				},
			},
			want: "5,2,6,",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			if got := this.serialize(tt.args.root); got != tt.want {
				t.Errorf("serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodec_deserialize(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "41 37 44 24 39 42 48 1",
			args: args{
				data: "41,37,24,1,39,44,42,48,",
			},
			want: &TreeNode{
				Val:   41,
				Left:  &utils.TreeNode{Val: 37, Left: &utils.TreeNode{Val: 24, Left: &utils.TreeNode{Val: 1}}, Right: &utils.TreeNode{Val: 39}},
				Right: &utils.TreeNode{Val: 44, Left: &utils.TreeNode{Val: 42}, Right: &TreeNode{Val: 48}},
			},
		},
		{
			name: "[1,2,3]",
			args: args{
				data: "1,2,3,",
			},
			want: &TreeNode{
				Val: 1,
				Right: &utils.TreeNode{Val: 2,
					Right: &TreeNode{Val: 3}},
			},
		},
		{
			name: "[5,2,6]",
			args: args{
				data: "5,2,6,",
			},
			want: &TreeNode{
				Val:   5,
				Left:  &utils.TreeNode{Val: 2},
				Right: &utils.TreeNode{Val: 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			if got := this.deserialize(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
