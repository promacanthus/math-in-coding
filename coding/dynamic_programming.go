package coding

import "math"

// ---------- 爬楼梯的最少成本 -----------
// 从第i级台阶往上爬的最少成本应该是: 从第i-1级台阶往上爬的最少成本 和 从第i-2级台阶往上爬的最少成本的 较小值 再加 上爬第i级台阶的成本。
// 这个关系可以用状态转移方程表示为 `f(i) = min(f(i-1),f(i-2)) + cost[i]`。
func minCostClimbingStairsTopDown(cost []int) int {
	dp := make([]int, len(cost))
	mccsHelperWithCache(cost, len(cost)-1, dp)
	return IntMin(dp[len(cost)-1], dp[len(cost)-2])
}

func mccsHelper(cost []int, i int) int {
	if i < 2 {
		return cost[i]
	}
	return IntMin(mccsHelper(cost, i-1), mccsHelper(cost, i-2)) + cost[i]
}

// mccsHelperWithCache will cache the intermediate results.
func mccsHelperWithCache(cost []int, i int, dp []int) {
	if i < 2 {
		dp[i] = cost[i]
	} else if dp[i] == 0 {
		mccsHelperWithCache(cost, i-2, dp)
		mccsHelperWithCache(cost, i-1, dp)
		dp[i] = IntMin(dp[i-1], dp[i-2]) + cost[i]
	}
}

// minCostClimbingStairsBottomUp is a method that from bottom to top.
func minCostClimbingStairsBottomUp(cost []int) int {
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = IntMin(dp[i-1], dp[i-2]) + cost[i]
	}
	return IntMin(dp[len(cost)-1], dp[len(cost)-2])
}

func minCostClimbingStairs(cost []int) int {
	dp := []int{cost[0], cost[1]}
	for i := 2; i < len(cost); i++ {
		dp[i%2] = IntMin(dp[0], dp[1]) + cost[i]
	}
	return IntMin(dp[0], dp[1])
}

// ---------- 房屋偷盗 ----------
func robTopDown(nums []int) int {
	dp := make([]int, len(nums))
	topDown(nums, len(nums)-1, dp)
	return dp[len(nums)-1]
}

func topDown(nums []int, i int, dp []int) {
	if i == 0 {
		dp[i] = nums[0]
		return
	}

	if i == 1 {
		dp[i] = IntMax(nums[0], nums[1])
		return
	}

	topDown(nums, i-2, dp)
	topDown(nums, i-1, dp)
	dp[i] = IntMax(dp[i-1], dp[i-2]+nums[i])
}

