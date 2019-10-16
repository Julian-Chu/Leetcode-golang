package leetcode66

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[0]",
			args: args{
				digits: []int{0},
			},
			want: []int{1},
		},
		{
			name: "[1]",
			args: args{
				digits: []int{1},
			},
			want: []int{2},
		},
		{
			name: "[9]",
			args: args{
				digits: []int{9},
			},
			want: []int{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
