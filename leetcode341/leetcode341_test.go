package leetcode341

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		nestedList []*NestedInteger
	}
	tests := []struct {
		name string
		args args
		want *NestedIterator
	}{
		{
			name: "[[1,1],2],[1,1]]",
			args: args{
				nestedList: []*NestedInteger{
					{
						List: []*NestedInteger{
							{Num: 1},
							{Num: 1},
						},
					},
					{Num: 2},
					{
						List: []*NestedInteger{
							{Num: 1},
							{Num: 1},
						},
					},
				},
			},
			want: &NestedIterator{
				idx:  -1,
				nums: []int{1, 1, 2, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.nestedList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName(t *testing.T) {
	var nums []int
	inc(nums)

	if nums[0] != 1 {
		t.Error("failed")
	}
}

func inc(nums []int) {
	nums = append(nums, 1)
}