func robBottomUp(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[1]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = IntMax(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = IntMax(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func robBottomUpOptimized(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := []int{nums[0], IntMax(nums[0], nums[1])}
	for i := 2; i < len(nums); i++ {
		dp[i%2] = IntMax(dp[(i-2)%2]+nums[i], dp[(i-1)%2])
	}
	return dp[(len(nums)-1)%2]
}

// f 不偷： f(i)=max(f(i-1),g(i-1)), f(0)=0
// g 偷：g(i)=f(i-1)+nums[i], g(0)=nums[0]
func robWithTwoDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	f := []int{0, 0}
	g := []int{nums[0], 0}
	for i := 1; i < len(nums); i++ {
		f[i%2] = IntMax(f[(i-1)%2], g[(i-1)%2])
		g[i%2] = f[(i-1)%2] + nums[i]
	}
	return IntMax(f[(len(nums)-1)%2], g[(len(nums)-1)%2])
}

// ---------- 环形房屋偷盗 ----------
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	f := make([]int, len(nums))
	g := make([]int, len(nums))
	robCycleTopDown(nums, 0, len(nums)-2, f)
	robCycleTopDown(nums, 1, len(nums)-1, g)
	return IntMax(f[len(nums)-2], g[len(nums)-1])
}

func robCycleTopDown(nums []int, start, end int, dp []int) {
	if end == start {
		dp[end] = nums[start]
		return
	}

	if end-1 == start {
		dp[end-1] = IntMax(nums[start], nums[end-1])
		return
	}

	robCycleTopDown(nums, start, end-2, dp)
	robCycleTopDown(nums, start, end-1, dp)
	dp[end] = IntMax(dp[end-2]+nums[end], dp[end-1])
}

func robCycle(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return IntMax(robCycleBottomUp(nums, 0, len(nums)-2), robCycleBottomUp(nums, 1, len(nums)-1))
}

func robCycleBottomUp(nums []int, start, end int) int {
	dp := []int{nums[start], IntMax(nums[start], nums[start+1])}
	for i := start + 2; i <= end; i++ {
		j := i - start
		dp[j%2] = IntMax(dp[(j-1)%2], dp[(j-2)%2]+nums[i])
	}
	return dp[(end-start)%2]
}

// ----------- 粉刷房子 ----------
// r 房子刷成红色的成本：r(i)=min(g(i-1),b(i-1))+cost[i][0]; r(0)=cost[0][0]
// g 房子刷成绿色的成本：g(i)=min(r(i-1),b(i-1))+cost[i][1]; g(0)=cost[0][1]
// b 房子刷成蓝色的成本：b(i)=min(r(i-1),g(i-1))+cost[i][2]; b(0)=cost[0][2]
func minCostTopDown(cost [][]int) int {
	r := make([]int, len(cost))
	g := make([]int, len(cost))
	b := make([]int, len(cost))
	minCostHelper(cost, len(cost)-1, r, g, b)
	return IntsMin(r[len(cost)-1], g[len(cost)-1], b[len(cost)-1])
}

func minCostHelper(cost [][]int, i int, r, g, b []int) {
	if i == 0 {
		r[0] = cost[0][0]
		g[0] = cost[0][1]
		b[0] = cost[0][2]
		return
	}
	minCostHelper(cost, i-1, r, g, b)
	r[i] = IntMin(g[i-1], b[i-1]) + cost[i][0]
	g[i] = IntMin(r[i-1], b[i-1]) + cost[i][1]
	b[i] = IntMin(r[i-1], g[i-1]) + cost[i][2]
}

func minCostBottomUp(cost [][]int) int {
	r := []int{cost[0][0], 0}
	g := []int{cost[0][1], 0}
	b := []int{cost[0][2], 0}

	for i := 1; i < len(cost); i++ {
		idx := (i - 1) % 2
		r[i%2] = IntMin(g[idx], b[idx]) + cost[i][0]
		g[i%2] = IntMin(r[idx], b[idx]) + cost[i][1]
		b[i%2] = IntMin(r[idx], g[idx]) + cost[i][2]
	}
	last := (len(cost) - 1) % 2
	return IntsMin(r[last], g[last], b[last])
}

// ---------- 反转字符串 ----------
// f 反转后最后一个字符为0的反转次数:
// - 第i个位置是0：f(i)=f(i-1); f(0)=0
// - 第i个位置是1：f(i)=f(i-1)+1; f(0)=1
// g 反转后最后一个字符为1的反转次数:
// - 第i个位置是0：g(i)=min(f(i-1),g(i-1))+1 ; g(0)=1
// - 第i个位置是1：g(i)=min(f(i-1),g(i-1)); g(0)=0
func minFlipsMonoIncrTopDown(str string) int {
	f := make([]int, len(str))
	g := make([]int, len(str))
	minFlipsHelper(str, len(str)-1, f, g)
	return IntMin(f[len(str)-1], g[len(str)-1])
}

func minFlipsHelper(str string, i int, f, g []int) {
	if i == 0 {
		switch str[0] {
		case '0':
			f[0] = 0
			g[0] = 1
		case '1':
			f[0] = 1
			g[0] = 0
		}
		return
	}
	minFlipsHelper(str, i-1, f, g)
	switch str[i] {
	case '0':
		f[i] = f[i-1]
		g[i] = IntMin(f[i-1], g[i-1]) + 1
	case '1':
		f[i] = f[i-1] + 1
		g[i] = IntMin(f[i-1], g[i-1])
	}
}

func minFlipsMonoIncrBottomUp(str string) int {
	f, g := make([]int, 2), make([]int, 2)
	switch str[0] {
	case '0':
		f[0] = 0
		g[0] = 1
	case '1':
		f[0] = 1
		g[0] = 0
	}

	for i := 1; i < len(str); i++ {
		idx := (i - 1) % 2
		switch str[i] {
		case '0':
			f[i%2] = f[idx]
			g[i%2] = IntMin(f[idx], g[idx]) + 1
		case '1':
			f[i%2] = f[idx] + 1
			g[i%2] = IntMin(g[idx], f[idx])
		}
	}
	last := (len(str) - 1) % 2
	return IntMin(f[last], g[last])
}

// ---------- 斐波那契列 ----------
// f(i,j)=f(j,k)+1; i表示当前数组，j表示前一个数字，k表示更前一个数字
func findFibonacci(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	set := make(map[int]int)
	for i, num := range nums {
		set[num] = i
	}
	res := 2
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			k, ok := set[nums[i]-nums[j]]
			if !ok || k >= j {
				dp[i][j] = 2
			}
			dp[i][j] = IntMax(dp[i][j], dp[j][k]+1)
			res = IntMax(res, dp[i][j])
		}
	}
	if res > 2 {
		return res
	}
	return 0
}

