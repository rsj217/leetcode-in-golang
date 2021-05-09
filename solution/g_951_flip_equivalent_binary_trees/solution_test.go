package g_951_flip_equivalent_binary_trees

import (
	. "github/rsj217/leetcode-in-golang/datastruct"

	"testing"
)

func Test_flipEquiv(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[1,2,3,4,5,6,nil,nil,nil,7,8] [1,3,2,nil,6,4,5,nil,nil,nil,nil,8,7]",
			args: args{
				root1: Create([]interface{}{1, 2, 3, 4, 5, 6, nil, nil, nil, 7, 8}),
				root2: Create([]interface{}{1, 3, 2, nil, 6, 4, 5, nil, nil, nil, nil, 8, 7}),
			},
			want: true,
		},
		{
			name: "[1,2,3,4,5,6,nil,nil,nil,7,8] [1,3,2,nil,6,4,5,nil,nil,nil,nil,9,7]",
			args: args{
				root1: Create([]interface{}{1, 2, 3, 4, 5, 6, nil, nil, nil, 7, 8}),
				root2: Create([]interface{}{1, 3, 2, nil, 6, 4, 5, nil, nil, nil, nil, 9, 7}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipEquiv(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("flipEquiv() = %v, want %v", got, tt.want)
			}
		})
	}
}
