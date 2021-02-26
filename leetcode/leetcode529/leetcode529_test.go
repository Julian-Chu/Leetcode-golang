package leetcode529

import (
	"reflect"
	"testing"
)

func Test_updateBoard(t *testing.T) {
	type args struct {
		board [][]byte
		click []int
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "case 0",
			args: args{
				board: [][]byte{{'E'}},
				click: []int{0, 0},
			},
			want: [][]byte{{'B'}},
		},
		{
			name: "case1",
			args: args{
				board: [][]byte{
					{'E', 'E', 'E', 'E', 'E'},
					{'E', 'E', 'M', 'E', 'E'},
					{'E', 'E', 'E', 'E', 'E'},
					{'E', 'E', 'E', 'E', 'E'}},
				click: []int{3, 0},
			},
			want: [][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'M', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			},
		},
		{
			name: "case2",
			args: args{
				board: [][]byte{
					{'B', '1', 'E', '1', 'B'},
					{'B', '1', 'M', '1', 'B'},
					{'B', '1', '1', '1', 'B'},
					{'B', 'B', 'B', 'B', 'B'},
				},
				click: []int{1, 2},
			},
			want: [][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'X', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			},
		},
		{
			name: "case3",
			args: args{
				board: [][]byte{
					{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
					{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'M'},
					{'E', 'E', 'M', 'E', 'E', 'E', 'E', 'E'},
					{'M', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
					{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
					{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
					{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
					{'E', 'E', 'M', 'M', 'E', 'E', 'E', 'E'}},
				click: []int{0, 0},
			},
			want: [][]byte{
				{'B', 'B', 'B', 'B', 'B', 'B', '1', 'E'},
				{'B', '1', '1', '1', 'B', 'B', '1', 'M'},
				{'1', '2', 'M', '1', 'B', 'B', '1', '1'},
				{'M', '2', '1', '1', 'B', 'B', 'B', 'B'},
				{'1', '1', 'B', 'B', 'B', 'B', 'B', 'B'},
				{'B', 'B', 'B', 'B', 'B', 'B', 'B', 'B'},
				{'B', '1', '2', '2', '1', 'B', 'B', 'B'},
				{'B', '1', 'M', 'M', '1', 'B', 'B', 'B'}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateBoard(tt.args.board, tt.args.click); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
