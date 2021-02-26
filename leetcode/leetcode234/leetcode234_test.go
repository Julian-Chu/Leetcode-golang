package leetcode234

import (
	"Leetcode-golang/utils"
	"testing"
)

func Test_isPalindrome(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want bool
	}{
		{
			name: "1->2",
			args: []int{1, 2},
			want: false,
		},
		{
			name: "1->2->2->1",
			args: []int{1, 2, 2, 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(utils.IntsToListNode(tt.args)); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
