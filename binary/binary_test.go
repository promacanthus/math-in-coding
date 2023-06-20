package binary

import (
	"testing"
)

func TestBinaryToDecimal(t *testing.T) {
	type args struct {
		binary int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "53 --> 110101",
			args: args{
				binary: 110101,
			},
			want: 53,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinaryToDecimal(tt.args.binary); got != tt.want {
				t.Errorf("binaryToDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimalToBinary(t *testing.T) {
	type args struct {
		decimal int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "110101 ---> 53",
			args: args{
				decimal: 53,
			},
			want: 110101,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecimalToBinary(tt.args.decimal); got != tt.want {
				t.Errorf("decimalToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeftShift(t *testing.T) {
	type args struct {
		num int64
		m   int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "53 << 1",
			args: args{
				num: 53,
				m:   1,
			},
			want: 106,
		},
		{
			name: "53 << 3",
			args: args{
				num: 53,
				m:   3,
			},
			want: 424,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeftShift(tt.args.num, tt.args.m); got != tt.want {
				t.Errorf("LeftShift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRightShift(t *testing.T) {
	type args struct {
		num int64
		m   int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "53 >> 1",
			args: args{
				num: 53,
				m:   1,
			},
			want: 26,
		},
		{
			name: "53 >> 3",
			args: args{
				num: 53,
				m:   3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RightShift(tt.args.num, tt.args.m); got != tt.want {
				t.Errorf("RightShift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOR(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "or",
			args: args{
				a: 53, // 110101
				b: 35, // 100011
			},
			want: 55, // 110111
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OR(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("OR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAND(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "and",
			args: args{
				a: 53, // 110101
				b: 35, // 100011
			},
			want: 33, // 100001
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AND(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("AND() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXOR(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "xor",
			args: args{
				a: 53, // 110101
				b: 35, // 100011
			},
			want: 22, // 010110
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XOR(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("XOR() = %v, want %v", got, tt.want)
			}
		})
	}
}