// ---------- 最少回文分割 ----------
func minCut(s string) int {
	// 列表示的是每个子串的长度
	// 行表示的是整个字符串的长度
	// isPal[j][i]：表示子字符串s[j...i]是否回文
	isPal := make([][]bool, len(s))
	for i := range isPal {
		isPal[i] = make([]bool, len(s))
	}
	// Checks if all substrings are palindromic.
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if s[j] == s[i] &&
				((j+1) >= i ||
					// 子字符串s[j+1...i-1]是回文
					isPal[j+1][i-1]) {
				isPal[j][i] = true
			}
		}
	}

	dp := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if isPal[0][i] {
			// s[0...i]是回文，不需要分割
			dp[i] = 0
		} else {
			dp[i] = i // s[0...1]最多分割i次，所有的子串都是一个字母，那么是回文
			for j := 1; j <= i; j++ {
				if isPal[j][i] {
					// s[j...i]是回文，比较s[0...j-1]的分割自出和最多分割次数哪个更小
					dp[i] = IntMin(dp[i], dp[j-1]+1)
				}
			}
		}
	}
	return dp[len(s)-1]
}

// ---------- 最长公共子序列 -----------
// 用函数f(i,j)表示第1个字符串中下标从0到i的子字符串(记为s1[0..i])
// 和第2个字符串中下标从0到j的子字符串(记为s2[0..j])的最长公共子序列的长度。
// 如果第1个字符串的长度是m，第2个字符串的长度是n，那么f(m-1,n-1)就是整个问题的解。
//
// 如果第1个字符串中下标为i的字符(记为s1[i])与第2个字符串中下标为j(记为s2[j])的字符相同，
// 那么f(i,j)相当于在s1[0..i-1]和s2[0..j-1]的最长公共子序列的后面添加一个公共字符，也就是f(i,j)=f(i-1,j-1)+1。
//
// 如果字符s1[i]与字符s2[j]不相同，则这两个字符不可能同时出现在s1[0..i]和s2[0..j]的公共子序列中。
// 此时s1[0..i]和s2[0..j]的最长公共子序列要么是s1[0..i-1]和s2[0..j]的最长公共子序列，
// 要么是s1[0..i]和s2[0..j-1]的最长公共子序列。也就是说，此时f(i,j)=max(f(i-1,j),f(i,j-1))。
func longestCommonSubsequence(s1, s2 string) int {
	len1, len2 := len(s1), len(s2)
	if len1 < len2 {
		return longestCommonSubsequence(s2, s1)
	}

	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}

	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if s1[i] == s2[j] {
				dp[(i+1)%2][j+1] = dp[i%2][j] + 1
			} else {
				dp[(i+1)%2][j+1] = IntMax(dp[(i+1)%2][j], dp[i%2][j+1])
			}
		}
	}
	return dp[len1%2][len2]
}

