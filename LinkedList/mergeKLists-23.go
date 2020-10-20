package LinkedList

//通过俩俩合并

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var rs *ListNode
	for _, value := range lists {
		rs = merge(rs, value)
	}
	return rs

}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	rs := &ListNode{}
	cur := rs

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next

		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l2 != nil {
		cur.Next = l2
	}
	if l1 != nil {
		cur.Next = l1
	}
	return rs.Next
}

//分治方法
func mergeKLists2(lists []*ListNode) *ListNode {

	return merge2(lists, 0, len(lists)-1)
}

func merge2(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	//
	if l > r {
		return nil
	}
	mid := l + (r-l)/2
	//假设只有一个链表元素，mid + 1就越界
	return merge(merge2(lists, l, mid), merge2(lists, mid+1, r))
}
