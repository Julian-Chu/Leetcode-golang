package leetcode1054

import (
	"reflect"
	"testing"
)

func Test_rearrangeBarcodes(t *testing.T) {
	type args struct {
		barcodes []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1,2,1]",
			args: args{
				barcodes: []int{1, 2, 1},
			},
			want: []int{1, 2, 1},
		},
		{
			name: "[1,1,1,2,2,2]",
			args: args{
				barcodes: []int{1, 1, 1, 2, 2, 2},
			},
			want: []int{1, 2, 1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rearrangeBarcodes(tt.args.barcodes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rearrangeBarcodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
