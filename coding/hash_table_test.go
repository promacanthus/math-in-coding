package coding

import (
	"reflect"
	"testing"
	"time"
)

func Test_isAnagram(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				s1: "anagram",
				s2: "nagaram",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagramV1(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isAnagramV1() = %v, want %v", got, tt.want)
			}
			if got := isAnagramV2(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isAnagramV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_groupAnagram(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				s: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{
				{"bat"},
				{"tan", "nat"},
				{"eat", "tea", "ate"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagramV1(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagramV1() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagramV2(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagramV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAlienSorted(t *testing.T) {
	type args struct {
		words []string
		order string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				words: []string{"offer", "is"},
				order: "zyxwvutsrqponmlkjihgfedcba",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAlienSorted(tt.args.words, tt.args.order); got != tt.want {
				t.Errorf("isAlienSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMinimalDifference(t *testing.T) {
	type args struct {
		times []string
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				times: []string{"23:50", "23:59", "00:00"},
			},
			want: time.Duration(9 * time.Minute),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinimalDifferenceV1(tt.args.times); got != tt.want {
				t.Errorf("findMinimalDifferenceV1() = %v, want %v", got, tt.want)
			}
			if got := findMinimalDifferenceV2(tt.args.times); got != tt.want {
				t.Errorf("findMinimalDifferenceV2() = %v, want %v", got, tt.want)
			}
			if got := findMinimalDifferenceV3(tt.args.times); got != tt.want {
				t.Errorf("findMinimalDifferenceV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
