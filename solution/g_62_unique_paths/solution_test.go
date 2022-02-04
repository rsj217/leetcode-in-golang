package g_62_unique_paths

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "m=1, n=2",
			args: args{m: 1, n: 2},
			want: 1,
		},
		{
			name: "m=2, n=1",
			args: args{m: 2, n: 1},
			want: 1,
		},
		{
			name: "m=3, n=7",
			args: args{m: 3, n: 7},
			want: 28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
