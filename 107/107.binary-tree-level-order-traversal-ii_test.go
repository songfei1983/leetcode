package binarytreelevelordertraversalii

import (
	"reflect"
	"testing"
)

func Test_levelOrderBottom(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}}, [][]int{[]int{15, 7}, []int{9, 20}, []int{3}}},
		{"2", args{&TreeNode{1, &TreeNode{2, &TreeNode{3, &TreeNode{4, &TreeNode{5, nil, nil}, nil}, nil}, nil}, nil}}, [][]int{[]int{5}, []int{4}, []int{3}, []int{2}, []int{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderBottom(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}
