package leetcode290

import "testing"

func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		str     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "abba, dog cat cat dog",
			args: args{
				pattern: "abba",
				str:     "dog cat cat dog",
			},
			want: true,
		},
		{
			name: "abba, dog cat cat fish",
			args: args{
				pattern: "abba",
				str:     "dog cat cat fish",
			},
			want: false,
		},
		{
			name: "abba,dog dog dog dog",
			args: args{
				pattern: "abba",
				str:     "dog dog dog dog",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.args.pattern, tt.args.str); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
