package leetcode621

import "testing"

func Test_leastInterval(t *testing.T) {
	type args struct {
		tasks []byte
		n     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 5",
			args: args{
				tasks: []byte{'A', 'A', 'B', 'B', 'C', 'C', 'D', 'D'},
				n:     2,
			},
			want: 8,
		},
		{
			name: "case 4",
			args: args{
				tasks: []byte{'A', 'A', 'B', 'B', 'C', 'C', 'D', 'D', 'E', 'E', 'F',
					'F', 'G', 'G', 'H', 'H', 'I', 'I', 'J', 'J', 'K', 'K', 'L', 'L',
					'M', 'M', 'N', 'N', 'O', 'O', 'P', 'P', 'Q', 'Q', 'R', 'R', 'S',
					'S', 'T', 'T', 'U', 'U', 'V', 'V', 'W', 'W', 'X', 'X', 'Y', 'Y', 'Z', 'Z'},
				n: 2,
			},
			want: 52,
		},
		{
			name: "case1",
			args: args{
				tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
				n:     2,
			},
			want: 8,
		},
		{
			name: "case2",
			args: args{
				tasks: []byte{'A', 'A', 'B', 'B', 'B'},
				n:     2,
			},
			want: 7,
		},
		{
			name: "case3",
			args: args{
				tasks: []byte{'A', 'A', 'A', 'B', 'B'},
				n:     2,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastInterval(tt.args.tasks, tt.args.n); got != tt.want {
				t.Errorf("leastInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
