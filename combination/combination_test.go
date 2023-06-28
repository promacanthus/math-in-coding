package combination

import "testing"

func Test_combine(t *testing.T) {
	type args struct {
		k      int
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
				k:      2,
				input:  []string{"t1", "t2", "t3"},
				result: []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			combine(tt.args.k, tt.args.input, tt.args.result)
		})
	}
}
