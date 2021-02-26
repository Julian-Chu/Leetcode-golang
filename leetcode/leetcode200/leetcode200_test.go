package leetcode200

import "testing"

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 0",
			args: args{
				grid: [][]byte{},
			},
			want: 0,
		},
		{
			name: "case 1",
			args: args{
				grid: [][]byte{
					[]byte("10"),
					[]byte("01")},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				grid: [][]byte{
					[]byte("10"),
					[]byte("10")},
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				grid: [][]byte{
					[]byte("11110"),
					[]byte("11010"),
					[]byte("11000"),
					[]byte("00000"),
				},
			},
			want: 1,
		},
		{
			name: "case 4",
			args: args{
				grid: [][]byte{
					[]byte("11000"),
					[]byte("11000"),
					[]byte("00100"),
					[]byte("00011"),
				},
			},
			want: 3,
		},
		{
			name: "case 5",
			args: args{
				grid: [][]byte{
					[]byte("111"),
					[]byte("010"),
					[]byte("111"),
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.args.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
