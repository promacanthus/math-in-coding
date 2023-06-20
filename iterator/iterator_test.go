package iterator

import (
	"sort"
	"testing"
)

func TestGetWheats(t *testing.T) {
	type args struct {
		grid int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "63",
			args: args{
				grid: 63,
			},
			want: 9223372036854775807,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWheats(tt.args.grid); got != tt.want {
				t.Errorf("GetWheats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSquareRoot(t *testing.T) {
	type args struct {
		root      int
		maxTry    int
		threshold float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "3^2 = 9",
			args: args{
				root:      9,
				maxTry:    100,
				threshold: 0.1,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSquareRoot(tt.args.root, tt.args.maxTry, tt.args.threshold); got != tt.want {
				t.Errorf("GetSquareRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchString(t *testing.T) {
	dir := []string{"One", "Two", "Three", "Four"}
	sort.Strings(dir) // ordered: Four One Three Two

	type args struct {
		dir    []string
		target string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Find",
			args: args{
				dir:    dir,
				target: "Three",
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchString(tt.args.dir, tt.args.target); got != tt.want {
				t.Errorf("SearchString() = %v, want %v", got, tt.want)
			}
		})
	}
}