// ---------- 字符串交织 ----------
func isInterleave(s1, s2, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l1+l2 != l3 {
		return false
	}

	dp := make([][]bool, l1+1)
	for i := range dp {
		dp[i] = make([]bool, l2+1)
	}
	dp[0][0] = true
	for i := 0; i < l1; i++ {
		dp[i+1][0] = s1[i] == s3[i] && dp[i][0]
	}
	for j := 0; j < l2; j++ {
		dp[0][j+1] = s2[j] == s3[j] && dp[0][j]
	}
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			dp[i+1][j+1] = (s1[i] == s3[i+j+1] && dp[i][j+1]) || (s2[j] == s3[i+j+1] && dp[i+1][j])
		}
	}
	return dp[l1][l2]
}

// ---------- 子序列的数目 ----------
// 用f(i,j)表示字符串S下标从0到i的子字符串(记为S[0..i])中等于字符串T下标从0到j的子字符串(记为T[0..j])的子序列的数目。
// 如果字符串S的长度是m，字符串T的长度是n，那么f(m-1,n-1)就是字符串S中等于字符串T的子序列的数目。
//
// 如果字符串S中下标为i的字符(记为S[i])等于字符串T中下标为j的字符(记为T[j])，那么对S[i]有两个选择：
//  1. 用S[i]去匹配T[j]，那么S[0..i]中等于T[0..j]的子序列的数目等于S[0..i-1]中等于T[0..j-1]的子序列的数目；
//  2. 舍去S[i]，那么S[0..i]中等于T[0..j]的子序列的数目等于S[0..i-1]中等于T[0..j]的子序列的数目。
//
// 因此，当S[i]等于T[j]时，f(i,j)=f(i-1,j-1)+f(i-1,j)。
// 如果S[i]和T[j]不相同，则只能舍去S[i]，此时f(i,j)=f(i-1,j)。
func numDistinct(s, t string) int {
	n, m := len(s), len(t)
	if n < m {
		return 0
	}
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	// An empty substring of s matches an empty substring of t.
	dp[0][0] = 1
	// If t is empty, the substring number is 1.
	for i := 1; i < n; i++ {
		dp[i%2][0] = 1
	}
	// If s is empty, the substring number is 0.
	for j := 1; j < m; j++ {
		dp[0][j] = 0
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= i && j < m; j++ {
			if s[i] == t[j] {
				dp[(i+1)%2][j+1] = dp[i%2][j] + dp[i%2][j+1]
			} else {
				dp[(i+1)%2][j+1] = dp[i%2][j+1]
			}
		}
	}
	return dp[n%2][m]
}

// ---------- 路径的数目 ----------
// 用函数f(i,j)表示从格子的左上角坐标为(0,0)的位置出发到达坐标为(i,j)的位置的路径的数目。
// 如果格子的大小为m×n，那么f(m-1,n-1)就是问题的解。
func uniquePathsTopDown(m, n int) int {
	// 记录结果避免重复计算
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	return upHelper(m-1, n-1, dp)
}

func upHelper(i, j int, dp [][]int) int {
	if dp[i][i] == 0 {
		// 判断是否已经计算过结果
		if i == 0 || j == 0 {
			// 在顶部或者左侧，只有一种走法
			dp[i][j] = 1
		} else {
			// 向下或者向右
			dp[i][j] = upHelper(i-1, j, dp) + upHelper(i, j-1, dp)
		}
	}
	return dp[i][j]
}

