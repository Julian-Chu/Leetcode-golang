package leetcode692

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		words []string
		k     int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{

		{
			name: `["a","aa","aaa"]`,
			args: args{
				words: []string{"a", "aa", "aaa"},
				k:     2,
			},
			want: []string{"a", "aa"},
		},
		{
			name: "[\"i\", \"love\", \"leetcode\", \"i\", \"love\", \"coding\"]\n2",
			args: args{
				words: []string{"i", "love", "leetcode", "i", "love", "coding"},
				k:     2,
			},
			want: []string{"i", "love"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.words, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
