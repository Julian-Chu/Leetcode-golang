package leetcode79

import "testing"

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ABC",
			args: args{
				board: [][]byte{{'A', 'B', 'C'}},
				word:  "ABC",
			},
			want: true,
		},
		{
			name: "AB",
			args: args{
				board: [][]byte{{'A'}, {'B'}},
				word:  "AB",
			},
			want: true,
		},
		{
			name: "ABD",
			args: args{
				board: [][]byte{{'A', 'B'}, {'C', 'D'}},
				word:  "ABD",
			},
			want: true,
		},
		{
			name: "ABDC",
			args: args{
				board: [][]byte{{'A', 'B'}, {'C', 'D'}},
				word:  "ABDC",
			},
			want: true,
		},
		{
			name: "CDB",
			args: args{
				board: [][]byte{{'A', 'B'}, {'C', 'D'}},
				word:  "CDB",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
