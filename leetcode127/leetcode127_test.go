package leetcode127

import "testing"

func Benchmark_ladderLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	}
}
func Test_ladderLength(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				beginWord: "hog",
				endWord:   "cog",
				wordList:  []string{"hot", "cog"},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			want: 5,
		},
		{
			name: "case 3",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			},
			want: 0,
		},
		{
			name: "case 4",
			args: args{
				beginWord: "hot",
				endWord:   "dog",
				wordList:  []string{"hot", "dog"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ladderLength(tt.args.beginWord, tt.args.endWord, tt.args.wordList); got != tt.want {
				t.Errorf("ladderLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
