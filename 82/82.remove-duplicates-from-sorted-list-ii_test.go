package removeduplicatesfromsortedlistii

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
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
			args{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}}},
			&ListNode{1, &ListNode{2, &ListNode{5, nil}}},
		},
		{
			"2",
			args{&ListNode{1, &ListNode{1, &ListNode{1, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{5, nil}}}}}}}},
			&ListNode{3, &ListNode{4, nil}},
		},
		{
			"3",
			args{&ListNode{1, nil}},
			&ListNode{1, nil},
		},
		{
			"3",
			args{&ListNode{1, &ListNode{1, nil}}},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				gots, _ := json.Marshal(got)
				wants, _ := json.Marshal(tt.want)
				t.Errorf("deleteDuplicates() = %v, want %v", string(gots), string(wants))
			}
		})
	}
}
