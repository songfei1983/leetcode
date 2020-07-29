package lettercombinationsofaphonenumber

import (
	"reflect"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"", args{""}, nil},
		{"2", args{"2"}, []string{"a", "b", "c"}},
		{"23", args{"23"}, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
