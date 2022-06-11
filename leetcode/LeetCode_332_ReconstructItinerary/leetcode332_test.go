package LeetCode_332_ReconstructItinerary

import (
	"reflect"
	"testing"
)

func Test_findItinerary(t *testing.T) {
	type args struct {
		tickets [][]string
	}
	var tests = []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case5",
			args: args{
				tickets: [][]string{{"JFK", "ATL"}, {"ORD", "PHL"}, {"JFK", "ORD"}, {"PHX", "LAX"}, {"LAX", "JFK"}, {"PHL", "ATL"}, {"ATL", "PHX"}},
			},
			want: []string{"JFK", "ATL", "PHX", "LAX", "JFK", "ORD", "PHL", "ATL"},
		},
		{
			name: "case4",
			args: args{
				tickets: [][]string{{"JFK", "ATL"}, {"ATL", "JFK"}},
			},
			want: []string{"JFK", "ATL", "JFK"},
		},
		{
			name: `[["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]`,
			args: args{
				tickets: [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}},
			},
			want: []string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		},
		{
			name: `[["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]`,
			args: args{
				tickets: [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}},
			},
			want: []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
		},
		{
			name: `[["JFK","KUL"],["JFK","NRT"],["NRT","JFK"]]`,
			args: args{
				tickets: [][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}},
			},
			want: []string{"JFK", "NRT", "JFK", "KUL"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findItinerary(tt.args.tickets); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findItinerary() = %v, want %v", got, tt.want)
			}
		})
	}
}
