package leetcode22

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				n: 1,
			},
			want: []string{"()"},
		},
		{
			name: "2",
			args: args{
				n: 2,
			},
			want: []string{"(())", "()()"},
		},
		{
			name: "3",
			args: args{
				n: 3,
			},
			want: []string{"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()"},
		},
		{
			name: "4",
			args: args{
				n: 4,
			},
			want: []string{
				"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
