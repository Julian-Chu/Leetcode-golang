package leetcode211

import "testing"

func TestWordDictionary_Search1(t *testing.T) {
	dic := Constructor()
	dic.AddWord("bad")
	dic.AddWord("dad")
	dic.AddWord("mad")
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "pad",
			args: args{
				word: "pad",
			},
			want: false,
		},
		{
			name: "bad",
			args: args{
				word: "bad",
			},
			want: true,
		},
		{
			name: ".ad",
			args: args{
				word: ".ad",
			},
			want: true,
		},
		{
			name: "b..",
			args: args{
				word: "b..",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dic.Search(tt.args.word); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestWordDictionary_Search2(t *testing.T) {
	dic := Constructor()
	dic.AddWord("a")
	dic.AddWord("ab")
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: ".",
			args: args{
				word: ".",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dic.Search(tt.args.word); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestWordDictionary_Search3(t *testing.T) {
	dic := Constructor()
	dic.AddWord("at")
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a",
			args: args{
				word: "a",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dic.Search(tt.args.word); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