func uniquePathsBottomUp(m, n int) int {
	dp := InitDP(m, n)
	for i := 0; i < m; i++ {
		dp[i%2][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i%2][j] = dp[(i-1)%2][j] + dp[i%2][j-1]
		}
	}
	return dp[(m-1)%2][n-1]
}

// ---------- 最小路径和 ----------
// 用函数f(i,j)表示从格子的左上角坐标为(0,0)的位置(用grid[0][0]表示)出发到达坐标为(i,j)的位置(用grid[i][j]表示)的路径的数字之和的最小值。
// 如果格子的大小为m×n，那么f(m-1,n-1)就是问题的解。
func minPathSumBottomUp(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := InitDP(2, n)
	dp[0][0] = grid[0][0]
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		dp[i%2][0] = dp[(i-1)%2][0] + grid[i][0]
		for j := 1; j < n; j++ {
			dp[i%2][j] = IntMin(dp[(i-1)%2][j], dp[i%2][j-1]) + grid[i][j]
		}
	}
	return dp[(m-1)%2][n-1]
}

func minPathSumTopDown(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := InitDP(m, n)
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	return mpsHelper(m-1, n-1, dp, grid)
}

func mpsHelper(i, j int, dp, grid [][]int) int {
	if dp[i][j] == 0 && i != 0 && j != 0 {
		dp[i][j] = IntMin(mpsHelper(i-1, j, dp, grid), mpsHelper(i, j-1, dp, grid)) + grid[i][j]
	}
	return dp[i][j]
}

// ---------- 三角形中最小路径之和 ----------
// 在左端对齐的三角形中，从一个数字出发，下一步要么前往下一行正下方的数字，要么前往右下方的数字。
// 可以用f(i,j)表示从三角形的顶部出发到达行号和列号分别为i和j(i≥j)的位置时路径数字之和的最小值，
// 同时用T[i][j]表示三角形行号和列号分别为i和j的数字。如果三角形中包含n行数字，那么f(n-1,j)的最小值就是整个问题的最优解。
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	dp := InitDP(m, m)
	dp[0][0] = triangle[0][0]
	for i := 1; i < m; i++ {
		dp[i%2][0] = dp[(i-1)%2][0] + triangle[i][0]
		for j := 1; j <= i; j++ {
			if j == i {
				dp[i%2][j] = dp[(i-1)%2][j-1] + triangle[i][j]
			} else {
				dp[i%2][j] = IntMin(dp[(i-1)%2][j], dp[(i-1)%2][j-1]) + triangle[i][j]
			}
		}
	}
	min := math.MaxInt
	for _, v := range dp[(m-1)%2] {
		if v < min {
			min = v
		}
	}
	return min
}

// ---------- 分割等和子集 ---------
// 用函数f(i,j)表示能否从前i个物品(物品标号分别为0,1,…,i-1)中选择若干物品放满容量为j的背包。
// 如果总共有n个物品，背包的容量为t，那么f(n,t)就是问题的解。
// 对标号为i-1的物品有两个选择。
// 1. 将标号为i-1的物品放入背包中，如果能从前i-1个物品(物品标号分别为0,1,…,i-2)中选择若干物品放满容量为j-nums[i-1]的背包(即f(i-1,j-nums[i-1])为true)，那么f(i,j)就为true。
// 2. 不将标号为i-1的物品放入背包中，如果从前i-1个物品中选择若干物品放满容量为j的背包(即f(i-1,j)为true)，那么f(i,j)也为true。
func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	return subsetSumTopDown(nums, sum/2) && subsetSumBottomUp(nums, sum/2)
}

func subsetSumTopDown(nums []int, target int) bool {
	dp := make([][]*bool, len(nums)+1)
	for i := range dp {
		dp[i] = make([]*bool, target+1)
	}
	return *sssHelper(nums, dp, len(nums), target)
}

