package leetcode524

import "testing"

func Test_findLongestWord(t *testing.T) {
	type args struct {
		s string
		d []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				s: "abpcplea",
				d: []string{"ale", "apple", "monkey", "plea"},
			},
			want: "apple",
		},
		{
			name: "abpcplea",
			args: args{
				s: "abpcplea",
				d: []string{"a", "b", "c"},
			},
			want: "a",
		},
		{
			name: "bab",
			args: args{
				s: "bab",
				d: []string{"ba", "ab", "a", "b"},
			},
			want: "ab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLongestWord(tt.args.s, tt.args.d); got != tt.want {
				t.Errorf("findLongestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
