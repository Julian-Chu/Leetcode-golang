package leetcode74

import "testing"

func Test_searchMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "target = 3",
			args: args{
				matrix: matrix,
				target: 3,
			},
			want: true,
		},
		{
			name: "target=13",
			args: args{
				matrix: matrix,
				target: 13,
			},
			want: false,
		},
		{
			name: "target=11",
			args: args{
				matrix: matrix,
				target: 11,
			},
			want: true,
		},
		{
			name: "{1,3} 3",
			args: args{
				matrix: [][]int{{1}, {3}},
				target: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
