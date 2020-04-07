package leetcode451

import "testing"

func Test_frequencySort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "aba",
			args: args{
				s: "aba",
			},
			want: "aab",
		},
		{
			name: "aaba",
			args: args{
				s: "aaba",
			},
			want: "aaab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := frequencySort(tt.args.s); got != tt.want {
				t.Errorf("frequencySort() = %v, want %v", got, tt.want)
			}
		})
	}
}
