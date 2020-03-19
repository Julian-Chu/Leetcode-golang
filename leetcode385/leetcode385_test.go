package leetcode385

import (
	"reflect"
	"testing"
)

func Test_deserialize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *NestedInteger
	}{
		{
			name: "[123, [456]]",
			args: args{
				s: "[123,[456]]",
			},
			want: &NestedInteger{
				val:  0,
				list: []*NestedInteger{{val: 123}, {list: []*NestedInteger{{456, nil}}}},
			},
		},
		{
			name: "[123]",
			args: args{
				s: "[123]",
			},
			want: &NestedInteger{
				val: 0,
				list: []*NestedInteger{
					{val: 123},
				},
			},
		},
		{
			name: "324",
			args: args{
				s: "324",
			},
			want: &NestedInteger{
				val:  324,
				list: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deserialize(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
