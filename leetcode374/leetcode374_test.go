package leetcode374

import "testing"

func Test_guessNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		pickup int
		name   string
		args   args
		want   int
	}{
		{
			pickup: 6,
			name:   "10 6",
			args: args{
				n: 10,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pickup = tt.pickup
			if got := guessNumber(tt.args.n); got != tt.want {
				t.Errorf("guessNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
