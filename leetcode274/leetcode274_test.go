package leetcode274

import "testing"

func Test_hIndex(t *testing.T) {
	type args struct {
		citations []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[10,8,5,4]",
			args: args{
				citations: []int{10, 8, 5, 4},
			},
			want: 4,
		},
		{
			name: "[0,1]",
			args: args{
				citations: []int{0, 1},
			},
			want: 1,
		},
		{
			name: "[100]",
			args: args{
				citations: []int{100},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				citations: []int{},
			},
			want: 0,
		},
		{
			name: "3,0,6,1,5",
			args: args{
				citations: []int{3, 0, 6, 1, 5},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex(tt.args.citations); got != tt.want {
				t.Errorf("hIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
