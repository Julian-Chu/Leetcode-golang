package leetcode721

import (
	"reflect"
	"testing"
)

func Test_accountsMerge_1(t *testing.T) {
	type args struct {
		accounts [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "case1",
			args: args{
				accounts: [][]string{{"David", "David0@m.co", "David4@m.co", "David3@m.co"}, {"David", "David5@m.co", "David5@m.co", "David0@m.co"}, {"David", "David1@m.co", "David4@m.co", "David0@m.co"}, {"David", "David0@m.co", "David1@m.co", "David3@m.co"}, {"David", "David4@m.co", "David1@m.co", "David3@m.co"}},
			},
			want: [][]string{{"David", "David0@m.co", "David1@m.co", "David3@m.co", "David4@m.co", "David5@m.co"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := accountsMerge_1(tt.args.accounts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("accountsMerge_1() = %v, want %v", got, tt.want)
			}
		})
	}
}
