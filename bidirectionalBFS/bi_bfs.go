package bidirectionalbfs

type Node struct {
	UserID  int
	Friends map[int]struct{}
	degree  map[int]int // 存放从不同用户出发，到当前用户是几度
}

func NewNode(userID int) *Node {
	return &Node{
		UserID:  userID,
		Friends: make(map[int]struct{}),
		degree:  map[int]int{userID: 0},
	}
}

// biBFS implements a bidirectional breadth-first search algorithm to find the shortest path
// between two nodes in a graph represented by a list of users.
//
// It takes three parameters:
//
//   - `users` , which is a slice of  `Node`  objects representing the users in the graph,
//   - `id1`  and  `id2` , which are the IDs of the two nodes between which the shortest path needs to be found.
//
// It first checks if the given IDs are valid, i.e., within the range of the  `users`  slice.
// If not, it returns -1 indicating an error. Then, two queues  `queueA`  and  `queueB`  are
// initialized with the starting nodes  `id1`  and  `id2`  respectively. Two visited maps
// `visitedA`  and  `visitedB`  are also initialized to keep track of the visited nodes in
// each direction. It then starts a loop that continues until the sum of the degrees of
// exploration in both directions exceeds a maximum degree limit (20 in this case).
// In each iteration of the loop, the degrees of exploration in both directions are incremented.
// For each degree of exploration, the function calls the  `getNextDegreeFriend`  function to
// find the next degree+1 friends of the starting node in each direction. These friends
// are added to the respective queues and marked as visited in the corresponding visited map.
// After each degree of exploration, the function checks if there is an overlap between
// the visited nodes in both directions using the  `hasOverlap`  function. If an overlap is
// found, it means a path between the two nodes has been found and the function returns the
// sum of the degrees of exploration as the shortest path length. If no path is found within
// the maximum degree limit, the function returns -1 indicating that there is no path between the two nodes.
func biBFS(users []*Node, id1, id2 int) int {
	if id1 > len(users) || id2 > len(users) {
		return -1
	}
	if id1 == id2 {
		return 0
	}

	queueA := make([]int, 0)
	queueA = append(queueA, id1)
	visitedA := make(map[int]bool)
	visitedA[id1] = true

	queueB := make([]int, 0)
	queueB = append(queueB, id2)
	visitedB := make(map[int]bool)
	visitedB[id2] = true

	degreeA, degreeB := 0, 0
	maxDegree := 20 // 防止两者之间不存在通路
	for degreeA+degreeB < maxDegree {
		degreeA++
		// 沿着A的方向广度优先搜索degree+1的好友
		getNextDegreeFriend(id1, degreeA, users, queueA, visitedA)
		// 判断到目前为止，被发现的a的好友，和被发现的b的好友，两个集合是否存在交集
		if hasOverlap(visitedA, visitedB) {
			return degreeA + degreeB
		}

		degreeB++
		// 沿着B的方向广度优先搜索degree+1的好友
		getNextDegreeFriend(id2, degreeB, users, queueB, visitedB)
		// 判断到目前为止，被发现的a的好友，和被发现的b的好友，两个集合是否存在交集
		if hasOverlap(visitedA, visitedB) {
			return degreeA + degreeB
		}
	}
	return -1
}

// getNextDegreeFriend takes in five parameters:
//
//   - `id`  (an integer representing the ID of the starting user),
//   - `degree`  (an integer representing the degree of separation),
//   - `users`  (a slice of pointers to  `Node`  objects representing all the users),
//   - `queue`  (a slice of integers representing the users to be processed),
//   - `visited`  (a map of integers to booleans representing the users that have been visited).
//
// It iterates over the  `queue`  until it is empty. For each user ID in the  `queue` ,
// it checks if the user has already been visited. If so, it continues to the next iteration.
// If not, it iterates over the friends of the current user and checks if the degree of
// separation between the friend and the current user is equal to the desired  `degree` .
// If so, it marks the friend as visited by setting its corresponding value in the  `visited`  map to true.
func getNextDegreeFriend(id, degree int, users []*Node, queue []int, visited map[int]bool) {
	for len(queue) > 0 {
		currentID := queue[0]
		queue = queue[1:]
		if visited[currentID] {
			continue
		}
		for friendID := range users[currentID].Friends {
			if users[friendID].degree[currentID] == degree {
				visited[friendID] = true
			}
		}
	}
}

// hasOverlap checks for any overlap between two sets of visited IDs.
// It takes two parameters: visitedA and visitedB, which are maps with
// IDs as keys and boolean values indicating whether the ID has been
// visited or not. It iterates over the keys in visitedA, and if it finds
// an ID that also exists in visitedB, it returns true to indicate an
// overlap. If no overlap is found after iterating through all the keys,
// it returns false.
//
// In other words, it determines if there are any common
// IDs between the two sets of visited IDs. If there is at least one ID
// that appears in both sets, it returns true; otherwise, it returns false.
func hasOverlap(visitedA, visitedB map[int]bool) bool {
	for id := range visitedA {
		if visitedB[id] {
			return true
		}
	}
	return false
}
