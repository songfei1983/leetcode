package swapnodesinpairs

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"1",
			args{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}},
			&ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, nil}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				gotS, _ := json.Marshal(got)
				wantS, _ := json.Marshal(tt.want)
				t.Errorf("swapPairs() = %v, want %v", string(gotS), string(wantS))
			}
		})
	}
}
