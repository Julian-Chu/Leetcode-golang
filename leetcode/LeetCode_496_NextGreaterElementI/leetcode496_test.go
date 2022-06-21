package LeetCode_496_NextGreaterElementI

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElement(t *testing.T) {
	type args struct {
		findNums []int
		nums     []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "num1=[4,1,2], num2=[1,3,4,2]",
			args: args{
				findNums: []int{4, 1, 2},
				nums:     []int{1, 3, 4, 2},
			},
			want: []int{-1, 3, -1},
		},
		{
			name: "num1=[2,4], num2=[1,2,3,4]",
			args: args{
				findNums: []int{2, 4},
				nums:     []int{1, 2, 3, 4},
			},
			want: []int{3, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement(tt.args.findNums, tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
