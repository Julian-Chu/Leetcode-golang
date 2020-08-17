package leetcode1290

import (
	"Leetcode-golang/utils"
	"testing"
)

func Test_getDecimalValue(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,0,1]",
			args: args{
				head: []int{1, 0, 1},
			},
			want: 5,
		},
		{
			name: "{0}",
			args: args{
				head: []int{0},
			},
			want: 0,
		},
		{
			name: "{1}",
			args: args{
				head: []int{1},
			},
			want: 1,
		},
		{
			name: "[1,0,0,1,0,0,1,1,1,0,0,0,0,0,0]",
			args: args{
				head: []int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0},
			},
			want: 18880,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := utils.IntsToListNode(tt.args.head)
			if got := getDecimalValue(head); got != tt.want {
				t.Errorf("getDecimalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
