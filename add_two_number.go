package leetcode

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{Val: 0, Next: nil}
	var cur *ListNode = list
	cur1 := l1
	cur2 := l2
	offset := 0
	var prev *ListNode = nil
	for ; cur1 != nil && cur2 != nil; cur1, cur2 = cur1.Next, cur2.Next {
		val := cur1.Val + cur2.Val + offset
		if val >= 10 {
			val -= 10
			offset = 1
		} else {
			offset = 0
		}
		if prev != nil {
			cur = &ListNode{Val: val, Next: nil}
			prev.Next = cur
		} else {
			cur.Val = val
		}
		prev = cur
	}

	for ; cur1 != nil; cur1 = cur1.Next {
		val := cur1.Val + offset
		if val >= 10 {
			val -= 10
			offset = 1
		} else {
			offset = 0
		}
		if prev != nil {
			cur = &ListNode{Val: val, Next: nil}
			prev.Next = cur
		} else {
			cur.Val = val
		}
		prev = cur
	}

	for ; cur2 != nil; cur2 = cur2.Next {
		val := cur2.Val + offset
		if val >= 10 {
			val -= 10
			offset = 1
		} else {
			offset = 0
		}
		if prev != nil {
			cur = &ListNode{Val: val, Next: nil}
			prev.Next = cur
		} else {
			cur.Val = val
		}
		prev = cur
	}

	if offset != 0 {
		cur = &ListNode{Val: offset, Next: nil}
		prev.Next = cur
	}

	return list
}
