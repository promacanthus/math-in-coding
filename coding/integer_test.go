package coding

import (
	"math"
	"reflect"
	"testing"
)

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "15/2",
			args: args{
				dividend: 15,
				divisor:  2,
			},
			want: 7,
		},
		{
			name: "minInt/-1",
			args: args{
				dividend: math.MinInt,
				divisor:  -1,
			},
			want: math.MaxInt,
		},
		{
			name: "maxInt/maxInt",
			args: args{
				dividend: math.MaxInt,
				divisor:  math.MaxInt,
			},
			want: 1,
		},
		{
			name: "performance optimization",
			args: args{
				dividend: math.MaxInt,
				divisor:  2,
			},
			want: 4611686018427387903,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				a: "10",
				b: "10",
			},
			want: "100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				n: 4,
			},
			want: []int{0, 1, 1, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBitsV1(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBitsV1() = %v, want %v", got, tt.want)
			}
			if got := countBitsV2(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBitsV2() = %v, want %v", got, tt.want)
			}
			if got := countBitsV3(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBitsV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_singleNumber(t *testing.T) {
	type args struct {
		numbers []int
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
				numbers: []int{0, 1, 0, 1, 0, 1, 100},
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.numbers); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProduct(t *testing.T) {
	type args struct {
		words []string
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
				words: []string{"abcw", "foo", "bar", "fxyz", "abcdef"},
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProductV1(tt.args.words); got != tt.want {
				t.Errorf("maxProductV1() = %v, want %v", got, tt.want)
			}
			if got := maxProductV2(tt.args.words); got != tt.want {
				t.Errorf("maxProductV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
