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
