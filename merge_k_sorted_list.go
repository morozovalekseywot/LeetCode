package leetcode

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	result := ListNode{-938129192, nil}
	cur := &result

	for true {
		idx := -1
		for i := 0; i < len(lists); i++ {
			if lists[i] != nil {
				if idx == -1 {
					idx = i
				} else if lists[i].Val < lists[idx].Val {
					idx = i
				}
			}
		}
		if idx == -1 {
			break
		}

		cur.Next = &ListNode{Val: lists[idx].Val, Next: nil}
		cur = cur.Next
		lists[idx] = lists[idx].Next
	}

	return result.Next
}
