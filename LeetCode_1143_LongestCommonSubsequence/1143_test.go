package leetcode1143

import "testing"

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a a",
			args: args{
				"a",
				"a",
			},
			want: 1,
		},
		{
			name: "ab a",
			args: args{
				"ab",
				"a",
			},
			want: 1,
		},
		{
			name: "ab ac",
			args: args{
				"ab",
				"ac",
			},
			want: 1,
		},
		{
			name: "abc ac",
			args: args{
				"abc",
				"ac",
			},
			want: 2,
		},
		{
			name: "abcde ace",
			args: args{
				"abcde",
				"ace",
			},
			want: 3,
		},

		{
			name: "abc def",
			args: args{
				"abc",
				"def",
			},
			want: 0,
		}, {
			name: "bsbininm jmjkbkjkv",
			args: args{
				"bsbininm",
				"jmjkbkjkv",
			},
			want: 1,
		},

		{
			name: "ezupkr ubmrapg",
			args: args{
				"ezupkr",
				"ubmrapg",
			},
			want: 2,
		},
		{
			name: "oxcpqrsvwf shmtulqrypy",
			args: args{
				"oxcpqrsvwf",
				"shmtulqrypy",
			},
			want: 2,
		},
		{
			name: "pmjghexybyrgzczy hafcdqbgncrcbihkd",
			args: args{
				"pmjghexybyrgzczy",
				"hafcdqbgncrcbihkd",
			},
			want: 4,
		},
		{
			name: "mhunuzqrkzsnidwbun szulspmhwpazoxijwbq",
			args: args{
				"mhunuzqrkzsnidwbun",
				"szulspmhwpazoxijwbq",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
