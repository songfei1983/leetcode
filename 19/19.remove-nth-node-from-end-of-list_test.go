package removenthnodefromendoflist

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"0", args{
			head: &ListNode{1, nil},
			n:    1,
		}, nil},
		{"1", args{
			head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			n:    2,
		}, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{5, nil}}}}},
		{"2", args{
			head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			n:    1,
		}, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}},
		{"3", args{
			head: &ListNode{1, &ListNode{2, nil}},
			n:    2,
		}, &ListNode{2, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				gotS, _ := json.Marshal(got)
				wantS, _ := json.Marshal(tt.want)
				t.Errorf("removeNthFromEnd() = %v, want %v", string(gotS), string(wantS))
			}
		})
	}
}
