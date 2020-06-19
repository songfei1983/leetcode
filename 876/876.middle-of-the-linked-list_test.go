package middleofthelinkedlist

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	tests := []struct {
		name string
		args *ListNode
		want *ListNode
	}{
		{"0", create([]int{1}), create([]int{1})},
		{"1", create([]int{1, 2, 3, 4, 5}), create([]int{3, 4, 5})},
		{"2", create([]int{1, 2, 3, 4, 5, 6}), create([]int{4, 5, 6})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func create(ns []int) *ListNode {
	h := &ListNode{}
	n := h
	for _, v := range ns {
		n.Val = v
		n.Next = &ListNode{}
		n = n.Next
	}
	return h.Next
}
