package LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//反转指定位置的链表内容
func ReverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	var prev *ListNode

	//head作为前一个值
	for m > 1 {
		prev = cur
		cur = cur.Next
		n--
		m--
	}

	con := prev
	tail := cur

	//反转的链表起始点
	var third *ListNode
	for n > 0 {
		third = cur.Next
		cur.Next = prev
		prev = cur
		cur = third
		n--
	}
	if con != nil {
		con.Next = prev
	} else {
		head = prev
	}
	tail.Next = cur
	return head
}
