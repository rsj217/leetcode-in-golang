package g_814_bianry_tree_pruning

import (
	. "github/rsj217/leetcode-in-golang/datastruct"

	"testing"
)

func Test_pruneTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		//want int
	}{
		{
			name: "[0, nil, 0, 0, 0]",
			args: args{Create([]interface{}{0, nil, 0, 0, 0})},
			//want: 4,
		},
	}
	// TODO
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := pruneTree(tt.args.root); got != tt.want {
			//	t.Errorf("pruneTree() = %v, want %v", got, tt.want)
			//}
		})
	}
}
