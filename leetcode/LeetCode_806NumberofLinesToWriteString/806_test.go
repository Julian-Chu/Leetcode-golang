package LeetCode_806NumberofLinesToWriteString

import (
	"reflect"
	"testing"
)

func Test_numberOfLines(t *testing.T) {
	type args struct {
		widths []int
		S      string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Case 1",
			args{[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
				"aaaaaaaaaaaa"},
			[]int{2, 20},
		},
		{"Case 2",
			args{[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
				"abcdefghijklmnopqrstuvwxyz"},
			[]int{3, 60},
		}, {"Case 3",
			args{[]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
				"bbbcccdddaaa"},
			[]int{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfLines(tt.args.widths, tt.args.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numberOfLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
