package leetcode165

import "testing"

func Test_compareVersion(t *testing.T) {
	type args struct {
		version1 string
		version2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0.1 vs 1.1",
			args: args{
				version1: "0.1",
				version2: "1.1",
			},
			want: -1,
		},
		{
			name: "1.0.1 vs 1",
			args: args{
				version1: "1.0.1",
				version2: "1",
			},
			want: 1,
		},
		{
			name: "7.5.2.4 vs 7.5.3",
			args: args{
				version1: "7.5.2.4",
				version2: "7.5.3",
			},
			want: -1,
		},
		{
			name: "1.01 vs 1.001",
			args: args{
				version1: "1.01",
				version2: "1.001",
			},
			want: 0,
		},
		{
			name: "1.0 vs 1.0.0",
			args: args{
				version1: "1.0",
				version2: "1.0.0",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareVersion(tt.args.version1, tt.args.version2); got != tt.want {
				t.Errorf("compareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
