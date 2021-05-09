package g_662_maximum_width_of_binary_tree

import (
	. "leetcode/datastruct"
	"testing"
)

func Test_widthOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,3,2,5,3,null,9]",
			args: args{Create([]interface{}{1, 3, 2, 5, 3, nil, 9})},
			want: 4,
		},
		{
			name: "[]",
			args: args{Create([]interface{}{})},
			want: 0,
		},
		{
			name: "[1]",
			args: args{Create([]interface{}{1})},
			want: 1,
		},
		{
			name: "[1, 3, nil, 5, 3]",
			args: args{Create([]interface{}{1, 3, nil, 5, 3})},
			want: 2,
		},
		{
			name: "[1, 3, 2, 5]",
			args: args{Create([]interface{}{1, 3, 2, 5})},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := widthOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("widthOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
