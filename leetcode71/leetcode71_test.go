package leetcode71

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "/home/",
			args: args{
				path: "/home/",
			},
			want: "/home",
		},
		{
			name: "/../",
			args: args{
				path: "/../",
			},
			want: "/",
		},
		{
			name: "/home//foo/",
			args: args{
				path: "/home//foo/",
			},
			want: "/home/foo",
		},
		{
			name: "/a/./b/../../c/",
			args: args{
				path: "/a/./b/../../c/",
			},
			want: "/c",
		},
		{
			name: "/a///b////c/d///././/..",
			args: args{
				path: "/a/b////c/d/././/..",
			},
			want: "/a/b/c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
