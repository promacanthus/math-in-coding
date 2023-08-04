package coding

import (
	"math"
	"strings"
)

type TrieNode struct {
	children []*TrieNode
	isWord   bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make([]*TrieNode, 26),
		isWord:   false,
	}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

type PrefixTree interface {
	Insert(word string)
	Search(word string) bool
	StartWith(word string) bool
}

func (t *Trie) Insert(word string) {
	node := t.root
	for i := 0; i < len(word); i++ {
		if node.children[word[i]-'a'] == nil {
			node.children[word[i]-'a'] = NewTrieNode()
		}
		node = node.children[word[i]-'a']
	}
	node.isWord = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for i := 0; i < len(word); i++ {
		if node.children[word[i]-'a'] == nil {
			return false
		}
		node = node.children[word[i]-'a']
	}
	return node.isWord
}

func (t *Trie) StartWith(prefix string) bool {
	node := t.root
	for i := 0; i < len(prefix); i++ {
		if node.children[prefix[i]-'a'] == nil {
			return false
		}
		node = node.children[prefix[i]-'a']
	}
	return true
}

// ---------- 替换单词 ----------
func replaceWord(dict []string, sentence string) string {
	trie := NewTrie()
	for _, word := range dict {
		trie.Insert(word)
	}
	words := strings.Split(sentence, " ")
	for i := 0; i < len(words); i++ {
		prefix := findPrefix(trie.root, words[i])
		if len(prefix) != 0 {
			words[i] = prefix
		}
	}
	return strings.Join(words, " ")
}

func findPrefix(root *TrieNode, word string) string {
	node := root
	var res []byte
	for i := 0; i < len(word); i++ {
		if node.isWord || node.children[word[i]-'a'] == nil {
			break
		}
		res = append(res, word[i])
		node = node.children[word[i]-'a']
	}
	if node.isWord {
		return string(res)
	} else {
		return ""
	}
}

// ---------- 神奇字典 ----------
type MagicDictionary struct {
	root *TrieNode
}

func NewMagicDictionary() *MagicDictionary {
	return &MagicDictionary{
		root: NewTrieNode(),
	}
}

func (d *MagicDictionary) buildDict(words []string) {
	for _, word := range words {
		node := d.root
		for i := 0; i < len(word); i++ {
			if node.children[word[i]-'a'] == nil {
				node.children[word[i]-'a'] = NewTrieNode()
			}
			node = node.children[word[i]-'a']
		}
		node.isWord = true
	}
}

func (d *MagicDictionary) search(word string) bool {
	return dfsSearch(d.root, word, 0, 0)
}

// root: 表示字段是当前的节点
// word: 当前需要搜索的单词
// i: 当前正在搜索单词的第几个字符
// edit: 当前的编辑距离
func dfsSearch(root *TrieNode, word string, i, edit int) bool {
	if root == nil {
		return false
	}
	if root.isWord && i == len(word) && edit == 1 {
		return true
	}
	if i < len(word) && edit <= 1 {
		found := false
		for j := 0; j < 26 && !found; j++ {
			var next int
			// 判断当前字符`word[i]`与字典中的字符`j`是否相同
			if j == int(word[i]-'a') {
				next = edit // 编辑距离不变
			} else {
				next = edit + 1 // 编辑距离加 1
			}
			found = dfsSearch(root.children[j], word, i+1, next)
		}
		return found
	}
	return false
}

// ---------- 最短的单词编码 ----------
func miniLengthEncoding(words []string) int {
	root := buildTireTree(words)
	var total int
	dfsSum(root, 1, &total)
	return total
}

// 如果找后缀，那么将单词反转后就可以用前缀树来表示了。
func buildTireTree(words []string) *TrieNode {
	root := NewTrieNode()
	for _, word := range words {
		node := root
		for i := len(word) - 1; i >= 0; i-- {
			if node.children[word[i]-'a'] == nil {
				node.children[word[i]-'a'] = NewTrieNode()
			}
			node = node.children[word[i]-'a']
		}
	}
	return root
}

func dfsSum(root *TrieNode, i int, total *int) {
	isLeaf := true
	for _, children := range root.children {
		if children != nil {
			isLeaf = false
			dfsSum(children, i+1, total)
		}
	}
	if isLeaf {
		*total += i
	}
}

// ---------- 单词之和 ----------
type TrieNodeWithValue struct {
	children []*TrieNodeWithValue
	value    int
}

func NewTrieNodeWithValue() *TrieNodeWithValue {
	return &TrieNodeWithValue{
		children: make([]*TrieNodeWithValue, 26),
		value:    0,
	}
}

type MapSum struct {
	root *TrieNodeWithValue
}

func NewMapSum() *MapSum {
	return &MapSum{root: NewTrieNodeWithValue()}
}

func (m *MapSum) insert(str string, n int) {
	node := m.root
	for i := 0; i < len(str); i++ {
		if node.children[str[i]-'a'] == nil {
			node.children[str[i]-'a'] = NewTrieNodeWithValue()
		}
		node = node.children[str[i]-'a']
	}
	node.value = n
}

func (m *MapSum) sum(prefix string) int {
	node := m.root
	for i := 0; i < len(prefix); i++ {
		if node.children[prefix[i]-'a'] == nil {
			return 0
		}
		node = node.children[prefix[i]-'a']
	}
	return getSum(node)
}

func getSum(root *TrieNodeWithValue) int {
	if root == nil {
		return 0
	}
	sum := root.value
	for _, children := range root.children {
		sum += getSum(children)
	}
	return sum
}

// ---------- 最大的异或 ----------
func findMaximumXORV1(nums []int) int {
	max := math.MinInt
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]^nums[j] > max {
				max = nums[i] ^ nums[j]
			}
		}
	}
	return max
}

// 前缀树的每个节点对应整数的一个数位，路径对应一个整数。
func findMaximumXORV2(nums []int) int {
	trie := NewBinaryTrie()
	for _, num := range nums {
		trie.insert(num)
	}
	max := math.MinInt
	for _, num := range nums {
		xor := trie.searchMax(num)
		max = int(math.Max(float64(max), float64(xor)))
	}
	return max
}

type BinaryTrieNode struct {
	children []*BinaryTrieNode
}

func NewBinaryTrieNode() *BinaryTrieNode {
	return &BinaryTrieNode{children: make([]*BinaryTrieNode, 2)}
}

type BinaryTrie struct {
	root *BinaryTrieNode
}

func NewBinaryTrie() *BinaryTrie {
	return &BinaryTrie{root: NewBinaryTrieNode()}
}

func (t *BinaryTrie) insert(num int) {
	node := t.root
	for i := 32; i >= 0; i-- {
		bit := (num >> i) & 1
		if node.children[bit] == nil {
			node.children[bit] = NewBinaryTrieNode()
		}
		node = node.children[bit]
	}
}

func (t *BinaryTrie) searchMax(num int) int {
	var xor int
	node := t.root
	for i := 32; i >= 0; i-- {
		bit := (num >> i) & 1
		if node.children[1-bit] != nil {
			xor = (xor << 1) + 1
			node = node.children[1-bit]
		} else {
			xor = xor << 1
			node = node.children[bit]
		}
	}
	return xor
}