func sssHelper(nums []int, dp [][]*bool, i, j int) *bool {
	if dp[i][j] == nil {
		if j == 0 {
			dp[i][j] = BooleanPointer(true)
		} else if i == 0 {
			dp[i][j] = BooleanPointer(false)
		} else {
			dp[i][j] = sssHelper(nums, dp, i-1, j) // 不选第i个物品
			if !(*dp[i][j]) && j >= nums[i-1] {    // 选第i个物品
				dp[i][j] = sssHelper(nums, dp, i-1, j-nums[i-1])
			}
		}
	}
	return dp[i][j]
}

func subsetSumBottomUp(nums []int, target int) bool {
	dp := make([][]bool, len(nums)+1)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	for i := 0; i < len(nums); i++ {
		dp[i%2][0] = true
	}
	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= target; j++ {
			dp[i%2][j] = dp[(i-1)%2][j]
			if !dp[i%2][j] && j >= nums[i-1] {
				dp[i%2][j] = dp[(i-1)%2][j-nums[i-1]]
			}
		}
	}
	return dp[(len(nums))%2][target]
}

// ---------- 加减的目标值 ----------
func findTargetSumWays(nums []int, s int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if (sum+s)%2 == 1 || sum < s {
		return 0
	}
	return subSetSumBottomUp(nums, (s+sum)/2)
}

func subSetSumBottomUp(nums []int, target int) int {
	dp := InitDP(len(nums)+1, target+1)
	for i := 0; i <= len(nums); i++ {
		dp[i%2][0] = 1
	}
	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= target; j++ {
			if j >= nums[i-1] {
				dp[i%2][j] = dp[(i-1)%2][j] + dp[(i-1)%2][j-nums[i-1]]
			}
		}
	}
	return dp[(len(nums))%2][target]
}

// ---------- 最少硬币数目 ----------
// 用函数f(i)表示凑出总额为i的硬币需要的最少数目。
// 这个函数只有一个参数，表示硬币的总额。如果目标总额为t，那么f(t)就是整个问题的解。
// - 在总额为i-coins[0]的硬币中添加1枚标号为0的硬币，此时f(i)等于f(i-coins[0])+1
// - 在总额为i-coins[1]的硬币中添加1枚标号为1的硬币，此时f(i)等于f(i-coins[1])+1
// 以此类推，
// - 在总额为i-coins[n-1]的硬币中添加1枚标号为n-1的硬币，此时f(i)等于f(i-coins[n-1])+1。
// 因为目标是计算凑出总额为i的硬币，所以f(i)是上述所有情况的最小值。
// 该状态转移方程可以表示为f(i)=min(f(i-coins[j])+1)(coins[j]≤i)。
func coinChange(coins []int, target int) int {
	dp := make([]int, target+1)
	for i := 1; i <= target; i++ {
		dp[i] = target + 1 // 初始化一个稍微大一点的值
		for _, coin := range coins {
			if i >= coin {
				dp[i] = IntMin(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[target] > target {
		return -1
	}
	return dp[target]
}

// ---------- 排列的数目 ----------
// 用f(i)表示和为i的排列的数目。
// 在和为i-nums[0]的排列中添加标号为0的数字，此时f(i)等于f(i-nums[0])；
// 在和为i-nums[1]的排列中添加标号为1的数字，此时f(i)等于f(i-nums[1])。
// 以此类推，
// 在和为i-nums[n-1]的排列中添加标号为n-1的数字(n为数组的长度)，此时f(i)等于f(i-nums[n-1])。
// 因为目标是求出所有和为i的排列的数目，所以将上述所有情况全部累加起来。
// 该状态转移方程可以表示为f (i)=∑f (i-nums[j])(nums[j]≤i)
// 由于只有一个空排列的数字之和等于0，因此f(0)等于1。
func permutationSum(nums []int, target int) int {
	dp := make([]int, len(nums)+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
