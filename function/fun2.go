package function

import "learn/common"

// 201 Bitwise AND of Numbers Range
func RangeBitwiseAnd(left int, right int) int {
	i := 0
	for left != right {
		left >>= 1
		right >>= 1
		i++
	}
	return left << i
}

// 202. Happy Number
func IsHappy(n int) bool {
	t := make(map[int]struct{})
	return IsHappySub(n, t)
}

func IsHappySub(n int, t map[int]struct{}) bool {
	if n == 1 {
		return true
	}
	if n == 0 {
		return false
	}
	if _, ok := t[n]; ok {
		return false
	}
	t[n] = struct{}{}

	next := 0
	for n > 0 {
		s := n % 10
		next += s * s
		n /= 10
	}
	return IsHappySub(next, t)
}

// 203. Remove Linked List Elements
func RemoveElements(head *common.ListNode, val int) *common.ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}

	if head != nil {
		node := head
		for node.Next != nil {
			if node.Next.Val == val {
				node.Next = node.Next.Next
			} else {
				node = node.Next
			}
		}
	}

	return head
}

// 204. Count Primes
func CountPrimes(n int) int {
	res := 0
	t := make([]bool, n)

	for i := 2; i < n; i++ {
		if t[i-1] {
			continue
		}

		for d := 2; i*d <= n; d++ {
			t[i*d-1] = true
		}

		res++
	}

	return res
}

// 205. Isomorphic Strings
func IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	tras := make([]byte, 256)
	retras := make([]byte, 256)
	for i, c1 := range s {
		c2 := t[i]
		if tras[c1] == 0 {
			tras[c1] = c2
		} else if tras[c1] != c2 {
			return false
		}

		if retras[c2] == 0 {
			retras[c2] = byte(c1)
		} else if retras[c2] != byte(c1) {
			return false
		}
	}
	return true
}

// 206. Reverse Linked List
func ReverseList(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}

	first := head
	second := head.Next
	first.Next = nil

	for second != nil {
		tmp := second.Next
		second.Next = first
		first = second
		second = tmp
	}
	return first
}

// 207. Course Schedule
func CanFinish(numCourses int, prerequisites [][]int) bool {
	visited := make([]byte, numCourses)
	limit := make([][]int, numCourses)

	for _, sub := range prerequisites {
		limit[sub[0]] = append(limit[sub[0]], sub[1])
	}

	for i := 0; i < numCourses; i++ {
		if !canFinishSub(limit, visited, i) {
			return false
		}
	}
	return true
}

func canFinishSub(limit [][]int, visited []byte, index int) bool {
	if visited[index] == 1 {
		return true
	}

	if visited[index] == 2 {
		return false
	}

	for _, sub := range limit[index] {
		visited[index] = 2
		if !canFinishSub(limit, visited, sub) {
			return false
		}
	}

	visited[index] = 1
	return true
}

// 209. Minimum Size Subarray Sum
func MinSubArrayLen(target int, nums []int) int {
	min := common.MAXINTNUM
	left := 0
	right := 0
	sum := 0

	for _, cur := range nums {
		sum += cur
		right++
		if sum < target {
			continue
		}

		for sum >= target {
			if right-left < min {
				min = right - left
			}
			sum -= nums[left]
			left++
		}
	}

	if min == common.MAXINTNUM {
		min = 0
	}

	return min
}

// 210. FindOrder
func FindOrder(numCourses int, prerequisites [][]int) []int {
	OrderBy := make(map[int]map[int]interface{})
	res := []int{}

	for i := 0; i < numCourses; i++ {
		OrderBy[i] = map[int]interface{}{}
	}

	for _, prerequisite := range prerequisites {
		if _, ok := OrderBy[prerequisite[0]]; ok {
			OrderBy[prerequisite[0]][prerequisite[1]] = struct{}{}
		} else {
			OrderBy[prerequisite[0]] = map[int]interface{}{prerequisite[1]: struct{}{}}
		}
	}

	ctn := true
	delCourses := make(map[int]interface{})
	for ctn {
		ctn = false
		for courses, orders := range OrderBy {
			for order := range orders {
				if _, ok := delCourses[order]; ok {
					delete(orders, order)
				}
			}

			if len(orders) == 0 {
				res = append(res, courses)
				ctn = true
				delCourses[courses] = struct{}{}
				delete(OrderBy, courses)
			}
		}
	}

	if len(res) != numCourses {
		res = []int{}
	}

	return res
}

