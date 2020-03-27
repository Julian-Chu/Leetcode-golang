package leetcode449

import (
	"Leetcode-golang/utils"
	"reflect"
	"testing"
)

func TestCodec_serialize(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "[1,2]",
			args: args{
				root: []int{1, 2, utils.NULL},
			},
			want: "1,2,NULL",
		},
		{
			name: "[1,2,NULL,NULL,3,NULL,NULL]",
			args: args{
				root: []int{1, 2, utils.NULL, utils.NULL, 3, utils.NULL, utils.NULL},
			},
			want: "1,2,NULL,NULL,3,NULL,NULL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			if got := this.serialize(utils.Ints2Tree(tt.args.root)); got != tt.want {
				t.Errorf("serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodec_deserialize(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "[1,2,NULL,NULL,3,NULL,NULL]",
			args: args{
				data: "1,2,NULL,NULL,3,NULL,NULL",
			},
			want: []int{1, 2, utils.NULL, utils.NULL, 3, utils.NULL, utils.NULL},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			want := utils.Ints2Tree(tt.want)
			if got := this.deserialize(tt.args.data); !reflect.DeepEqual(got, want) {
				t.Errorf("deserialize() = %v, want %v", tt.name, tt.want)
			}
		})
	}
}
