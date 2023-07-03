package dynamicprogramming

import "testing"

func Test_getStrDistance(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				a: "abc",
				b: "acd",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStrDistance(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getStrDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
