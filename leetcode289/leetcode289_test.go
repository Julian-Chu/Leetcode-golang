package leetcode289

import "testing"

func Test_gameOfLife(t *testing.T) {
	type args struct {
		board     [][]int
		wantBoard [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				board:     [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}},
				wantBoard: [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameOfLife(tt.args.board)

			for i := range tt.args.board {
				for j := range tt.args.board[i] {
					if tt.args.board[i][j] != tt.args.wantBoard[i][j] {
						t.Errorf("Failed in (i:%v,j:%v)", i, j)
					}
				}

			}
		})
	}
}
