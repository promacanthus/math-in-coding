package prefixtree

import "fmt"

type Node struct {
	Label       byte           // 节点的名称，在前缀树里是单个字母
	Sons        map[byte]*Node // 使用哈希表存放子节点，便于确认是否已经添加过某个字母对应的节点
	Prefix      string         // 从树的根到当前节点这条通路上，全部字母所组成的前缀
	Explanation string         // 词条的解释
}

func NewNode(label byte, prefix, exp string) *Node {
	return &Node{
		Label:       label,
		Sons:        make(map[byte]*Node),
		Prefix:      prefix,
		Explanation: exp,
	}
}

// DFSByStack implements a depth-first search (DFS) algorithm using a stack data structure. It takes a 
// root node as input and performs the DFS traversal. It starts by creating an empty stack and pushing 
// the root node onto it. Then, it enters a loop that continues until the stack is empty. In each iteration, 
// it takes the top node from the stack and removes it. If the current node has no sons (i.e., it is a leaf 
// node), it prints the concatenation of the node's prefix and label, and then returns. Otherwise, if the 
// current node has sons, it pushes each son onto the stack. The loop continues until all nodes have been 
// visited and processed. It is important to note that this implementation assumes that the Node struct 
// has a "Prefix" and "Label" field, as well as a "Sons" field which is a slice of Node pointers. 
func DFSByStack(root *Node) {
	stack := make([]*Node, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(node.Sons) == 0 {
			fmt.Println(node.Prefix + string(node.Label))
			return
		} else {
			for _, node := range node.Sons {
				stack = append(stack, node)
			}
		}
	}
}

// DFSByRecursion is a depth-first search algorithm implemented using recursion. It takes a root node
// as input and performs a depth-first traversal of the tree starting from the root. It first checks
// if the root node is nil, which indicates an empty tree. If the tree is empty, the function simply
// returns. Next, it checks if the root node has any child nodes. If the root node has no child nodes,
// it means it is a leaf node. In this case, it prints the prefix followed by the label of the leaf
// node and returns. If the root node has child nodes, it iterates over each child node using a for
// loop. For each child node, it recursively calls the DFSByRecursion function. This recursive call
// allows the function to traverse deeper into the tree. Overall, it performs a depth-first search
// by recursively visiting each node in the tree, starting from the root node and traversing through
// its child nodes. It prints the labels of the leaf nodes in the order they are visited.
func DFSByRecursion(root *Node) {
	if root == nil {
		return
	}

	if len(root.Sons) == 0 {
		fmt.Println(root.Prefix + string(root.Label))
		return
	}

	for _, node := range root.Sons {
		DFSByRecursion(node)
	}
}
