package permutation

import (
	"testing"
)

func Test_permutation(t *testing.T) {
	type args struct {
		horses []string
		result []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				horses: []string{"t1", "t2", "t3"},
				result: []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			permutation(tt.args.horses, tt.args.result)
		})
	}
}

func TestPassword(t *testing.T) {
	type args struct {
		input  []string
		result []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				input:  []string{"a", "b", "c", "d", "e"},
				result: []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GeneratePassword(tt.args.input, tt.args.result)
		})
	}
}
