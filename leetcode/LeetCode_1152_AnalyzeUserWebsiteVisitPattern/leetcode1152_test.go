package LeetCode_1152_AnalyzeUserWebsiteVisitPattern

import (
	"reflect"
	"testing"
)

func Test_mostVisitedPattern(t *testing.T) {
	type args struct {
		username  []string
		timestamp []int
		website   []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case 1",
			args: args{
				[]string{"joe", "joe", "joe", "james", "james", "james", "james", "mary", "mary", "mary"},
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				[]string{"home", "about", "career", "home", "cart", "maps", "home", "home", "about", "career"},
			},
			want: []string{
				"home",
				"about",
				"career",
			},
		},
		{
			name: "case 2",
			args: args{
				[]string{"ua", "ua", "ua", "ub", "ub", "ub"},
				[]int{1, 2, 3, 4, 5, 6},
				[]string{"a", "b", "a", "a", "b", "c"},
			},
			want: []string{"a", "b", "a"},
		},
		{
			name: "case 3",
			args: args{
				[]string{"ua", "ua", "ua", "ub", "ub", "ub"},
				[]int{1, 2, 3, 4, 5, 6},
				[]string{"a", "b", "c", "a", "b", "a"},
			},
			want: []string{"a", "b", "a"},
		},
		{
			name: "case 4",
			args: args{
				[]string{"dowg", "dowg", "dowg"},
				[]int{158931262, 562600350, 148438945},
				[]string{"y", "loedo", "y"},
			},
			want: []string{"y", "y", "loedo"},
		},
		{
			name: "case 5",
			args: args{
				[]string{"zkiikgv", "zkiikgv", "zkiikgv", "zkiikgv"},
				[]int{436363475, 710406388, 386655081, 797150921},
				[]string{"wnaaxbfhxp", "mryxsjc", "oz", "wlarkzzqht"},
			},
			want: []string{"oz", "mryxsjc", "wlarkzzqht"},
		},
		{
			name: "case 6",
			args: args{
				[]string{"pdtkrjhd", "pdtkrjhd", "pdtkrjhd", "pdtkrjhd", "pdtkrjhd", "pdtkrjhd"},
				[]int{210984153, 262799291, 958396687, 605779010, 373702273, 205190519},
				[]string{"xgriygexlk", "qs", "rugydl", "bkrok", "canlv", "cahgsobjjs"},
			},
			want: []string{"cahgsobjjs", "bkrok", "rugydl"},
		},
		{
			name: "case 7",
			args: args{
				[]string{"h", "eiy", "cq", "h", "cq", "txldsscx", "cq", "txldsscx", "h", "cq", "cq"},
				[]int{527896567, 334462937, 517687281, 134127993, 859112386, 159548699, 51100299, 444082139, 926837079, 317455832, 411747930},
				[]string{"hibympufi", "hibympufi", "hibympufi", "hibympufi", "hibympufi", "hibympufi", "hibympufi", "hibympufi", "yljmntrclw", "hibympufi", "yljmntrclw"},
			},
			want: []string{"hibympufi", "hibympufi", "yljmntrclw"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostVisitedPattern(tt.args.username, tt.args.timestamp, tt.args.website); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mostVisitedPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
