package g_509_fibonacci_number

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{0},
			want: 0,
		},
		{
			name: "1",
			args: args{1},
			want: 1,
		},
		{
			name: "2",
			args: args{2},
			want: 1,
		},
		{
			name: "3",
			args: args{3},
			want: 2,
		},
		{
			name: "4",
			args: args{4},
			want: 3,
		}, {
			name: "5",
			args: args{5},
			want: 5,
		},
		{
			name: "6",
			args: args{6},
			want: 8,
		},
		{
			name: "10",
			args: args{10},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
