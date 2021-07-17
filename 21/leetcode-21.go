package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummy := &ListNode{}
	head := dummy

	for l1 != nil && l2 != nil {
		//可以将l1或者l2直接赋值给head.next
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	if l1 != nil {
		head.Next = l1
	}

	if l2 != nil {
		head.Next = l2
	}
	return dummy.Next

}
