package mathematical_induction

import "testing"

func TestProve(t *testing.T) {
	type args struct {
		k      int
		result *Result
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				k:      1,
				result: &Result{},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				k:      2,
				result: &Result{},
			},
			want: true,
		},
		{
			name: "63",
			args: args{
				k:      63,
				result: &Result{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prove(tt.args.k, tt.args.result); got != tt.want {
				t.Errorf("Prove() = %v, want %v", got, tt.want)
			}
		})
	}
}