// // FindWords 单次搜索（时间复杂度较高，废弃）
// func FindWordsAb(board [][]byte, words []string) []string {
// 	var res []string
// 	var exist [][]bool

// 	root := &common.DictNode{
// 		Childs: make([]*common.DictNode, common.BIG_ENGLISH_CHAR_NUM),
// 	}

// 	exist = make([][]bool, len(board))
// 	for i, line := range board {
// 		exist[i] = make([]bool, len(line))
// 	}

// 	for _, word := range words {
// 		if findWordByDictAb(word, root) {
// 			res = append(res, word)
// 			continue
// 		}

// 		for i, line := range board {
// 			for j := range line {
// 				if findWordSubAb(board, exist, word, root, i, j) {
// 					res = append(res, word)
// 					goto end
// 				}
// 			}
// 		}
// 	end:
// 	}

// 	return res
// }

// // findWordByDict 查询字典是否存在（时间复杂度较高，废弃）
// func findWordByDictAb(word string, root *common.DictNode) bool {
// 	for _, byt := range word {
// 		if root.Childs[byt-'a'] == nil {
// 			return false
// 		}
// 		root = root.Childs[byt-'a']
// 	}

// 	return true
// }

// // findWordSub 递归查询 （时间复杂度较高，废弃）
// func findWordSubAb(board [][]byte, exist [][]bool, word string, root *common.DictNode, i, j int) bool {
// 	if len(word) == 0 {
// 		return true
// 	}

// 	if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) || exist[i][j] {
// 		return false
// 	}

// 	if board[i][j] == word[0] {
// 		root.Childs[word[0]-'a'] = &common.DictNode{
// 			Childs: make([]*common.DictNode, common.BIG_ENGLISH_CHAR_NUM),
// 		}

// 		exist[i][j] = true
// 		ok := findWordSubAb(board, exist, word[1:], root.Childs[word[0]-'a'], i+1, j) ||
// 			findWordSubAb(board, exist, word[1:], root.Childs[word[0]-'a'], i, j+1) ||
// 			findWordSubAb(board, exist, word[1:], root.Childs[word[0]-'a'], i-1, j) ||
// 			findWordSubAb(board, exist, word[1:], root.Childs[word[0]-'a'], i, j-1)

// 		exist[i][j] = false
// 		return ok
// 	}

// 	return false
// }

// FindWords 单次搜索
func FindWords(board [][]byte, words []string) []string {
	var res []string
	var exist [][]bool

	root := &common.DictNode{
		Childs: make([]*common.DictNode, 26),
	}

	exist = make([][]bool, len(board))
	for i, line := range board {
		exist[i] = make([]bool, len(line))
	}

	maxLenth := 0
	for _, word := range words {
		if len(word) > maxLenth {
			maxLenth = len(word)
		}
	}

	for i, line := range board {
		for j := range line {
			buildWordTree(board, exist, root, i, j, 0, maxLenth)
		}
	}

	for _, word := range words {
		if findWordByDict(word, root) {
			res = append(res, word)
			continue
		}
	}

	return res
}

// findWordByDict 查询字典是否存在
func findWordByDict(word string, root *common.DictNode) bool {
	for _, byt := range word {
		if root.Childs[byt-'a'] == nil {
			return false
		}
		root = root.Childs[byt-'a']
	}

	return true
}

// buildWordTree 字典树搭建
func buildWordTree(board [][]byte, exist [][]bool, root *common.DictNode, i, j, high, maxLenth int) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) || exist[i][j] || high >= maxLenth {
		return
	}

	curNode := root.Childs[board[i][j]-'a']
	if curNode == nil {
		curNode = &common.DictNode{
			Childs: make([]*common.DictNode, 26),
		}

		root.Childs[board[i][j]-'a'] = curNode
	}

	exist[i][j] = true
	buildWordTree(board, exist, curNode, i+1, j, high+1, maxLenth)
	buildWordTree(board, exist, curNode, i, j+1, high+1, maxLenth)
	buildWordTree(board, exist, curNode, i-1, j, high+1, maxLenth)
	buildWordTree(board, exist, curNode, i, j-1, high+1, maxLenth)
	exist[i][j] = false
	return
}
