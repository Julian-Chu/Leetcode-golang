package leetcode365

import "testing"

func Test_canMeasureWater(t *testing.T) {
	type args struct {
		x int
		y int
		z int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "x=3,y=5,z=4",
			args: args{
				x: 3,
				y: 5,
				z: 4,
			},
			want: true,
		},
		{
			name: "x=0,y=2,z=1",
			args: args{
				x: 0,
				y: 2,
				z: 1,
			},
			want: false,
		},
		{
			name: "x=3,y=5,z=8",
			args: args{
				x: 3,
				y: 5,
				z: 8,
			},
			want: true,
		},
		{
			name: "x=3,y=5,z=1",
			args: args{
				x: 3,
				y: 5,
				z: 1,
			},
			want: true,
		},
		{
			name: "x=2,y=6,z=5",
			args: args{
				x: 2,
				y: 6,
				z: 5,
			},
			want: false,
		},
		{
			name: "x=13,y=11,z=1",
			args: args{
				x: 13,
				y: 11,
				z: 1,
			},
			want: true,
		},
		{
			name: "x=34,y=5,z=6",
			args: args{
				x: 34,
				y: 5,
				z: 6,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMeasureWater(tt.args.x, tt.args.y, tt.args.z); got != tt.want {
				t.Errorf("canMeasureWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
