package leetcode93

import (
	"reflect"
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
			name: "0000",
			args: args{
				s: "0000",
			},
			want: []string{"0.0.0.0"},
		},
		{
			name: "1111",
			args: args{
				s: "1111",
			},
			want: []string{"1.1.1.1"},
		},
		{
			name: "01010",
			args: args{
				s: "01010",
			},
			want: []string{"0.1.0.10", "0.10.1.0"},
		},
		{
			name: "25525511135",
			args: args{
				s: "25525511135",
			},
			want: []string{"255.255.11.135", "255.255.111.35"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddresses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
