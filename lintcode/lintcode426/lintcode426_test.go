package lintcode426

import (
	"reflect"
	"sort"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "25525511135",
			args: args{
				s: "25525511135",
			},
			want: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			name: "1111",
			args: args{
				s: "1111",
			},
			want: []string{"1.1.1.1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := restoreIpAddresses(tt.args.s)
			sort.Strings(got)
			sort.Strings(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
