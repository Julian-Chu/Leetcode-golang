package LeetCode_383_RansomNote

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "fihjjjjei hjibagacbhadfaefdjaeaebgi",
			args: args{
				ransomNote: "fihjjjjei",
				magazine:   "hjibagacbhadfaefdjaeaebgi",
			},
			want: false,
		},
		{
			name: "djfjfhjf, aaigcbiafifghhdihcfdjjej",
			args: args{
				ransomNote: "djfjfhjf",
				magazine:   "aaigcbiafifghhdihcfdjjej",
			},
			want: true,
		},
		{
			name: "aa aab",
			args: args{
				ransomNote: "aa",
				magazine:   "aab",
			},
			want: true,
		},
		{
			name: "a, b",
			args: args{
				ransomNote: "a",
				magazine:   "b",
			},
			want: false,
		},
		{
			name: "aa, ab",
			args: args{
				ransomNote: "aa",
				magazine:   "ab",
			},
			want: false,
		},
		{
			name: "aaa, aab",
			args: args{
				ransomNote: "aaa",
				magazine:   "aab",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
