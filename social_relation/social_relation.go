package socialrelation

import (
	"fmt"
	"math/rand"
)

type Node struct {
	UserID  int              // 节点名称，使用用户ID
	Friends map[int]struct{} // 存放相连的朋友节点，便于确认和某个用户是否相连
	Degree  int              // 存放和给定的用户节点是几度好友
}

func NewNode(id int) *Node {
	return &Node{
		UserID:  id,
		Friends: make(map[int]struct{}),
		Degree:  0,
	}
}

func generate(userNum, relationNum int) []*Node {
	userNodes := make([]*Node, userNum)
	for i := 0; i < userNum; i++ {
		userNodes[i] = NewNode(i)
	}

	for i := 0; i < relationNum; i++ {
		friendAID := rand.Intn(userNum)
		friendBID := rand.Intn(userNum)
		if friendAID == friendBID {
			continue
		}
		friendA := userNodes[friendAID]
		friendB := userNodes[friendBID]
		friendA.Friends[friendBID] = struct{}{}
		friendB.Friends[friendAID] = struct{}{}
	}
	return userNodes
}

func BFS(users []*Node, id int) {
	// 防止数组越界
	if id > len(users) {
		return
	}

	// 广度优先搜索使用队列
	queue := make([]int, 0)
	// 添加初始节点
	queue = append(queue, id)
	// 存放已经被访问过的节点，防止回路
	visited := make(map[int]struct{})
	visited[id] = struct{}{}

	for len(queue) > 0 {
		// 取出队列头部的节点
		currentID := queue[0]
		queue = queue[1:]
		if users[currentID] == nil {
			continue
		}

		// 遍历当前节点所有直接相连的节点，并加入队列中
		for friendID := range users[currentID].Friends {
			if users[friendID] == nil {
				continue
			}
			if _, ok := visited[friendID]; ok {
				continue
			}
			queue = append(queue, friendID)
			visited[friendID] = struct{}{}
			users[friendID].Degree = users[currentID].Degree + 1
			fmt.Printf("%d 度好友: %d\n", users[friendID].Degree, friendID)
		}
	}
}
