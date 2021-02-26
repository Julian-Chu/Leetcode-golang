package LeetCode_125_ValidPalindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"A man, a plan, a canal: Panama", args{"A man, a plan, a canal: Panama"}, true},
		{"aaa", args{"aaa"}, true},
		{"A,aa", args{"A,aa"}, true},
		{"A,ba", args{"A,ba"}, true},
		{"race a car", args{"race a car"}, false},
		{" ", args{" "}, true},
		{"", args{""}, true},
		{name: "ab", args: args{"ab"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
