package LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//插入排序

func insertionSortList(head *ListNode) *ListNode {
	/*	dummy := &ListNode{}
		curr := dummy
		curr.Next = head*/
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	tail := head
	sort := head.Next

	for sort != nil {
		if sort.Val < tail.Val {
			pos := dummy
			for pos.Next.Val < sort.Val {
				pos = pos.Next
			}
			tail.Next = sort.Next
			sort.Next = pos.Next
			pos.Next = sort
			sort = tail.Next
		} else {
			tail = tail.Next
			sort = sort.Next
		}
	}
	return dummy.Next
}
