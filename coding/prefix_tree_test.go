package coding

import (
	"testing"
)

func Test_replaceWord(t *testing.T) {
	type args struct {
		dict     []string
		sentence string
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
				dict:     []string{"cat", "bat", "rat"},
				sentence: "the cattle was rattled by the battery",
			},
			want: "the cat was rat by the bat",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceWord(tt.args.dict, tt.args.sentence); got != tt.want {
				t.Errorf("replaceWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MagicDictionary(t *testing.T) {
	tests := []struct {
		name        string
		buildWords  []string
		searchWords []string
		want        []bool
	}{
		// TODO: Add test cases.
		{
			name:        "",
			buildWords:  []string{"happy", "new", "year"},
			searchWords: []string{"now", "new"},
			want:        []bool{true, false},
		},
	}
	for _, tt := range tests {
		md := NewMagicDictionary()
		t.Run(tt.name, func(t *testing.T) {
			md.buildDict(tt.buildWords)
			for i, word := range tt.searchWords {
				if got := md.search(word); got != tt.want[i] {
					t.Errorf("search() = %v, want %v", got, tt.want[i])
				}
			}
		})
	}
}

func Test_miniLengthEncoding(t *testing.T) {
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
				words: []string{"time", "me", "bell"},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := miniLengthEncoding(tt.args.words); got != tt.want {
				t.Errorf("miniLengthEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MapSum(t *testing.T) {
	tests := []struct {
		name   string
		key    string
		value  int
		prefix string
		want   int
	}{
		{
			name:   "",
			key:    "happy",
			value:  3,
			prefix: "hap",
			want:   3,
		},
		{
			name:   "",
			key:    "happen",
			value:  2,
			prefix: "hap",
			want:   5,
		},
	}
	ms := NewMapSum()
	for _, tt := range tests {
		ms.insert(tt.key, tt.value)
		if got := ms.sum(tt.prefix); got != tt.want {
			t.Errorf("sum() = %v, want %v", got, tt.want)
		}
	}
}

func Test_findMaximumXOR(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{1, 3, 4, 7},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaximumXORV1(tt.args.nums); got != tt.want {
				t.Errorf("findMaximumXORV1() = %v, want %v", got, tt.want)
			}
			if got := findMaximumXORV2(tt.args.nums); got != tt.want {
				t.Errorf("findMaximumXORV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
