package LeetCode_685_RedundantConnectionII

import (
	"reflect"
	"testing"
)

func Test_findRedundantDirectedConnection(t *testing.T) {
	type args struct {
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case3",
			args: args{
				edges: [][]int{{2, 1}, {3, 1}, {4, 2}, {1, 4}},
			},
			want: []int{2, 1},
		},
		{
			name: "case1",
			args: args{
				edges: [][]int{{1, 2}, {1, 3}, {2, 3}},
			},
			want: []int{2, 3},
		},
		{
			name: "case2",
			args: args{
				edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}},
			},
			want: []int{4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRedundantDirectedConnection(tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRedundantDirectedConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
