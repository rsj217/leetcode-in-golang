package g_96_unique_binary_search_trees

import "testing"

func Test_numTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "2",
			args: args{n: 2},
			want: 2,
		},
		{
			name: "3",
			args: args{n: 3},
			want: 5,
		},
		{
			name: "5",
			args: args{n: 5},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTrees(tt.args.n); got != tt.want {
				t.Errorf("numTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
