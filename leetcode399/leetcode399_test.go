package leetcode399

import (
	"math"
	"testing"
)

func Test_calcEquation(t *testing.T) {
	type args struct {
		equations [][]string
		values    []float64
		queries   [][]string
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "case3",
			args: args{
				equations: [][]string{{"x1", "x2"}, {"x2", "x3"}, {"x3", "x4"}, {"x4", "x5"}},
				values:    []float64{3, 4, 5, 6},
				queries:   [][]string{{"x1", "x5"}, {"x5", "x2"}, {"x2", "x4"}, {"x2", "x2"}, {"x2", "x9"}, {"x9", "x9"}},
			},
			want: []float64{360, 0.008333, 20, 1, -1, -1},
		},
		{
			name: "case1",
			args: args{
				equations: [][]string{{"a", "b"}, {"b", "c"}},
				values:    []float64{2.0, 3.0},
				queries:   [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			},
			want: []float64{6.0, 0.5, -1.0, 1.0, -1.0},
		},
		{
			name: "case2",
			args: args{
				equations: [][]string{{"a", "b"}, {"c", "d"}},
				values:    []float64{1.0, 1.0},
				queries:   [][]string{{"a", "c"}, {"b", "d"}, {"b", "a"}, {"d", "c"}},
			},
			want: []float64{-1, -1, 1, 1},
		},
	}

	compare := func(actual, expect []float64) bool {
		tolerance := 0.0001

		for i := range actual {
			if math.Abs(actual[i]-expect[i]) > tolerance {
				return false
			}
		}
		return true
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcEquation(tt.args.equations, tt.args.values, tt.args.queries)

			if !compare(got, tt.want) {
				t.Errorf("calcEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}
