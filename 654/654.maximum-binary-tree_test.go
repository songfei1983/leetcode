package maximumbinarytree

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_constructMaximumBinaryTree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"1", args{[]int{3, 2, 1, 6, 0, 5}}, &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 1},
				},
			},
			Right: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 0},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructMaximumBinaryTree(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				gotB, _ := json.Marshal(got)
				wantB, _ := json.Marshal(tt.want)
				t.Errorf("constructMaximumBinaryTree() = %v, want %v", string(gotB), string(wantB))
			}
		})
	}
}
